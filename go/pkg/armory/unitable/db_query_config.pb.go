// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: armory/unitable/db_query_config.proto

package unitable

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

type DbQueryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queries []*DbQuery `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty"`
}

func (x *DbQueryConfig) Reset() {
	*x = DbQueryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_armory_unitable_db_query_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbQueryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbQueryConfig) ProtoMessage() {}

func (x *DbQueryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_armory_unitable_db_query_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbQueryConfig.ProtoReflect.Descriptor instead.
func (*DbQueryConfig) Descriptor() ([]byte, []int) {
	return file_armory_unitable_db_query_config_proto_rawDescGZIP(), []int{0}
}

func (x *DbQueryConfig) GetQueries() []*DbQuery {
	if x != nil {
		return x.Queries
	}
	return nil
}

var File_armory_unitable_db_query_config_proto protoreflect.FileDescriptor

var file_armory_unitable_db_query_config_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2f, 0x64, 0x62, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79, 0x2e,
	0x75, 0x6e, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x1e, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79,
	0x2f, 0x75, 0x6e, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x64, 0x62, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0d, 0x44, 0x62, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a, 0x07, 0x71, 0x75, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x72, 0x6d,
	0x6f, 0x72, 0x79, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x62, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x6e, 0x0a,
	0x19, 0x69, 0x6f, 0x2e, 0x6e, 0x63, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72,
	0x79, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x12, 0x44, 0x62, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x63, 0x72,
	0x61, 0x66, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79, 0x2f, 0x75, 0x6e, 0x69, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x3b, 0x75, 0x6e, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_armory_unitable_db_query_config_proto_rawDescOnce sync.Once
	file_armory_unitable_db_query_config_proto_rawDescData = file_armory_unitable_db_query_config_proto_rawDesc
)

func file_armory_unitable_db_query_config_proto_rawDescGZIP() []byte {
	file_armory_unitable_db_query_config_proto_rawDescOnce.Do(func() {
		file_armory_unitable_db_query_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_armory_unitable_db_query_config_proto_rawDescData)
	})
	return file_armory_unitable_db_query_config_proto_rawDescData
}

var file_armory_unitable_db_query_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_armory_unitable_db_query_config_proto_goTypes = []interface{}{
	(*DbQueryConfig)(nil), // 0: armory.unitable.DbQueryConfig
	(*DbQuery)(nil),       // 1: armory.unitable.DbQuery
}
var file_armory_unitable_db_query_config_proto_depIdxs = []int32{
	1, // 0: armory.unitable.DbQueryConfig.queries:type_name -> armory.unitable.DbQuery
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_armory_unitable_db_query_config_proto_init() }
func file_armory_unitable_db_query_config_proto_init() {
	if File_armory_unitable_db_query_config_proto != nil {
		return
	}
	file_armory_unitable_db_query_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_armory_unitable_db_query_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DbQueryConfig); i {
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
			RawDescriptor: file_armory_unitable_db_query_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_armory_unitable_db_query_config_proto_goTypes,
		DependencyIndexes: file_armory_unitable_db_query_config_proto_depIdxs,
		MessageInfos:      file_armory_unitable_db_query_config_proto_msgTypes,
	}.Build()
	File_armory_unitable_db_query_config_proto = out.File
	file_armory_unitable_db_query_config_proto_rawDesc = nil
	file_armory_unitable_db_query_config_proto_goTypes = nil
	file_armory_unitable_db_query_config_proto_depIdxs = nil
}