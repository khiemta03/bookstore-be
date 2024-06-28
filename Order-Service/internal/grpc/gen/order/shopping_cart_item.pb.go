// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: shopping_cart_item.proto

package pb

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

type ShoppingCartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartItemId string  `protobuf:"bytes,1,opt,name=cart_item_id,json=cartItemId,proto3" json:"cart_item_id,omitempty"`
	UserId     string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId     string  `protobuf:"bytes,3,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Quantity   int32   `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UnitPrice  float64 `protobuf:"fixed64,5,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
}

func (x *ShoppingCartItem) Reset() {
	*x = ShoppingCartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_cart_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShoppingCartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShoppingCartItem) ProtoMessage() {}

func (x *ShoppingCartItem) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_cart_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShoppingCartItem.ProtoReflect.Descriptor instead.
func (*ShoppingCartItem) Descriptor() ([]byte, []int) {
	return file_shopping_cart_item_proto_rawDescGZIP(), []int{0}
}

func (x *ShoppingCartItem) GetCartItemId() string {
	if x != nil {
		return x.CartItemId
	}
	return ""
}

func (x *ShoppingCartItem) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ShoppingCartItem) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *ShoppingCartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ShoppingCartItem) GetUnitPrice() float64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

var File_shopping_cart_item_proto protoreflect.FileDescriptor

var file_shopping_cart_item_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x61, 0x72, 0x74, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xa1,
	0x01, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x68, 0x69, 0x65, 0x6d, 0x74, 0x61, 0x30, 0x33, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2d, 0x62, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shopping_cart_item_proto_rawDescOnce sync.Once
	file_shopping_cart_item_proto_rawDescData = file_shopping_cart_item_proto_rawDesc
)

func file_shopping_cart_item_proto_rawDescGZIP() []byte {
	file_shopping_cart_item_proto_rawDescOnce.Do(func() {
		file_shopping_cart_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_shopping_cart_item_proto_rawDescData)
	})
	return file_shopping_cart_item_proto_rawDescData
}

var file_shopping_cart_item_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shopping_cart_item_proto_goTypes = []interface{}{
	(*ShoppingCartItem)(nil), // 0: pb.ShoppingCartItem
}
var file_shopping_cart_item_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shopping_cart_item_proto_init() }
func file_shopping_cart_item_proto_init() {
	if File_shopping_cart_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shopping_cart_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShoppingCartItem); i {
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
			RawDescriptor: file_shopping_cart_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shopping_cart_item_proto_goTypes,
		DependencyIndexes: file_shopping_cart_item_proto_depIdxs,
		MessageInfos:      file_shopping_cart_item_proto_msgTypes,
	}.Build()
	File_shopping_cart_item_proto = out.File
	file_shopping_cart_item_proto_rawDesc = nil
	file_shopping_cart_item_proto_goTypes = nil
	file_shopping_cart_item_proto_depIdxs = nil
}
