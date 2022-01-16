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
// source: google/ads/googleads/v8/services/ad_group_ad_label_service.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	resources "github.com/dictav/go-genproto-googleads/pb/v8/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [AdGroupAdLabelService.GetAdGroupAdLabel][google.ads.googleads.v8.services.AdGroupAdLabelService.GetAdGroupAdLabel].
type GetAdGroupAdLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the ad group ad label to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetAdGroupAdLabelRequest) Reset() {
	*x = GetAdGroupAdLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdGroupAdLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdGroupAdLabelRequest) ProtoMessage() {}

func (x *GetAdGroupAdLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdGroupAdLabelRequest.ProtoReflect.Descriptor instead.
func (*GetAdGroupAdLabelRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAdGroupAdLabelRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for [AdGroupAdLabelService.MutateAdGroupAdLabels][google.ads.googleads.v8.services.AdGroupAdLabelService.MutateAdGroupAdLabels].
type MutateAdGroupAdLabelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the customer whose ad group ad labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on ad group ad labels.
	Operations []*AdGroupAdLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateAdGroupAdLabelsRequest) Reset() {
	*x = MutateAdGroupAdLabelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupAdLabelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupAdLabelsRequest) ProtoMessage() {}

func (x *MutateAdGroupAdLabelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupAdLabelsRequest.ProtoReflect.Descriptor instead.
func (*MutateAdGroupAdLabelsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateAdGroupAdLabelsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateAdGroupAdLabelsRequest) GetOperations() []*AdGroupAdLabelOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateAdGroupAdLabelsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateAdGroupAdLabelsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an ad group ad label.
type AdGroupAdLabelOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*AdGroupAdLabelOperation_Create
	//	*AdGroupAdLabelOperation_Remove
	Operation isAdGroupAdLabelOperation_Operation `protobuf_oneof:"operation"`
}

func (x *AdGroupAdLabelOperation) Reset() {
	*x = AdGroupAdLabelOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupAdLabelOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupAdLabelOperation) ProtoMessage() {}

func (x *AdGroupAdLabelOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupAdLabelOperation.ProtoReflect.Descriptor instead.
func (*AdGroupAdLabelOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescGZIP(), []int{2}
}

func (m *AdGroupAdLabelOperation) GetOperation() isAdGroupAdLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *AdGroupAdLabelOperation) GetCreate() *resources.AdGroupAdLabel {
	if x, ok := x.GetOperation().(*AdGroupAdLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *AdGroupAdLabelOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*AdGroupAdLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isAdGroupAdLabelOperation_Operation interface {
	isAdGroupAdLabelOperation_Operation()
}

type AdGroupAdLabelOperation_Create struct {
	// Create operation: No resource name is expected for the new ad group ad
	// label.
	Create *resources.AdGroupAdLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupAdLabelOperation_Remove struct {
	// Remove operation: A resource name for the ad group ad label
	// being removed, in this format:
	//
	// `customers/{customer_id}/adGroupAdLabels/{ad_group_id}~{ad_id}
	// _{label_id}`
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*AdGroupAdLabelOperation_Create) isAdGroupAdLabelOperation_Operation() {}

func (*AdGroupAdLabelOperation_Remove) isAdGroupAdLabelOperation_Operation() {}

// Response message for an ad group ad labels mutate.
type MutateAdGroupAdLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateAdGroupAdLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateAdGroupAdLabelsResponse) Reset() {
	*x = MutateAdGroupAdLabelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupAdLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupAdLabelsResponse) ProtoMessage() {}

func (x *MutateAdGroupAdLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupAdLabelsResponse.ProtoReflect.Descriptor instead.
func (*MutateAdGroupAdLabelsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateAdGroupAdLabelsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateAdGroupAdLabelsResponse) GetResults() []*MutateAdGroupAdLabelResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for an ad group ad label mutate.
type MutateAdGroupAdLabelResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateAdGroupAdLabelResult) Reset() {
	*x = MutateAdGroupAdLabelResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupAdLabelResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupAdLabelResult) ProtoMessage() {}

func (x *MutateAdGroupAdLabelResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupAdLabelResult.ProtoReflect.Descriptor instead.
func (*MutateAdGroupAdLabelResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateAdGroupAdLabelResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v8_services_ad_group_ad_label_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x64, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x61, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f,
	0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf2, 0x01,
	0x0a, 0x1c, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x5e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e,
	0x6c, 0x79, 0x22, 0x8d, 0x01, 0x0a, 0x17, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b,
	0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xbf, 0x01, 0x0a, 0x1d, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x56, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x22, 0x41, 0x0a, 0x1a, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xa3, 0x04, 0x0a, 0x15, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xcd, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x49, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x12, 0x31,
	0x2f, 0x76, 0x38, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x61,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x2a,
	0x7d, 0xda, 0x41, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0xf2, 0x01, 0x0a, 0x15, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x39, 0x22, 0x34, 0x2f, 0x76, 0x38, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x3d, 0x2a, 0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x81, 0x02,
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1a, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x38, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescData = file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDesc
)

func file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDescData
}

var file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_goTypes = []interface{}{
	(*GetAdGroupAdLabelRequest)(nil),      // 0: google.ads.googleads.v8.services.GetAdGroupAdLabelRequest
	(*MutateAdGroupAdLabelsRequest)(nil),  // 1: google.ads.googleads.v8.services.MutateAdGroupAdLabelsRequest
	(*AdGroupAdLabelOperation)(nil),       // 2: google.ads.googleads.v8.services.AdGroupAdLabelOperation
	(*MutateAdGroupAdLabelsResponse)(nil), // 3: google.ads.googleads.v8.services.MutateAdGroupAdLabelsResponse
	(*MutateAdGroupAdLabelResult)(nil),    // 4: google.ads.googleads.v8.services.MutateAdGroupAdLabelResult
	(*resources.AdGroupAdLabel)(nil),      // 5: google.ads.googleads.v8.resources.AdGroupAdLabel
	(*status.Status)(nil),                 // 6: google.rpc.Status
}
var file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v8.services.MutateAdGroupAdLabelsRequest.operations:type_name -> google.ads.googleads.v8.services.AdGroupAdLabelOperation
	5, // 1: google.ads.googleads.v8.services.AdGroupAdLabelOperation.create:type_name -> google.ads.googleads.v8.resources.AdGroupAdLabel
	6, // 2: google.ads.googleads.v8.services.MutateAdGroupAdLabelsResponse.partial_failure_error:type_name -> google.rpc.Status
	4, // 3: google.ads.googleads.v8.services.MutateAdGroupAdLabelsResponse.results:type_name -> google.ads.googleads.v8.services.MutateAdGroupAdLabelResult
	0, // 4: google.ads.googleads.v8.services.AdGroupAdLabelService.GetAdGroupAdLabel:input_type -> google.ads.googleads.v8.services.GetAdGroupAdLabelRequest
	1, // 5: google.ads.googleads.v8.services.AdGroupAdLabelService.MutateAdGroupAdLabels:input_type -> google.ads.googleads.v8.services.MutateAdGroupAdLabelsRequest
	5, // 6: google.ads.googleads.v8.services.AdGroupAdLabelService.GetAdGroupAdLabel:output_type -> google.ads.googleads.v8.resources.AdGroupAdLabel
	3, // 7: google.ads.googleads.v8.services.AdGroupAdLabelService.MutateAdGroupAdLabels:output_type -> google.ads.googleads.v8.services.MutateAdGroupAdLabelsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_init() }
func file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_init() {
	if File_google_ads_googleads_v8_services_ad_group_ad_label_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdGroupAdLabelRequest); i {
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
		file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupAdLabelsRequest); i {
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
		file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupAdLabelOperation); i {
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
		file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupAdLabelsResponse); i {
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
		file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupAdLabelResult); i {
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
	file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AdGroupAdLabelOperation_Create)(nil),
		(*AdGroupAdLabelOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_services_ad_group_ad_label_service_proto = out.File
	file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_rawDesc = nil
	file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_goTypes = nil
	file_google_ads_googleads_v8_services_ad_group_ad_label_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupAdLabelServiceClient is the client API for AdGroupAdLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupAdLabelServiceClient interface {
	// Returns the requested ad group ad label in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdGroupAdLabel(ctx context.Context, in *GetAdGroupAdLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupAdLabel, error)
	// Creates and removes ad group ad labels.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [LabelError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateAdGroupAdLabels(ctx context.Context, in *MutateAdGroupAdLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdLabelsResponse, error)
}

type adGroupAdLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupAdLabelServiceClient(cc grpc.ClientConnInterface) AdGroupAdLabelServiceClient {
	return &adGroupAdLabelServiceClient{cc}
}

func (c *adGroupAdLabelServiceClient) GetAdGroupAdLabel(ctx context.Context, in *GetAdGroupAdLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupAdLabel, error) {
	out := new(resources.AdGroupAdLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.AdGroupAdLabelService/GetAdGroupAdLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupAdLabelServiceClient) MutateAdGroupAdLabels(ctx context.Context, in *MutateAdGroupAdLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdLabelsResponse, error) {
	out := new(MutateAdGroupAdLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.AdGroupAdLabelService/MutateAdGroupAdLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAdLabelServiceServer is the server API for AdGroupAdLabelService service.
type AdGroupAdLabelServiceServer interface {
	// Returns the requested ad group ad label in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdGroupAdLabel(context.Context, *GetAdGroupAdLabelRequest) (*resources.AdGroupAdLabel, error)
	// Creates and removes ad group ad labels.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [LabelError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateAdGroupAdLabels(context.Context, *MutateAdGroupAdLabelsRequest) (*MutateAdGroupAdLabelsResponse, error)
}

// UnimplementedAdGroupAdLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupAdLabelServiceServer struct {
}

func (*UnimplementedAdGroupAdLabelServiceServer) GetAdGroupAdLabel(context.Context, *GetAdGroupAdLabelRequest) (*resources.AdGroupAdLabel, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupAdLabel not implemented")
}
func (*UnimplementedAdGroupAdLabelServiceServer) MutateAdGroupAdLabels(context.Context, *MutateAdGroupAdLabelsRequest) (*MutateAdGroupAdLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupAdLabels not implemented")
}

func RegisterAdGroupAdLabelServiceServer(s *grpc.Server, srv AdGroupAdLabelServiceServer) {
	s.RegisterService(&_AdGroupAdLabelService_serviceDesc, srv)
}

func _AdGroupAdLabelService_GetAdGroupAdLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupAdLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdLabelServiceServer).GetAdGroupAdLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.AdGroupAdLabelService/GetAdGroupAdLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdLabelServiceServer).GetAdGroupAdLabel(ctx, req.(*GetAdGroupAdLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupAdLabelService_MutateAdGroupAdLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupAdLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdLabelServiceServer).MutateAdGroupAdLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.AdGroupAdLabelService/MutateAdGroupAdLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdLabelServiceServer).MutateAdGroupAdLabels(ctx, req.(*MutateAdGroupAdLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupAdLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v8.services.AdGroupAdLabelService",
	HandlerType: (*AdGroupAdLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupAdLabel",
			Handler:    _AdGroupAdLabelService_GetAdGroupAdLabel_Handler,
		},
		{
			MethodName: "MutateAdGroupAdLabels",
			Handler:    _AdGroupAdLabelService_MutateAdGroupAdLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v8/services/ad_group_ad_label_service.proto",
}
