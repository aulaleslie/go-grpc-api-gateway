// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: pkg/asset_sub_group/pb/asset_sub_group.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AssetSubGroupService_Create_FullMethodName = "/asset_sub_group.AssetSubGroupService/Create"
	AssetSubGroupService_Update_FullMethodName = "/asset_sub_group.AssetSubGroupService/Update"
	AssetSubGroupService_Delete_FullMethodName = "/asset_sub_group.AssetSubGroupService/Delete"
	AssetSubGroupService_Read_FullMethodName   = "/asset_sub_group.AssetSubGroupService/Read"
)

// AssetSubGroupServiceClient is the client API for AssetSubGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetSubGroupServiceClient interface {
	Create(ctx context.Context, in *CreateUpdateRequest, opts ...grpc.CallOption) (*CUDResponse, error)
	Update(ctx context.Context, in *CreateUpdateRequest, opts ...grpc.CallOption) (*CUDResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*CUDResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
}

type assetSubGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetSubGroupServiceClient(cc grpc.ClientConnInterface) AssetSubGroupServiceClient {
	return &assetSubGroupServiceClient{cc}
}

func (c *assetSubGroupServiceClient) Create(ctx context.Context, in *CreateUpdateRequest, opts ...grpc.CallOption) (*CUDResponse, error) {
	out := new(CUDResponse)
	err := c.cc.Invoke(ctx, AssetSubGroupService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetSubGroupServiceClient) Update(ctx context.Context, in *CreateUpdateRequest, opts ...grpc.CallOption) (*CUDResponse, error) {
	out := new(CUDResponse)
	err := c.cc.Invoke(ctx, AssetSubGroupService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetSubGroupServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*CUDResponse, error) {
	out := new(CUDResponse)
	err := c.cc.Invoke(ctx, AssetSubGroupService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetSubGroupServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, AssetSubGroupService_Read_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetSubGroupServiceServer is the server API for AssetSubGroupService service.
// All implementations must embed UnimplementedAssetSubGroupServiceServer
// for forward compatibility
type AssetSubGroupServiceServer interface {
	Create(context.Context, *CreateUpdateRequest) (*CUDResponse, error)
	Update(context.Context, *CreateUpdateRequest) (*CUDResponse, error)
	Delete(context.Context, *DeleteRequest) (*CUDResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	mustEmbedUnimplementedAssetSubGroupServiceServer()
}

// UnimplementedAssetSubGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetSubGroupServiceServer struct {
}

func (UnimplementedAssetSubGroupServiceServer) Create(context.Context, *CreateUpdateRequest) (*CUDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAssetSubGroupServiceServer) Update(context.Context, *CreateUpdateRequest) (*CUDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAssetSubGroupServiceServer) Delete(context.Context, *DeleteRequest) (*CUDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAssetSubGroupServiceServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedAssetSubGroupServiceServer) mustEmbedUnimplementedAssetSubGroupServiceServer() {}

// UnsafeAssetSubGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetSubGroupServiceServer will
// result in compilation errors.
type UnsafeAssetSubGroupServiceServer interface {
	mustEmbedUnimplementedAssetSubGroupServiceServer()
}

func RegisterAssetSubGroupServiceServer(s grpc.ServiceRegistrar, srv AssetSubGroupServiceServer) {
	s.RegisterService(&AssetSubGroupService_ServiceDesc, srv)
}

func _AssetSubGroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetSubGroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetSubGroupService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetSubGroupServiceServer).Create(ctx, req.(*CreateUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetSubGroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetSubGroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetSubGroupService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetSubGroupServiceServer).Update(ctx, req.(*CreateUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetSubGroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetSubGroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetSubGroupService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetSubGroupServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetSubGroupService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetSubGroupServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetSubGroupService_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetSubGroupServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetSubGroupService_ServiceDesc is the grpc.ServiceDesc for AssetSubGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetSubGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "asset_sub_group.AssetSubGroupService",
	HandlerType: (*AssetSubGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AssetSubGroupService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AssetSubGroupService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AssetSubGroupService_Delete_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _AssetSubGroupService_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/asset_sub_group/pb/asset_sub_group.proto",
}
