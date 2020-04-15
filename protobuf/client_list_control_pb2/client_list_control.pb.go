// Copyright 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.6.1
// source: protobuf/client_list_control_pb2/client_list_control.proto

package client_list_control_pb2

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Paging controls to be sent with List requests.
// Attributes:
//     start: The id of a resource to start the page with
//     limit: The number of results per page, defaults to 100 and maxes out at 1000
type ClientPagingControls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	Limit int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ClientPagingControls) Reset() {
	*x = ClientPagingControls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_client_list_control_pb2_client_list_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPagingControls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPagingControls) ProtoMessage() {}

func (x *ClientPagingControls) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_client_list_control_pb2_client_list_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPagingControls.ProtoReflect.Descriptor instead.
func (*ClientPagingControls) Descriptor() ([]byte, []int) {
	return file_protobuf_client_list_control_pb2_client_list_control_proto_rawDescGZIP(), []int{0}
}

func (x *ClientPagingControls) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *ClientPagingControls) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// Information about the pagination used, sent back with List responses.
// Attributes:
//     next: The id of the first resource in the next page
//     start: The id of the first resource in the returned page
//     limit: The number of results per page, defaults to 100 and maxes out at 1000
type ClientPagingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  string `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Start string `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	Limit int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ClientPagingResponse) Reset() {
	*x = ClientPagingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_client_list_control_pb2_client_list_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPagingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPagingResponse) ProtoMessage() {}

func (x *ClientPagingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_client_list_control_pb2_client_list_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPagingResponse.ProtoReflect.Descriptor instead.
func (*ClientPagingResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_client_list_control_pb2_client_list_control_proto_rawDescGZIP(), []int{1}
}

func (x *ClientPagingResponse) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

func (x *ClientPagingResponse) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *ClientPagingResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// Sorting controls to be sent with List requests. More than one can be sent.
// If so, the first is used, and additional controls are tie-breakers.
// Attributes:
//     keys: Nested set of keys to sort by (i.e. ['default, block_num'])
//     reverse: Whether or not to reverse the sort (i.e. descending order)
type ClientSortControls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys    []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Reverse bool     `protobuf:"varint,2,opt,name=reverse,proto3" json:"reverse,omitempty"`
}

func (x *ClientSortControls) Reset() {
	*x = ClientSortControls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_client_list_control_pb2_client_list_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientSortControls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientSortControls) ProtoMessage() {}

func (x *ClientSortControls) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_client_list_control_pb2_client_list_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientSortControls.ProtoReflect.Descriptor instead.
func (*ClientSortControls) Descriptor() ([]byte, []int) {
	return file_protobuf_client_list_control_pb2_client_list_control_proto_rawDescGZIP(), []int{2}
}

func (x *ClientSortControls) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *ClientSortControls) GetReverse() bool {
	if x != nil {
		return x.Reverse
	}
	return false
}

var File_protobuf_client_list_control_pb2_client_list_control_proto protoreflect.FileDescriptor

var file_protobuf_client_list_control_pb2_client_list_control_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70,
	0x62, 0x32, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x14,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x56, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x42, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x42, 0x32, 0x0a, 0x15,
	0x73, 0x61, 0x77, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x01, 0x5a, 0x17, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x62, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_client_list_control_pb2_client_list_control_proto_rawDescOnce sync.Once
	file_protobuf_client_list_control_pb2_client_list_control_proto_rawDescData = file_protobuf_client_list_control_pb2_client_list_control_proto_rawDesc
)

func file_protobuf_client_list_control_pb2_client_list_control_proto_rawDescGZIP() []byte {
	file_protobuf_client_list_control_pb2_client_list_control_proto_rawDescOnce.Do(func() {
		file_protobuf_client_list_control_pb2_client_list_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_client_list_control_pb2_client_list_control_proto_rawDescData)
	})
	return file_protobuf_client_list_control_pb2_client_list_control_proto_rawDescData
}

var file_protobuf_client_list_control_pb2_client_list_control_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuf_client_list_control_pb2_client_list_control_proto_goTypes = []interface{}{
	(*ClientPagingControls)(nil), // 0: ClientPagingControls
	(*ClientPagingResponse)(nil), // 1: ClientPagingResponse
	(*ClientSortControls)(nil),   // 2: ClientSortControls
}
var file_protobuf_client_list_control_pb2_client_list_control_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_client_list_control_pb2_client_list_control_proto_init() }
func file_protobuf_client_list_control_pb2_client_list_control_proto_init() {
	if File_protobuf_client_list_control_pb2_client_list_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_client_list_control_pb2_client_list_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPagingControls); i {
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
		file_protobuf_client_list_control_pb2_client_list_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPagingResponse); i {
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
		file_protobuf_client_list_control_pb2_client_list_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientSortControls); i {
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
			RawDescriptor: file_protobuf_client_list_control_pb2_client_list_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_client_list_control_pb2_client_list_control_proto_goTypes,
		DependencyIndexes: file_protobuf_client_list_control_pb2_client_list_control_proto_depIdxs,
		MessageInfos:      file_protobuf_client_list_control_pb2_client_list_control_proto_msgTypes,
	}.Build()
	File_protobuf_client_list_control_pb2_client_list_control_proto = out.File
	file_protobuf_client_list_control_pb2_client_list_control_proto_rawDesc = nil
	file_protobuf_client_list_control_pb2_client_list_control_proto_goTypes = nil
	file_protobuf_client_list_control_pb2_client_list_control_proto_depIdxs = nil
}
