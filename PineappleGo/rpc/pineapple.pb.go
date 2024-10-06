// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: pineapple.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pineapple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pineapple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_pineapple_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag   uint64 `protobuf:"varint,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pineapple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pineapple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_pineapple_proto_rawDescGZIP(), []int{1}
}

func (x *ReadResponse) GetTag() uint64 {
	if x != nil {
		return x.Tag
	}
	return 0
}

func (x *ReadResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type PeekRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PeekRequest) Reset() {
	*x = PeekRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pineapple_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeekRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeekRequest) ProtoMessage() {}

func (x *PeekRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pineapple_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeekRequest.ProtoReflect.Descriptor instead.
func (*PeekRequest) Descriptor() ([]byte, []int) {
	return file_pineapple_proto_rawDescGZIP(), []int{2}
}

func (x *PeekRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type PeekResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag uint64 `protobuf:"varint,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *PeekResponse) Reset() {
	*x = PeekResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pineapple_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeekResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeekResponse) ProtoMessage() {}

func (x *PeekResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pineapple_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeekResponse.ProtoReflect.Descriptor instead.
func (*PeekResponse) Descriptor() ([]byte, []int) {
	return file_pineapple_proto_rawDescGZIP(), []int{3}
}

func (x *PeekResponse) GetTag() uint64 {
	if x != nil {
		return x.Tag
	}
	return 0
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Tag   uint64 `protobuf:"varint,2,opt,name=tag,proto3" json:"tag,omitempty"`
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pineapple_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pineapple_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_pineapple_proto_rawDescGZIP(), []int{4}
}

func (x *WriteRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *WriteRequest) GetTag() uint64 {
	if x != nil {
		return x.Tag
	}
	return 0
}

func (x *WriteRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pineapple_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pineapple_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_pineapple_proto_rawDescGZIP(), []int{5}
}

type ModifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Request []byte `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *ModifyRequest) Reset() {
	*x = ModifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pineapple_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyRequest) ProtoMessage() {}

func (x *ModifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pineapple_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyRequest.ProtoReflect.Descriptor instead.
func (*ModifyRequest) Descriptor() ([]byte, []int) {
	return file_pineapple_proto_rawDescGZIP(), []int{6}
}

func (x *ModifyRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ModifyRequest) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

type ModifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ModifyResponse) Reset() {
	*x = ModifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pineapple_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyResponse) ProtoMessage() {}

func (x *ModifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pineapple_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyResponse.ProtoReflect.Descriptor instead.
func (*ModifyResponse) Descriptor() ([]byte, []int) {
	return file_pineapple_proto_rawDescGZIP(), []int{7}
}

var File_pineapple_proto protoreflect.FileDescriptor

var file_pineapple_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x69, 0x6e, 0x65, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1f, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x36, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x74, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1f, 0x0a, 0x0b, 0x50, 0x65,
	0x65, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x20, 0x0a, 0x0c, 0x50,
	0x65, 0x65, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x48, 0x0a,
	0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xab, 0x01, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x25, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0c, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x6b, 0x12,
	0x0c, 0x2e, 0x50, 0x65, 0x65, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x50, 0x65, 0x65, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28,
	0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x12, 0x0e, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pineapple_proto_rawDescOnce sync.Once
	file_pineapple_proto_rawDescData = file_pineapple_proto_rawDesc
)

func file_pineapple_proto_rawDescGZIP() []byte {
	file_pineapple_proto_rawDescOnce.Do(func() {
		file_pineapple_proto_rawDescData = protoimpl.X.CompressGZIP(file_pineapple_proto_rawDescData)
	})
	return file_pineapple_proto_rawDescData
}

var file_pineapple_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pineapple_proto_goTypes = []interface{}{
	(*ReadRequest)(nil),    // 0: ReadRequest
	(*ReadResponse)(nil),   // 1: ReadResponse
	(*PeekRequest)(nil),    // 2: PeekRequest
	(*PeekResponse)(nil),   // 3: PeekResponse
	(*WriteRequest)(nil),   // 4: WriteRequest
	(*WriteResponse)(nil),  // 5: WriteResponse
	(*ModifyRequest)(nil),  // 6: ModifyRequest
	(*ModifyResponse)(nil), // 7: ModifyResponse
}
var file_pineapple_proto_depIdxs = []int32{
	0, // 0: Node.Read:input_type -> ReadRequest
	2, // 1: Node.Peek:input_type -> PeekRequest
	4, // 2: Node.Write:input_type -> WriteRequest
	6, // 3: Node.Modify:input_type -> ModifyRequest
	1, // 4: Node.Read:output_type -> ReadResponse
	3, // 5: Node.Peek:output_type -> PeekResponse
	5, // 6: Node.Write:output_type -> WriteResponse
	7, // 7: Node.Modify:output_type -> ModifyResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pineapple_proto_init() }
func file_pineapple_proto_init() {
	if File_pineapple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pineapple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_pineapple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_pineapple_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeekRequest); i {
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
		file_pineapple_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeekResponse); i {
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
		file_pineapple_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_pineapple_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResponse); i {
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
		file_pineapple_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyRequest); i {
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
		file_pineapple_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyResponse); i {
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
			RawDescriptor: file_pineapple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pineapple_proto_goTypes,
		DependencyIndexes: file_pineapple_proto_depIdxs,
		MessageInfos:      file_pineapple_proto_msgTypes,
	}.Build()
	File_pineapple_proto = out.File
	file_pineapple_proto_rawDesc = nil
	file_pineapple_proto_goTypes = nil
	file_pineapple_proto_depIdxs = nil
}
