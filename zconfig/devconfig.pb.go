// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devconfig.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SWAdapterType int32

const (
	SWAdapterType_IGNORE SWAdapterType = 0
	SWAdapterType_VLAN   SWAdapterType = 1
	SWAdapterType_BOND   SWAdapterType = 2
)

var SWAdapterType_name = map[int32]string{
	0: "IGNORE",
	1: "VLAN",
	2: "BOND",
}
var SWAdapterType_value = map[string]int32{
	"IGNORE": 0,
	"VLAN":   1,
	"BOND":   2,
}

func (x SWAdapterType) String() string {
	return proto.EnumName(SWAdapterType_name, int32(x))
}
func (SWAdapterType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type MapServer struct {
	NameOrIp   string `protobuf:"bytes,1,opt,name=NameOrIp" json:"NameOrIp,omitempty"`
	Credential string `protobuf:"bytes,2,opt,name=Credential" json:"Credential,omitempty"`
}

func (m *MapServer) Reset()                    { *m = MapServer{} }
func (m *MapServer) String() string            { return proto.CompactTextString(m) }
func (*MapServer) ProtoMessage()               {}
func (*MapServer) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *MapServer) GetNameOrIp() string {
	if m != nil {
		return m.NameOrIp
	}
	return ""
}

func (m *MapServer) GetCredential() string {
	if m != nil {
		return m.Credential
	}
	return ""
}

type ZedServer struct {
	HostName string   `protobuf:"bytes,1,opt,name=HostName" json:"HostName,omitempty"`
	EID      []string `protobuf:"bytes,2,rep,name=EID" json:"EID,omitempty"`
}

func (m *ZedServer) Reset()                    { *m = ZedServer{} }
func (m *ZedServer) String() string            { return proto.CompactTextString(m) }
func (*ZedServer) ProtoMessage()               {}
func (*ZedServer) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *ZedServer) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *ZedServer) GetEID() []string {
	if m != nil {
		return m.EID
	}
	return nil
}

type DeviceLispDetails struct {
	LispMapServers         []*MapServer `protobuf:"bytes,1,rep,name=LispMapServers" json:"LispMapServers,omitempty"`
	LispInstance           uint32       `protobuf:"varint,2,opt,name=LispInstance" json:"LispInstance,omitempty"`
	EID                    string       `protobuf:"bytes,4,opt,name=EID" json:"EID,omitempty"`
	EIDHashLen             uint32       `protobuf:"varint,5,opt,name=EIDHashLen" json:"EIDHashLen,omitempty"`
	ZedServers             []*ZedServer `protobuf:"bytes,6,rep,name=ZedServers" json:"ZedServers,omitempty"`
	EidAllocationPrefix    []byte       `protobuf:"bytes,8,opt,name=EidAllocationPrefix,proto3" json:"EidAllocationPrefix,omitempty"`
	EidAllocationPrefixLen uint32       `protobuf:"varint,9,opt,name=EidAllocationPrefixLen" json:"EidAllocationPrefixLen,omitempty"`
	ClientAddr             string       `protobuf:"bytes,10,opt,name=ClientAddr" json:"ClientAddr,omitempty"`
	Experimental           bool         `protobuf:"varint,20,opt,name=Experimental" json:"Experimental,omitempty"`
}

func (m *DeviceLispDetails) Reset()                    { *m = DeviceLispDetails{} }
func (m *DeviceLispDetails) String() string            { return proto.CompactTextString(m) }
func (*DeviceLispDetails) ProtoMessage()               {}
func (*DeviceLispDetails) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *DeviceLispDetails) GetLispMapServers() []*MapServer {
	if m != nil {
		return m.LispMapServers
	}
	return nil
}

func (m *DeviceLispDetails) GetLispInstance() uint32 {
	if m != nil {
		return m.LispInstance
	}
	return 0
}

func (m *DeviceLispDetails) GetEID() string {
	if m != nil {
		return m.EID
	}
	return ""
}

func (m *DeviceLispDetails) GetEIDHashLen() uint32 {
	if m != nil {
		return m.EIDHashLen
	}
	return 0
}

func (m *DeviceLispDetails) GetZedServers() []*ZedServer {
	if m != nil {
		return m.ZedServers
	}
	return nil
}

func (m *DeviceLispDetails) GetEidAllocationPrefix() []byte {
	if m != nil {
		return m.EidAllocationPrefix
	}
	return nil
}

func (m *DeviceLispDetails) GetEidAllocationPrefixLen() uint32 {
	if m != nil {
		return m.EidAllocationPrefixLen
	}
	return 0
}

func (m *DeviceLispDetails) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

func (m *DeviceLispDetails) GetExperimental() bool {
	if m != nil {
		return m.Experimental
	}
	return false
}

// Device Operational Commands Semantic
// For rebooting device,  command=Reset, counter = counter+delta, desiredState = on
// For poweroff device,  command=Reset, counter = counter+delta, desiredState = off
// For backup at midnight, command=Backup, counter = counter+delta, desiredState=n/a, opsTime = mm/dd/yy:hh:ss
// Current implementation does support only single command outstanding for each type
// In future can be extended to have more scheduled events
//
type DeviceOpsCmd struct {
	Counter      uint32 `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
	DesiredState bool   `protobuf:"varint,3,opt,name=desiredState" json:"desiredState,omitempty"`
	// FIXME: change to timestamp, once we move to gogo proto
	OpsTime string `protobuf:"bytes,4,opt,name=opsTime" json:"opsTime,omitempty"`
}

func (m *DeviceOpsCmd) Reset()                    { *m = DeviceOpsCmd{} }
func (m *DeviceOpsCmd) String() string            { return proto.CompactTextString(m) }
func (*DeviceOpsCmd) ProtoMessage()               {}
func (*DeviceOpsCmd) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *DeviceOpsCmd) GetCounter() uint32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *DeviceOpsCmd) GetDesiredState() bool {
	if m != nil {
		return m.DesiredState
	}
	return false
}

func (m *DeviceOpsCmd) GetOpsTime() string {
	if m != nil {
		return m.OpsTime
	}
	return ""
}

type SWAdapterParams struct {
	AType SWAdapterType `protobuf:"varint,1,opt,name=aType,enum=SWAdapterType" json:"aType,omitempty"`
	// vlan
	UnderlayInterface string `protobuf:"bytes,8,opt,name=underlayInterface" json:"underlayInterface,omitempty"`
	VlanId            uint32 `protobuf:"varint,9,opt,name=vlanId" json:"vlanId,omitempty"`
	// OR : repeated physical interfaces for bond0
	Bondgroup []string `protobuf:"bytes,10,rep,name=bondgroup" json:"bondgroup,omitempty"`
}

func (m *SWAdapterParams) Reset()                    { *m = SWAdapterParams{} }
func (m *SWAdapterParams) String() string            { return proto.CompactTextString(m) }
func (*SWAdapterParams) ProtoMessage()               {}
func (*SWAdapterParams) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *SWAdapterParams) GetAType() SWAdapterType {
	if m != nil {
		return m.AType
	}
	return SWAdapterType_IGNORE
}

func (m *SWAdapterParams) GetUnderlayInterface() string {
	if m != nil {
		return m.UnderlayInterface
	}
	return ""
}

func (m *SWAdapterParams) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

func (m *SWAdapterParams) GetBondgroup() []string {
	if m != nil {
		return m.Bondgroup
	}
	return nil
}

type SystemAdapter struct {
	// name of the adapter
	Name         string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	AllocDetails *SWAdapterParams `protobuf:"bytes,20,opt,name=allocDetails" json:"allocDetails,omitempty"`
	// this is part of the freelink group
	FreeUplink bool `protobuf:"varint,2,opt,name=freeUplink" json:"freeUplink,omitempty"`
	// this is part of the uplink group
	Uplink bool `protobuf:"varint,3,opt,name=uplink" json:"uplink,omitempty"`
	// attach this network config for this adapter
	NetworkUUID string `protobuf:"bytes,4,opt,name=networkUUID" json:"networkUUID,omitempty"`
	// if its static network we need ip address
	Addr string `protobuf:"bytes,5,opt,name=addr" json:"addr,omitempty"`
}

func (m *SystemAdapter) Reset()                    { *m = SystemAdapter{} }
func (m *SystemAdapter) String() string            { return proto.CompactTextString(m) }
func (*SystemAdapter) ProtoMessage()               {}
func (*SystemAdapter) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *SystemAdapter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SystemAdapter) GetAllocDetails() *SWAdapterParams {
	if m != nil {
		return m.AllocDetails
	}
	return nil
}

func (m *SystemAdapter) GetFreeUplink() bool {
	if m != nil {
		return m.FreeUplink
	}
	return false
}

func (m *SystemAdapter) GetUplink() bool {
	if m != nil {
		return m.Uplink
	}
	return false
}

func (m *SystemAdapter) GetNetworkUUID() string {
	if m != nil {
		return m.NetworkUUID
	}
	return ""
}

func (m *SystemAdapter) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type EdgeDevConfig struct {
	Id                 *UUIDandVersion          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DevConfigSha256    []byte                   `protobuf:"bytes,2,opt,name=devConfigSha256,proto3" json:"devConfigSha256,omitempty"`
	DevConfigSignature []byte                   `protobuf:"bytes,3,opt,name=devConfigSignature,proto3" json:"devConfigSignature,omitempty"`
	Apps               []*AppInstanceConfig     `protobuf:"bytes,4,rep,name=apps" json:"apps,omitempty"`
	Networks           []*NetworkConfig         `protobuf:"bytes,5,rep,name=networks" json:"networks,omitempty"`
	Datastores         []*DatastoreConfig       `protobuf:"bytes,6,rep,name=datastores" json:"datastores,omitempty"`
	LispInfo           *DeviceLispDetails       `protobuf:"bytes,7,opt,name=lispInfo" json:"lispInfo,omitempty"`
	Base               []*BaseOSConfig          `protobuf:"bytes,8,rep,name=base" json:"base,omitempty"`
	Reboot             *DeviceOpsCmd            `protobuf:"bytes,9,opt,name=reboot" json:"reboot,omitempty"`
	Backup             *DeviceOpsCmd            `protobuf:"bytes,10,opt,name=backup" json:"backup,omitempty"`
	ConfigItems        []*ConfigItem            `protobuf:"bytes,11,rep,name=configItems" json:"configItems,omitempty"`
	SystemAdapterList  []*SystemAdapter         `protobuf:"bytes,12,rep,name=systemAdapterList" json:"systemAdapterList,omitempty"`
	Services           []*ServiceInstanceConfig `protobuf:"bytes,13,rep,name=services" json:"services,omitempty"`
	// Override dmidecode info if set
	Manufacturer string `protobuf:"bytes,14,opt,name=manufacturer" json:"manufacturer,omitempty"`
	ProductName  string `protobuf:"bytes,15,opt,name=productName" json:"productName,omitempty"`
}

func (m *EdgeDevConfig) Reset()                    { *m = EdgeDevConfig{} }
func (m *EdgeDevConfig) String() string            { return proto.CompactTextString(m) }
func (*EdgeDevConfig) ProtoMessage()               {}
func (*EdgeDevConfig) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *EdgeDevConfig) GetId() *UUIDandVersion {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *EdgeDevConfig) GetDevConfigSha256() []byte {
	if m != nil {
		return m.DevConfigSha256
	}
	return nil
}

func (m *EdgeDevConfig) GetDevConfigSignature() []byte {
	if m != nil {
		return m.DevConfigSignature
	}
	return nil
}

func (m *EdgeDevConfig) GetApps() []*AppInstanceConfig {
	if m != nil {
		return m.Apps
	}
	return nil
}

func (m *EdgeDevConfig) GetNetworks() []*NetworkConfig {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *EdgeDevConfig) GetDatastores() []*DatastoreConfig {
	if m != nil {
		return m.Datastores
	}
	return nil
}

func (m *EdgeDevConfig) GetLispInfo() *DeviceLispDetails {
	if m != nil {
		return m.LispInfo
	}
	return nil
}

func (m *EdgeDevConfig) GetBase() []*BaseOSConfig {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *EdgeDevConfig) GetReboot() *DeviceOpsCmd {
	if m != nil {
		return m.Reboot
	}
	return nil
}

func (m *EdgeDevConfig) GetBackup() *DeviceOpsCmd {
	if m != nil {
		return m.Backup
	}
	return nil
}

func (m *EdgeDevConfig) GetConfigItems() []*ConfigItem {
	if m != nil {
		return m.ConfigItems
	}
	return nil
}

func (m *EdgeDevConfig) GetSystemAdapterList() []*SystemAdapter {
	if m != nil {
		return m.SystemAdapterList
	}
	return nil
}

func (m *EdgeDevConfig) GetServices() []*ServiceInstanceConfig {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *EdgeDevConfig) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *EdgeDevConfig) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

// Timers and other per-device policy which relates to the interaction
// with zedcloud. Note that the timers are randomized on the device
// to avoid synchronization with other devices. Random range is between
// between .5 and 1.5 of these nominal values. If not set (i.e. zero),
// it means the default value of 60 seconds.
// Initially we'll have a "configinterval" and a "metricinterval" item.
// We'll also need a "resetIfCloudGoneTime" and a "fallbackIfCloudGoneTime"
// to control a normal operation and an upgrade inprogress check of
// cloud connectivity.
type ConfigItem struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Types that are valid to be assigned to ConfigItemValue:
	//	*ConfigItem_BoolValue
	//	*ConfigItem_Uint32Value
	//	*ConfigItem_Uint64Value
	//	*ConfigItem_FloatValue
	//	*ConfigItem_StringValue
	ConfigItemValue isConfigItem_ConfigItemValue `protobuf_oneof:"configItemValue"`
}

func (m *ConfigItem) Reset()                    { *m = ConfigItem{} }
func (m *ConfigItem) String() string            { return proto.CompactTextString(m) }
func (*ConfigItem) ProtoMessage()               {}
func (*ConfigItem) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

type isConfigItem_ConfigItemValue interface {
	isConfigItem_ConfigItemValue()
}

type ConfigItem_BoolValue struct {
	BoolValue bool `protobuf:"varint,3,opt,name=boolValue,oneof"`
}
type ConfigItem_Uint32Value struct {
	Uint32Value uint32 `protobuf:"varint,4,opt,name=uint32Value,oneof"`
}
type ConfigItem_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,5,opt,name=uint64Value,oneof"`
}
type ConfigItem_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,6,opt,name=floatValue,oneof"`
}
type ConfigItem_StringValue struct {
	StringValue string `protobuf:"bytes,7,opt,name=stringValue,oneof"`
}

func (*ConfigItem_BoolValue) isConfigItem_ConfigItemValue()   {}
func (*ConfigItem_Uint32Value) isConfigItem_ConfigItemValue() {}
func (*ConfigItem_Uint64Value) isConfigItem_ConfigItemValue() {}
func (*ConfigItem_FloatValue) isConfigItem_ConfigItemValue()  {}
func (*ConfigItem_StringValue) isConfigItem_ConfigItemValue() {}

func (m *ConfigItem) GetConfigItemValue() isConfigItem_ConfigItemValue {
	if m != nil {
		return m.ConfigItemValue
	}
	return nil
}

func (m *ConfigItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigItem) GetBoolValue() bool {
	if x, ok := m.GetConfigItemValue().(*ConfigItem_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *ConfigItem) GetUint32Value() uint32 {
	if x, ok := m.GetConfigItemValue().(*ConfigItem_Uint32Value); ok {
		return x.Uint32Value
	}
	return 0
}

func (m *ConfigItem) GetUint64Value() uint64 {
	if x, ok := m.GetConfigItemValue().(*ConfigItem_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *ConfigItem) GetFloatValue() float32 {
	if x, ok := m.GetConfigItemValue().(*ConfigItem_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *ConfigItem) GetStringValue() string {
	if x, ok := m.GetConfigItemValue().(*ConfigItem_StringValue); ok {
		return x.StringValue
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConfigItem) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConfigItem_OneofMarshaler, _ConfigItem_OneofUnmarshaler, _ConfigItem_OneofSizer, []interface{}{
		(*ConfigItem_BoolValue)(nil),
		(*ConfigItem_Uint32Value)(nil),
		(*ConfigItem_Uint64Value)(nil),
		(*ConfigItem_FloatValue)(nil),
		(*ConfigItem_StringValue)(nil),
	}
}

func _ConfigItem_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConfigItem)
	// configItemValue
	switch x := m.ConfigItemValue.(type) {
	case *ConfigItem_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *ConfigItem_Uint32Value:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint32Value))
	case *ConfigItem_Uint64Value:
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint64Value))
	case *ConfigItem_FloatValue:
		b.EncodeVarint(6<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FloatValue)))
	case *ConfigItem_StringValue:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case nil:
	default:
		return fmt.Errorf("ConfigItem.ConfigItemValue has unexpected type %T", x)
	}
	return nil
}

func _ConfigItem_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConfigItem)
	switch tag {
	case 3: // configItemValue.boolValue
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConfigItemValue = &ConfigItem_BoolValue{x != 0}
		return true, err
	case 4: // configItemValue.uint32Value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConfigItemValue = &ConfigItem_Uint32Value{uint32(x)}
		return true, err
	case 5: // configItemValue.uint64Value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConfigItemValue = &ConfigItem_Uint64Value{x}
		return true, err
	case 6: // configItemValue.floatValue
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.ConfigItemValue = &ConfigItem_FloatValue{math.Float32frombits(uint32(x))}
		return true, err
	case 7: // configItemValue.stringValue
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ConfigItemValue = &ConfigItem_StringValue{x}
		return true, err
	default:
		return false, nil
	}
}

func _ConfigItem_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConfigItem)
	// configItemValue
	switch x := m.ConfigItemValue.(type) {
	case *ConfigItem_BoolValue:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += 1
	case *ConfigItem_Uint32Value:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Uint32Value))
	case *ConfigItem_Uint64Value:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Uint64Value))
	case *ConfigItem_FloatValue:
		n += proto.SizeVarint(6<<3 | proto.WireFixed32)
		n += 4
	case *ConfigItem_StringValue:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ConfigRequest struct {
	ConfigHash string `protobuf:"bytes,1,opt,name=configHash" json:"configHash,omitempty"`
}

func (m *ConfigRequest) Reset()                    { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()               {}
func (*ConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *ConfigRequest) GetConfigHash() string {
	if m != nil {
		return m.ConfigHash
	}
	return ""
}

type ConfigResponse struct {
	Config     *EdgeDevConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	ConfigHash string         `protobuf:"bytes,2,opt,name=configHash" json:"configHash,omitempty"`
}

func (m *ConfigResponse) Reset()                    { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()               {}
func (*ConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *ConfigResponse) GetConfig() *EdgeDevConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ConfigResponse) GetConfigHash() string {
	if m != nil {
		return m.ConfigHash
	}
	return ""
}

func init() {
	proto.RegisterType((*MapServer)(nil), "MapServer")
	proto.RegisterType((*ZedServer)(nil), "ZedServer")
	proto.RegisterType((*DeviceLispDetails)(nil), "DeviceLispDetails")
	proto.RegisterType((*DeviceOpsCmd)(nil), "DeviceOpsCmd")
	proto.RegisterType((*SWAdapterParams)(nil), "sWAdapterParams")
	proto.RegisterType((*SystemAdapter)(nil), "SystemAdapter")
	proto.RegisterType((*EdgeDevConfig)(nil), "EdgeDevConfig")
	proto.RegisterType((*ConfigItem)(nil), "ConfigItem")
	proto.RegisterType((*ConfigRequest)(nil), "ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "ConfigResponse")
	proto.RegisterEnum("SWAdapterType", SWAdapterType_name, SWAdapterType_value)
}

func init() { proto.RegisterFile("devconfig.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 1054 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x8e, 0x1d, 0xc7, 0xb5, 0x8f, 0x7f, 0xe2, 0x70, 0x43, 0x21, 0x14, 0x5b, 0xeb, 0x09, 0xdb,
	0x60, 0x14, 0x9b, 0x52, 0xb8, 0x5d, 0x80, 0x01, 0xbb, 0x58, 0x12, 0x1b, 0x8d, 0x81, 0x2c, 0x29,
	0xe8, 0x26, 0x1b, 0x7a, 0x47, 0x4b, 0xb4, 0x23, 0x44, 0x22, 0x35, 0x92, 0xca, 0x9a, 0x3e, 0xc9,
	0x80, 0x3d, 0xd0, 0x1e, 0x63, 0x8f, 0xb1, 0xdb, 0x81, 0x3f, 0x92, 0x65, 0x27, 0xbb, 0xe3, 0xf9,
	0xbe, 0x8f, 0x47, 0x87, 0xe7, 0xcf, 0x86, 0xfd, 0x88, 0xde, 0x85, 0x9c, 0x2d, 0xe3, 0x55, 0x90,
	0x09, 0xae, 0xf8, 0x33, 0x0b, 0xa4, 0x29, 0x67, 0x05, 0x40, 0xb2, 0x6c, 0x43, 0x81, 0x16, 0x44,
	0x52, 0x2e, 0x37, 0x6f, 0x31, 0xaa, 0x36, 0x80, 0x9e, 0x54, 0x5c, 0x90, 0x15, 0x2d, 0x4d, 0x2a,
	0xee, 0xe2, 0xd0, 0x99, 0xfe, 0x5b, 0x68, 0xff, 0x42, 0xb2, 0x39, 0x15, 0x77, 0x54, 0xa0, 0x67,
	0xd0, 0xba, 0x20, 0x29, 0xbd, 0x14, 0xb3, 0xcc, 0xab, 0x0d, 0x6b, 0xa3, 0x36, 0x2e, 0x6d, 0xf4,
	0x1c, 0xe0, 0x54, 0xd0, 0x88, 0x32, 0x15, 0x93, 0xc4, 0xab, 0x1b, 0xb6, 0x82, 0xf8, 0x3f, 0x42,
	0xfb, 0x03, 0x8d, 0xd6, 0x8e, 0xce, 0xb8, 0x54, 0xfa, 0x72, 0xe1, 0xa8, 0xb0, 0xd1, 0x00, 0x76,
	0xa7, 0xb3, 0x89, 0x57, 0x1f, 0xee, 0x8e, 0xda, 0x58, 0x1f, 0xfd, 0x7f, 0xeb, 0x70, 0x30, 0xa1,
	0x3a, 0xa8, 0xf3, 0x58, 0x66, 0x13, 0xaa, 0x48, 0x9c, 0x48, 0x34, 0x86, 0xbe, 0x36, 0xcb, 0xe8,
	0xa4, 0x57, 0x1b, 0xee, 0x8e, 0x3a, 0x63, 0x08, 0x4a, 0x08, 0x6f, 0x29, 0x90, 0x0f, 0x5d, 0x8d,
	0xcc, 0x98, 0x54, 0x84, 0x85, 0xd4, 0x84, 0xd9, 0xc3, 0x1b, 0x58, 0xf1, 0xfd, 0x86, 0x09, 0x4b,
	0x1f, 0xf5, 0xd3, 0xa6, 0xb3, 0xc9, 0x19, 0x91, 0x37, 0xe7, 0x94, 0x79, 0x7b, 0xe6, 0x4e, 0x05,
	0x41, 0x2f, 0x01, 0xca, 0xa7, 0x49, 0xaf, 0xe9, 0xa2, 0x28, 0x21, 0x5c, 0x61, 0xd1, 0x2b, 0xf8,
	0x6c, 0x1a, 0x47, 0xc7, 0x49, 0xc2, 0x43, 0xa2, 0x62, 0xce, 0xde, 0x09, 0xba, 0x8c, 0x3f, 0x7a,
	0xad, 0x61, 0x6d, 0xd4, 0xc5, 0x8f, 0x51, 0xe8, 0x08, 0x9e, 0x3e, 0x02, 0xeb, 0x48, 0xda, 0x26,
	0x92, 0xff, 0x61, 0x4d, 0x41, 0x92, 0x98, 0x32, 0x75, 0x1c, 0x45, 0xc2, 0x03, 0x57, 0x90, 0x12,
	0xd1, 0xb9, 0x98, 0x7e, 0xcc, 0xa8, 0x88, 0x53, 0xca, 0x14, 0x49, 0xbc, 0xcf, 0x87, 0xb5, 0x51,
	0x0b, 0x6f, 0x60, 0xfe, 0x12, 0xba, 0x36, 0xf1, 0x97, 0x99, 0x3c, 0x4d, 0x23, 0xe4, 0xc1, 0x93,
	0x90, 0xe7, 0x4c, 0x51, 0xe1, 0x52, 0x57, 0x98, 0xda, 0x5b, 0x44, 0x65, 0x2c, 0x68, 0x34, 0x57,
	0x44, 0x51, 0x6f, 0xd7, 0x7a, 0xab, 0x62, 0xfa, 0x36, 0xcf, 0xe4, 0xfb, 0x38, 0xa5, 0x2e, 0xbb,
	0x85, 0xe9, 0xff, 0x55, 0x83, 0x7d, 0xf9, 0xeb, 0x71, 0x44, 0x32, 0x45, 0xc5, 0x3b, 0x22, 0x48,
	0x2a, 0xd1, 0xd7, 0xb0, 0x47, 0xde, 0xdf, 0x67, 0xb6, 0x41, 0xfa, 0xe3, 0x7e, 0x50, 0x0a, 0x34,
	0x8a, 0x2d, 0x89, 0xbe, 0x83, 0x83, 0x9c, 0x45, 0x54, 0x24, 0xe4, 0x7e, 0xa6, 0x03, 0x59, 0x92,
	0x90, 0x9a, 0x6c, 0xb6, 0xf1, 0x43, 0x02, 0x3d, 0x85, 0xe6, 0x5d, 0x42, 0xd8, 0x2c, 0x72, 0xb9,
	0x73, 0x16, 0xfa, 0x02, 0xda, 0x0b, 0xce, 0xa2, 0x95, 0xe0, 0x79, 0xe6, 0x81, 0xe9, 0xbc, 0x35,
	0xe0, 0xff, 0x5d, 0x83, 0xde, 0xfc, 0x5e, 0x2a, 0x9a, 0xba, 0x00, 0x10, 0x82, 0x06, 0x5b, 0xf7,
	0xae, 0x39, 0xa3, 0x37, 0xd0, 0x25, 0xba, 0x0c, 0xae, 0x3f, 0x4d, 0x3e, 0x3b, 0xe3, 0x41, 0xb0,
	0xf5, 0x2e, 0xbc, 0xa1, 0xd2, 0x55, 0x5a, 0x0a, 0x4a, 0xaf, 0xb2, 0x24, 0x66, 0xb7, 0x26, 0xa9,
	0x2d, 0x5c, 0x41, 0x74, 0xc4, 0xb9, 0xe5, 0x6c, 0x46, 0x9d, 0x85, 0x86, 0xd0, 0x61, 0x54, 0xfd,
	0xc1, 0xc5, 0xed, 0xd5, 0x55, 0xd9, 0xad, 0x55, 0x48, 0xc7, 0x48, 0x74, 0xe5, 0xf7, 0x6c, 0x8c,
	0xfa, 0xec, 0xff, 0xb9, 0x07, 0xbd, 0x69, 0xb4, 0xa2, 0x13, 0x7a, 0x77, 0x6a, 0x76, 0x00, 0x7a,
	0x01, 0xf5, 0x38, 0x32, 0xef, 0xe8, 0x8c, 0xf7, 0x03, 0x7d, 0x91, 0xb0, 0xe8, 0x9a, 0x0a, 0x19,
	0x73, 0x86, 0xeb, 0x71, 0x84, 0x46, 0x66, 0xf1, 0x58, 0xf5, 0xfc, 0x86, 0x8c, 0x7f, 0x38, 0x32,
	0x51, 0x76, 0xf1, 0x36, 0x8c, 0x02, 0x40, 0x6b, 0x28, 0x5e, 0x31, 0xa2, 0x72, 0x61, 0x1b, 0xa1,
	0x8b, 0x1f, 0x61, 0xd0, 0xb7, 0xd0, 0x20, 0x59, 0x26, 0xbd, 0x86, 0x19, 0x18, 0x14, 0x1c, 0x67,
	0xe5, 0x10, 0x5a, 0x29, 0x36, 0x3c, 0x7a, 0x09, 0x2d, 0xf7, 0x2e, 0xe9, 0xed, 0x19, 0x6d, 0x3f,
	0xb8, 0xb0, 0x80, 0xd3, 0x95, 0x3c, 0x7a, 0x05, 0x10, 0x11, 0x45, 0xf4, 0x4a, 0xa3, 0xc5, 0x28,
	0x0e, 0x82, 0x49, 0x01, 0x39, 0x7d, 0x45, 0x83, 0x02, 0x68, 0x25, 0x66, 0xfc, 0x97, 0xdc, 0x7b,
	0x62, 0xd2, 0x80, 0x82, 0x07, 0xcb, 0x06, 0x97, 0x1a, 0xf4, 0x15, 0x34, 0xf4, 0x56, 0xf5, 0x5a,
	0xc6, 0x77, 0x2f, 0x38, 0x21, 0x92, 0x5e, 0xce, 0x8b, 0x80, 0x35, 0x85, 0xbe, 0x81, 0xa6, 0xa0,
	0x0b, 0xce, 0x95, 0xe9, 0x32, 0x2d, 0xaa, 0x0e, 0x11, 0x76, 0xa4, 0x96, 0x2d, 0x48, 0x78, 0x6b,
	0x3a, 0xee, 0x31, 0x99, 0x25, 0xd1, 0xf7, 0xd0, 0xb1, 0xfb, 0x7a, 0xa6, 0x68, 0x2a, 0xbd, 0x8e,
	0xf9, 0x6e, 0x27, 0x38, 0x2d, 0x31, 0x5c, 0xe5, 0xd1, 0x4f, 0x70, 0x20, 0xab, 0xbd, 0x7a, 0x1e,
	0x4b, 0xe5, 0x75, 0x5d, 0xda, 0x36, 0xba, 0x18, 0x3f, 0x14, 0xa2, 0x31, 0xb4, 0xdc, 0xfe, 0x97,
	0x5e, 0xcf, 0x5c, 0x7a, 0x1a, 0xcc, 0x2d, 0xb0, 0x55, 0x9b, 0x52, 0xa7, 0x47, 0x3f, 0x25, 0x2c,
	0x5f, 0x92, 0x50, 0x97, 0x55, 0x78, 0x7d, 0xd3, 0x70, 0x1b, 0x98, 0x6e, 0xd7, 0x4c, 0xf0, 0x28,
	0x0f, 0xed, 0xce, 0xdf, 0xb7, 0xed, 0x5a, 0x81, 0xfc, 0x7f, 0x6a, 0x00, 0xeb, 0x37, 0xe9, 0x2d,
	0x7c, 0x4b, 0xef, 0xdd, 0x80, 0xe9, 0x23, 0x7a, 0xae, 0x67, 0x94, 0x27, 0xd7, 0x24, 0xc9, 0xdd,
	0x7a, 0x39, 0xdb, 0xc1, 0x6b, 0x08, 0xf9, 0xd0, 0xc9, 0x63, 0xa6, 0x5e, 0x8f, 0xad, 0x42, 0x4f,
	0x44, 0xef, 0x6c, 0x07, 0x57, 0xc1, 0x42, 0x73, 0xf4, 0xc6, 0x6a, 0xf4, 0x68, 0x34, 0x0a, 0x8d,
	0x03, 0xd1, 0x10, 0x60, 0x99, 0x70, 0xa2, 0xac, 0xa4, 0x39, 0xac, 0x8d, 0xea, 0x67, 0x3b, 0xb8,
	0x82, 0x69, 0x2f, 0x52, 0x89, 0x98, 0xad, 0xac, 0x44, 0x77, 0x4d, 0x5b, 0x7b, 0xa9, 0x80, 0x27,
	0x07, 0xb0, 0xbf, 0xae, 0x8a, 0x81, 0xfc, 0x43, 0xe8, 0xb9, 0xdc, 0xd1, 0xdf, 0x73, 0x2a, 0x95,
	0x9e, 0x7d, 0xab, 0xd1, 0x3f, 0x24, 0xee, 0xa9, 0x15, 0xc4, 0xff, 0x0d, 0xfa, 0xc5, 0x05, 0x99,
	0x71, 0x26, 0xf5, 0xc8, 0x34, 0x2d, 0xef, 0x26, 0xb6, 0x1f, 0x6c, 0x4c, 0x33, 0x76, 0xec, 0x96,
	0xe7, 0xfa, 0xb6, 0xe7, 0x97, 0x87, 0xd0, 0xdb, 0xd8, 0xa6, 0x08, 0xa0, 0x39, 0x7b, 0x7b, 0x71,
	0x89, 0xa7, 0x83, 0x1d, 0xd4, 0x82, 0xc6, 0xf5, 0xf9, 0xf1, 0xc5, 0xa0, 0xa6, 0x4f, 0x27, 0x97,
	0x17, 0x93, 0x41, 0xfd, 0xe4, 0x67, 0x78, 0x11, 0xf2, 0x34, 0xf8, 0x44, 0x23, 0x1a, 0x91, 0x20,
	0x4c, 0x78, 0x1e, 0x05, 0xf9, 0xc6, 0x3f, 0x85, 0x0f, 0x5f, 0xae, 0x62, 0x75, 0x93, 0x2f, 0x82,
	0x90, 0xa7, 0x87, 0x56, 0x77, 0x48, 0xb2, 0xf8, 0xf0, 0x93, 0xfd, 0xec, 0xa2, 0x69, 0x54, 0xaf,
	0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x67, 0xc7, 0x65, 0xc7, 0x08, 0x00, 0x00,
}
