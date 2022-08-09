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
// source: google/ads/googleads/v11/resources/experiment.proto

package resources

import (
	common "github.com/fresh8gaming/go-genproto-googleads/pb/v11/common"
	enums "github.com/fresh8gaming/go-genproto-googleads/pb/v11/enums"
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

// A Google ads experiment for users to experiment changes on multiple
// campaigns, compare the performance, and apply the effective changes.
type Experiment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the experiment.
	// Experiment resource names have the form:
	//
	// `customers/{customer_id}/experiments/{experiment_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the experiment. Read only.
	ExperimentId *int64 `protobuf:"varint,9,opt,name=experiment_id,json=experimentId,proto3,oneof" json:"experiment_id,omitempty"`
	// Required. The name of the experiment. It must have a minimum length of 1 and
	// maximum length of 1024. It must be unique under a customer.
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the experiment. It must have a minimum length of 1 and
	// maximum length of 2048.
	Description string `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	// For system managed experiments, the advertiser must provide a suffix during
	// construction, in the setup stage before moving to initiated. The suffix
	// will be appended to the in-design and experiment campaign names so that the
	// name is base campaign name + suffix.
	Suffix string `protobuf:"bytes,12,opt,name=suffix,proto3" json:"suffix,omitempty"`
	// The product/feature that uses this experiment.
	Type enums.ExperimentTypeEnum_ExperimentType `protobuf:"varint,13,opt,name=type,proto3,enum=google.ads.googleads.v11.enums.ExperimentTypeEnum_ExperimentType" json:"type,omitempty"`
	// The Advertiser-desired status of this experiment.
	Status enums.ExperimentStatusEnum_ExperimentStatus `protobuf:"varint,14,opt,name=status,proto3,enum=google.ads.googleads.v11.enums.ExperimentStatusEnum_ExperimentStatus" json:"status,omitempty"`
	// Date when the experiment starts. By default, the experiment starts
	// now or on the campaign's start date, whichever is later. If this field is
	// set, then the experiment starts at the beginning of the specified date in
	// the customer's time zone.
	//
	// Format: YYYY-MM-DD
	// Example: 2019-03-14
	StartDate *string `protobuf:"bytes,15,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
	// Date when the experiment ends. By default, the experiment ends on
	// the campaign's end date. If this field is set, then the experiment ends at
	// the end of the specified date in the customer's time zone.
	//
	// Format: YYYY-MM-DD
	// Example: 2019-04-18
	EndDate *string `protobuf:"bytes,16,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	// The goals of this experiment.
	Goals []*common.MetricGoal `protobuf:"bytes,17,rep,name=goals,proto3" json:"goals,omitempty"`
	// Output only. The resource name of the long-running operation that can be used to poll
	// for completion of experiment schedule or promote. The most recent long
	// running operation is returned.
	LongRunningOperation *string `protobuf:"bytes,18,opt,name=long_running_operation,json=longRunningOperation,proto3,oneof" json:"long_running_operation,omitempty"`
	// Output only. The status of the experiment promotion process.
	PromoteStatus enums.AsyncActionStatusEnum_AsyncActionStatus `protobuf:"varint,19,opt,name=promote_status,json=promoteStatus,proto3,enum=google.ads.googleads.v11.enums.AsyncActionStatusEnum_AsyncActionStatus" json:"promote_status,omitempty"`
}

func (x *Experiment) Reset() {
	*x = Experiment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_resources_experiment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Experiment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Experiment) ProtoMessage() {}

func (x *Experiment) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_resources_experiment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Experiment.ProtoReflect.Descriptor instead.
func (*Experiment) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_resources_experiment_proto_rawDescGZIP(), []int{0}
}

func (x *Experiment) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *Experiment) GetExperimentId() int64 {
	if x != nil && x.ExperimentId != nil {
		return *x.ExperimentId
	}
	return 0
}

func (x *Experiment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Experiment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Experiment) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

func (x *Experiment) GetType() enums.ExperimentTypeEnum_ExperimentType {
	if x != nil {
		return x.Type
	}
	return enums.ExperimentTypeEnum_ExperimentType(0)
}

func (x *Experiment) GetStatus() enums.ExperimentStatusEnum_ExperimentStatus {
	if x != nil {
		return x.Status
	}
	return enums.ExperimentStatusEnum_ExperimentStatus(0)
}

func (x *Experiment) GetStartDate() string {
	if x != nil && x.StartDate != nil {
		return *x.StartDate
	}
	return ""
}

func (x *Experiment) GetEndDate() string {
	if x != nil && x.EndDate != nil {
		return *x.EndDate
	}
	return ""
}

func (x *Experiment) GetGoals() []*common.MetricGoal {
	if x != nil {
		return x.Goals
	}
	return nil
}

func (x *Experiment) GetLongRunningOperation() string {
	if x != nil && x.LongRunningOperation != nil {
		return *x.LongRunningOperation
	}
	return ""
}

func (x *Experiment) GetPromoteStatus() enums.AsyncActionStatusEnum_AsyncActionStatus {
	if x != nil {
		return x.PromoteStatus
	}
	return enums.AsyncActionStatusEnum_AsyncActionStatus(0)
}

var File_google_ads_googleads_v11_resources_experiment_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v11_resources_experiment_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf5, 0x06, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x50, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x25, 0x0a, 0x23,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52,
	0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66,
	0x66, 0x69, 0x78, 0x12, 0x55, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x5d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a,
	0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73,
	0x12, 0x3e, 0x0a, 0x16, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x14, 0x6c, 0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x73, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x41, 0x73, 0x79, 0x6e, 0x63, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x58, 0xea, 0x41, 0x55, 0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x7d, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x19, 0x0a,
	0x17, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x81, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x42, 0x0f, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x31, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v11_resources_experiment_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v11_resources_experiment_proto_rawDescData = file_google_ads_googleads_v11_resources_experiment_proto_rawDesc
)

func file_google_ads_googleads_v11_resources_experiment_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v11_resources_experiment_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v11_resources_experiment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v11_resources_experiment_proto_rawDescData)
	})
	return file_google_ads_googleads_v11_resources_experiment_proto_rawDescData
}

var file_google_ads_googleads_v11_resources_experiment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v11_resources_experiment_proto_goTypes = []interface{}{
	(*Experiment)(nil),                                 // 0: google.ads.googleads.v11.resources.Experiment
	(enums.ExperimentTypeEnum_ExperimentType)(0),       // 1: google.ads.googleads.v11.enums.ExperimentTypeEnum.ExperimentType
	(enums.ExperimentStatusEnum_ExperimentStatus)(0),   // 2: google.ads.googleads.v11.enums.ExperimentStatusEnum.ExperimentStatus
	(*common.MetricGoal)(nil),                          // 3: google.ads.googleads.v11.common.MetricGoal
	(enums.AsyncActionStatusEnum_AsyncActionStatus)(0), // 4: google.ads.googleads.v11.enums.AsyncActionStatusEnum.AsyncActionStatus
}
var file_google_ads_googleads_v11_resources_experiment_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v11.resources.Experiment.type:type_name -> google.ads.googleads.v11.enums.ExperimentTypeEnum.ExperimentType
	2, // 1: google.ads.googleads.v11.resources.Experiment.status:type_name -> google.ads.googleads.v11.enums.ExperimentStatusEnum.ExperimentStatus
	3, // 2: google.ads.googleads.v11.resources.Experiment.goals:type_name -> google.ads.googleads.v11.common.MetricGoal
	4, // 3: google.ads.googleads.v11.resources.Experiment.promote_status:type_name -> google.ads.googleads.v11.enums.AsyncActionStatusEnum.AsyncActionStatus
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v11_resources_experiment_proto_init() }
func file_google_ads_googleads_v11_resources_experiment_proto_init() {
	if File_google_ads_googleads_v11_resources_experiment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v11_resources_experiment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Experiment); i {
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
	file_google_ads_googleads_v11_resources_experiment_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v11_resources_experiment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v11_resources_experiment_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v11_resources_experiment_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v11_resources_experiment_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v11_resources_experiment_proto = out.File
	file_google_ads_googleads_v11_resources_experiment_proto_rawDesc = nil
	file_google_ads_googleads_v11_resources_experiment_proto_goTypes = nil
	file_google_ads_googleads_v11_resources_experiment_proto_depIdxs = nil
}
