// Generated by github.com/macro-funs/tabkit
// DO NOT EDIT!!
// Version: 3.0.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: table.proto

package main

import (
	proto "google.golang.org/protobuf/proto"
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

type ActorType int32

const (
	ActorType_None    ActorType = 0 //
	ActorType_Pharah  ActorType = 1 // 法鸡
	ActorType_Junkrat ActorType = 2 // 狂鼠
	ActorType_Genji   ActorType = 3 // 源氏
	ActorType_Mercy   ActorType = 4 // 天使
)

// Enum value maps for ActorType.
var (
	ActorType_name = map[int32]string{
		0: "None",
		1: "Pharah",
		2: "Junkrat",
		3: "Genji",
		4: "Mercy",
	}
	ActorType_value = map[string]int32{
		"None":    0,
		"Pharah":  1,
		"Junkrat": 2,
		"Genji":   3,
		"Mercy":   4,
	}
)

func (x ActorType) Enum() *ActorType {
	p := new(ActorType)
	*p = x
	return p
}

func (x ActorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActorType) Descriptor() protoreflect.EnumDescriptor {
	return file_table_proto_enumTypes[0].Descriptor()
}

func (ActorType) Type() protoreflect.EnumType {
	return &file_table_proto_enumTypes[0]
}

func (x ActorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActorType.Descriptor instead.
func (ActorType) EnumDescriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{0}
}

type ExampleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int32     `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string    `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Rate     float32   `protobuf:"fixed32,3,opt,name=Rate,proto3" json:"Rate,omitempty"`
	Accuracy float64   `protobuf:"fixed64,4,opt,name=Accuracy,proto3" json:"Accuracy,omitempty"`
	Type     ActorType `protobuf:"varint,5,opt,name=Type,proto3,enum=main.ActorType" json:"Type,omitempty"`
	Skill    []int32   `protobuf:"varint,6,rep,packed,name=Skill,proto3" json:"Skill,omitempty"`
	Buff     int32     `protobuf:"varint,7,opt,name=Buff,proto3" json:"Buff,omitempty"`
	TagList  []string  `protobuf:"bytes,8,rep,name=TagList,proto3" json:"TagList,omitempty"`
	Multi    []int32   `protobuf:"varint,9,rep,packed,name=Multi,proto3" json:"Multi,omitempty"`
}

func (x *ExampleData) Reset() {
	*x = ExampleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleData) ProtoMessage() {}

func (x *ExampleData) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleData.ProtoReflect.Descriptor instead.
func (*ExampleData) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleData) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ExampleData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExampleData) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *ExampleData) GetAccuracy() float64 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

func (x *ExampleData) GetType() ActorType {
	if x != nil {
		return x.Type
	}
	return ActorType_None
}

func (x *ExampleData) GetSkill() []int32 {
	if x != nil {
		return x.Skill
	}
	return nil
}

func (x *ExampleData) GetBuff() int32 {
	if x != nil {
		return x.Buff
	}
	return 0
}

func (x *ExampleData) GetTagList() []string {
	if x != nil {
		return x.TagList
	}
	return nil
}

func (x *ExampleData) GetMulti() []int32 {
	if x != nil {
		return x.Multi
	}
	return nil
}

type ExtendData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Additive float32 `protobuf:"fixed32,1,opt,name=Additive,proto3" json:"Additive,omitempty"`
}

func (x *ExtendData) Reset() {
	*x = ExtendData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendData) ProtoMessage() {}

func (x *ExtendData) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendData.ProtoReflect.Descriptor instead.
func (*ExtendData) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{1}
}

func (x *ExtendData) GetAdditive() float32 {
	if x != nil {
		return x.Additive
	}
	return 0
}

type ExampleKV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerIP   string  `protobuf:"bytes,1,opt,name=ServerIP,proto3" json:"ServerIP,omitempty"`
	ServerPort uint32  `protobuf:"varint,2,opt,name=ServerPort,proto3" json:"ServerPort,omitempty"`
	GroupID    []int32 `protobuf:"varint,3,rep,packed,name=GroupID,proto3" json:"GroupID,omitempty"`
}

func (x *ExampleKV) Reset() {
	*x = ExampleKV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleKV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleKV) ProtoMessage() {}

func (x *ExampleKV) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleKV.ProtoReflect.Descriptor instead.
func (*ExampleKV) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{2}
}

func (x *ExampleKV) GetServerIP() string {
	if x != nil {
		return x.ServerIP
	}
	return ""
}

func (x *ExampleKV) GetServerPort() uint32 {
	if x != nil {
		return x.ServerPort
	}
	return 0
}

func (x *ExampleKV) GetGroupID() []int32 {
	if x != nil {
		return x.GroupID
	}
	return nil
}

// Combine struct
type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExampleData []*ExampleData `protobuf:"bytes,1,rep,name=ExampleData,proto3" json:"ExampleData,omitempty"` // table: ExampleData
	ExtendData  []*ExtendData  `protobuf:"bytes,2,rep,name=ExtendData,proto3" json:"ExtendData,omitempty"`   // table: ExtendData
	ExampleKV   []*ExampleKV   `protobuf:"bytes,3,rep,name=ExampleKV,proto3" json:"ExampleKV,omitempty"`     // table: ExampleKV
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{3}
}

func (x *Table) GetExampleData() []*ExampleData {
	if x != nil {
		return x.ExampleData
	}
	return nil
}

func (x *Table) GetExtendData() []*ExtendData {
	if x != nil {
		return x.ExtendData
	}
	return nil
}

func (x *Table) GetExampleKV() []*ExampleKV {
	if x != nil {
		return x.ExampleKV
	}
	return nil
}

var File_table_proto protoreflect.FileDescriptor

var file_table_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d,
	0x61, 0x69, 0x6e, 0x22, 0xe0, 0x01, 0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41,
	0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x41,
	0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x23, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x74,
	0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x75, 0x66, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x42, 0x75, 0x66, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x18, 0x09, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x05, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x22, 0x28, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x22, 0x61, 0x0a, 0x09, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4b, 0x56, 0x12, 0x1a, 0x0a,
	0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x50, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x44, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x44, 0x22, 0x9d, 0x01, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x33, 0x0a,
	0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x30, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x09, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4b,
	0x56, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4b, 0x56, 0x52, 0x09, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x4b, 0x56, 0x2a, 0x44, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x68,
	0x61, 0x72, 0x61, 0x68, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4a, 0x75, 0x6e, 0x6b, 0x72, 0x61,
	0x74, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x6a, 0x69, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x4d, 0x65, 0x72, 0x63, 0x79, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_table_proto_rawDescOnce sync.Once
	file_table_proto_rawDescData = file_table_proto_rawDesc
)

func file_table_proto_rawDescGZIP() []byte {
	file_table_proto_rawDescOnce.Do(func() {
		file_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_table_proto_rawDescData)
	})
	return file_table_proto_rawDescData
}

var file_table_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_table_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_table_proto_goTypes = []interface{}{
	(ActorType)(0),      // 0: main.ActorType
	(*ExampleData)(nil), // 1: main.ExampleData
	(*ExtendData)(nil),  // 2: main.ExtendData
	(*ExampleKV)(nil),   // 3: main.ExampleKV
	(*Table)(nil),       // 4: main.Table
}
var file_table_proto_depIdxs = []int32{
	0, // 0: main.ExampleData.Type:type_name -> main.ActorType
	1, // 1: main.Table.ExampleData:type_name -> main.ExampleData
	2, // 2: main.Table.ExtendData:type_name -> main.ExtendData
	3, // 3: main.Table.ExampleKV:type_name -> main.ExampleKV
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_table_proto_init() }
func file_table_proto_init() {
	if File_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleData); i {
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
		file_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendData); i {
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
		file_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleKV); i {
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
		file_table_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
			RawDescriptor: file_table_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_table_proto_goTypes,
		DependencyIndexes: file_table_proto_depIdxs,
		EnumInfos:         file_table_proto_enumTypes,
		MessageInfos:      file_table_proto_msgTypes,
	}.Build()
	File_table_proto = out.File
	file_table_proto_rawDesc = nil
	file_table_proto_goTypes = nil
	file_table_proto_depIdxs = nil
}
