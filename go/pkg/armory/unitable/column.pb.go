// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: armory/unitable/column.proto

package unitable

import (
	_ "github.com/mojo-lang/core/go/pkg/mojo"
	core "github.com/mojo-lang/core/go/pkg/mojo/core"
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

type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Database         string          `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	TableId          string          `protobuf:"bytes,3,opt,name=table_id,json=tableId,proto3" json:"tableId,omitempty"`
	Name             string          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName      string          `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"displayName,omitempty"`
	ExportName       string          `protobuf:"bytes,6,opt,name=export_name,json=exportName,proto3" json:"exportName,omitempty"`
	GroupDisplayName string          `protobuf:"bytes,7,opt,name=group_display_name,json=groupDisplayName,proto3" json:"groupDisplayName,omitempty"`
	Type             string          `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Format           string          `protobuf:"bytes,9,opt,name=format,proto3" json:"format,omitempty"`
	Indexed          bool            `protobuf:"varint,10,opt,name=indexed,proto3" json:"indexed,omitempty"`
	Unique           bool            `protobuf:"varint,11,opt,name=unique,proto3" json:"unique,omitempty"`
	Show             bool            `protobuf:"varint,12,opt,name=show,proto3" json:"show,omitempty"`
	Editable         bool            `protobuf:"varint,13,opt,name=editable,proto3" json:"editable,omitempty"`
	Filterable       bool            `protobuf:"varint,14,opt,name=filterable,proto3" json:"filterable,omitempty"`
	Temporal         bool            `protobuf:"varint,15,opt,name=temporal,proto3" json:"temporal,omitempty"`
	Dimensional      bool            `protobuf:"varint,16,opt,name=dimensional,proto3" json:"dimensional,omitempty"`
	Referenced       string          `protobuf:"bytes,20,opt,name=referenced,proto3" json:"referenced,omitempty"`
	CreateTime       *core.Timestamp `protobuf:"bytes,100,opt,name=create_time,json=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime       *core.Timestamp `protobuf:"bytes,101,opt,name=update_time,json=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_armory_unitable_column_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_armory_unitable_column_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_armory_unitable_column_proto_rawDescGZIP(), []int{0}
}

func (x *Column) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Column) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *Column) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

func (x *Column) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Column) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Column) GetExportName() string {
	if x != nil {
		return x.ExportName
	}
	return ""
}

func (x *Column) GetGroupDisplayName() string {
	if x != nil {
		return x.GroupDisplayName
	}
	return ""
}

func (x *Column) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Column) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *Column) GetIndexed() bool {
	if x != nil {
		return x.Indexed
	}
	return false
}

func (x *Column) GetUnique() bool {
	if x != nil {
		return x.Unique
	}
	return false
}

func (x *Column) GetShow() bool {
	if x != nil {
		return x.Show
	}
	return false
}

func (x *Column) GetEditable() bool {
	if x != nil {
		return x.Editable
	}
	return false
}

func (x *Column) GetFilterable() bool {
	if x != nil {
		return x.Filterable
	}
	return false
}

func (x *Column) GetTemporal() bool {
	if x != nil {
		return x.Temporal
	}
	return false
}

func (x *Column) GetDimensional() bool {
	if x != nil {
		return x.Dimensional
	}
	return false
}

func (x *Column) GetReferenced() string {
	if x != nil {
		return x.Referenced
	}
	return ""
}

func (x *Column) GetCreateTime() *core.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Column) GetUpdateTime() *core.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_armory_unitable_column_proto protoreflect.FileDescriptor

var file_armory_unitable_column_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x1a,
	0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x04, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x20, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0x8a, 0xd6, 0x24, 0x00, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x8a, 0xd6, 0x24, 0x00, 0x52, 0x07, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0x8a, 0xd6, 0x24, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x68, 0x6f, 0x77, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x68,
	0x6f, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x67, 0x0a, 0x19, 0x69, 0x6f,
	0x2e, 0x6e, 0x63, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x75,
	0x6e, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x0b, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x63, 0x72, 0x61, 0x66, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x72, 0x6d,
	0x6f, 0x72, 0x79, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x72, 0x6d, 0x6f, 0x72,
	0x79, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x3b, 0x75, 0x6e, 0x69, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_armory_unitable_column_proto_rawDescOnce sync.Once
	file_armory_unitable_column_proto_rawDescData = file_armory_unitable_column_proto_rawDesc
)

func file_armory_unitable_column_proto_rawDescGZIP() []byte {
	file_armory_unitable_column_proto_rawDescOnce.Do(func() {
		file_armory_unitable_column_proto_rawDescData = protoimpl.X.CompressGZIP(file_armory_unitable_column_proto_rawDescData)
	})
	return file_armory_unitable_column_proto_rawDescData
}

var file_armory_unitable_column_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_armory_unitable_column_proto_goTypes = []interface{}{
	(*Column)(nil),         // 0: armory.unitable.Column
	(*core.Timestamp)(nil), // 1: mojo.core.Timestamp
}
var file_armory_unitable_column_proto_depIdxs = []int32{
	1, // 0: armory.unitable.Column.create_time:type_name -> mojo.core.Timestamp
	1, // 1: armory.unitable.Column.update_time:type_name -> mojo.core.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_armory_unitable_column_proto_init() }
func file_armory_unitable_column_proto_init() {
	if File_armory_unitable_column_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_armory_unitable_column_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column); i {
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
			RawDescriptor: file_armory_unitable_column_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_armory_unitable_column_proto_goTypes,
		DependencyIndexes: file_armory_unitable_column_proto_depIdxs,
		MessageInfos:      file_armory_unitable_column_proto_msgTypes,
	}.Build()
	File_armory_unitable_column_proto = out.File
	file_armory_unitable_column_proto_rawDesc = nil
	file_armory_unitable_column_proto_goTypes = nil
	file_armory_unitable_column_proto_depIdxs = nil
}
