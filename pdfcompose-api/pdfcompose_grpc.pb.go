// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pdfcompose_api

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

// PdfComposeServiceClient is the client API for PdfComposeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PdfComposeServiceClient interface {
	CreatePdf(ctx context.Context, opts ...grpc.CallOption) (PdfComposeService_CreatePdfClient, error)
}

type pdfComposeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPdfComposeServiceClient(cc grpc.ClientConnInterface) PdfComposeServiceClient {
	return &pdfComposeServiceClient{cc}
}

func (c *pdfComposeServiceClient) CreatePdf(ctx context.Context, opts ...grpc.CallOption) (PdfComposeService_CreatePdfClient, error) {
	stream, err := c.cc.NewStream(ctx, &PdfComposeService_ServiceDesc.Streams[0], "/pdfcompose.PdfComposeService/CreatePdf", opts...)
	if err != nil {
		return nil, err
	}
	x := &pdfComposeServiceCreatePdfClient{stream}
	return x, nil
}

type PdfComposeService_CreatePdfClient interface {
	Send(*CreatePdfRequest) error
	Recv() (*CreatePdfResponse, error)
	grpc.ClientStream
}

type pdfComposeServiceCreatePdfClient struct {
	grpc.ClientStream
}

func (x *pdfComposeServiceCreatePdfClient) Send(m *CreatePdfRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pdfComposeServiceCreatePdfClient) Recv() (*CreatePdfResponse, error) {
	m := new(CreatePdfResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PdfComposeServiceServer is the server API for PdfComposeService service.
// All implementations must embed UnimplementedPdfComposeServiceServer
// for forward compatibility
type PdfComposeServiceServer interface {
	CreatePdf(PdfComposeService_CreatePdfServer) error
	mustEmbedUnimplementedPdfComposeServiceServer()
}

// UnimplementedPdfComposeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPdfComposeServiceServer struct {
}

func (UnimplementedPdfComposeServiceServer) CreatePdf(PdfComposeService_CreatePdfServer) error {
	return status.Errorf(codes.Unimplemented, "method CreatePdf not implemented")
}
func (UnimplementedPdfComposeServiceServer) mustEmbedUnimplementedPdfComposeServiceServer() {}

// UnsafePdfComposeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PdfComposeServiceServer will
// result in compilation errors.
type UnsafePdfComposeServiceServer interface {
	mustEmbedUnimplementedPdfComposeServiceServer()
}

func RegisterPdfComposeServiceServer(s grpc.ServiceRegistrar, srv PdfComposeServiceServer) {
	s.RegisterService(&PdfComposeService_ServiceDesc, srv)
}

func _PdfComposeService_CreatePdf_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PdfComposeServiceServer).CreatePdf(&pdfComposeServiceCreatePdfServer{stream})
}

type PdfComposeService_CreatePdfServer interface {
	Send(*CreatePdfResponse) error
	Recv() (*CreatePdfRequest, error)
	grpc.ServerStream
}

type pdfComposeServiceCreatePdfServer struct {
	grpc.ServerStream
}

func (x *pdfComposeServiceCreatePdfServer) Send(m *CreatePdfResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pdfComposeServiceCreatePdfServer) Recv() (*CreatePdfRequest, error) {
	m := new(CreatePdfRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PdfComposeService_ServiceDesc is the grpc.ServiceDesc for PdfComposeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PdfComposeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pdfcompose.PdfComposeService",
	HandlerType: (*PdfComposeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreatePdf",
			Handler:       _PdfComposeService_CreatePdf_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pdfcompose.proto",
}
