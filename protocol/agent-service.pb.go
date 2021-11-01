// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: agent-service.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnrollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	CsrPEM string   `protobuf:"bytes,2,opt,name=csrPEM,proto3" json:"csrPEM,omitempty"`
	Groups []string `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
	Name   string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *EnrollRequest) Reset() {
	*x = EnrollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollRequest) ProtoMessage() {}

func (x *EnrollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollRequest.ProtoReflect.Descriptor instead.
func (*EnrollRequest) Descriptor() ([]byte, []int) {
	return file_agent_service_proto_rawDescGZIP(), []int{0}
}

func (x *EnrollRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *EnrollRequest) GetCsrPEM() string {
	if x != nil {
		return x.CsrPEM
	}
	return ""
}

func (x *EnrollRequest) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *EnrollRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EnrollResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnrollResponse) Reset() {
	*x = EnrollResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollResponse) ProtoMessage() {}

func (x *EnrollResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollResponse.ProtoReflect.Descriptor instead.
func (*EnrollResponse) Descriptor() ([]byte, []int) {
	return file_agent_service_proto_rawDescGZIP(), []int{1}
}

type GetEnrollStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsEnrolled             bool                    `protobuf:"varint,1,opt,name=isEnrolled,proto3" json:"isEnrolled,omitempty"`
	IsEnrollmentRequested  bool                    `protobuf:"varint,2,opt,name=isEnrollmentRequested,proto3" json:"isEnrollmentRequested,omitempty"`
	SignedPEM              string                  `protobuf:"bytes,10,opt,name=signedPEM,proto3" json:"signedPEM,omitempty"`
	IssuedAt               *timestamppb.Timestamp  `protobuf:"bytes,11,opt,name=issuedAt,proto3" json:"issuedAt,omitempty"`
	ExpiresAt              *timestamppb.Timestamp  `protobuf:"bytes,12,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
	CertificateAuthorities []*CertificateAuthority `protobuf:"bytes,20,rep,name=certificateAuthorities,proto3" json:"certificateAuthorities,omitempty"`
}

func (x *GetEnrollStatusResponse) Reset() {
	*x = GetEnrollStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnrollStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnrollStatusResponse) ProtoMessage() {}

func (x *GetEnrollStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnrollStatusResponse.ProtoReflect.Descriptor instead.
func (*GetEnrollStatusResponse) Descriptor() ([]byte, []int) {
	return file_agent_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetEnrollStatusResponse) GetIsEnrolled() bool {
	if x != nil {
		return x.IsEnrolled
	}
	return false
}

func (x *GetEnrollStatusResponse) GetIsEnrollmentRequested() bool {
	if x != nil {
		return x.IsEnrollmentRequested
	}
	return false
}

func (x *GetEnrollStatusResponse) GetSignedPEM() string {
	if x != nil {
		return x.SignedPEM
	}
	return ""
}

func (x *GetEnrollStatusResponse) GetIssuedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IssuedAt
	}
	return nil
}

func (x *GetEnrollStatusResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *GetEnrollStatusResponse) GetCertificateAuthorities() []*CertificateAuthority {
	if x != nil {
		return x.CertificateAuthorities
	}
	return nil
}

type GetCertificateAuthorityByNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkNames []string `protobuf:"bytes,1,rep,name=networkNames,proto3" json:"networkNames,omitempty"`
}

func (x *GetCertificateAuthorityByNetworkRequest) Reset() {
	*x = GetCertificateAuthorityByNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertificateAuthorityByNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertificateAuthorityByNetworkRequest) ProtoMessage() {}

func (x *GetCertificateAuthorityByNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertificateAuthorityByNetworkRequest.ProtoReflect.Descriptor instead.
func (*GetCertificateAuthorityByNetworkRequest) Descriptor() ([]byte, []int) {
	return file_agent_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCertificateAuthorityByNetworkRequest) GetNetworkNames() []string {
	if x != nil {
		return x.NetworkNames
	}
	return nil
}

type GetCertificateAuthorityByNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertificateAuthorities []*CertificateAuthority `protobuf:"bytes,1,rep,name=certificateAuthorities,proto3" json:"certificateAuthorities,omitempty"`
}

func (x *GetCertificateAuthorityByNetworkResponse) Reset() {
	*x = GetCertificateAuthorityByNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertificateAuthorityByNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertificateAuthorityByNetworkResponse) ProtoMessage() {}

func (x *GetCertificateAuthorityByNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertificateAuthorityByNetworkResponse.ProtoReflect.Descriptor instead.
func (*GetCertificateAuthorityByNetworkResponse) Descriptor() ([]byte, []int) {
	return file_agent_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetCertificateAuthorityByNetworkResponse) GetCertificateAuthorities() []*CertificateAuthority {
	if x != nil {
		return x.CertificateAuthorities
	}
	return nil
}

var File_agent_service_proto protoreflect.FileDescriptor

var file_agent_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a,
	0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0d, 0x45,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x73, 0x72, 0x50, 0x45, 0x4d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x73, 0x72, 0x50, 0x45, 0x4d, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd7, 0x02, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x45, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x15, 0x69, 0x73, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x73, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x50, 0x45, 0x4d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x45, 0x4d, 0x12, 0x36, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x56, 0x0a, 0x16, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x16, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x22, 0x4d, 0x0a, 0x27, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42, 0x79, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x22, 0x82, 0x01, 0x0a, 0x28, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42, 0x79, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56,
	0x0a, 0x16, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x16,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x32, 0xab, 0x02, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x45, 0x6e, 0x72, 0x6f, 0x6c,
	0x6c, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x72,
	0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72,
	0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x42, 0x79, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42, 0x79,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x42, 0x79, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6c, 0x79, 0x6e, 0x67, 0x64, 0x6b, 0x2f, 0x6e, 0x65, 0x62, 0x75, 0x6c,
	0x61, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_service_proto_rawDescOnce sync.Once
	file_agent_service_proto_rawDescData = file_agent_service_proto_rawDesc
)

func file_agent_service_proto_rawDescGZIP() []byte {
	file_agent_service_proto_rawDescOnce.Do(func() {
		file_agent_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_service_proto_rawDescData)
	})
	return file_agent_service_proto_rawDescData
}

var file_agent_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_agent_service_proto_goTypes = []interface{}{
	(*EnrollRequest)(nil),                            // 0: protocol.EnrollRequest
	(*EnrollResponse)(nil),                           // 1: protocol.EnrollResponse
	(*GetEnrollStatusResponse)(nil),                  // 2: protocol.GetEnrollStatusResponse
	(*GetCertificateAuthorityByNetworkRequest)(nil),  // 3: protocol.GetCertificateAuthorityByNetworkRequest
	(*GetCertificateAuthorityByNetworkResponse)(nil), // 4: protocol.GetCertificateAuthorityByNetworkResponse
	(*timestamppb.Timestamp)(nil),                    // 5: google.protobuf.Timestamp
	(*CertificateAuthority)(nil),                     // 6: protocol.CertificateAuthority
	(*emptypb.Empty)(nil),                            // 7: google.protobuf.Empty
}
var file_agent_service_proto_depIdxs = []int32{
	5, // 0: protocol.GetEnrollStatusResponse.issuedAt:type_name -> google.protobuf.Timestamp
	5, // 1: protocol.GetEnrollStatusResponse.expiresAt:type_name -> google.protobuf.Timestamp
	6, // 2: protocol.GetEnrollStatusResponse.certificateAuthorities:type_name -> protocol.CertificateAuthority
	6, // 3: protocol.GetCertificateAuthorityByNetworkResponse.certificateAuthorities:type_name -> protocol.CertificateAuthority
	0, // 4: protocol.AgentService.Enroll:input_type -> protocol.EnrollRequest
	7, // 5: protocol.AgentService.GetEnrollStatus:input_type -> google.protobuf.Empty
	3, // 6: protocol.AgentService.GetCertificateAuthorityByNetwork:input_type -> protocol.GetCertificateAuthorityByNetworkRequest
	1, // 7: protocol.AgentService.Enroll:output_type -> protocol.EnrollResponse
	2, // 8: protocol.AgentService.GetEnrollStatus:output_type -> protocol.GetEnrollStatusResponse
	4, // 9: protocol.AgentService.GetCertificateAuthorityByNetwork:output_type -> protocol.GetCertificateAuthorityByNetworkResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_agent_service_proto_init() }
func file_agent_service_proto_init() {
	if File_agent_service_proto != nil {
		return
	}
	file_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_agent_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollRequest); i {
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
		file_agent_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollResponse); i {
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
		file_agent_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnrollStatusResponse); i {
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
		file_agent_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCertificateAuthorityByNetworkRequest); i {
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
		file_agent_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCertificateAuthorityByNetworkResponse); i {
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
			RawDescriptor: file_agent_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_service_proto_goTypes,
		DependencyIndexes: file_agent_service_proto_depIdxs,
		MessageInfos:      file_agent_service_proto_msgTypes,
	}.Build()
	File_agent_service_proto = out.File
	file_agent_service_proto_rawDesc = nil
	file_agent_service_proto_goTypes = nil
	file_agent_service_proto_depIdxs = nil
}