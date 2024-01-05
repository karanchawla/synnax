// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: v1/device.proto

package apiv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Rack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  uint32 `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Rack) Reset() {
	*x = Rack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rack) ProtoMessage() {}

func (x *Rack) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rack.ProtoReflect.Descriptor instead.
func (*Rack) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{0}
}

func (x *Rack) GetKey() uint32 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *Rack) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeviceCreateRackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Racks []*Rack `protobuf:"bytes,1,rep,name=racks,proto3" json:"racks,omitempty"`
}

func (x *DeviceCreateRackRequest) Reset() {
	*x = DeviceCreateRackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceCreateRackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCreateRackRequest) ProtoMessage() {}

func (x *DeviceCreateRackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCreateRackRequest.ProtoReflect.Descriptor instead.
func (*DeviceCreateRackRequest) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceCreateRackRequest) GetRacks() []*Rack {
	if x != nil {
		return x.Racks
	}
	return nil
}

type DeviceCreateRackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Racks []*Rack `protobuf:"bytes,1,rep,name=racks,proto3" json:"racks,omitempty"`
}

func (x *DeviceCreateRackResponse) Reset() {
	*x = DeviceCreateRackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceCreateRackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCreateRackResponse) ProtoMessage() {}

func (x *DeviceCreateRackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCreateRackResponse.ProtoReflect.Descriptor instead.
func (*DeviceCreateRackResponse) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceCreateRackResponse) GetRacks() []*Rack {
	if x != nil {
		return x.Racks
	}
	return nil
}

type DeviceRetrieveRackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []uint32 `protobuf:"varint,1,rep,packed,name=keys,proto3" json:"keys,omitempty"`
}

func (x *DeviceRetrieveRackRequest) Reset() {
	*x = DeviceRetrieveRackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRetrieveRackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRetrieveRackRequest) ProtoMessage() {}

func (x *DeviceRetrieveRackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRetrieveRackRequest.ProtoReflect.Descriptor instead.
func (*DeviceRetrieveRackRequest) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceRetrieveRackRequest) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

type DeviceRetrieveRackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Racks []*Rack `protobuf:"bytes,1,rep,name=racks,proto3" json:"racks,omitempty"`
}

func (x *DeviceRetrieveRackResponse) Reset() {
	*x = DeviceRetrieveRackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRetrieveRackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRetrieveRackResponse) ProtoMessage() {}

func (x *DeviceRetrieveRackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRetrieveRackResponse.ProtoReflect.Descriptor instead.
func (*DeviceRetrieveRackResponse) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{4}
}

func (x *DeviceRetrieveRackResponse) GetRacks() []*Rack {
	if x != nil {
		return x.Racks
	}
	return nil
}

type DeviceDeleteRackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []uint32 `protobuf:"varint,1,rep,packed,name=keys,proto3" json:"keys,omitempty"`
}

func (x *DeviceDeleteRackRequest) Reset() {
	*x = DeviceDeleteRackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceDeleteRackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDeleteRackRequest) ProtoMessage() {}

func (x *DeviceDeleteRackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDeleteRackRequest.ProtoReflect.Descriptor instead.
func (*DeviceDeleteRackRequest) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{5}
}

func (x *DeviceDeleteRackRequest) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    uint64 `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type   string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Config string `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{6}
}

func (x *Module) GetKey() uint64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *Module) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Module) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Module) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type DeviceCreateModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modules []*Module `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *DeviceCreateModuleRequest) Reset() {
	*x = DeviceCreateModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceCreateModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCreateModuleRequest) ProtoMessage() {}

func (x *DeviceCreateModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCreateModuleRequest.ProtoReflect.Descriptor instead.
func (*DeviceCreateModuleRequest) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{7}
}

func (x *DeviceCreateModuleRequest) GetModules() []*Module {
	if x != nil {
		return x.Modules
	}
	return nil
}

type DeviceCreateModuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modules []*Module `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *DeviceCreateModuleResponse) Reset() {
	*x = DeviceCreateModuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceCreateModuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCreateModuleResponse) ProtoMessage() {}

func (x *DeviceCreateModuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCreateModuleResponse.ProtoReflect.Descriptor instead.
func (*DeviceCreateModuleResponse) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{8}
}

func (x *DeviceCreateModuleResponse) GetModules() []*Module {
	if x != nil {
		return x.Modules
	}
	return nil
}

type DeviceRetrieveModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rack uint32   `protobuf:"varint,1,opt,name=rack,proto3" json:"rack,omitempty"`
	Keys []uint64 `protobuf:"varint,2,rep,packed,name=keys,proto3" json:"keys,omitempty"`
}

func (x *DeviceRetrieveModuleRequest) Reset() {
	*x = DeviceRetrieveModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRetrieveModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRetrieveModuleRequest) ProtoMessage() {}

func (x *DeviceRetrieveModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRetrieveModuleRequest.ProtoReflect.Descriptor instead.
func (*DeviceRetrieveModuleRequest) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{9}
}

func (x *DeviceRetrieveModuleRequest) GetRack() uint32 {
	if x != nil {
		return x.Rack
	}
	return 0
}

func (x *DeviceRetrieveModuleRequest) GetKeys() []uint64 {
	if x != nil {
		return x.Keys
	}
	return nil
}

type DeviceRetrieveModuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modules []*Module `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *DeviceRetrieveModuleResponse) Reset() {
	*x = DeviceRetrieveModuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRetrieveModuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRetrieveModuleResponse) ProtoMessage() {}

func (x *DeviceRetrieveModuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRetrieveModuleResponse.ProtoReflect.Descriptor instead.
func (*DeviceRetrieveModuleResponse) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{10}
}

func (x *DeviceRetrieveModuleResponse) GetModules() []*Module {
	if x != nil {
		return x.Modules
	}
	return nil
}

type DeviceDeleteModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []uint64 `protobuf:"varint,1,rep,packed,name=keys,proto3" json:"keys,omitempty"`
}

func (x *DeviceDeleteModuleRequest) Reset() {
	*x = DeviceDeleteModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_device_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceDeleteModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDeleteModuleRequest) ProtoMessage() {}

func (x *DeviceDeleteModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_device_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDeleteModuleRequest.ProtoReflect.Descriptor instead.
func (*DeviceDeleteModuleRequest) Descriptor() ([]byte, []int) {
	return file_v1_device_proto_rawDescGZIP(), []int{11}
}

func (x *DeviceDeleteModuleRequest) GetKeys() []uint64 {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_v1_device_proto protoreflect.FileDescriptor

var file_v1_device_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x04, 0x52, 0x61, 0x63, 0x6b, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x17, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x05, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x72, 0x61,
	0x63, 0x6b, 0x73, 0x22, 0x3e, 0x0a, 0x18, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x72, 0x61,
	0x63, 0x6b, 0x73, 0x22, 0x2f, 0x0a, 0x19, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x22, 0x40, 0x0a, 0x1a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x63, 0x6b, 0x52,
	0x05, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x2d, 0x0a, 0x17, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x5a, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x45, 0x0a, 0x19, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x46, 0x0a, 0x1a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x22, 0x45, 0x0a, 0x1b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72,
	0x61, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x48, 0x0a, 0x1c, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x22, 0x2f, 0x0a, 0x19, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x32, 0x6c, 0x0a, 0x19, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x74, 0x0a, 0x1b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x55, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x60, 0x0a, 0x19, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x21, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x66, 0x0a, 0x17, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x6e, 0x0a, 0x19, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a,
	0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x5c, 0x0a, 0x17, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x86,
	0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x6e, 0x6e, 0x61, 0x78, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x6e, 0x61, 0x78, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x06, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_device_proto_rawDescOnce sync.Once
	file_v1_device_proto_rawDescData = file_v1_device_proto_rawDesc
)

func file_v1_device_proto_rawDescGZIP() []byte {
	file_v1_device_proto_rawDescOnce.Do(func() {
		file_v1_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_device_proto_rawDescData)
	})
	return file_v1_device_proto_rawDescData
}

var file_v1_device_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_v1_device_proto_goTypes = []interface{}{
	(*Rack)(nil),                         // 0: api.v1.Rack
	(*DeviceCreateRackRequest)(nil),      // 1: api.v1.DeviceCreateRackRequest
	(*DeviceCreateRackResponse)(nil),     // 2: api.v1.DeviceCreateRackResponse
	(*DeviceRetrieveRackRequest)(nil),    // 3: api.v1.DeviceRetrieveRackRequest
	(*DeviceRetrieveRackResponse)(nil),   // 4: api.v1.DeviceRetrieveRackResponse
	(*DeviceDeleteRackRequest)(nil),      // 5: api.v1.DeviceDeleteRackRequest
	(*Module)(nil),                       // 6: api.v1.Module
	(*DeviceCreateModuleRequest)(nil),    // 7: api.v1.DeviceCreateModuleRequest
	(*DeviceCreateModuleResponse)(nil),   // 8: api.v1.DeviceCreateModuleResponse
	(*DeviceRetrieveModuleRequest)(nil),  // 9: api.v1.DeviceRetrieveModuleRequest
	(*DeviceRetrieveModuleResponse)(nil), // 10: api.v1.DeviceRetrieveModuleResponse
	(*DeviceDeleteModuleRequest)(nil),    // 11: api.v1.DeviceDeleteModuleRequest
	(*emptypb.Empty)(nil),                // 12: google.protobuf.Empty
}
var file_v1_device_proto_depIdxs = []int32{
	0,  // 0: api.v1.DeviceCreateRackRequest.racks:type_name -> api.v1.Rack
	0,  // 1: api.v1.DeviceCreateRackResponse.racks:type_name -> api.v1.Rack
	0,  // 2: api.v1.DeviceRetrieveRackResponse.racks:type_name -> api.v1.Rack
	6,  // 3: api.v1.DeviceCreateModuleRequest.modules:type_name -> api.v1.Module
	6,  // 4: api.v1.DeviceCreateModuleResponse.modules:type_name -> api.v1.Module
	6,  // 5: api.v1.DeviceRetrieveModuleResponse.modules:type_name -> api.v1.Module
	7,  // 6: api.v1.DeviceCreateModuleService.Create:input_type -> api.v1.DeviceCreateModuleRequest
	9,  // 7: api.v1.DeviceRetrieveModuleService.Retrieve:input_type -> api.v1.DeviceRetrieveModuleRequest
	11, // 8: api.v1.DeviceDeleteModuleService.Delete:input_type -> api.v1.DeviceDeleteModuleRequest
	1,  // 9: api.v1.DeviceCreateRackService.Create:input_type -> api.v1.DeviceCreateRackRequest
	3,  // 10: api.v1.DeviceRetrieveRackService.Retrieve:input_type -> api.v1.DeviceRetrieveRackRequest
	5,  // 11: api.v1.DeviceDeleteRackService.Delete:input_type -> api.v1.DeviceDeleteRackRequest
	8,  // 12: api.v1.DeviceCreateModuleService.Create:output_type -> api.v1.DeviceCreateModuleResponse
	10, // 13: api.v1.DeviceRetrieveModuleService.Retrieve:output_type -> api.v1.DeviceRetrieveModuleResponse
	12, // 14: api.v1.DeviceDeleteModuleService.Delete:output_type -> google.protobuf.Empty
	2,  // 15: api.v1.DeviceCreateRackService.Create:output_type -> api.v1.DeviceCreateRackResponse
	4,  // 16: api.v1.DeviceRetrieveRackService.Retrieve:output_type -> api.v1.DeviceRetrieveRackResponse
	12, // 17: api.v1.DeviceDeleteRackService.Delete:output_type -> google.protobuf.Empty
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_v1_device_proto_init() }
func file_v1_device_proto_init() {
	if File_v1_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rack); i {
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
		file_v1_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceCreateRackRequest); i {
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
		file_v1_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceCreateRackResponse); i {
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
		file_v1_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceRetrieveRackRequest); i {
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
		file_v1_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceRetrieveRackResponse); i {
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
		file_v1_device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceDeleteRackRequest); i {
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
		file_v1_device_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
		file_v1_device_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceCreateModuleRequest); i {
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
		file_v1_device_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceCreateModuleResponse); i {
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
		file_v1_device_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceRetrieveModuleRequest); i {
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
		file_v1_device_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceRetrieveModuleResponse); i {
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
		file_v1_device_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceDeleteModuleRequest); i {
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
			RawDescriptor: file_v1_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   6,
		},
		GoTypes:           file_v1_device_proto_goTypes,
		DependencyIndexes: file_v1_device_proto_depIdxs,
		MessageInfos:      file_v1_device_proto_msgTypes,
	}.Build()
	File_v1_device_proto = out.File
	file_v1_device_proto_rawDesc = nil
	file_v1_device_proto_goTypes = nil
	file_v1_device_proto_depIdxs = nil
}
