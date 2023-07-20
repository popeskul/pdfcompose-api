// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: pdfcompose.proto

package pdfcompose_api

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

type CreatePdfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File1 []byte `protobuf:"bytes,1,opt,name=file1,proto3" json:"file1,omitempty"`
	File2 []byte `protobuf:"bytes,2,opt,name=file2,proto3" json:"file2,omitempty"`
	File3 []byte `protobuf:"bytes,3,opt,name=file3,proto3" json:"file3,omitempty"`
}

func (x *CreatePdfRequest) Reset() {
	*x = CreatePdfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pdfcompose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePdfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePdfRequest) ProtoMessage() {}

func (x *CreatePdfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pdfcompose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePdfRequest.ProtoReflect.Descriptor instead.
func (*CreatePdfRequest) Descriptor() ([]byte, []int) {
	return file_pdfcompose_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePdfRequest) GetFile1() []byte {
	if x != nil {
		return x.File1
	}
	return nil
}

func (x *CreatePdfRequest) GetFile2() []byte {
	if x != nil {
		return x.File2
	}
	return nil
}

func (x *CreatePdfRequest) GetFile3() []byte {
	if x != nil {
		return x.File3
	}
	return nil
}

type CreatePdfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pdf []byte `protobuf:"bytes,1,opt,name=pdf,proto3" json:"pdf,omitempty"`
}

func (x *CreatePdfResponse) Reset() {
	*x = CreatePdfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pdfcompose_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePdfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePdfResponse) ProtoMessage() {}

func (x *CreatePdfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pdfcompose_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePdfResponse.ProtoReflect.Descriptor instead.
func (*CreatePdfResponse) Descriptor() ([]byte, []int) {
	return file_pdfcompose_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePdfResponse) GetPdf() []byte {
	if x != nil {
		return x.Pdf
	}
	return nil
}

var File_pdfcompose_proto protoreflect.FileDescriptor

var file_pdfcompose_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x54,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x64, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x33, 0x22, 0x25, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x64,
	0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x64, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x64, 0x66, 0x32, 0x5f, 0x0a, 0x11, 0x50,
	0x64, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4a, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x64, 0x66, 0x12, 0x1c, 0x2e,
	0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x64, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x64,
	0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x64, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e,
	0x2f, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x3b,
	0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pdfcompose_proto_rawDescOnce sync.Once
	file_pdfcompose_proto_rawDescData = file_pdfcompose_proto_rawDesc
)

func file_pdfcompose_proto_rawDescGZIP() []byte {
	file_pdfcompose_proto_rawDescOnce.Do(func() {
		file_pdfcompose_proto_rawDescData = protoimpl.X.CompressGZIP(file_pdfcompose_proto_rawDescData)
	})
	return file_pdfcompose_proto_rawDescData
}

var file_pdfcompose_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pdfcompose_proto_goTypes = []interface{}{
	(*CreatePdfRequest)(nil),  // 0: pdfcompose.CreatePdfRequest
	(*CreatePdfResponse)(nil), // 1: pdfcompose.CreatePdfResponse
}
var file_pdfcompose_proto_depIdxs = []int32{
	0, // 0: pdfcompose.PdfComposeService.CreatePdf:input_type -> pdfcompose.CreatePdfRequest
	1, // 1: pdfcompose.PdfComposeService.CreatePdf:output_type -> pdfcompose.CreatePdfResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pdfcompose_proto_init() }
func file_pdfcompose_proto_init() {
	if File_pdfcompose_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pdfcompose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePdfRequest); i {
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
		file_pdfcompose_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePdfResponse); i {
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
			RawDescriptor: file_pdfcompose_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pdfcompose_proto_goTypes,
		DependencyIndexes: file_pdfcompose_proto_depIdxs,
		MessageInfos:      file_pdfcompose_proto_msgTypes,
	}.Build()
	File_pdfcompose_proto = out.File
	file_pdfcompose_proto_rawDesc = nil
	file_pdfcompose_proto_goTypes = nil
	file_pdfcompose_proto_depIdxs = nil
}
