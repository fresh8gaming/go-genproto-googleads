// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v8/services/invoice_service.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	enums "github.com/dictav/go-genproto-googleads/pb/v8/enums"
	resources "github.com/dictav/go-genproto-googleads/pb/v8/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request message for fetching the invoices of a given billing setup that were
// issued during a given month.
type ListInvoicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer to fetch invoices for.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The billing setup resource name of the requested invoices.
	//
	// `customers/{customer_id}/billingSetups/{billing_setup_id}`
	BillingSetup string `protobuf:"bytes,2,opt,name=billing_setup,json=billingSetup,proto3" json:"billing_setup,omitempty"`
	// Required. The issue year to retrieve invoices, in yyyy format. Only
	// invoices issued in 2019 or later can be retrieved.
	IssueYear string `protobuf:"bytes,3,opt,name=issue_year,json=issueYear,proto3" json:"issue_year,omitempty"`
	// Required. The issue month to retrieve invoices.
	IssueMonth enums.MonthOfYearEnum_MonthOfYear `protobuf:"varint,4,opt,name=issue_month,json=issueMonth,proto3,enum=google.ads.googleads.v8.enums.MonthOfYearEnum_MonthOfYear" json:"issue_month,omitempty"`
}

func (x *ListInvoicesRequest) Reset() {
	*x = ListInvoicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_services_invoice_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvoicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvoicesRequest) ProtoMessage() {}

func (x *ListInvoicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_services_invoice_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvoicesRequest.ProtoReflect.Descriptor instead.
func (*ListInvoicesRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_services_invoice_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListInvoicesRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *ListInvoicesRequest) GetBillingSetup() string {
	if x != nil {
		return x.BillingSetup
	}
	return ""
}

func (x *ListInvoicesRequest) GetIssueYear() string {
	if x != nil {
		return x.IssueYear
	}
	return ""
}

func (x *ListInvoicesRequest) GetIssueMonth() enums.MonthOfYearEnum_MonthOfYear {
	if x != nil {
		return x.IssueMonth
	}
	return enums.MonthOfYearEnum_UNSPECIFIED
}

// Response message for [InvoiceService.ListInvoices][google.ads.googleads.v8.services.InvoiceService.ListInvoices].
type ListInvoicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of invoices that match the billing setup and time period.
	Invoices []*resources.Invoice `protobuf:"bytes,1,rep,name=invoices,proto3" json:"invoices,omitempty"`
}

func (x *ListInvoicesResponse) Reset() {
	*x = ListInvoicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_services_invoice_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvoicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvoicesResponse) ProtoMessage() {}

func (x *ListInvoicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_services_invoice_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvoicesResponse.ProtoReflect.Descriptor instead.
func (*ListInvoicesResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_services_invoice_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListInvoicesResponse) GetInvoices() []*resources.Invoice {
	if x != nil {
		return x.Invoices
	}
	return nil
}

var File_google_ads_googleads_v8_services_invoice_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_services_invoice_service_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x31, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x5f,
	0x6f, 0x66, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x74, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0c, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x22,
	0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x59, 0x65,
	0x61, 0x72, 0x12, 0x60, 0x0a, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x4f, 0x66, 0x59,
	0x65, 0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x4f, 0x66, 0x59,
	0x65, 0x61, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x22, 0x5e, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x08,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x32, 0xba, 0x02, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe0, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12,
	0x26, 0x2f, 0x76, 0x38, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0xda, 0x41, 0x30, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65,
	0x74, 0x75, 0x70, 0x2c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x2c, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x42, 0xfa, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x13, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64,
	0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_services_invoice_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_services_invoice_service_proto_rawDescData = file_google_ads_googleads_v8_services_invoice_service_proto_rawDesc
)

func file_google_ads_googleads_v8_services_invoice_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_services_invoice_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_services_invoice_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_services_invoice_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_services_invoice_service_proto_rawDescData
}

var file_google_ads_googleads_v8_services_invoice_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v8_services_invoice_service_proto_goTypes = []interface{}{
	(*ListInvoicesRequest)(nil),            // 0: google.ads.googleads.v8.services.ListInvoicesRequest
	(*ListInvoicesResponse)(nil),           // 1: google.ads.googleads.v8.services.ListInvoicesResponse
	(enums.MonthOfYearEnum_MonthOfYear)(0), // 2: google.ads.googleads.v8.enums.MonthOfYearEnum.MonthOfYear
	(*resources.Invoice)(nil),              // 3: google.ads.googleads.v8.resources.Invoice
}
var file_google_ads_googleads_v8_services_invoice_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v8.services.ListInvoicesRequest.issue_month:type_name -> google.ads.googleads.v8.enums.MonthOfYearEnum.MonthOfYear
	3, // 1: google.ads.googleads.v8.services.ListInvoicesResponse.invoices:type_name -> google.ads.googleads.v8.resources.Invoice
	0, // 2: google.ads.googleads.v8.services.InvoiceService.ListInvoices:input_type -> google.ads.googleads.v8.services.ListInvoicesRequest
	1, // 3: google.ads.googleads.v8.services.InvoiceService.ListInvoices:output_type -> google.ads.googleads.v8.services.ListInvoicesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_services_invoice_service_proto_init() }
func file_google_ads_googleads_v8_services_invoice_service_proto_init() {
	if File_google_ads_googleads_v8_services_invoice_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_services_invoice_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvoicesRequest); i {
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
		file_google_ads_googleads_v8_services_invoice_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvoicesResponse); i {
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
			RawDescriptor: file_google_ads_googleads_v8_services_invoice_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v8_services_invoice_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_services_invoice_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v8_services_invoice_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_services_invoice_service_proto = out.File
	file_google_ads_googleads_v8_services_invoice_service_proto_rawDesc = nil
	file_google_ads_googleads_v8_services_invoice_service_proto_goTypes = nil
	file_google_ads_googleads_v8_services_invoice_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InvoiceServiceClient is the client API for InvoiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InvoiceServiceClient interface {
	// Returns all invoices associated with a billing setup, for a given month.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [InvoiceError]()
	//   [QuotaError]()
	//   [RequestError]()
	ListInvoices(ctx context.Context, in *ListInvoicesRequest, opts ...grpc.CallOption) (*ListInvoicesResponse, error)
}

type invoiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceServiceClient(cc grpc.ClientConnInterface) InvoiceServiceClient {
	return &invoiceServiceClient{cc}
}

func (c *invoiceServiceClient) ListInvoices(ctx context.Context, in *ListInvoicesRequest, opts ...grpc.CallOption) (*ListInvoicesResponse, error) {
	out := new(ListInvoicesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.InvoiceService/ListInvoices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServiceServer is the server API for InvoiceService service.
type InvoiceServiceServer interface {
	// Returns all invoices associated with a billing setup, for a given month.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [InvoiceError]()
	//   [QuotaError]()
	//   [RequestError]()
	ListInvoices(context.Context, *ListInvoicesRequest) (*ListInvoicesResponse, error)
}

// UnimplementedInvoiceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInvoiceServiceServer struct {
}

func (*UnimplementedInvoiceServiceServer) ListInvoices(context.Context, *ListInvoicesRequest) (*ListInvoicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInvoices not implemented")
}

func RegisterInvoiceServiceServer(s *grpc.Server, srv InvoiceServiceServer) {
	s.RegisterService(&_InvoiceService_serviceDesc, srv)
}

func _InvoiceService_ListInvoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInvoicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).ListInvoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.InvoiceService/ListInvoices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).ListInvoices(ctx, req.(*ListInvoicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InvoiceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v8.services.InvoiceService",
	HandlerType: (*InvoiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListInvoices",
			Handler:    _InvoiceService_ListInvoices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v8/services/invoice_service.proto",
}
