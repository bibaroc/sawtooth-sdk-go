// Copyright 2018 Intel Corporation
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
// source: protobuf/merkle_pb2/merkle.proto

package merkle_pb2

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

// An Entry in the change log for a given state root.
type ChangeLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A root hash of a merkle trie this tree was based off.
	Parent []byte `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The hashes that were added for this root. These may be deleted during
	// pruning, if this root is being abandoned.
	Additions [][]byte `protobuf:"bytes,2,rep,name=additions,proto3" json:"additions,omitempty"`
	// The list of successors.
	Successors []*ChangeLogEntry_Successor `protobuf:"bytes,3,rep,name=successors,proto3" json:"successors,omitempty"`
}

func (x *ChangeLogEntry) Reset() {
	*x = ChangeLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_merkle_pb2_merkle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeLogEntry) ProtoMessage() {}

func (x *ChangeLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_merkle_pb2_merkle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeLogEntry.ProtoReflect.Descriptor instead.
func (*ChangeLogEntry) Descriptor() ([]byte, []int) {
	return file_protobuf_merkle_pb2_merkle_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeLogEntry) GetParent() []byte {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *ChangeLogEntry) GetAdditions() [][]byte {
	if x != nil {
		return x.Additions
	}
	return nil
}

func (x *ChangeLogEntry) GetSuccessors() []*ChangeLogEntry_Successor {
	if x != nil {
		return x.Successors
	}
	return nil
}

// A state root that succeed this root.
type ChangeLogEntry_Successor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A root hash of a merkle trie based of off this root.
	Successor []byte `protobuf:"bytes,1,opt,name=successor,proto3" json:"successor,omitempty"`
	// The keys (i.e. hashes) that were replaced (i.e. deleted) by this
	// successor.  These may be deleted during pruning.
	Deletions [][]byte `protobuf:"bytes,2,rep,name=deletions,proto3" json:"deletions,omitempty"`
}

func (x *ChangeLogEntry_Successor) Reset() {
	*x = ChangeLogEntry_Successor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_merkle_pb2_merkle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeLogEntry_Successor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeLogEntry_Successor) ProtoMessage() {}

func (x *ChangeLogEntry_Successor) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_merkle_pb2_merkle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeLogEntry_Successor.ProtoReflect.Descriptor instead.
func (*ChangeLogEntry_Successor) Descriptor() ([]byte, []int) {
	return file_protobuf_merkle_pb2_merkle_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ChangeLogEntry_Successor) GetSuccessor() []byte {
	if x != nil {
		return x.Successor
	}
	return nil
}

func (x *ChangeLogEntry_Successor) GetDeletions() [][]byte {
	if x != nil {
		return x.Deletions
	}
	return nil
}

var File_protobuf_merkle_pb2_merkle_proto protoreflect.FileDescriptor

var file_protobuf_merkle_pb2_merkle_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x65, 0x72, 0x6b, 0x6c,
	0x65, 0x5f, 0x70, 0x62, 0x32, 0x2f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x09, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x1a, 0x47, 0x0a, 0x09, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x25, 0x0a, 0x15, 0x73, 0x61, 0x77, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x01, 0x5a, 0x0a, 0x6d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x5f, 0x70, 0x62, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_merkle_pb2_merkle_proto_rawDescOnce sync.Once
	file_protobuf_merkle_pb2_merkle_proto_rawDescData = file_protobuf_merkle_pb2_merkle_proto_rawDesc
)

func file_protobuf_merkle_pb2_merkle_proto_rawDescGZIP() []byte {
	file_protobuf_merkle_pb2_merkle_proto_rawDescOnce.Do(func() {
		file_protobuf_merkle_pb2_merkle_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_merkle_pb2_merkle_proto_rawDescData)
	})
	return file_protobuf_merkle_pb2_merkle_proto_rawDescData
}

var file_protobuf_merkle_pb2_merkle_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_merkle_pb2_merkle_proto_goTypes = []interface{}{
	(*ChangeLogEntry)(nil),           // 0: ChangeLogEntry
	(*ChangeLogEntry_Successor)(nil), // 1: ChangeLogEntry.Successor
}
var file_protobuf_merkle_pb2_merkle_proto_depIdxs = []int32{
	1, // 0: ChangeLogEntry.successors:type_name -> ChangeLogEntry.Successor
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_merkle_pb2_merkle_proto_init() }
func file_protobuf_merkle_pb2_merkle_proto_init() {
	if File_protobuf_merkle_pb2_merkle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_merkle_pb2_merkle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeLogEntry); i {
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
		file_protobuf_merkle_pb2_merkle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeLogEntry_Successor); i {
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
			RawDescriptor: file_protobuf_merkle_pb2_merkle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_merkle_pb2_merkle_proto_goTypes,
		DependencyIndexes: file_protobuf_merkle_pb2_merkle_proto_depIdxs,
		MessageInfos:      file_protobuf_merkle_pb2_merkle_proto_msgTypes,
	}.Build()
	File_protobuf_merkle_pb2_merkle_proto = out.File
	file_protobuf_merkle_pb2_merkle_proto_rawDesc = nil
	file_protobuf_merkle_pb2_merkle_proto_goTypes = nil
	file_protobuf_merkle_pb2_merkle_proto_depIdxs = nil
}
