// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: discount.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Discount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscountId    string               `protobuf:"bytes,1,opt,name=discount_id,json=discountId,proto3" json:"discount_id,omitempty"`
	DiscountCode  string               `protobuf:"bytes,2,opt,name=discount_code,json=discountCode,proto3" json:"discount_code,omitempty"`
	DiscountValue float64              `protobuf:"fixed64,3,opt,name=discount_value,json=discountValue,proto3" json:"discount_value,omitempty"`
	StartDate     *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate       *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *Discount) Reset() {
	*x = Discount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Discount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discount) ProtoMessage() {}

func (x *Discount) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discount.ProtoReflect.Descriptor instead.
func (*Discount) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{0}
}

func (x *Discount) GetDiscountId() string {
	if x != nil {
		return x.DiscountId
	}
	return ""
}

func (x *Discount) GetDiscountCode() string {
	if x != nil {
		return x.DiscountCode
	}
	return ""
}

func (x *Discount) GetDiscountValue() float64 {
	if x != nil {
		return x.DiscountValue
	}
	return 0
}

func (x *Discount) GetStartDate() *timestamp.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Discount) GetEndDate() *timestamp.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

var File_discount_proto protoreflect.FileDescriptor

var file_discount_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x68, 0x69, 0x65, 0x6d, 0x74, 0x61, 0x30, 0x33, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2d, 0x62, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discount_proto_rawDescOnce sync.Once
	file_discount_proto_rawDescData = file_discount_proto_rawDesc
)

func file_discount_proto_rawDescGZIP() []byte {
	file_discount_proto_rawDescOnce.Do(func() {
		file_discount_proto_rawDescData = protoimpl.X.CompressGZIP(file_discount_proto_rawDescData)
	})
	return file_discount_proto_rawDescData
}

var file_discount_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_discount_proto_goTypes = []interface{}{
	(*Discount)(nil),            // 0: pb.Discount
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_discount_proto_depIdxs = []int32{
	1, // 0: pb.Discount.start_date:type_name -> google.protobuf.Timestamp
	1, // 1: pb.Discount.end_date:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_discount_proto_init() }
func file_discount_proto_init() {
	if File_discount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Discount); i {
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
			RawDescriptor: file_discount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_discount_proto_goTypes,
		DependencyIndexes: file_discount_proto_depIdxs,
		MessageInfos:      file_discount_proto_msgTypes,
	}.Build()
	File_discount_proto = out.File
	file_discount_proto_rawDesc = nil
	file_discount_proto_goTypes = nil
	file_discount_proto_depIdxs = nil
}
