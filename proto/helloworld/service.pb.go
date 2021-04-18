// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: helloworld/service.proto

package helloworld

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrganisationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc    string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Contact string `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
	Image   string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Video   string `protobuf:"bytes,6,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *OrganisationRequest) Reset() {
	*x = OrganisationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganisationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganisationRequest) ProtoMessage() {}

func (x *OrganisationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganisationRequest.ProtoReflect.Descriptor instead.
func (*OrganisationRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_service_proto_rawDescGZIP(), []int{0}
}

func (x *OrganisationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrganisationRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *OrganisationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OrganisationRequest) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *OrganisationRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *OrganisationRequest) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

// The request message containing the user's name
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_service_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_helloworld_service_proto_rawDescGZIP(), []int{2}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_helloworld_service_proto protoreflect.FileDescriptor

var file_helloworld_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x99, 0x01, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x22, 0x0a, 0x0c,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4e, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x0d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x01, 0x2a, 0x32, 0x69, 0x0a, 0x16, 0x48, 0x74, 0x74, 0x70,
	0x42, 0x6f, 0x64, 0x79, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x13,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x32, 0x6d, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x5d, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x13, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x16, 0x5a, 0x14, 0x79, 0x75, 0x74, 0x68, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_helloworld_service_proto_rawDescOnce sync.Once
	file_helloworld_service_proto_rawDescData = file_helloworld_service_proto_rawDesc
)

func file_helloworld_service_proto_rawDescGZIP() []byte {
	file_helloworld_service_proto_rawDescOnce.Do(func() {
		file_helloworld_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_service_proto_rawDescData)
	})
	return file_helloworld_service_proto_rawDescData
}

var file_helloworld_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_helloworld_service_proto_goTypes = []interface{}{
	(*OrganisationRequest)(nil), // 0: organisationRequest
	(*HelloRequest)(nil),        // 1: HelloRequest
	(*HelloReply)(nil),          // 2: HelloReply
	(*emptypb.Empty)(nil),       // 3: google.protobuf.Empty
	(*httpbody.HttpBody)(nil),   // 4: google.api.HttpBody
}
var file_helloworld_service_proto_depIdxs = []int32{
	1, // 0: Greeter.SayHello:input_type -> HelloRequest
	3, // 1: HttpBodyExampleService.HelloWorld:input_type -> google.protobuf.Empty
	0, // 2: Organisation.createOrganisation:input_type -> organisationRequest
	2, // 3: Greeter.SayHello:output_type -> HelloReply
	4, // 4: HttpBodyExampleService.HelloWorld:output_type -> google.api.HttpBody
	4, // 5: Organisation.createOrganisation:output_type -> google.api.HttpBody
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_service_proto_init() }
func file_helloworld_service_proto_init() {
	if File_helloworld_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganisationRequest); i {
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
		file_helloworld_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_helloworld_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_helloworld_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_helloworld_service_proto_goTypes,
		DependencyIndexes: file_helloworld_service_proto_depIdxs,
		MessageInfos:      file_helloworld_service_proto_msgTypes,
	}.Build()
	File_helloworld_service_proto = out.File
	file_helloworld_service_proto_rawDesc = nil
	file_helloworld_service_proto_goTypes = nil
	file_helloworld_service_proto_depIdxs = nil
}
