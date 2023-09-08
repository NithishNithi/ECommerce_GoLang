// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/customer.proto

package ecommerce

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

type CustomerDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId      string             `protobuf:"bytes,1,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	Firstname       string             `protobuf:"bytes,2,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname        string             `protobuf:"bytes,3,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	Password        string             `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	Email           string             `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Address         []*Address         `protobuf:"bytes,6,rep,name=Address,proto3" json:"Address,omitempty"`
	ShippingAddress []*ShippingAddress `protobuf:"bytes,7,rep,name=ShippingAddress,proto3" json:"ShippingAddress,omitempty"`
}

func (x *CustomerDetails) Reset() {
	*x = CustomerDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDetails) ProtoMessage() {}

func (x *CustomerDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDetails.ProtoReflect.Descriptor instead.
func (*CustomerDetails) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerDetails) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CustomerDetails) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *CustomerDetails) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *CustomerDetails) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CustomerDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CustomerDetails) GetAddress() []*Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *CustomerDetails) GetShippingAddress() []*ShippingAddress {
	if x != nil {
		return x.ShippingAddress
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,1,opt,name=Country,proto3" json:"Country,omitempty"`
	Street1 string `protobuf:"bytes,2,opt,name=Street1,proto3" json:"Street1,omitempty"`
	Street2 string `protobuf:"bytes,3,opt,name=Street2,proto3" json:"Street2,omitempty"`
	City    string `protobuf:"bytes,4,opt,name=City,proto3" json:"City,omitempty"`
	State   string `protobuf:"bytes,5,opt,name=State,proto3" json:"State,omitempty"`
	Zip     string `protobuf:"bytes,6,opt,name=Zip,proto3" json:"Zip,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetStreet1() string {
	if x != nil {
		return x.Street1
	}
	return ""
}

func (x *Address) GetStreet2() string {
	if x != nil {
		return x.Street2
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

type ShippingAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street1 string `protobuf:"bytes,1,opt,name=Street1,proto3" json:"Street1,omitempty"`
	Street2 string `protobuf:"bytes,2,opt,name=Street2,proto3" json:"Street2,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	State   string `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *ShippingAddress) Reset() {
	*x = ShippingAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShippingAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingAddress) ProtoMessage() {}

func (x *ShippingAddress) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippingAddress.ProtoReflect.Descriptor instead.
func (*ShippingAddress) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{2}
}

func (x *ShippingAddress) GetStreet1() string {
	if x != nil {
		return x.Street1
	}
	return ""
}

func (x *ShippingAddress) GetStreet2() string {
	if x != nil {
		return x.Street2
	}
	return ""
}

func (x *ShippingAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ShippingAddress) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type CustomerLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *CustomerLoginRequest) Reset() {
	*x = CustomerLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerLoginRequest) ProtoMessage() {}

func (x *CustomerLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerLoginRequest.ProtoReflect.Descriptor instead.
func (*CustomerLoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{3}
}

func (x *CustomerLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CustomerLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer_ID string `protobuf:"bytes,1,opt,name=Customer_ID,json=CustomerID,proto3" json:"Customer_ID,omitempty"`
}

func (x *CustomerResponse) Reset() {
	*x = CustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerResponse) ProtoMessage() {}

func (x *CustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerResponse.ProtoReflect.Descriptor instead.
func (*CustomerResponse) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{4}
}

func (x *CustomerResponse) GetCustomer_ID() string {
	if x != nil {
		return x.Customer_ID
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{5}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	Customerid string `protobuf:"bytes,3,opt,name=Customerid,proto3" json:"Customerid,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{6}
}

func (x *Token) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Token) GetCustomerid() string {
	if x != nil {
		return x.Customerid
	}
	return ""
}

type PasswordDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	OldPassword string `protobuf:"bytes,2,opt,name=OldPassword,proto3" json:"OldPassword,omitempty"`
	NewPassword string `protobuf:"bytes,3,opt,name=NewPassword,proto3" json:"NewPassword,omitempty"`
}

func (x *PasswordDetails) Reset() {
	*x = PasswordDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordDetails) ProtoMessage() {}

func (x *PasswordDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordDetails.ProtoReflect.Descriptor instead.
func (*PasswordDetails) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{7}
}

func (x *PasswordDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PasswordDetails) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *PasswordDetails) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type UpdateDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	Field      string `protobuf:"bytes,2,opt,name=Field,proto3" json:"Field,omitempty"`
	OldValue   string `protobuf:"bytes,3,opt,name=OldValue,proto3" json:"OldValue,omitempty"`
	NewValue   string `protobuf:"bytes,4,opt,name=NewValue,proto3" json:"NewValue,omitempty"`
}

func (x *UpdateDetails) Reset() {
	*x = UpdateDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDetails) ProtoMessage() {}

func (x *UpdateDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDetails.ProtoReflect.Descriptor instead.
func (*UpdateDetails) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateDetails) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *UpdateDetails) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *UpdateDetails) GetOldValue() string {
	if x != nil {
		return x.OldValue
	}
	return ""
}

func (x *UpdateDetails) GetNewValue() string {
	if x != nil {
		return x.NewValue
	}
	return ""
}

type DeleteDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *DeleteDetails) Reset() {
	*x = DeleteDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDetails) ProtoMessage() {}

func (x *DeleteDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDetails.ProtoReflect.Descriptor instead.
func (*DeleteDetails) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteDetails) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetbyId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *GetbyId) Reset() {
	*x = GetbyId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetbyId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetbyId) ProtoMessage() {}

func (x *GetbyId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetbyId.ProtoReflect.Descriptor instead.
func (*GetbyId) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{10}
}

func (x *GetbyId) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_proto_customer_proto protoreflect.FileDescriptor

var file_proto_customer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x22, 0x8f, 0x02, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x2b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x43, 0x0a,
	0x0f, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x0f, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x32, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x32, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x5a, 0x69, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x5a, 0x69, 0x70, 0x22, 0x6f, 0x0a, 0x0f, 0x53, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x32, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x48, 0x0a, 0x14, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x33, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x22, 0x07, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x53, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x22, 0x6b, 0x0a, 0x0f, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x7d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x6c,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x6c,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x65, 0x77, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1f, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x62, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x99, 0x03, 0x0a, 0x0f, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x12, 0x19, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x1a, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x19, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x1a, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x1a, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x1a, 0x0f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x62, 0x79, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x73, 0x68, 0x6f, 0x72, 0x65, 0x6e, 0x73, 0x31, 0x38,
	0x2f, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_customer_proto_rawDescOnce sync.Once
	file_proto_customer_proto_rawDescData = file_proto_customer_proto_rawDesc
)

func file_proto_customer_proto_rawDescGZIP() []byte {
	file_proto_customer_proto_rawDescOnce.Do(func() {
		file_proto_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_customer_proto_rawDescData)
	})
	return file_proto_customer_proto_rawDescData
}

var file_proto_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_customer_proto_goTypes = []interface{}{
	(*CustomerDetails)(nil),      // 0: customer.CustomerDetails
	(*Address)(nil),              // 1: customer.Address
	(*ShippingAddress)(nil),      // 2: customer.ShippingAddress
	(*CustomerLoginRequest)(nil), // 3: customer.CustomerLoginRequest
	(*CustomerResponse)(nil),     // 4: customer.CustomerResponse
	(*Empty)(nil),                // 5: customer.empty
	(*Token)(nil),                // 6: customer.Token
	(*PasswordDetails)(nil),      // 7: customer.PasswordDetails
	(*UpdateDetails)(nil),        // 8: customer.UpdateDetails
	(*DeleteDetails)(nil),        // 9: customer.DeleteDetails
	(*GetbyId)(nil),              // 10: customer.GetbyId
}
var file_proto_customer_proto_depIdxs = []int32{
	1,  // 0: customer.CustomerDetails.Address:type_name -> customer.Address
	2,  // 1: customer.CustomerDetails.ShippingAddress:type_name -> customer.ShippingAddress
	0,  // 2: customer.CustomerService.CreateCustomer:input_type -> customer.CustomerDetails
	6,  // 3: customer.CustomerService.CreateTokens:input_type -> customer.Token
	7,  // 4: customer.CustomerService.UpdatePassword:input_type -> customer.PasswordDetails
	8,  // 5: customer.CustomerService.UpdateCustomer:input_type -> customer.UpdateDetails
	9,  // 6: customer.CustomerService.DeleteCustomer:input_type -> customer.DeleteDetails
	10, // 7: customer.CustomerService.GetByCustomerId:input_type -> customer.GetbyId
	4,  // 8: customer.CustomerService.CreateCustomer:output_type -> customer.CustomerResponse
	5,  // 9: customer.CustomerService.CreateTokens:output_type -> customer.empty
	4,  // 10: customer.CustomerService.UpdatePassword:output_type -> customer.CustomerResponse
	4,  // 11: customer.CustomerService.UpdateCustomer:output_type -> customer.CustomerResponse
	5,  // 12: customer.CustomerService.DeleteCustomer:output_type -> customer.empty
	0,  // 13: customer.CustomerService.GetByCustomerId:output_type -> customer.CustomerDetails
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_customer_proto_init() }
func file_proto_customer_proto_init() {
	if File_proto_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerDetails); i {
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
		file_proto_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_proto_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShippingAddress); i {
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
		file_proto_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerLoginRequest); i {
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
		file_proto_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerResponse); i {
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
		file_proto_customer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_customer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_proto_customer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordDetails); i {
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
		file_proto_customer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDetails); i {
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
		file_proto_customer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDetails); i {
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
		file_proto_customer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetbyId); i {
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
			RawDescriptor: file_proto_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_customer_proto_goTypes,
		DependencyIndexes: file_proto_customer_proto_depIdxs,
		MessageInfos:      file_proto_customer_proto_msgTypes,
	}.Build()
	File_proto_customer_proto = out.File
	file_proto_customer_proto_rawDesc = nil
	file_proto_customer_proto_goTypes = nil
	file_proto_customer_proto_depIdxs = nil
}
