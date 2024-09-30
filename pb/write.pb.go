// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: write.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WriteType int32

const (
	WriteType_WRITE_FULL WriteType = 0
)

// Enum value maps for WriteType.
var (
	WriteType_name = map[int32]string{
		0: "WRITE_FULL",
	}
	WriteType_value = map[string]int32{
		"WRITE_FULL": 0,
	}
)

func (x WriteType) Enum() *WriteType {
	p := new(WriteType)
	*p = x
	return p
}

func (x WriteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WriteType) Descriptor() protoreflect.EnumDescriptor {
	return file_write_proto_enumTypes[0].Descriptor()
}

func (WriteType) Type() protoreflect.EnumType {
	return &file_write_proto_enumTypes[0]
}

func (x WriteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WriteType.Descriptor instead.
func (WriteType) EnumDescriptor() ([]byte, []int) {
	return file_write_proto_rawDescGZIP(), []int{0}
}

type Write struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WriteType     WriteType              `protobuf:"varint,1,opt,name=write_type,json=writeType,proto3,enum=synk.WriteType" json:"write_type,omitempty"`
	WriteData     *WriteData             `protobuf:"bytes,2,opt,name=write_data,json=writeData,proto3" json:"write_data,omitempty"`
	AccessToken   string                 `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	SheetName     string                 `protobuf:"bytes,4,opt,name=sheet_name,json=sheetName,proto3" json:"sheet_name,omitempty"`
	SpreadsheetId string                 `protobuf:"bytes,5,opt,name=spreadsheet_id,json=spreadsheetId,proto3" json:"spreadsheet_id,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Write) Reset() {
	*x = Write{}
	if protoimpl.UnsafeEnabled {
		mi := &file_write_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Write) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Write) ProtoMessage() {}

func (x *Write) ProtoReflect() protoreflect.Message {
	mi := &file_write_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Write.ProtoReflect.Descriptor instead.
func (*Write) Descriptor() ([]byte, []int) {
	return file_write_proto_rawDescGZIP(), []int{0}
}

func (x *Write) GetWriteType() WriteType {
	if x != nil {
		return x.WriteType
	}
	return WriteType_WRITE_FULL
}

func (x *Write) GetWriteData() *WriteData {
	if x != nil {
		return x.WriteData
	}
	return nil
}

func (x *Write) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Write) GetSheetName() string {
	if x != nil {
		return x.SheetName
	}
	return ""
}

func (x *Write) GetSpreadsheetId() string {
	if x != nil {
		return x.SpreadsheetId
	}
	return ""
}

func (x *Write) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_write_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_write_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_write_proto_rawDescGZIP(), []int{1}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type WriteData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DynamicFields []*KeyValue `protobuf:"bytes,1,rep,name=dynamic_fields,json=dynamicFields,proto3" json:"dynamic_fields,omitempty"`
}

func (x *WriteData) Reset() {
	*x = WriteData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_write_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteData) ProtoMessage() {}

func (x *WriteData) ProtoReflect() protoreflect.Message {
	mi := &file_write_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteData.ProtoReflect.Descriptor instead.
func (*WriteData) Descriptor() ([]byte, []int) {
	return file_write_proto_rawDescGZIP(), []int{2}
}

func (x *WriteData) GetDynamicFields() []*KeyValue {
	if x != nil {
		return x.DynamicFields
	}
	return nil
}

var File_write_proto protoreflect.FileDescriptor

var file_write_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73,
	0x79, 0x6e, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x2e,
	0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x79, 0x6e, 0x6b, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e,
	0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x79, 0x6e, 0x6b, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x65, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64,
	0x73, 0x68, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x42, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x35, 0x0a, 0x0e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x79, 0x6e,
	0x6b, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2a, 0x1b, 0x0a, 0x09, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f,
	0x46, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_write_proto_rawDescOnce sync.Once
	file_write_proto_rawDescData = file_write_proto_rawDesc
)

func file_write_proto_rawDescGZIP() []byte {
	file_write_proto_rawDescOnce.Do(func() {
		file_write_proto_rawDescData = protoimpl.X.CompressGZIP(file_write_proto_rawDescData)
	})
	return file_write_proto_rawDescData
}

var file_write_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_write_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_write_proto_goTypes = []any{
	(WriteType)(0),                // 0: synk.WriteType
	(*Write)(nil),                 // 1: synk.Write
	(*KeyValue)(nil),              // 2: synk.KeyValue
	(*WriteData)(nil),             // 3: synk.WriteData
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_write_proto_depIdxs = []int32{
	0, // 0: synk.Write.write_type:type_name -> synk.WriteType
	3, // 1: synk.Write.write_data:type_name -> synk.WriteData
	4, // 2: synk.Write.timestamp:type_name -> google.protobuf.Timestamp
	2, // 3: synk.WriteData.dynamic_fields:type_name -> synk.KeyValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_write_proto_init() }
func file_write_proto_init() {
	if File_write_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_write_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Write); i {
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
		file_write_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*KeyValue); i {
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
		file_write_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WriteData); i {
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
			RawDescriptor: file_write_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_write_proto_goTypes,
		DependencyIndexes: file_write_proto_depIdxs,
		EnumInfos:         file_write_proto_enumTypes,
		MessageInfos:      file_write_proto_msgTypes,
	}.Build()
	File_write_proto = out.File
	file_write_proto_rawDesc = nil
	file_write_proto_goTypes = nil
	file_write_proto_depIdxs = nil
}
