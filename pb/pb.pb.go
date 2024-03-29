// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: pb/pb.proto

package pb

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

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{0}
}

func (x *GetReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []byte `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *GetResp) Reset() {
	*x = GetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResp) ProtoMessage() {}

func (x *GetResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResp.ProtoReflect.Descriptor instead.
func (*GetResp) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{1}
}

func (x *GetResp) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *GetResp) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type SetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val []byte `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *SetReq) Reset() {
	*x = SetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetReq) ProtoMessage() {}

func (x *SetReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetReq.ProtoReflect.Descriptor instead.
func (*SetReq) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{2}
}

func (x *SetReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetReq) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

type SetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *SetResp) Reset() {
	*x = SetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResp) ProtoMessage() {}

func (x *SetResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResp.ProtoReflect.Descriptor instead.
func (*SetResp) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{3}
}

func (x *SetResp) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type DelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DelReq) Reset() {
	*x = DelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelReq) ProtoMessage() {}

func (x *DelReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelReq.ProtoReflect.Descriptor instead.
func (*DelReq) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{4}
}

func (x *DelReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *DelResp) Reset() {
	*x = DelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelResp) ProtoMessage() {}

func (x *DelResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelResp.ProtoReflect.Descriptor instead.
func (*DelResp) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{5}
}

func (x *DelResp) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_pb_pb_proto protoreflect.FileDescriptor

var file_pb_pb_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x22, 0x1a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22,
	0x2c, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x1b, 0x0a,
	0x07, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x1a, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x1b, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x65, 0x72, 0x72, 0x32, 0x79, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x24, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x24, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x03, 0x44, 0x65, 0x6c, 0x12,
	0x0d, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x0e,
	0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_pb_proto_rawDescOnce sync.Once
	file_pb_pb_proto_rawDescData = file_pb_pb_proto_rawDesc
)

func file_pb_pb_proto_rawDescGZIP() []byte {
	file_pb_pb_proto_rawDescOnce.Do(func() {
		file_pb_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_pb_proto_rawDescData)
	})
	return file_pb_pb_proto_rawDescData
}

var file_pb_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_pb_proto_goTypes = []interface{}{
	(*GetReq)(nil),  // 0: cache.GetReq
	(*GetResp)(nil), // 1: cache.GetResp
	(*SetReq)(nil),  // 2: cache.SetReq
	(*SetResp)(nil), // 3: cache.SetResp
	(*DelReq)(nil),  // 4: cache.DelReq
	(*DelResp)(nil), // 5: cache.DelResp
}
var file_pb_pb_proto_depIdxs = []int32{
	0, // 0: cache.Cache.Get:input_type -> cache.GetReq
	2, // 1: cache.Cache.Set:input_type -> cache.SetReq
	4, // 2: cache.Cache.Del:input_type -> cache.DelReq
	1, // 3: cache.Cache.Get:output_type -> cache.GetResp
	3, // 4: cache.Cache.Set:output_type -> cache.SetResp
	5, // 5: cache.Cache.Del:output_type -> cache.DelResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_pb_proto_init() }
func file_pb_pb_proto_init() {
	if File_pb_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_pb_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResp); i {
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
		file_pb_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetReq); i {
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
		file_pb_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetResp); i {
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
		file_pb_pb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelReq); i {
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
		file_pb_pb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelResp); i {
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
			RawDescriptor: file_pb_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_pb_proto_goTypes,
		DependencyIndexes: file_pb_pb_proto_depIdxs,
		MessageInfos:      file_pb_pb_proto_msgTypes,
	}.Build()
	File_pb_pb_proto = out.File
	file_pb_pb_proto_rawDesc = nil
	file_pb_pb_proto_goTypes = nil
	file_pb_pb_proto_depIdxs = nil
}
