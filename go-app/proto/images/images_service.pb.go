// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: images/images_service.proto

package __

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ResizeImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageData []byte `protobuf:"bytes,1,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
	Width     int32  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height    int32  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *ResizeImageRequest) Reset() {
	*x = ResizeImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_images_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResizeImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResizeImageRequest) ProtoMessage() {}

func (x *ResizeImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_images_images_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResizeImageRequest.ProtoReflect.Descriptor instead.
func (*ResizeImageRequest) Descriptor() ([]byte, []int) {
	return file_images_images_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResizeImageRequest) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

func (x *ResizeImageRequest) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ResizeImageRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type ResizeImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResizedImageData []byte `protobuf:"bytes,1,opt,name=resized_image_data,json=resizedImageData,proto3" json:"resized_image_data,omitempty"`
}

func (x *ResizeImageResponse) Reset() {
	*x = ResizeImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_images_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResizeImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResizeImageResponse) ProtoMessage() {}

func (x *ResizeImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_images_images_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResizeImageResponse.ProtoReflect.Descriptor instead.
func (*ResizeImageResponse) Descriptor() ([]byte, []int) {
	return file_images_images_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResizeImageResponse) GetResizedImageData() []byte {
	if x != nil {
		return x.ResizedImageData
	}
	return nil
}

type ConvertToGrayscaleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageData []byte `protobuf:"bytes,1,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
}

func (x *ConvertToGrayscaleRequest) Reset() {
	*x = ConvertToGrayscaleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_images_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertToGrayscaleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertToGrayscaleRequest) ProtoMessage() {}

func (x *ConvertToGrayscaleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_images_images_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertToGrayscaleRequest.ProtoReflect.Descriptor instead.
func (*ConvertToGrayscaleRequest) Descriptor() ([]byte, []int) {
	return file_images_images_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConvertToGrayscaleRequest) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

type ConvertToGrayscaleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrayscaleImageData []byte `protobuf:"bytes,1,opt,name=grayscale_image_data,json=grayscaleImageData,proto3" json:"grayscale_image_data,omitempty"`
}

func (x *ConvertToGrayscaleResponse) Reset() {
	*x = ConvertToGrayscaleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_images_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertToGrayscaleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertToGrayscaleResponse) ProtoMessage() {}

func (x *ConvertToGrayscaleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_images_images_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertToGrayscaleResponse.ProtoReflect.Descriptor instead.
func (*ConvertToGrayscaleResponse) Descriptor() ([]byte, []int) {
	return file_images_images_service_proto_rawDescGZIP(), []int{3}
}

func (x *ConvertToGrayscaleResponse) GetGrayscaleImageData() []byte {
	if x != nil {
		return x.GrayscaleImageData
	}
	return nil
}

var File_images_images_service_proto protoreflect.FileDescriptor

var file_images_images_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x43, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x12, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x72, 0x65, 0x73, 0x69, 0x7a,
	0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x19, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x47, 0x72, 0x61, 0x79, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x54, 0x6f, 0x47, 0x72, 0x61, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x67, 0x72, 0x61, 0x79, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x12, 0x67, 0x72, 0x61, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x32, 0xf0, 0x01, 0x0a, 0x0f, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x64, 0x0a, 0x0b, 0x52,
	0x65, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x77, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x47, 0x72,
	0x61, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x47, 0x72, 0x61, 0x79, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x47, 0x72, 0x61,
	0x79, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x79, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_images_images_service_proto_rawDescOnce sync.Once
	file_images_images_service_proto_rawDescData = file_images_images_service_proto_rawDesc
)

func file_images_images_service_proto_rawDescGZIP() []byte {
	file_images_images_service_proto_rawDescOnce.Do(func() {
		file_images_images_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_images_images_service_proto_rawDescData)
	})
	return file_images_images_service_proto_rawDescData
}

var file_images_images_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_images_images_service_proto_goTypes = []interface{}{
	(*ResizeImageRequest)(nil),         // 0: images.ResizeImageRequest
	(*ResizeImageResponse)(nil),        // 1: images.ResizeImageResponse
	(*ConvertToGrayscaleRequest)(nil),  // 2: images.ConvertToGrayscaleRequest
	(*ConvertToGrayscaleResponse)(nil), // 3: images.ConvertToGrayscaleResponse
}
var file_images_images_service_proto_depIdxs = []int32{
	0, // 0: images.ImageProcessing.ResizeImage:input_type -> images.ResizeImageRequest
	2, // 1: images.ImageProcessing.ConvertToGrayscale:input_type -> images.ConvertToGrayscaleRequest
	1, // 2: images.ImageProcessing.ResizeImage:output_type -> images.ResizeImageResponse
	3, // 3: images.ImageProcessing.ConvertToGrayscale:output_type -> images.ConvertToGrayscaleResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_images_images_service_proto_init() }
func file_images_images_service_proto_init() {
	if File_images_images_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_images_images_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResizeImageRequest); i {
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
		file_images_images_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResizeImageResponse); i {
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
		file_images_images_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertToGrayscaleRequest); i {
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
		file_images_images_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertToGrayscaleResponse); i {
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
			RawDescriptor: file_images_images_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_images_images_service_proto_goTypes,
		DependencyIndexes: file_images_images_service_proto_depIdxs,
		MessageInfos:      file_images_images_service_proto_msgTypes,
	}.Build()
	File_images_images_service_proto = out.File
	file_images_images_service_proto_rawDesc = nil
	file_images_images_service_proto_goTypes = nil
	file_images_images_service_proto_depIdxs = nil
}