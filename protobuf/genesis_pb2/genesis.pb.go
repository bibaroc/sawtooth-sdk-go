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
// source: protobuf/genesis_pb2/genesis.proto

package genesis_pb2

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	batch_pb2 "github.com/hyperledger/sawtooth-sdk-go/protobuf/batch_pb2"
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

type GenesisData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of batches that will be applied during the genesis process
	Batches []*batch_pb2.Batch `protobuf:"bytes,1,rep,name=batches,proto3" json:"batches,omitempty"`
}

func (x *GenesisData) Reset() {
	*x = GenesisData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_genesis_pb2_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisData) ProtoMessage() {}

func (x *GenesisData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_genesis_pb2_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisData.ProtoReflect.Descriptor instead.
func (*GenesisData) Descriptor() ([]byte, []int) {
	return file_protobuf_genesis_pb2_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisData) GetBatches() []*batch_pb2.Batch {
	if x != nil {
		return x.Batches
	}
	return nil
}

var File_protobuf_genesis_pb2_genesis_proto protoreflect.FileDescriptor

var file_protobuf_genesis_pb2_genesis_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73,
	0x69, 0x73, 0x5f, 0x70, 0x62, 0x32, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x70, 0x62, 0x32, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x42, 0x26, 0x0a, 0x15, 0x73, 0x61, 0x77, 0x74, 0x6f, 0x6f, 0x74,
	0x68, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x01,
	0x5a, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x70, 0x62, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_genesis_pb2_genesis_proto_rawDescOnce sync.Once
	file_protobuf_genesis_pb2_genesis_proto_rawDescData = file_protobuf_genesis_pb2_genesis_proto_rawDesc
)

func file_protobuf_genesis_pb2_genesis_proto_rawDescGZIP() []byte {
	file_protobuf_genesis_pb2_genesis_proto_rawDescOnce.Do(func() {
		file_protobuf_genesis_pb2_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_genesis_pb2_genesis_proto_rawDescData)
	})
	return file_protobuf_genesis_pb2_genesis_proto_rawDescData
}

var file_protobuf_genesis_pb2_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_genesis_pb2_genesis_proto_goTypes = []interface{}{
	(*GenesisData)(nil),     // 0: GenesisData
	(*batch_pb2.Batch)(nil), // 1: Batch
}
var file_protobuf_genesis_pb2_genesis_proto_depIdxs = []int32{
	1, // 0: GenesisData.batches:type_name -> Batch
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_genesis_pb2_genesis_proto_init() }
func file_protobuf_genesis_pb2_genesis_proto_init() {
	if File_protobuf_genesis_pb2_genesis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_genesis_pb2_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisData); i {
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
			RawDescriptor: file_protobuf_genesis_pb2_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_genesis_pb2_genesis_proto_goTypes,
		DependencyIndexes: file_protobuf_genesis_pb2_genesis_proto_depIdxs,
		MessageInfos:      file_protobuf_genesis_pb2_genesis_proto_msgTypes,
	}.Build()
	File_protobuf_genesis_pb2_genesis_proto = out.File
	file_protobuf_genesis_pb2_genesis_proto_rawDesc = nil
	file_protobuf_genesis_pb2_genesis_proto_goTypes = nil
	file_protobuf_genesis_pb2_genesis_proto_depIdxs = nil
}
