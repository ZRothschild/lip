// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: company/company.proto

package company

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

type CompanyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId uint64 `protobuf:"varint,1,opt,name=CompanyId,proto3" json:"CompanyId,omitempty"`
}

func (x *CompanyReq) Reset() {
	*x = CompanyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyReq) ProtoMessage() {}

func (x *CompanyReq) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyReq.ProtoReflect.Descriptor instead.
func (*CompanyReq) Descriptor() ([]byte, []int) {
	return file_company_company_proto_rawDescGZIP(), []int{0}
}

func (x *CompanyReq) GetCompanyId() uint64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type CompanyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message   string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Detail    *CompanyDetail `protobuf:"bytes,3,opt,name=Detail,proto3" json:"Detail,omitempty"`
	RequestId string         `protobuf:"bytes,4,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *CompanyResp) Reset() {
	*x = CompanyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyResp) ProtoMessage() {}

func (x *CompanyResp) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyResp.ProtoReflect.Descriptor instead.
func (*CompanyResp) Descriptor() ([]byte, []int) {
	return file_company_company_proto_rawDescGZIP(), []int{1}
}

func (x *CompanyResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CompanyResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CompanyResp) GetDetail() *CompanyDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *CompanyResp) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type CompanyDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname    string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Username    string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Mobile      string `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Phone       string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	IdCardFront string `protobuf:"bytes,7,opt,name=idCardFront,proto3" json:"idCardFront,omitempty"`
	IdCardBack  string `protobuf:"bytes,8,opt,name=idCardBack,proto3" json:"idCardBack,omitempty"`
	Seniority   uint64 `protobuf:"varint,9,opt,name=seniority,proto3" json:"seniority,omitempty"`
	Profile     string `protobuf:"bytes,10,opt,name=profile,proto3" json:"profile,omitempty"`
	Location    string `protobuf:"bytes,11,opt,name=location,proto3" json:"location,omitempty"`
	CompanyName string `protobuf:"bytes,12,opt,name=companyName,proto3" json:"companyName,omitempty"`
	Avatar      string `protobuf:"bytes,13,opt,name=avatar,proto3" json:"avatar,omitempty"`
	License     string `protobuf:"bytes,14,opt,name=license,proto3" json:"license,omitempty"`
	UserId      uint64 `protobuf:"varint,15,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CompanyDetail) Reset() {
	*x = CompanyDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyDetail) ProtoMessage() {}

func (x *CompanyDetail) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyDetail.ProtoReflect.Descriptor instead.
func (*CompanyDetail) Descriptor() ([]byte, []int) {
	return file_company_company_proto_rawDescGZIP(), []int{2}
}

func (x *CompanyDetail) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *CompanyDetail) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CompanyDetail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CompanyDetail) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *CompanyDetail) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CompanyDetail) GetIdCardFront() string {
	if x != nil {
		return x.IdCardFront
	}
	return ""
}

func (x *CompanyDetail) GetIdCardBack() string {
	if x != nil {
		return x.IdCardBack
	}
	return ""
}

func (x *CompanyDetail) GetSeniority() uint64 {
	if x != nil {
		return x.Seniority
	}
	return 0
}

func (x *CompanyDetail) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *CompanyDetail) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CompanyDetail) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *CompanyDetail) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *CompanyDetail) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

func (x *CompanyDetail) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_company_company_proto protoreflect.FileDescriptor

var file_company_company_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x8d, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x43, 0x61, 0x72,
	0x64, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64,
	0x42, 0x61, 0x63, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x7b, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22,
	0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x64, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_company_company_proto_rawDescOnce sync.Once
	file_company_company_proto_rawDescData = file_company_company_proto_rawDesc
)

func file_company_company_proto_rawDescGZIP() []byte {
	file_company_company_proto_rawDescOnce.Do(func() {
		file_company_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_company_proto_rawDescData)
	})
	return file_company_company_proto_rawDescData
}

var file_company_company_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_company_company_proto_goTypes = []interface{}{
	(*CompanyReq)(nil),    // 0: company.service.CompanyReq
	(*CompanyResp)(nil),   // 1: company.service.CompanyResp
	(*CompanyDetail)(nil), // 2: company.service.CompanyDetail
}
var file_company_company_proto_depIdxs = []int32{
	2, // 0: company.service.CompanyResp.Detail:type_name -> company.service.CompanyDetail
	0, // 1: company.service.CompanyService.CompanyDetail:input_type -> company.service.CompanyReq
	1, // 2: company.service.CompanyService.CompanyDetail:output_type -> company.service.CompanyResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_company_company_proto_init() }
func file_company_company_proto_init() {
	if File_company_company_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_company_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyReq); i {
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
		file_company_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyResp); i {
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
		file_company_company_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyDetail); i {
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
			RawDescriptor: file_company_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_company_company_proto_goTypes,
		DependencyIndexes: file_company_company_proto_depIdxs,
		MessageInfos:      file_company_company_proto_msgTypes,
	}.Build()
	File_company_company_proto = out.File
	file_company_company_proto_rawDesc = nil
	file_company_company_proto_goTypes = nil
	file_company_company_proto_depIdxs = nil
}