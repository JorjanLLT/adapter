// Code generated by entproto. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: entpb/entpb.proto

package entpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Battery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sn             string                  `protobuf:"bytes,2,opt,name=sn,proto3" json:"sn,omitempty"`
	SoftVersion    *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=soft_version,json=softVersion,proto3" json:"soft_version,omitempty"`
	HardVersion    *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=hard_version,json=hardVersion,proto3" json:"hard_version,omitempty"`
	Soft_4GVersion *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=soft_4g_version,json=soft4gVersion,proto3" json:"soft_4g_version,omitempty"`
	Hard_4GVersion *wrapperspb.UInt32Value `protobuf:"bytes,6,opt,name=hard_4g_version,json=hard4gVersion,proto3" json:"hard_4g_version,omitempty"`
	Sn_4G          *wrapperspb.UInt64Value `protobuf:"bytes,7,opt,name=sn_4g,json=sn4g,proto3" json:"sn_4g,omitempty"`
	Iccid          *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=iccid,proto3" json:"iccid,omitempty"`
	Soc            *wrapperspb.UInt32Value `protobuf:"bytes,9,opt,name=soc,proto3" json:"soc,omitempty"`
	Heartbeats     []*Heartbeat            `protobuf:"bytes,10,rep,name=heartbeats,proto3" json:"heartbeats,omitempty"`
	Reigns         []*Reign                `protobuf:"bytes,11,rep,name=reigns,proto3" json:"reigns,omitempty"`
}

func (x *Battery) Reset() {
	*x = Battery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Battery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Battery) ProtoMessage() {}

func (x *Battery) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Battery.ProtoReflect.Descriptor instead.
func (*Battery) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{0}
}

func (x *Battery) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Battery) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

func (x *Battery) GetSoftVersion() *wrapperspb.UInt32Value {
	if x != nil {
		return x.SoftVersion
	}
	return nil
}

func (x *Battery) GetHardVersion() *wrapperspb.UInt32Value {
	if x != nil {
		return x.HardVersion
	}
	return nil
}

func (x *Battery) GetSoft_4GVersion() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Soft_4GVersion
	}
	return nil
}

func (x *Battery) GetHard_4GVersion() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Hard_4GVersion
	}
	return nil
}

func (x *Battery) GetSn_4G() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Sn_4G
	}
	return nil
}

func (x *Battery) GetIccid() *wrapperspb.StringValue {
	if x != nil {
		return x.Iccid
	}
	return nil
}

func (x *Battery) GetSoc() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Soc
	}
	return nil
}

func (x *Battery) GetHeartbeats() []*Heartbeat {
	if x != nil {
		return x.Heartbeats
	}
	return nil
}

func (x *Battery) GetReigns() []*Reign {
	if x != nil {
		return x.Reigns
	}
	return nil
}

type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sn                   string                 `protobuf:"bytes,2,opt,name=sn,proto3" json:"sn,omitempty"`
	BatteryId            int64                  `protobuf:"varint,3,opt,name=battery_id,json=batteryId,proto3" json:"battery_id,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Voltage              float64                `protobuf:"fixed64,5,opt,name=voltage,proto3" json:"voltage,omitempty"`
	Current              float64                `protobuf:"fixed64,6,opt,name=current,proto3" json:"current,omitempty"`
	Soc                  uint32                 `protobuf:"varint,7,opt,name=soc,proto3" json:"soc,omitempty"`
	Soh                  uint32                 `protobuf:"varint,8,opt,name=soh,proto3" json:"soh,omitempty"`
	InCabinet            bool                   `protobuf:"varint,9,opt,name=in_cabinet,json=inCabinet,proto3" json:"in_cabinet,omitempty"`
	Capacity             float64                `protobuf:"fixed64,10,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MonMaxVoltage        uint32                 `protobuf:"varint,11,opt,name=mon_max_voltage,json=monMaxVoltage,proto3" json:"mon_max_voltage,omitempty"`
	MonMaxVoltagePos     uint32                 `protobuf:"varint,12,opt,name=mon_max_voltage_pos,json=monMaxVoltagePos,proto3" json:"mon_max_voltage_pos,omitempty"`
	MonMinVoltage        uint32                 `protobuf:"varint,13,opt,name=mon_min_voltage,json=monMinVoltage,proto3" json:"mon_min_voltage,omitempty"`
	MonMinVoltagePos     uint32                 `protobuf:"varint,14,opt,name=mon_min_voltage_pos,json=monMinVoltagePos,proto3" json:"mon_min_voltage_pos,omitempty"`
	MaxTemp              uint32                 `protobuf:"varint,15,opt,name=max_temp,json=maxTemp,proto3" json:"max_temp,omitempty"`
	MinTemp              uint32                 `protobuf:"varint,16,opt,name=min_temp,json=minTemp,proto3" json:"min_temp,omitempty"`
	Faults               []byte                 `protobuf:"bytes,17,opt,name=faults,proto3" json:"faults,omitempty"`
	MosStatus            []byte                 `protobuf:"bytes,18,opt,name=mos_status,json=mosStatus,proto3" json:"mos_status,omitempty"`
	MonVoltage           []byte                 `protobuf:"bytes,19,opt,name=mon_voltage,json=monVoltage,proto3" json:"mon_voltage,omitempty"`
	Temp                 []byte                 `protobuf:"bytes,20,opt,name=temp,proto3" json:"temp,omitempty"`
	MosTemp              uint32                 `protobuf:"varint,21,opt,name=mos_temp,json=mosTemp,proto3" json:"mos_temp,omitempty"`
	EnvTemp              uint32                 `protobuf:"varint,22,opt,name=env_temp,json=envTemp,proto3" json:"env_temp,omitempty"`
	Geom                 []byte                 `protobuf:"bytes,23,opt,name=geom,proto3" json:"geom,omitempty"`
	Gps                  int64                  `protobuf:"varint,24,opt,name=gps,proto3" json:"gps,omitempty"`
	Strength             uint32                 `protobuf:"varint,25,opt,name=strength,proto3" json:"strength,omitempty"`
	Cycles               uint32                 `protobuf:"varint,26,opt,name=cycles,proto3" json:"cycles,omitempty"`
	ChargingTime         uint32                 `protobuf:"varint,27,opt,name=charging_time,json=chargingTime,proto3" json:"charging_time,omitempty"`
	DisChargingTime      uint32                 `protobuf:"varint,28,opt,name=dis_charging_time,json=disChargingTime,proto3" json:"dis_charging_time,omitempty"`
	UsingTime            uint32                 `protobuf:"varint,29,opt,name=using_time,json=usingTime,proto3" json:"using_time,omitempty"`
	TotalChargingTime    uint32                 `protobuf:"varint,30,opt,name=total_charging_time,json=totalChargingTime,proto3" json:"total_charging_time,omitempty"`
	TotalDisChargingTime uint32                 `protobuf:"varint,31,opt,name=total_dis_charging_time,json=totalDisChargingTime,proto3" json:"total_dis_charging_time,omitempty"`
	TotalUsingTime       uint32                 `protobuf:"varint,32,opt,name=total_using_time,json=totalUsingTime,proto3" json:"total_using_time,omitempty"`
	Battery              *Battery               `protobuf:"bytes,33,opt,name=battery,proto3" json:"battery,omitempty"`
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{1}
}

func (x *Heartbeat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Heartbeat) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

func (x *Heartbeat) GetBatteryId() int64 {
	if x != nil {
		return x.BatteryId
	}
	return 0
}

func (x *Heartbeat) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Heartbeat) GetVoltage() float64 {
	if x != nil {
		return x.Voltage
	}
	return 0
}

func (x *Heartbeat) GetCurrent() float64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *Heartbeat) GetSoc() uint32 {
	if x != nil {
		return x.Soc
	}
	return 0
}

func (x *Heartbeat) GetSoh() uint32 {
	if x != nil {
		return x.Soh
	}
	return 0
}

func (x *Heartbeat) GetInCabinet() bool {
	if x != nil {
		return x.InCabinet
	}
	return false
}

func (x *Heartbeat) GetCapacity() float64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Heartbeat) GetMonMaxVoltage() uint32 {
	if x != nil {
		return x.MonMaxVoltage
	}
	return 0
}

func (x *Heartbeat) GetMonMaxVoltagePos() uint32 {
	if x != nil {
		return x.MonMaxVoltagePos
	}
	return 0
}

func (x *Heartbeat) GetMonMinVoltage() uint32 {
	if x != nil {
		return x.MonMinVoltage
	}
	return 0
}

func (x *Heartbeat) GetMonMinVoltagePos() uint32 {
	if x != nil {
		return x.MonMinVoltagePos
	}
	return 0
}

func (x *Heartbeat) GetMaxTemp() uint32 {
	if x != nil {
		return x.MaxTemp
	}
	return 0
}

func (x *Heartbeat) GetMinTemp() uint32 {
	if x != nil {
		return x.MinTemp
	}
	return 0
}

func (x *Heartbeat) GetFaults() []byte {
	if x != nil {
		return x.Faults
	}
	return nil
}

func (x *Heartbeat) GetMosStatus() []byte {
	if x != nil {
		return x.MosStatus
	}
	return nil
}

func (x *Heartbeat) GetMonVoltage() []byte {
	if x != nil {
		return x.MonVoltage
	}
	return nil
}

func (x *Heartbeat) GetTemp() []byte {
	if x != nil {
		return x.Temp
	}
	return nil
}

func (x *Heartbeat) GetMosTemp() uint32 {
	if x != nil {
		return x.MosTemp
	}
	return 0
}

func (x *Heartbeat) GetEnvTemp() uint32 {
	if x != nil {
		return x.EnvTemp
	}
	return 0
}

func (x *Heartbeat) GetGeom() []byte {
	if x != nil {
		return x.Geom
	}
	return nil
}

func (x *Heartbeat) GetGps() int64 {
	if x != nil {
		return x.Gps
	}
	return 0
}

func (x *Heartbeat) GetStrength() uint32 {
	if x != nil {
		return x.Strength
	}
	return 0
}

func (x *Heartbeat) GetCycles() uint32 {
	if x != nil {
		return x.Cycles
	}
	return 0
}

func (x *Heartbeat) GetChargingTime() uint32 {
	if x != nil {
		return x.ChargingTime
	}
	return 0
}

func (x *Heartbeat) GetDisChargingTime() uint32 {
	if x != nil {
		return x.DisChargingTime
	}
	return 0
}

func (x *Heartbeat) GetUsingTime() uint32 {
	if x != nil {
		return x.UsingTime
	}
	return 0
}

func (x *Heartbeat) GetTotalChargingTime() uint32 {
	if x != nil {
		return x.TotalChargingTime
	}
	return 0
}

func (x *Heartbeat) GetTotalDisChargingTime() uint32 {
	if x != nil {
		return x.TotalDisChargingTime
	}
	return 0
}

func (x *Heartbeat) GetTotalUsingTime() uint32 {
	if x != nil {
		return x.TotalUsingTime
	}
	return 0
}

func (x *Heartbeat) GetBattery() *Battery {
	if x != nil {
		return x.Battery
	}
	return nil
}

type Reign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Action      int64                   `protobuf:"varint,2,opt,name=action,proto3" json:"action,omitempty"`
	Sn          string                  `protobuf:"bytes,3,opt,name=sn,proto3" json:"sn,omitempty"`
	BatteryId   int64                   `protobuf:"varint,4,opt,name=battery_id,json=batteryId,proto3" json:"battery_id,omitempty"`
	CreatedAt   *timestamppb.Timestamp  `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Serial      string                  `protobuf:"bytes,6,opt,name=serial,proto3" json:"serial,omitempty"`
	Ordinal     int64                   `protobuf:"varint,7,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	CabinetName *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=cabinet_name,json=cabinetName,proto3" json:"cabinet_name,omitempty"`
	Remark      *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Geom        []byte                  `protobuf:"bytes,10,opt,name=geom,proto3" json:"geom,omitempty"`
	Battery     *Battery                `protobuf:"bytes,11,opt,name=battery,proto3" json:"battery,omitempty"`
}

func (x *Reign) Reset() {
	*x = Reign{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entpb_entpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reign) ProtoMessage() {}

func (x *Reign) ProtoReflect() protoreflect.Message {
	mi := &file_entpb_entpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reign.ProtoReflect.Descriptor instead.
func (*Reign) Descriptor() ([]byte, []int) {
	return file_entpb_entpb_proto_rawDescGZIP(), []int{2}
}

func (x *Reign) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Reign) GetAction() int64 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *Reign) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

func (x *Reign) GetBatteryId() int64 {
	if x != nil {
		return x.BatteryId
	}
	return 0
}

func (x *Reign) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Reign) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *Reign) GetOrdinal() int64 {
	if x != nil {
		return x.Ordinal
	}
	return 0
}

func (x *Reign) GetCabinetName() *wrapperspb.StringValue {
	if x != nil {
		return x.CabinetName
	}
	return nil
}

func (x *Reign) GetRemark() *wrapperspb.StringValue {
	if x != nil {
		return x.Remark
	}
	return nil
}

func (x *Reign) GetGeom() []byte {
	if x != nil {
		return x.Geom
	}
	return nil
}

func (x *Reign) GetBattery() *Battery {
	if x != nil {
		return x.Battery
	}
	return nil
}

var File_entpb_entpb_proto protoreflect.FileDescriptor

var file_entpb_entpb_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x04, 0x0a, 0x07,
	0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x6f, 0x66, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x73, 0x6f, 0x66,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0c, 0x68, 0x61, 0x72, 0x64,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x68, 0x61,
	0x72, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0f, 0x73, 0x6f, 0x66,
	0x74, 0x5f, 0x34, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0d, 0x73, 0x6f, 0x66, 0x74, 0x34, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x44, 0x0a, 0x0f, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x34, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x68, 0x61, 0x72, 0x64, 0x34, 0x67, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x6e, 0x5f, 0x34, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x04, 0x73, 0x6e, 0x34, 0x67, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x63, 0x63, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x69, 0x63, 0x63, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x03,
	0x73, 0x6f, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x73, 0x6f, 0x63, 0x12, 0x30, 0x0a, 0x0a,
	0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x12, 0x24,
	0x0a, 0x06, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x69, 0x67, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x69, 0x67, 0x6e, 0x73, 0x22, 0xa3, 0x08, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x73, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x76,
	0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x6f, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73,
	0x6f, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6f, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x73, 0x6f, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x62, 0x69, 0x6e,
	0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x6e, 0x43, 0x61, 0x62, 0x69,
	0x6e, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x26, 0x0a, 0x0f, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x6f, 0x6c, 0x74, 0x61,
	0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d, 0x6f, 0x6e, 0x4d, 0x61, 0x78,
	0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d,
	0x61, 0x78, 0x5f, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6d, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x56, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x50, 0x6f, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x69,
	0x6e, 0x5f, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x6d, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x2d,
	0x0a, 0x13, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6d, 0x6f, 0x6e,
	0x4d, 0x69, 0x6e, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x6d, 0x61, 0x78, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x54,
	0x65, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x6f, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x6d, 0x6f, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f,
	0x6e, 0x5f, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x6d, 0x6f, 0x6e, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x6d, 0x70, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x65, 0x6d, 0x70, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x6d, 0x6f, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e,
	0x76, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e,
	0x76, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x65, 0x6f, 0x6d, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x67, 0x65, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x70, 0x73,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x70, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73,
	0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x73, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x64, 0x69, 0x73, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x2e, 0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e,
	0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x17, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x73, 0x5f, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69,
	0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x75, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x18, 0x21, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x79, 0x52, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x22, 0x80, 0x03, 0x0a, 0x05, 0x52,
	0x65, 0x69, 0x67, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x73, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x3f, 0x0a, 0x0c, 0x63, 0x61, 0x62, 0x69,
	0x6e, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x63, 0x61,
	0x62, 0x69, 0x6e, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x67, 0x65, 0x6f, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x67,
	0x65, 0x6f, 0x6d, 0x12, 0x28, 0x0a, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x79, 0x52, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x42, 0x36, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x72, 0x6f,
	0x72, 0x61, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x64,
	0x65, 0x66, 0x73, 0x2f, 0x78, 0x63, 0x64, 0x65, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entpb_entpb_proto_rawDescOnce sync.Once
	file_entpb_entpb_proto_rawDescData = file_entpb_entpb_proto_rawDesc
)

func file_entpb_entpb_proto_rawDescGZIP() []byte {
	file_entpb_entpb_proto_rawDescOnce.Do(func() {
		file_entpb_entpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_entpb_entpb_proto_rawDescData)
	})
	return file_entpb_entpb_proto_rawDescData
}

var file_entpb_entpb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_entpb_entpb_proto_goTypes = []interface{}{
	(*Battery)(nil),                // 0: entpb.Battery
	(*Heartbeat)(nil),              // 1: entpb.Heartbeat
	(*Reign)(nil),                  // 2: entpb.Reign
	(*wrapperspb.UInt32Value)(nil), // 3: google.protobuf.UInt32Value
	(*wrapperspb.UInt64Value)(nil), // 4: google.protobuf.UInt64Value
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_entpb_entpb_proto_depIdxs = []int32{
	3,  // 0: entpb.Battery.soft_version:type_name -> google.protobuf.UInt32Value
	3,  // 1: entpb.Battery.hard_version:type_name -> google.protobuf.UInt32Value
	3,  // 2: entpb.Battery.soft_4g_version:type_name -> google.protobuf.UInt32Value
	3,  // 3: entpb.Battery.hard_4g_version:type_name -> google.protobuf.UInt32Value
	4,  // 4: entpb.Battery.sn_4g:type_name -> google.protobuf.UInt64Value
	5,  // 5: entpb.Battery.iccid:type_name -> google.protobuf.StringValue
	3,  // 6: entpb.Battery.soc:type_name -> google.protobuf.UInt32Value
	1,  // 7: entpb.Battery.heartbeats:type_name -> entpb.Heartbeat
	2,  // 8: entpb.Battery.reigns:type_name -> entpb.Reign
	6,  // 9: entpb.Heartbeat.created_at:type_name -> google.protobuf.Timestamp
	0,  // 10: entpb.Heartbeat.battery:type_name -> entpb.Battery
	6,  // 11: entpb.Reign.created_at:type_name -> google.protobuf.Timestamp
	5,  // 12: entpb.Reign.cabinet_name:type_name -> google.protobuf.StringValue
	5,  // 13: entpb.Reign.remark:type_name -> google.protobuf.StringValue
	0,  // 14: entpb.Reign.battery:type_name -> entpb.Battery
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_entpb_entpb_proto_init() }
func file_entpb_entpb_proto_init() {
	if File_entpb_entpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entpb_entpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Battery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entpb_entpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_entpb_entpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reign); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entpb_entpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entpb_entpb_proto_goTypes,
		DependencyIndexes: file_entpb_entpb_proto_depIdxs,
		MessageInfos:      file_entpb_entpb_proto_msgTypes,
	}.Build()
	File_entpb_entpb_proto = out.File
	file_entpb_entpb_proto_rawDesc = nil
	file_entpb_entpb_proto_goTypes = nil
	file_entpb_entpb_proto_depIdxs = nil
}
