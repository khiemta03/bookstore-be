// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: book_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (BookService_ListBooksClient, error)
	GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	GetPublisher(ctx context.Context, in *GetPublisherRequest, opts ...grpc.CallOption) (*Publisher, error)
	AddNewBook(ctx context.Context, in *AddNewBookRequest, opts ...grpc.CallOption) (*Book, error)
	AddNewAuthor(ctx context.Context, in *AddNewAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	AddNewPublisher(ctx context.Context, in *AddNewPublisherRequest, opts ...grpc.CallOption) (*Publisher, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdatePublisher(ctx context.Context, in *UpdatePublisherRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/pb.BookService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (BookService_ListBooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookService_ServiceDesc.Streams[0], "/pb.BookService/ListBooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceListBooksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookService_ListBooksClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type bookServiceListBooksClient struct {
	grpc.ClientStream
}

func (x *bookServiceListBooksClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookServiceClient) GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/pb.BookService/GetAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetPublisher(ctx context.Context, in *GetPublisherRequest, opts ...grpc.CallOption) (*Publisher, error) {
	out := new(Publisher)
	err := c.cc.Invoke(ctx, "/pb.BookService/GetPublisher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) AddNewBook(ctx context.Context, in *AddNewBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/pb.BookService/AddNewBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) AddNewAuthor(ctx context.Context, in *AddNewAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/pb.BookService/AddNewAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) AddNewPublisher(ctx context.Context, in *AddNewPublisherRequest, opts ...grpc.CallOption) (*Publisher, error) {
	out := new(Publisher)
	err := c.cc.Invoke(ctx, "/pb.BookService/AddNewPublisher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.BookService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.BookService/UpdateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdatePublisher(ctx context.Context, in *UpdatePublisherRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.BookService/UpdatePublisher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	ListBooks(*ListBooksRequest, BookService_ListBooksServer) error
	GetAuthor(context.Context, *GetAuthorRequest) (*Author, error)
	GetPublisher(context.Context, *GetPublisherRequest) (*Publisher, error)
	AddNewBook(context.Context, *AddNewBookRequest) (*Book, error)
	AddNewAuthor(context.Context, *AddNewAuthorRequest) (*Author, error)
	AddNewPublisher(context.Context, *AddNewPublisherRequest) (*Publisher, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*emptypb.Empty, error)
	UpdateAuthor(context.Context, *UpdateAuthorRequest) (*emptypb.Empty, error)
	UpdatePublisher(context.Context, *UpdatePublisherRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) GetBook(context.Context, *GetBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookServiceServer) ListBooks(*ListBooksRequest, BookService_ListBooksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedBookServiceServer) GetAuthor(context.Context, *GetAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}
func (UnimplementedBookServiceServer) GetPublisher(context.Context, *GetPublisherRequest) (*Publisher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublisher not implemented")
}
func (UnimplementedBookServiceServer) AddNewBook(context.Context, *AddNewBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewBook not implemented")
}
func (UnimplementedBookServiceServer) AddNewAuthor(context.Context, *AddNewAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewAuthor not implemented")
}
func (UnimplementedBookServiceServer) AddNewPublisher(context.Context, *AddNewPublisherRequest) (*Publisher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewPublisher not implemented")
}
func (UnimplementedBookServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookServiceServer) UpdateAuthor(context.Context, *UpdateAuthorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthor not implemented")
}
func (UnimplementedBookServiceServer) UpdatePublisher(context.Context, *UpdatePublisherRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePublisher not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_ListBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBooksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServiceServer).ListBooks(m, &bookServiceListBooksServer{stream})
}

type BookService_ListBooksServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type bookServiceListBooksServer struct {
	grpc.ServerStream
}

func (x *bookServiceListBooksServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

func _BookService_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookService/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetAuthor(ctx, req.(*GetAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetPublisher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetPublisher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookService/GetPublisher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetPublisher(ctx, req.(*GetPublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_AddNewBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).AddNewBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookService/AddNewBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).AddNewBook(ctx, req.(*AddNewBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_AddNewAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).AddNewAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookService/AddNewAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).AddNewAuthor(ctx, req.(*AddNewAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_AddNewPublisher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewPublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).AddNewPublisher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookService/AddNewPublisher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).AddNewPublisher(ctx, req.(*AddNewPublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookService/UpdateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateAuthor(ctx, req.(*UpdateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdatePublisher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdatePublisher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookService/UpdatePublisher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdatePublisher(ctx, req.(*UpdatePublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
		{
			MethodName: "GetAuthor",
			Handler:    _BookService_GetAuthor_Handler,
		},
		{
			MethodName: "GetPublisher",
			Handler:    _BookService_GetPublisher_Handler,
		},
		{
			MethodName: "AddNewBook",
			Handler:    _BookService_AddNewBook_Handler,
		},
		{
			MethodName: "AddNewAuthor",
			Handler:    _BookService_AddNewAuthor_Handler,
		},
		{
			MethodName: "AddNewPublisher",
			Handler:    _BookService_AddNewPublisher_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookService_UpdateBook_Handler,
		},
		{
			MethodName: "UpdateAuthor",
			Handler:    _BookService_UpdateAuthor_Handler,
		},
		{
			MethodName: "UpdatePublisher",
			Handler:    _BookService_UpdatePublisher_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBooks",
			Handler:       _BookService_ListBooks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "book_service.proto",
}
