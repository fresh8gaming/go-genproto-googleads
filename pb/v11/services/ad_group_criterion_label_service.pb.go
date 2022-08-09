// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: google/ads/googleads/v11/services/ad_group_criterion_label_service.proto

package services

import (
	context "context"
	resources "github.com/fresh8gaming/go-genproto-googleads/pb/v11/resources"
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

// Request message for
// [AdGroupCriterionLabelService.MutateAdGroupCriterionLabels][google.ads.googleads.v11.services.AdGroupCriterionLabelService.MutateAdGroupCriterionLabels].
type MutateAdGroupCriterionLabelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the customer whose ad group criterion labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on ad group criterion labels.
	Operations []*AdGroupCriterionLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateAdGroupCriterionLabelsRequest) Reset() {
	*x = MutateAdGroupCriterionLabelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupCriterionLabelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupCriterionLabelsRequest) ProtoMessage() {}

func (x *MutateAdGroupCriterionLabelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupCriterionLabelsRequest.ProtoReflect.Descriptor instead.
func (*MutateAdGroupCriterionLabelsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateAdGroupCriterionLabelsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateAdGroupCriterionLabelsRequest) GetOperations() []*AdGroupCriterionLabelOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateAdGroupCriterionLabelsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateAdGroupCriterionLabelsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an ad group criterion label.
type AdGroupCriterionLabelOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*AdGroupCriterionLabelOperation_Create
	//	*AdGroupCriterionLabelOperation_Remove
	Operation isAdGroupCriterionLabelOperation_Operation `protobuf_oneof:"operation"`
}

func (x *AdGroupCriterionLabelOperation) Reset() {
	*x = AdGroupCriterionLabelOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupCriterionLabelOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupCriterionLabelOperation) ProtoMessage() {}

func (x *AdGroupCriterionLabelOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupCriterionLabelOperation.ProtoReflect.Descriptor instead.
func (*AdGroupCriterionLabelOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDescGZIP(), []int{1}
}

func (m *AdGroupCriterionLabelOperation) GetOperation() isAdGroupCriterionLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *AdGroupCriterionLabelOperation) GetCreate() *resources.AdGroupCriterionLabel {
	if x, ok := x.GetOperation().(*AdGroupCriterionLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *AdGroupCriterionLabelOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*AdGroupCriterionLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isAdGroupCriterionLabelOperation_Operation interface {
	isAdGroupCriterionLabelOperation_Operation()
}

type AdGroupCriterionLabelOperation_Create struct {
	// Create operation: No resource name is expected for the new ad group
	// label.
	Create *resources.AdGroupCriterionLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupCriterionLabelOperation_Remove struct {
	// Remove operation: A resource name for the ad group criterion label
	// being removed, in this format:
	//
	// `customers/{customer_id}/adGroupCriterionLabels/{ad_group_id}~{criterion_id}~{label_id}`
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*AdGroupCriterionLabelOperation_Create) isAdGroupCriterionLabelOperation_Operation() {}

func (*AdGroupCriterionLabelOperation_Remove) isAdGroupCriterionLabelOperation_Operation() {}

// Response message for an ad group criterion labels mutate.
type MutateAdGroupCriterionLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateAdGroupCriterionLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateAdGroupCriterionLabelsResponse) Reset() {
	*x = MutateAdGroupCriterionLabelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupCriterionLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupCriterionLabelsResponse) ProtoMessage() {}

func (x *MutateAdGroupCriterionLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupCriterionLabelsResponse.ProtoReflect.Descriptor instead.
func (*MutateAdGroupCriterionLabelsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateAdGroupCriterionLabelsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateAdGroupCriterionLabelsResponse) GetResults() []*MutateAdGroupCriterionLabelResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for an ad group criterion label mutate.
type MutateAdGroupCriterionLabelResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateAdGroupCriterionLabelResult) Reset() {
	*x = MutateAdGroupCriterionLabelResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupCriterionLabelResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupCriterionLabelResult) ProtoMessage() {}

func (x *MutateAdGroupCriterionLabelResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupCriterionLabelResult.ProtoReflect.Descriptor instead.
func (*MutateAdGroupCriterionLabelResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateAdGroupCriterionLabelResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x41, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x02, 0x0a,
	0x23, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x66, 0x0a, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79,
	0x22, 0xd1, 0x01, 0x0a, 0x1e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x48, 0x00,
	0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x48, 0x00, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xce, 0x01, 0x0a, 0x24, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f,
	0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x7d, 0x0a, 0x21, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x33, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f,
	0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x32, 0xf9, 0x02, 0x0a, 0x1c, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x91, 0x02, 0x0a, 0x1c, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f,
	0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x22,
	0x3c, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f,
	0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x42, 0x8d, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x21, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64,
	0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDescData = file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDesc
)

func file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDescData
}

var file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_goTypes = []interface{}{
	(*MutateAdGroupCriterionLabelsRequest)(nil),  // 0: google.ads.googleads.v11.services.MutateAdGroupCriterionLabelsRequest
	(*AdGroupCriterionLabelOperation)(nil),       // 1: google.ads.googleads.v11.services.AdGroupCriterionLabelOperation
	(*MutateAdGroupCriterionLabelsResponse)(nil), // 2: google.ads.googleads.v11.services.MutateAdGroupCriterionLabelsResponse
	(*MutateAdGroupCriterionLabelResult)(nil),    // 3: google.ads.googleads.v11.services.MutateAdGroupCriterionLabelResult
	(*resources.AdGroupCriterionLabel)(nil),      // 4: google.ads.googleads.v11.resources.AdGroupCriterionLabel
	(*status.Status)(nil),                        // 5: google.rpc.Status
}
var file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v11.services.MutateAdGroupCriterionLabelsRequest.operations:type_name -> google.ads.googleads.v11.services.AdGroupCriterionLabelOperation
	4, // 1: google.ads.googleads.v11.services.AdGroupCriterionLabelOperation.create:type_name -> google.ads.googleads.v11.resources.AdGroupCriterionLabel
	5, // 2: google.ads.googleads.v11.services.MutateAdGroupCriterionLabelsResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 3: google.ads.googleads.v11.services.MutateAdGroupCriterionLabelsResponse.results:type_name -> google.ads.googleads.v11.services.MutateAdGroupCriterionLabelResult
	0, // 4: google.ads.googleads.v11.services.AdGroupCriterionLabelService.MutateAdGroupCriterionLabels:input_type -> google.ads.googleads.v11.services.MutateAdGroupCriterionLabelsRequest
	2, // 5: google.ads.googleads.v11.services.AdGroupCriterionLabelService.MutateAdGroupCriterionLabels:output_type -> google.ads.googleads.v11.services.MutateAdGroupCriterionLabelsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_init() }
func file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_init() {
	if File_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupCriterionLabelsRequest); i {
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
		file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupCriterionLabelOperation); i {
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
		file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupCriterionLabelsResponse); i {
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
		file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupCriterionLabelResult); i {
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
	file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AdGroupCriterionLabelOperation_Create)(nil),
		(*AdGroupCriterionLabelOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto = out.File
	file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_rawDesc = nil
	file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_goTypes = nil
	file_google_ads_googleads_v11_services_ad_group_criterion_label_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupCriterionLabelServiceClient is the client API for AdGroupCriterionLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupCriterionLabelServiceClient interface {
	// Creates and removes ad group criterion labels.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateAdGroupCriterionLabels(ctx context.Context, in *MutateAdGroupCriterionLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionLabelsResponse, error)
}

type adGroupCriterionLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupCriterionLabelServiceClient(cc grpc.ClientConnInterface) AdGroupCriterionLabelServiceClient {
	return &adGroupCriterionLabelServiceClient{cc}
}

func (c *adGroupCriterionLabelServiceClient) MutateAdGroupCriterionLabels(ctx context.Context, in *MutateAdGroupCriterionLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionLabelsResponse, error) {
	out := new(MutateAdGroupCriterionLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.AdGroupCriterionLabelService/MutateAdGroupCriterionLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCriterionLabelServiceServer is the server API for AdGroupCriterionLabelService service.
type AdGroupCriterionLabelServiceServer interface {
	// Creates and removes ad group criterion labels.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateAdGroupCriterionLabels(context.Context, *MutateAdGroupCriterionLabelsRequest) (*MutateAdGroupCriterionLabelsResponse, error)
}

// UnimplementedAdGroupCriterionLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupCriterionLabelServiceServer struct {
}

func (*UnimplementedAdGroupCriterionLabelServiceServer) MutateAdGroupCriterionLabels(context.Context, *MutateAdGroupCriterionLabelsRequest) (*MutateAdGroupCriterionLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupCriterionLabels not implemented")
}

func RegisterAdGroupCriterionLabelServiceServer(s *grpc.Server, srv AdGroupCriterionLabelServiceServer) {
	s.RegisterService(&_AdGroupCriterionLabelService_serviceDesc, srv)
}

func _AdGroupCriterionLabelService_MutateAdGroupCriterionLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupCriterionLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionLabelServiceServer).MutateAdGroupCriterionLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.AdGroupCriterionLabelService/MutateAdGroupCriterionLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionLabelServiceServer).MutateAdGroupCriterionLabels(ctx, req.(*MutateAdGroupCriterionLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupCriterionLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.AdGroupCriterionLabelService",
	HandlerType: (*AdGroupCriterionLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupCriterionLabels",
			Handler:    _AdGroupCriterionLabelService_MutateAdGroupCriterionLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/ad_group_criterion_label_service.proto",
}
