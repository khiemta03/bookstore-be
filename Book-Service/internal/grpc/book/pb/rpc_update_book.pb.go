// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: rpc_update_book.proto

package pb

import (
	date "google.golang.org/genproto/googleapis/type/date"
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

type UpdateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title           string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	FullTitle       string     `protobuf:"bytes,3,opt,name=full_title,json=fullTitle,proto3" json:"full_title,omitempty"`
	Publisher       string     `protobuf:"bytes,4,opt,name=publisher,proto3" json:"publisher,omitempty"`
	PublicationDate *date.Date `protobuf:"bytes,5,opt,name=publication_date,json=publicationDate,proto3" json:"publication_date,omitempty"`
	Isbn            string     `protobuf:"bytes,6,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Description     string     `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Price           float64    `protobuf:"fixed64,8,opt,name=price,proto3" json:"price,omitempty"`
	StockQuantity   int32      `protobuf:"varint,9,opt,name=stock_quantity,json=stockQuantity,proto3" json:"stock_quantity,omitempty"`
	FrontCoverImage string     `protobuf:"bytes,10,opt,name=front_cover_image,json=frontCoverImage,proto3" json:"front_cover_image,omitempty"`
	BackCoverImage  string     `protobuf:"bytes,11,opt,name=back_cover_image,json=backCoverImage,proto3" json:"back_cover_image,omitempty"`
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_update_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_update_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_rpc_update_book_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateBookRequest) GetFullTitle() string {
	if x != nil {
		return x.FullTitle
	}
	return ""
}

func (x *UpdateBookRequest) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *UpdateBookRequest) GetPublicationDate() *date.Date {
	if x != nil {
		return x.PublicationDate
	}
	return nil
}

func (x *UpdateBookRequest) GetIsbn() string {
	if x != nil {
		return x.Isbn
	}
	return ""
}

func (x *UpdateBookRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateBookRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateBookRequest) GetStockQuantity() int32 {
	if x != nil {
		return x.StockQuantity
	}
	return 0
}

func (x *UpdateBookRequest) GetFrontCoverImage() string {
	if x != nil {
		return x.FrontCoverImage
	}
	return ""
}

func (x *UpdateBookRequest) GetBackCoverImage() string {
	if x != nil {
		return x.BackCoverImage
	}
	return ""
}

var File_rpc_update_book_proto protoreflect.FileDescriptor

var file_rpc_update_book_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x02, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x12, 0x3c, 0x0a, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x73,
	0x62, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68, 0x69, 0x65, 0x6d, 0x74, 0x61, 0x30, 0x33, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x62, 0x65, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x62, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_update_book_proto_rawDescOnce sync.Once
	file_rpc_update_book_proto_rawDescData = file_rpc_update_book_proto_rawDesc
)

func file_rpc_update_book_proto_rawDescGZIP() []byte {
	file_rpc_update_book_proto_rawDescOnce.Do(func() {
		file_rpc_update_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_update_book_proto_rawDescData)
	})
	return file_rpc_update_book_proto_rawDescData
}

var file_rpc_update_book_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rpc_update_book_proto_goTypes = []interface{}{
	(*UpdateBookRequest)(nil), // 0: pb.UpdateBookRequest
	(*date.Date)(nil),         // 1: google.type.Date
}
var file_rpc_update_book_proto_depIdxs = []int32{
	1, // 0: pb.UpdateBookRequest.publication_date:type_name -> google.type.Date
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_update_book_proto_init() }
func file_rpc_update_book_proto_init() {
	if File_rpc_update_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_update_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookRequest); i {
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
			RawDescriptor: file_rpc_update_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_update_book_proto_goTypes,
		DependencyIndexes: file_rpc_update_book_proto_depIdxs,
		MessageInfos:      file_rpc_update_book_proto_msgTypes,
	}.Build()
	File_rpc_update_book_proto = out.File
	file_rpc_update_book_proto_rawDesc = nil
	file_rpc_update_book_proto_goTypes = nil
	file_rpc_update_book_proto_depIdxs = nil
}
