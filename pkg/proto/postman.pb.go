// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: postman.proto

package proto

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

type DeliverNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider    string         `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	DeviceToken string         `protobuf:"bytes,2,opt,name=device_token,json=deviceToken,proto3" json:"device_token,omitempty"`
	Notify      *NotifyRequest `protobuf:"bytes,3,opt,name=notify,proto3" json:"notify,omitempty"`
}

func (x *DeliverNotificationRequest) Reset() {
	*x = DeliverNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverNotificationRequest) ProtoMessage() {}

func (x *DeliverNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postman_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverNotificationRequest.ProtoReflect.Descriptor instead.
func (*DeliverNotificationRequest) Descriptor() ([]byte, []int) {
	return file_postman_proto_rawDescGZIP(), []int{0}
}

func (x *DeliverNotificationRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *DeliverNotificationRequest) GetDeviceToken() string {
	if x != nil {
		return x.DeviceToken
	}
	return ""
}

func (x *DeliverNotificationRequest) GetNotify() *NotifyRequest {
	if x != nil {
		return x.Notify
	}
	return nil
}

type DeliverNotificationBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Providers    []string       `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty"`
	DeviceTokens []string       `protobuf:"bytes,2,rep,name=device_tokens,json=deviceTokens,proto3" json:"device_tokens,omitempty"`
	Notify       *NotifyRequest `protobuf:"bytes,3,opt,name=notify,proto3" json:"notify,omitempty"`
}

func (x *DeliverNotificationBatchRequest) Reset() {
	*x = DeliverNotificationBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverNotificationBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverNotificationBatchRequest) ProtoMessage() {}

func (x *DeliverNotificationBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postman_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverNotificationBatchRequest.ProtoReflect.Descriptor instead.
func (*DeliverNotificationBatchRequest) Descriptor() ([]byte, []int) {
	return file_postman_proto_rawDescGZIP(), []int{1}
}

func (x *DeliverNotificationBatchRequest) GetProviders() []string {
	if x != nil {
		return x.Providers
	}
	return nil
}

func (x *DeliverNotificationBatchRequest) GetDeviceTokens() []string {
	if x != nil {
		return x.DeviceTokens
	}
	return nil
}

func (x *DeliverNotificationBatchRequest) GetNotify() *NotifyRequest {
	if x != nil {
		return x.Notify
	}
	return nil
}

type DeliverEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To    string        `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Email *EmailRequest `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *DeliverEmailRequest) Reset() {
	*x = DeliverEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverEmailRequest) ProtoMessage() {}

func (x *DeliverEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postman_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverEmailRequest.ProtoReflect.Descriptor instead.
func (*DeliverEmailRequest) Descriptor() ([]byte, []int) {
	return file_postman_proto_rawDescGZIP(), []int{2}
}

func (x *DeliverEmailRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *DeliverEmailRequest) GetEmail() *EmailRequest {
	if x != nil {
		return x.Email
	}
	return nil
}

type DeliverEmailBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To    []string      `protobuf:"bytes,1,rep,name=to,proto3" json:"to,omitempty"`
	Email *EmailRequest `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *DeliverEmailBatchRequest) Reset() {
	*x = DeliverEmailBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverEmailBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverEmailBatchRequest) ProtoMessage() {}

func (x *DeliverEmailBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postman_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverEmailBatchRequest.ProtoReflect.Descriptor instead.
func (*DeliverEmailBatchRequest) Descriptor() ([]byte, []int) {
	return file_postman_proto_rawDescGZIP(), []int{3}
}

func (x *DeliverEmailBatchRequest) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *DeliverEmailBatchRequest) GetEmail() *EmailRequest {
	if x != nil {
		return x.Email
	}
	return nil
}

type EmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject  string  `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	TextBody *string `protobuf:"bytes,2,opt,name=text_body,json=textBody,proto3,oneof" json:"text_body,omitempty"`
	HtmlBody *string `protobuf:"bytes,3,opt,name=html_body,json=htmlBody,proto3,oneof" json:"html_body,omitempty"`
}

func (x *EmailRequest) Reset() {
	*x = EmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequest) ProtoMessage() {}

func (x *EmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postman_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequest.ProtoReflect.Descriptor instead.
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return file_postman_proto_rawDescGZIP(), []int{4}
}

func (x *EmailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailRequest) GetTextBody() string {
	if x != nil && x.TextBody != nil {
		return *x.TextBody
	}
	return ""
}

func (x *EmailRequest) GetHtmlBody() string {
	if x != nil && x.HtmlBody != nil {
		return *x.HtmlBody
	}
	return ""
}

type DeliverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeliverResponse) Reset() {
	*x = DeliverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postman_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverResponse) ProtoMessage() {}

func (x *DeliverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postman_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverResponse.ProtoReflect.Descriptor instead.
func (*DeliverResponse) Descriptor() ([]byte, []int) {
	return file_postman_proto_rawDescGZIP(), []int{5}
}

var File_postman_proto protoreflect.FileDescriptor

var file_postman_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x22, 0x92, 0x01, 0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0x50, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x29, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x55, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x88,
	0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x74, 0x65, 0x78, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x68,
	0x74, 0x6d, 0x6c, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x08, 0x68, 0x74, 0x6d, 0x6c, 0x42, 0x6f, 0x64, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd1, 0x02, 0x0a,
	0x07, 0x50, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x12, 0x52, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x18,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_postman_proto_rawDescOnce sync.Once
	file_postman_proto_rawDescData = file_postman_proto_rawDesc
)

func file_postman_proto_rawDescGZIP() []byte {
	file_postman_proto_rawDescOnce.Do(func() {
		file_postman_proto_rawDescData = protoimpl.X.CompressGZIP(file_postman_proto_rawDescData)
	})
	return file_postman_proto_rawDescData
}

var file_postman_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_postman_proto_goTypes = []any{
	(*DeliverNotificationRequest)(nil),      // 0: proto.DeliverNotificationRequest
	(*DeliverNotificationBatchRequest)(nil), // 1: proto.DeliverNotificationBatchRequest
	(*DeliverEmailRequest)(nil),             // 2: proto.DeliverEmailRequest
	(*DeliverEmailBatchRequest)(nil),        // 3: proto.DeliverEmailBatchRequest
	(*EmailRequest)(nil),                    // 4: proto.EmailRequest
	(*DeliverResponse)(nil),                 // 5: proto.DeliverResponse
	(*NotifyRequest)(nil),                   // 6: proto.NotifyRequest
}
var file_postman_proto_depIdxs = []int32{
	6, // 0: proto.DeliverNotificationRequest.notify:type_name -> proto.NotifyRequest
	6, // 1: proto.DeliverNotificationBatchRequest.notify:type_name -> proto.NotifyRequest
	4, // 2: proto.DeliverEmailRequest.email:type_name -> proto.EmailRequest
	4, // 3: proto.DeliverEmailBatchRequest.email:type_name -> proto.EmailRequest
	0, // 4: proto.Postman.DeliverNotification:input_type -> proto.DeliverNotificationRequest
	1, // 5: proto.Postman.DeliverNotificationBatch:input_type -> proto.DeliverNotificationBatchRequest
	2, // 6: proto.Postman.DeliverEmail:input_type -> proto.DeliverEmailRequest
	3, // 7: proto.Postman.DeliverEmailBatch:input_type -> proto.DeliverEmailBatchRequest
	5, // 8: proto.Postman.DeliverNotification:output_type -> proto.DeliverResponse
	5, // 9: proto.Postman.DeliverNotificationBatch:output_type -> proto.DeliverResponse
	5, // 10: proto.Postman.DeliverEmail:output_type -> proto.DeliverResponse
	5, // 11: proto.Postman.DeliverEmailBatch:output_type -> proto.DeliverResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_postman_proto_init() }
func file_postman_proto_init() {
	if File_postman_proto != nil {
		return
	}
	file_notify_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_postman_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DeliverNotificationRequest); i {
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
		file_postman_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeliverNotificationBatchRequest); i {
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
		file_postman_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeliverEmailRequest); i {
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
		file_postman_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeliverEmailBatchRequest); i {
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
		file_postman_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*EmailRequest); i {
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
		file_postman_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeliverResponse); i {
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
	file_postman_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_postman_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_postman_proto_goTypes,
		DependencyIndexes: file_postman_proto_depIdxs,
		MessageInfos:      file_postman_proto_msgTypes,
	}.Build()
	File_postman_proto = out.File
	file_postman_proto_rawDesc = nil
	file_postman_proto_goTypes = nil
	file_postman_proto_depIdxs = nil
}
