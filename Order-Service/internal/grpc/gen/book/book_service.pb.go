// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: book_service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_book_service_proto protoreflect.FileDescriptor

var file_book_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x72, 0x70, 0x63,
	0x5f, 0x64, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfc,
	0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x15, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x64, 0x61, 0x70, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6f, 0x6f,
	0x6b, 0x41, 0x64, 0x61, 0x70, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61,
	0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x15, 0x44, 0x65, 0x63,
	0x72, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61,
	0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x4b, 0x5a,
	0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x68, 0x69, 0x65,
	0x6d, 0x74, 0x61, 0x30, 0x33, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d,
	0x62, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_book_service_proto_goTypes = []interface{}{
	(*GetBookRequest)(nil),                // 0: pb.GetBookRequest
	(*CheckBookAdaptabilityRequest)(nil),  // 1: pb.CheckBookAdaptabilityRequest
	(*DecreaseStockQuantityRequest)(nil),  // 2: pb.DecreaseStockQuantityRequest
	(*Book)(nil),                          // 3: pb.Book
	(*DecreaseStockQuantityResponse)(nil), // 4: pb.DecreaseStockQuantityResponse
}
var file_book_service_proto_depIdxs = []int32{
	0, // 0: pb.BookService.GetBook:input_type -> pb.GetBookRequest
	1, // 1: pb.BookService.CheckBookAdaptability:input_type -> pb.CheckBookAdaptabilityRequest
	2, // 2: pb.BookService.DecreaseStockQuantity:input_type -> pb.DecreaseStockQuantityRequest
	3, // 3: pb.BookService.GetBook:output_type -> pb.Book
	4, // 4: pb.BookService.CheckBookAdaptability:output_type -> pb.DecreaseStockQuantityResponse
	4, // 5: pb.BookService.DecreaseStockQuantity:output_type -> pb.DecreaseStockQuantityResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_book_service_proto_init() }
func file_book_service_proto_init() {
	if File_book_service_proto != nil {
		return
	}
	file_book_proto_init()
	file_rpc_get_book_proto_init()
	file_rpc_check_book_adaptability_proto_init()
	file_rpc_decrease_stock_quantity_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_book_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_service_proto_goTypes,
		DependencyIndexes: file_book_service_proto_depIdxs,
	}.Build()
	File_book_service_proto = out.File
	file_book_service_proto_rawDesc = nil
	file_book_service_proto_goTypes = nil
	file_book_service_proto_depIdxs = nil
}
