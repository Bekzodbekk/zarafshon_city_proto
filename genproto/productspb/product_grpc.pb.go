// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: protos/product.proto

package productspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProdcutService_CreateProduct_FullMethodName  = "/ProdcutService/CreateProduct"
	ProdcutService_UpdateProduct_FullMethodName  = "/ProdcutService/UpdateProduct"
	ProdcutService_DeleteProduct_FullMethodName  = "/ProdcutService/DeleteProduct"
	ProdcutService_GetProductById_FullMethodName = "/ProdcutService/GetProductById"
	ProdcutService_GetProducts_FullMethodName    = "/ProdcutService/GetProducts"
)

// ProdcutServiceClient is the client API for ProdcutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProdcutServiceClient interface {
	CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error)
	UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductResp, error)
	DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error)
	GetProductById(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdResp, error)
	GetProducts(ctx context.Context, in *GetProductsReq, opts ...grpc.CallOption) (*GetProductsResp, error)
}

type prodcutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProdcutServiceClient(cc grpc.ClientConnInterface) ProdcutServiceClient {
	return &prodcutServiceClient{cc}
}

func (c *prodcutServiceClient) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProductResp)
	err := c.cc.Invoke(ctx, ProdcutService_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodcutServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProductResp)
	err := c.cc.Invoke(ctx, ProdcutService_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodcutServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProductResp)
	err := c.cc.Invoke(ctx, ProdcutService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodcutServiceClient) GetProductById(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductByIdResp)
	err := c.cc.Invoke(ctx, ProdcutService_GetProductById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodcutServiceClient) GetProducts(ctx context.Context, in *GetProductsReq, opts ...grpc.CallOption) (*GetProductsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductsResp)
	err := c.cc.Invoke(ctx, ProdcutService_GetProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProdcutServiceServer is the server API for ProdcutService service.
// All implementations must embed UnimplementedProdcutServiceServer
// for forward compatibility.
type ProdcutServiceServer interface {
	CreateProduct(context.Context, *CreateProductReq) (*CreateProductResp, error)
	UpdateProduct(context.Context, *UpdateProductReq) (*UpdateProductResp, error)
	DeleteProduct(context.Context, *DeleteProductReq) (*DeleteProductResp, error)
	GetProductById(context.Context, *GetProductByIdReq) (*GetProductByIdResp, error)
	GetProducts(context.Context, *GetProductsReq) (*GetProductsResp, error)
	mustEmbedUnimplementedProdcutServiceServer()
}

// UnimplementedProdcutServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProdcutServiceServer struct{}

func (UnimplementedProdcutServiceServer) CreateProduct(context.Context, *CreateProductReq) (*CreateProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProdcutServiceServer) UpdateProduct(context.Context, *UpdateProductReq) (*UpdateProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProdcutServiceServer) DeleteProduct(context.Context, *DeleteProductReq) (*DeleteProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProdcutServiceServer) GetProductById(context.Context, *GetProductByIdReq) (*GetProductByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedProdcutServiceServer) GetProducts(context.Context, *GetProductsReq) (*GetProductsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProdcutServiceServer) mustEmbedUnimplementedProdcutServiceServer() {}
func (UnimplementedProdcutServiceServer) testEmbeddedByValue()                        {}

// UnsafeProdcutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProdcutServiceServer will
// result in compilation errors.
type UnsafeProdcutServiceServer interface {
	mustEmbedUnimplementedProdcutServiceServer()
}

func RegisterProdcutServiceServer(s grpc.ServiceRegistrar, srv ProdcutServiceServer) {
	// If the following call pancis, it indicates UnimplementedProdcutServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProdcutService_ServiceDesc, srv)
}

func _ProdcutService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdcutServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProdcutService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdcutServiceServer).CreateProduct(ctx, req.(*CreateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdcutService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdcutServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProdcutService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdcutServiceServer).UpdateProduct(ctx, req.(*UpdateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdcutService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdcutServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProdcutService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdcutServiceServer).DeleteProduct(ctx, req.(*DeleteProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdcutService_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdcutServiceServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProdcutService_GetProductById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdcutServiceServer).GetProductById(ctx, req.(*GetProductByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdcutService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdcutServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProdcutService_GetProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdcutServiceServer).GetProducts(ctx, req.(*GetProductsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProdcutService_ServiceDesc is the grpc.ServiceDesc for ProdcutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProdcutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProdcutService",
	HandlerType: (*ProdcutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProdcutService_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProdcutService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProdcutService_DeleteProduct_Handler,
		},
		{
			MethodName: "GetProductById",
			Handler:    _ProdcutService_GetProductById_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _ProdcutService_GetProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/product.proto",
}
