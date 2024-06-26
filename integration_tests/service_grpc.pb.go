// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

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

// CounterServiceClient is the client API for CounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CounterServiceClient interface {
	Increment(ctx context.Context, in *UnaryRequest, opts ...grpc.CallOption) (*UnaryResponse, error)
	StreamingIncrements(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (CounterService_StreamingIncrementsClient, error)
	FailingIncrement(ctx context.Context, in *UnaryRequest, opts ...grpc.CallOption) (*UnaryResponse, error)
	EchoBinary(ctx context.Context, in *BinaryRequest, opts ...grpc.CallOption) (*BinaryResponse, error)
	HTTPGet(ctx context.Context, in *HttpGetRequest, opts ...grpc.CallOption) (*HttpGetResponse, error)
	HTTPPostWithNestedBodyPath(ctx context.Context, in *HttpPostRequest, opts ...grpc.CallOption) (*HttpPostResponse, error)
	HTTPPostWithStarBodyPath(ctx context.Context, in *HttpPostRequest, opts ...grpc.CallOption) (*HttpPostResponse, error)
	HTTPPatch(ctx context.Context, in *HttpPatchRequest, opts ...grpc.CallOption) (*HttpPatchResponse, error)
	HTTPDelete(ctx context.Context, in *HttpDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ExternalMessage(ctx context.Context, in *ExternalRequest, opts ...grpc.CallOption) (*ExternalResponse, error)
	HTTPGetWithURLSearchParams(ctx context.Context, in *HTTPGetWithURLSearchParamsRequest, opts ...grpc.CallOption) (*HTTPGetWithURLSearchParamsResponse, error)
	HTTPGetWithZeroValueURLSearchParams(ctx context.Context, in *HTTPGetWithZeroValueURLSearchParamsRequest, opts ...grpc.CallOption) (*HTTPGetWithZeroValueURLSearchParamsResponse, error)
	HTTPGetWithPathSegments(ctx context.Context, in *HTTPGetWithPathSegmentsRequest, opts ...grpc.CallOption) (*HTTPGetWithPathSegmentsResponse, error)
	HTTPPostWithFieldPath(ctx context.Context, in *HTTPPostWithFieldPathRequest, opts ...grpc.CallOption) (*HTTPPostWithFieldPathResponse, error)
	HTTPPostWithFieldPathAndSegments(ctx context.Context, in *HTTPPostWithFieldPathRequest, opts ...grpc.CallOption) (*HTTPPostWithFieldPathResponse, error)
}

type counterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterServiceClient(cc grpc.ClientConnInterface) CounterServiceClient {
	return &counterServiceClient{cc}
}

func (c *counterServiceClient) Increment(ctx context.Context, in *UnaryRequest, opts ...grpc.CallOption) (*UnaryResponse, error) {
	out := new(UnaryResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/Increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) StreamingIncrements(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (CounterService_StreamingIncrementsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CounterService_ServiceDesc.Streams[0], "/main.CounterService/StreamingIncrements", opts...)
	if err != nil {
		return nil, err
	}
	x := &counterServiceStreamingIncrementsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CounterService_StreamingIncrementsClient interface {
	Recv() (*StreamingResponse, error)
	grpc.ClientStream
}

type counterServiceStreamingIncrementsClient struct {
	grpc.ClientStream
}

func (x *counterServiceStreamingIncrementsClient) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *counterServiceClient) FailingIncrement(ctx context.Context, in *UnaryRequest, opts ...grpc.CallOption) (*UnaryResponse, error) {
	out := new(UnaryResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/FailingIncrement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) EchoBinary(ctx context.Context, in *BinaryRequest, opts ...grpc.CallOption) (*BinaryResponse, error) {
	out := new(BinaryResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/EchoBinary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) HTTPGet(ctx context.Context, in *HttpGetRequest, opts ...grpc.CallOption) (*HttpGetResponse, error) {
	out := new(HttpGetResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/HTTPGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) HTTPPostWithNestedBodyPath(ctx context.Context, in *HttpPostRequest, opts ...grpc.CallOption) (*HttpPostResponse, error) {
	out := new(HttpPostResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/HTTPPostWithNestedBodyPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) HTTPPostWithStarBodyPath(ctx context.Context, in *HttpPostRequest, opts ...grpc.CallOption) (*HttpPostResponse, error) {
	out := new(HttpPostResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/HTTPPostWithStarBodyPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) HTTPPatch(ctx context.Context, in *HttpPatchRequest, opts ...grpc.CallOption) (*HttpPatchResponse, error) {
	out := new(HttpPatchResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/HTTPPatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) HTTPDelete(ctx context.Context, in *HttpDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/main.CounterService/HTTPDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) ExternalMessage(ctx context.Context, in *ExternalRequest, opts ...grpc.CallOption) (*ExternalResponse, error) {
	out := new(ExternalResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/ExternalMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) HTTPGetWithURLSearchParams(ctx context.Context, in *HTTPGetWithURLSearchParamsRequest, opts ...grpc.CallOption) (*HTTPGetWithURLSearchParamsResponse, error) {
	out := new(HTTPGetWithURLSearchParamsResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/HTTPGetWithURLSearchParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) HTTPGetWithZeroValueURLSearchParams(ctx context.Context, in *HTTPGetWithZeroValueURLSearchParamsRequest, opts ...grpc.CallOption) (*HTTPGetWithZeroValueURLSearchParamsResponse, error) {
	out := new(HTTPGetWithZeroValueURLSearchParamsResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/HTTPGetWithZeroValueURLSearchParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) HTTPGetWithPathSegments(ctx context.Context, in *HTTPGetWithPathSegmentsRequest, opts ...grpc.CallOption) (*HTTPGetWithPathSegmentsResponse, error) {
	out := new(HTTPGetWithPathSegmentsResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/HTTPGetWithPathSegments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) HTTPPostWithFieldPath(ctx context.Context, in *HTTPPostWithFieldPathRequest, opts ...grpc.CallOption) (*HTTPPostWithFieldPathResponse, error) {
	out := new(HTTPPostWithFieldPathResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/HTTPPostWithFieldPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) HTTPPostWithFieldPathAndSegments(ctx context.Context, in *HTTPPostWithFieldPathRequest, opts ...grpc.CallOption) (*HTTPPostWithFieldPathResponse, error) {
	out := new(HTTPPostWithFieldPathResponse)
	err := c.cc.Invoke(ctx, "/main.CounterService/HTTPPostWithFieldPathAndSegments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServiceServer is the server API for CounterService service.
// All implementations must embed UnimplementedCounterServiceServer
// for forward compatibility
type CounterServiceServer interface {
	Increment(context.Context, *UnaryRequest) (*UnaryResponse, error)
	StreamingIncrements(*StreamingRequest, CounterService_StreamingIncrementsServer) error
	FailingIncrement(context.Context, *UnaryRequest) (*UnaryResponse, error)
	EchoBinary(context.Context, *BinaryRequest) (*BinaryResponse, error)
	HTTPGet(context.Context, *HttpGetRequest) (*HttpGetResponse, error)
	HTTPPostWithNestedBodyPath(context.Context, *HttpPostRequest) (*HttpPostResponse, error)
	HTTPPostWithStarBodyPath(context.Context, *HttpPostRequest) (*HttpPostResponse, error)
	HTTPPatch(context.Context, *HttpPatchRequest) (*HttpPatchResponse, error)
	HTTPDelete(context.Context, *HttpDeleteRequest) (*emptypb.Empty, error)
	ExternalMessage(context.Context, *ExternalRequest) (*ExternalResponse, error)
	HTTPGetWithURLSearchParams(context.Context, *HTTPGetWithURLSearchParamsRequest) (*HTTPGetWithURLSearchParamsResponse, error)
	HTTPGetWithZeroValueURLSearchParams(context.Context, *HTTPGetWithZeroValueURLSearchParamsRequest) (*HTTPGetWithZeroValueURLSearchParamsResponse, error)
	HTTPGetWithPathSegments(context.Context, *HTTPGetWithPathSegmentsRequest) (*HTTPGetWithPathSegmentsResponse, error)
	HTTPPostWithFieldPath(context.Context, *HTTPPostWithFieldPathRequest) (*HTTPPostWithFieldPathResponse, error)
	HTTPPostWithFieldPathAndSegments(context.Context, *HTTPPostWithFieldPathRequest) (*HTTPPostWithFieldPathResponse, error)
	mustEmbedUnimplementedCounterServiceServer()
}

// UnimplementedCounterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCounterServiceServer struct {
}

func (UnimplementedCounterServiceServer) Increment(context.Context, *UnaryRequest) (*UnaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (UnimplementedCounterServiceServer) StreamingIncrements(*StreamingRequest, CounterService_StreamingIncrementsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingIncrements not implemented")
}
func (UnimplementedCounterServiceServer) FailingIncrement(context.Context, *UnaryRequest) (*UnaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FailingIncrement not implemented")
}
func (UnimplementedCounterServiceServer) EchoBinary(context.Context, *BinaryRequest) (*BinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoBinary not implemented")
}
func (UnimplementedCounterServiceServer) HTTPGet(context.Context, *HttpGetRequest) (*HttpGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTTPGet not implemented")
}
func (UnimplementedCounterServiceServer) HTTPPostWithNestedBodyPath(context.Context, *HttpPostRequest) (*HttpPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTTPPostWithNestedBodyPath not implemented")
}
func (UnimplementedCounterServiceServer) HTTPPostWithStarBodyPath(context.Context, *HttpPostRequest) (*HttpPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTTPPostWithStarBodyPath not implemented")
}
func (UnimplementedCounterServiceServer) HTTPPatch(context.Context, *HttpPatchRequest) (*HttpPatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTTPPatch not implemented")
}
func (UnimplementedCounterServiceServer) HTTPDelete(context.Context, *HttpDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTTPDelete not implemented")
}
func (UnimplementedCounterServiceServer) ExternalMessage(context.Context, *ExternalRequest) (*ExternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExternalMessage not implemented")
}
func (UnimplementedCounterServiceServer) HTTPGetWithURLSearchParams(context.Context, *HTTPGetWithURLSearchParamsRequest) (*HTTPGetWithURLSearchParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTTPGetWithURLSearchParams not implemented")
}
func (UnimplementedCounterServiceServer) HTTPGetWithZeroValueURLSearchParams(context.Context, *HTTPGetWithZeroValueURLSearchParamsRequest) (*HTTPGetWithZeroValueURLSearchParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTTPGetWithZeroValueURLSearchParams not implemented")
}
func (UnimplementedCounterServiceServer) HTTPGetWithPathSegments(context.Context, *HTTPGetWithPathSegmentsRequest) (*HTTPGetWithPathSegmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTTPGetWithPathSegments not implemented")
}
func (UnimplementedCounterServiceServer) HTTPPostWithFieldPath(context.Context, *HTTPPostWithFieldPathRequest) (*HTTPPostWithFieldPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTTPPostWithFieldPath not implemented")
}
func (UnimplementedCounterServiceServer) HTTPPostWithFieldPathAndSegments(context.Context, *HTTPPostWithFieldPathRequest) (*HTTPPostWithFieldPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTTPPostWithFieldPathAndSegments not implemented")
}
func (UnimplementedCounterServiceServer) mustEmbedUnimplementedCounterServiceServer() {}

// UnsafeCounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CounterServiceServer will
// result in compilation errors.
type UnsafeCounterServiceServer interface {
	mustEmbedUnimplementedCounterServiceServer()
}

func RegisterCounterServiceServer(s grpc.ServiceRegistrar, srv CounterServiceServer) {
	s.RegisterService(&CounterService_ServiceDesc, srv)
}

func _CounterService_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/Increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).Increment(ctx, req.(*UnaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_StreamingIncrements_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CounterServiceServer).StreamingIncrements(m, &counterServiceStreamingIncrementsServer{stream})
}

type CounterService_StreamingIncrementsServer interface {
	Send(*StreamingResponse) error
	grpc.ServerStream
}

type counterServiceStreamingIncrementsServer struct {
	grpc.ServerStream
}

func (x *counterServiceStreamingIncrementsServer) Send(m *StreamingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CounterService_FailingIncrement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).FailingIncrement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/FailingIncrement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).FailingIncrement(ctx, req.(*UnaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_EchoBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).EchoBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/EchoBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).EchoBinary(ctx, req.(*BinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_HTTPGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).HTTPGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/HTTPGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).HTTPGet(ctx, req.(*HttpGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_HTTPPostWithNestedBodyPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).HTTPPostWithNestedBodyPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/HTTPPostWithNestedBodyPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).HTTPPostWithNestedBodyPath(ctx, req.(*HttpPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_HTTPPostWithStarBodyPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).HTTPPostWithStarBodyPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/HTTPPostWithStarBodyPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).HTTPPostWithStarBodyPath(ctx, req.(*HttpPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_HTTPPatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).HTTPPatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/HTTPPatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).HTTPPatch(ctx, req.(*HttpPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_HTTPDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).HTTPDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/HTTPDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).HTTPDelete(ctx, req.(*HttpDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_ExternalMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).ExternalMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/ExternalMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).ExternalMessage(ctx, req.(*ExternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_HTTPGetWithURLSearchParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPGetWithURLSearchParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).HTTPGetWithURLSearchParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/HTTPGetWithURLSearchParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).HTTPGetWithURLSearchParams(ctx, req.(*HTTPGetWithURLSearchParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_HTTPGetWithZeroValueURLSearchParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPGetWithZeroValueURLSearchParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).HTTPGetWithZeroValueURLSearchParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/HTTPGetWithZeroValueURLSearchParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).HTTPGetWithZeroValueURLSearchParams(ctx, req.(*HTTPGetWithZeroValueURLSearchParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_HTTPGetWithPathSegments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPGetWithPathSegmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).HTTPGetWithPathSegments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/HTTPGetWithPathSegments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).HTTPGetWithPathSegments(ctx, req.(*HTTPGetWithPathSegmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_HTTPPostWithFieldPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPPostWithFieldPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).HTTPPostWithFieldPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/HTTPPostWithFieldPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).HTTPPostWithFieldPath(ctx, req.(*HTTPPostWithFieldPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_HTTPPostWithFieldPathAndSegments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPPostWithFieldPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).HTTPPostWithFieldPathAndSegments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CounterService/HTTPPostWithFieldPathAndSegments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).HTTPPostWithFieldPathAndSegments(ctx, req.(*HTTPPostWithFieldPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CounterService_ServiceDesc is the grpc.ServiceDesc for CounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.CounterService",
	HandlerType: (*CounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increment",
			Handler:    _CounterService_Increment_Handler,
		},
		{
			MethodName: "FailingIncrement",
			Handler:    _CounterService_FailingIncrement_Handler,
		},
		{
			MethodName: "EchoBinary",
			Handler:    _CounterService_EchoBinary_Handler,
		},
		{
			MethodName: "HTTPGet",
			Handler:    _CounterService_HTTPGet_Handler,
		},
		{
			MethodName: "HTTPPostWithNestedBodyPath",
			Handler:    _CounterService_HTTPPostWithNestedBodyPath_Handler,
		},
		{
			MethodName: "HTTPPostWithStarBodyPath",
			Handler:    _CounterService_HTTPPostWithStarBodyPath_Handler,
		},
		{
			MethodName: "HTTPPatch",
			Handler:    _CounterService_HTTPPatch_Handler,
		},
		{
			MethodName: "HTTPDelete",
			Handler:    _CounterService_HTTPDelete_Handler,
		},
		{
			MethodName: "ExternalMessage",
			Handler:    _CounterService_ExternalMessage_Handler,
		},
		{
			MethodName: "HTTPGetWithURLSearchParams",
			Handler:    _CounterService_HTTPGetWithURLSearchParams_Handler,
		},
		{
			MethodName: "HTTPGetWithZeroValueURLSearchParams",
			Handler:    _CounterService_HTTPGetWithZeroValueURLSearchParams_Handler,
		},
		{
			MethodName: "HTTPGetWithPathSegments",
			Handler:    _CounterService_HTTPGetWithPathSegments_Handler,
		},
		{
			MethodName: "HTTPPostWithFieldPath",
			Handler:    _CounterService_HTTPPostWithFieldPath_Handler,
		},
		{
			MethodName: "HTTPPostWithFieldPathAndSegments",
			Handler:    _CounterService_HTTPPostWithFieldPathAndSegments_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingIncrements",
			Handler:       _CounterService_StreamingIncrements_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
