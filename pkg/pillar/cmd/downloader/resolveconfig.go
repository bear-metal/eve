// Copyright (c) 2019-2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package downloader

import (
	"fmt"
	"strings"
	"time"

	zconfig "github.com/lf-edge/eve/api/go/config"
	"github.com/lf-edge/eve/libs/zedUpload"
	"github.com/lf-edge/eve/pkg/pillar/flextimer"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/lf-edge/eve/pkg/pillar/utils"
	uuid "github.com/satori/go.uuid"
)

func runResolveHandler(ctx *downloaderContext, key string, c <-chan Notify) {

	log.Functionf("runResolveHandler starting")

	max := float64(retryTime)
	min := max * 0.3
	ticker := flextimer.NewRangeTicker(time.Duration(min),
		time.Duration(max))
	closed := false
	for !closed {
		select {
		case _, ok := <-c:
			if ok {
				rc := lookupResolveConfig(ctx, key)
				resolveTagsToHash(ctx, *rc)
				// XXX if err start timer
			} else {
				// Closed
				rs := lookupResolveStatus(ctx, key)
				if rs != nil {
					unpublishResolveStatus(ctx, rs)
				}
				closed = true
				// XXX stop timer
			}
		case <-ticker.C:
			log.Tracef("runResolveHandler(%s) timer", key)
			rs := lookupResolveStatus(ctx, key)
			if rs != nil {
				maybeRetryResolve(ctx, rs)
			}
		}
	}
	log.Functionf("runResolveHandler(%s) DONE", key)
}

func maybeRetryResolve(ctx *downloaderContext, status *types.ResolveStatus) {

	// object is either in download progress or,
	// successfully downloaded, nothing to do
	if !status.HasError() {
		return
	}
	t := time.Now()
	elapsed := t.Sub(status.ErrorTime)
	if elapsed < retryTime {
		log.Functionf("maybeRetryResolve(%s) %d remaining",
			status.Key(),
			(retryTime-elapsed)/time.Second)
		return
	}
	log.Functionf("maybeRetryResolve(%s) after %s at %v",
		status.Key(), status.Error, status.ErrorTime)

	config := lookupResolveConfig(ctx, status.Key())
	if config == nil {
		log.Functionf("maybeRetryResolve(%s) no config",
			status.Key())
		return
	}

	// reset Error, to start download again
	status.RetryCount++
	status.ClearError()
	publishResolveStatus(ctx, status)

	resolveTagsToHash(ctx, *config)
}

func publishResolveStatus(ctx *downloaderContext,
	status *types.ResolveStatus) {

	key := status.Key()
	log.Tracef("publishResolveStatus(%s)", key)
	pub := ctx.pubResolveStatus
	pub.Publish(key, *status)
	log.Tracef("publishResolveStatus(%s) Done", key)
}

func unpublishResolveStatus(ctx *downloaderContext,
	status *types.ResolveStatus) {

	key := status.Key()
	log.Tracef("unpublishResolveStatus(%s)", key)
	pub := ctx.pubResolveStatus
	pub.Unpublish(key)
	log.Tracef("unpublishResolveStatus(%s) Done", key)
}

func lookupResolveConfig(ctx *downloaderContext, key string) *types.ResolveConfig {

	sub := ctx.subResolveConfig
	c, _ := sub.Get(key)
	if c == nil {
		log.Functionf("lookupResolveConfig(%s) not found", key)
		return nil
	}
	config := c.(types.ResolveConfig)
	return &config
}

func lookupResolveStatus(ctx *downloaderContext,
	key string) *types.ResolveStatus {

	pub := ctx.pubResolveStatus
	c, _ := pub.Get(key)
	if c == nil {
		log.Functionf("lookupResolveStatus(%s) not found", key)
		return nil
	}
	status := c.(types.ResolveStatus)
	return &status
}

func resolveTagsToHash(ctx *downloaderContext, rc types.ResolveConfig) {
	var (
		err                           error
		errStr, remoteName, serverURL string
		syncOp                        zedUpload.SyncOpType = zedUpload.SyncOpGetObjectMetaData
		trType                        zedUpload.SyncTransportType
		auth                          *zedUpload.AuthInput
	)

	rs := lookupResolveStatus(ctx, rc.Key())
	if rs == nil {
		rs = &types.ResolveStatus{
			DatastoreID: rc.DatastoreID,
			Name:        rc.Name,
			Counter:     rc.Counter,
		}
	}
	rs.ClearError()

	sha := maybeNameHasSha(rc.Name)
	if sha != "" {
		rs.ImageSha256 = sha
		publishResolveStatus(ctx, rs)
		return
	}

	dst, err := utils.LookupDatastoreConfig(ctx.subDatastoreConfig, rc.DatastoreID)
	if err != nil {
		rs.SetErrorNow(err.Error())
		publishResolveStatus(ctx, rs)
		return
	}
	log.Tracef("Found datastore(%s) for %s", rc.DatastoreID.String(), rc.Name)

	// construct the datastore context
	dsCtx, err := constructDatastoreContext(ctx, rc.Name, false, *dst)
	if err != nil {
		errStr := fmt.Sprintf("%s, Datastore construction failed, %s", rc.Name, err)
		rs.SetErrorNow(errStr)
		publishResolveStatus(ctx, rs)
		return
	}

	downloadMaxPortCost := ctx.downloadMaxPortCost
	log.Functionf("Resolving config <%s> using %d downloadMaxPortCost",
		rc.Name, downloadMaxPortCost)

	addrCount := types.CountLocalAddrNoLinkLocalWithCost(ctx.deviceNetworkStatus,
		downloadMaxPortCost)
	log.Functionf("Have %d management port addresses for cost %d",
		addrCount, downloadMaxPortCost)
	if addrCount == 0 {
		err := fmt.Errorf("No IP management port addresses for resolve with cost %d",
			downloadMaxPortCost)
		log.Error(err.Error())
		rs.SetErrorNow(err.Error())
		publishResolveStatus(ctx, rs)
		return
	}

	switch dsCtx.TransportMethod {
	case zconfig.DsType_DsContainerRegistry.String():
		auth = &zedUpload.AuthInput{
			AuthType: "apikey",
			Uname:    dsCtx.APIKey,
			Password: dsCtx.Password,
		}
		trType = zedUpload.SyncOCIRegistryTr
		// get the name of the repository and the URL for the registry
		serverURL, remoteName, err = ociRepositorySplit(dsCtx.DownloadURL)
		if err != nil {
			errStr = fmt.Sprintf("invalid OCI registry URL: %s", serverURL)
		}

	default:
		errStr = "unsupported transport method " + dsCtx.TransportMethod

	}

	// if there were any errors, do not bother continuing
	// ideally in go we would have this as a check for error
	// and return, but we will get to it later
	if errStr != "" {
		log.Errorf("Error preparing to download. All errors:%s", errStr)
		rs.SetErrorNow(errStr)
		publishResolveStatus(ctx, rs)
		return
	}

	// Loop through all interfaces until a success
	for addrIndex := 0; addrIndex < addrCount; addrIndex++ {
		ipSrc, err := types.GetLocalAddrNoLinkLocalWithCost(ctx.deviceNetworkStatus,
			addrIndex, "", downloadMaxPortCost)
		if err != nil {
			log.Errorf("GetLocalAddr failed: %s", err)
			errStr = errStr + "\n" + err.Error()
			continue
		}
		ifname := types.GetMgmtPortFromAddr(ctx.deviceNetworkStatus, ipSrc)
		log.Functionf("Using IP source %v if %s transport %v",
			ipSrc, ifname, dsCtx.TransportMethod)

		sha256, err := objectMetadata(ctx, trType, syncOp, serverURL, auth,
			dsCtx.Dpath, dsCtx.Region,
			ifname, ipSrc, remoteName)
		if err != nil {
			errStr = errStr + "\n" + err.Error()
			continue
		}
		rs.ImageSha256 = sha256
		publishResolveStatus(ctx, rs)
		return

	}
	log.Errorf("All source IP addresses failed. All errors:%s", errStr)
	rs.SetErrorNow(errStr)
	publishResolveStatus(ctx, rs)
}

func maybeNameHasSha(name string) string {
	if strings.Contains(name, "@sha256:") {
		parts := strings.Split(name, "@sha256:")
		return strings.ToUpper(parts[1])
	}
	return ""
}

//checkAndUpdateResolveConfig fires modify handler for ResolveConfig
//we need to call it in case of no DatastoreConfig found
func checkAndUpdateResolveConfig(ctx *downloaderContext, dsID uuid.UUID) {
	log.Functionf("checkAndUpdateResolveConfig for %s", dsID)
	resolveStatuses := ctx.pubResolveStatus.GetAll()
	for _, v := range resolveStatuses {
		status := v.(types.ResolveStatus)
		if status.DatastoreID == dsID {
			config := lookupResolveConfig(ctx, status.Key())
			if config != nil {
				resHandler.modify(ctx, status.Key(), *config, *config)
			}
		}
	}
	log.Functionf("checkAndUpdateResolveConfig for %s, done", dsID)
}
