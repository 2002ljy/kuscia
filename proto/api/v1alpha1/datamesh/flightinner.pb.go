// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: kuscia/proto/api/v1alpha1/datamesh/flightinner.proto

package datamesh

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

type CommandDataMeshQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query      *CommandDomainDataQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Domaindata *DomainData             `protobuf:"bytes,2,opt,name=domaindata,proto3" json:"domaindata,omitempty"`
	Datasource *DomainDataSource       `protobuf:"bytes,3,opt,name=datasource,proto3" json:"datasource,omitempty"`
}

func (x *CommandDataMeshQuery) Reset() {
	*x = CommandDataMeshQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandDataMeshQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandDataMeshQuery) ProtoMessage() {}

func (x *CommandDataMeshQuery) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandDataMeshQuery.ProtoReflect.Descriptor instead.
func (*CommandDataMeshQuery) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDescGZIP(), []int{0}
}

func (x *CommandDataMeshQuery) GetQuery() *CommandDomainDataQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CommandDataMeshQuery) GetDomaindata() *DomainData {
	if x != nil {
		return x.Domaindata
	}
	return nil
}

func (x *CommandDataMeshQuery) GetDatasource() *DomainDataSource {
	if x != nil {
		return x.Datasource
	}
	return nil
}

type CommandDataMeshUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Update     *CommandDomainDataUpdate `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
	Domaindata *DomainData              `protobuf:"bytes,2,opt,name=domaindata,proto3" json:"domaindata,omitempty"`
	Datasource *DomainDataSource        `protobuf:"bytes,3,opt,name=datasource,proto3" json:"datasource,omitempty"`
}

func (x *CommandDataMeshUpdate) Reset() {
	*x = CommandDataMeshUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandDataMeshUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandDataMeshUpdate) ProtoMessage() {}

func (x *CommandDataMeshUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandDataMeshUpdate.ProtoReflect.Descriptor instead.
func (*CommandDataMeshUpdate) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDescGZIP(), []int{1}
}

func (x *CommandDataMeshUpdate) GetUpdate() *CommandDomainDataUpdate {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *CommandDataMeshUpdate) GetDomaindata() *DomainData {
	if x != nil {
		return x.Domaindata
	}
	return nil
}

func (x *CommandDataMeshUpdate) GetDatasource() *DomainDataSource {
	if x != nil {
		return x.Datasource
	}
	return nil
}

var File_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto protoreflect.FileDescriptor

var file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x68, 0x1a, 0x33, 0x6b, 0x75, 0x73, 0x63,
	0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x39, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6b, 0x75, 0x73, 0x63,
	0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x64, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02,
	0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73,
	0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x50, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x4e, 0x0a, 0x0a, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b,
	0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x12, 0x54, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6b,
	0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x92,
	0x02, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x65,
	0x73, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x53, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4e, 0x0a,
	0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x12, 0x54, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x5c, 0x0a, 0x20, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x68, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6b,
	0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDescOnce sync.Once
	file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDescData = file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDesc
)

func file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDescGZIP() []byte {
	file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDescOnce.Do(func() {
		file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDescData = protoimpl.X.CompressGZIP(file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDescData)
	})
	return file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDescData
}

var file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_goTypes = []interface{}{
	(*CommandDataMeshQuery)(nil),    // 0: kuscia.proto.api.v1alpha1.datamesh.CommandDataMeshQuery
	(*CommandDataMeshUpdate)(nil),   // 1: kuscia.proto.api.v1alpha1.datamesh.CommandDataMeshUpdate
	(*CommandDomainDataQuery)(nil),  // 2: kuscia.proto.api.v1alpha1.datamesh.CommandDomainDataQuery
	(*DomainData)(nil),              // 3: kuscia.proto.api.v1alpha1.datamesh.DomainData
	(*DomainDataSource)(nil),        // 4: kuscia.proto.api.v1alpha1.datamesh.DomainDataSource
	(*CommandDomainDataUpdate)(nil), // 5: kuscia.proto.api.v1alpha1.datamesh.CommandDomainDataUpdate
}
var file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_depIdxs = []int32{
	2, // 0: kuscia.proto.api.v1alpha1.datamesh.CommandDataMeshQuery.query:type_name -> kuscia.proto.api.v1alpha1.datamesh.CommandDomainDataQuery
	3, // 1: kuscia.proto.api.v1alpha1.datamesh.CommandDataMeshQuery.domaindata:type_name -> kuscia.proto.api.v1alpha1.datamesh.DomainData
	4, // 2: kuscia.proto.api.v1alpha1.datamesh.CommandDataMeshQuery.datasource:type_name -> kuscia.proto.api.v1alpha1.datamesh.DomainDataSource
	5, // 3: kuscia.proto.api.v1alpha1.datamesh.CommandDataMeshUpdate.update:type_name -> kuscia.proto.api.v1alpha1.datamesh.CommandDomainDataUpdate
	3, // 4: kuscia.proto.api.v1alpha1.datamesh.CommandDataMeshUpdate.domaindata:type_name -> kuscia.proto.api.v1alpha1.datamesh.DomainData
	4, // 5: kuscia.proto.api.v1alpha1.datamesh.CommandDataMeshUpdate.datasource:type_name -> kuscia.proto.api.v1alpha1.datamesh.DomainDataSource
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_init() }
func file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_init() {
	if File_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto != nil {
		return
	}
	file_kuscia_proto_api_v1alpha1_datamesh_domaindata_proto_init()
	file_kuscia_proto_api_v1alpha1_datamesh_domaindatasource_proto_init()
	file_kuscia_proto_api_v1alpha1_datamesh_flightdm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandDataMeshQuery); i {
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
		file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandDataMeshUpdate); i {
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
			RawDescriptor: file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_goTypes,
		DependencyIndexes: file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_depIdxs,
		MessageInfos:      file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_msgTypes,
	}.Build()
	File_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto = out.File
	file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_rawDesc = nil
	file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_goTypes = nil
	file_kuscia_proto_api_v1alpha1_datamesh_flightinner_proto_depIdxs = nil
}
