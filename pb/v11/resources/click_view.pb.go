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
// source: google/ads/googleads/v11/resources/click_view.proto

package resources

import (
	common "github.com/fresh8/go-genproto-googleads/pb/v11/common"
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

// A click view with metrics aggregated at each click level, including both
// valid and invalid clicks. For non-Search campaigns, metrics.clicks
// represents the number of valid and invalid interactions.
// Queries including ClickView must have a filter limiting the results to one
// day and can be requested for dates back to 90 days before the time of the
// request.
type ClickView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the click view.
	// Click view resource names have the form:
	//
	// `customers/{customer_id}/clickViews/{date (yyyy-MM-dd)}~{gclid}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The Google Click ID.
	Gclid *string `protobuf:"bytes,8,opt,name=gclid,proto3,oneof" json:"gclid,omitempty"`
	// Output only. The location criteria matching the area of interest associated with the
	// impression.
	AreaOfInterest *common.ClickLocation `protobuf:"bytes,3,opt,name=area_of_interest,json=areaOfInterest,proto3" json:"area_of_interest,omitempty"`
	// Output only. The location criteria matching the location of presence associated with the
	// impression.
	LocationOfPresence *common.ClickLocation `protobuf:"bytes,4,opt,name=location_of_presence,json=locationOfPresence,proto3" json:"location_of_presence,omitempty"`
	// Output only. Page number in search results where the ad was shown.
	PageNumber *int64 `protobuf:"varint,9,opt,name=page_number,json=pageNumber,proto3,oneof" json:"page_number,omitempty"`
	// Output only. The associated ad.
	AdGroupAd *string `protobuf:"bytes,10,opt,name=ad_group_ad,json=adGroupAd,proto3,oneof" json:"ad_group_ad,omitempty"`
	// Output only. The associated campaign location target, if one exists.
	CampaignLocationTarget *string `protobuf:"bytes,11,opt,name=campaign_location_target,json=campaignLocationTarget,proto3,oneof" json:"campaign_location_target,omitempty"`
	// Output only. The associated user list, if one exists.
	UserList *string `protobuf:"bytes,12,opt,name=user_list,json=userList,proto3,oneof" json:"user_list,omitempty"`
	// Output only. The associated keyword, if one exists and the click corresponds to the
	// SEARCH channel.
	Keyword string `protobuf:"bytes,13,opt,name=keyword,proto3" json:"keyword,omitempty"`
	// Output only. Basic information about the associated keyword, if it exists.
	KeywordInfo *common.KeywordInfo `protobuf:"bytes,14,opt,name=keyword_info,json=keywordInfo,proto3" json:"keyword_info,omitempty"`
}

func (x *ClickView) Reset() {
	*x = ClickView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_resources_click_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickView) ProtoMessage() {}

func (x *ClickView) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_resources_click_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickView.ProtoReflect.Descriptor instead.
func (*ClickView) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_resources_click_view_proto_rawDescGZIP(), []int{0}
}

func (x *ClickView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ClickView) GetGclid() string {
	if x != nil && x.Gclid != nil {
		return *x.Gclid
	}
	return ""
}

func (x *ClickView) GetAreaOfInterest() *common.ClickLocation {
	if x != nil {
		return x.AreaOfInterest
	}
	return nil
}

func (x *ClickView) GetLocationOfPresence() *common.ClickLocation {
	if x != nil {
		return x.LocationOfPresence
	}
	return nil
}

func (x *ClickView) GetPageNumber() int64 {
	if x != nil && x.PageNumber != nil {
		return *x.PageNumber
	}
	return 0
}

func (x *ClickView) GetAdGroupAd() string {
	if x != nil && x.AdGroupAd != nil {
		return *x.AdGroupAd
	}
	return ""
}

func (x *ClickView) GetCampaignLocationTarget() string {
	if x != nil && x.CampaignLocationTarget != nil {
		return *x.CampaignLocationTarget
	}
	return ""
}

func (x *ClickView) GetUserList() string {
	if x != nil && x.UserList != nil {
		return *x.UserList
	}
	return ""
}

func (x *ClickView) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ClickView) GetKeywordInfo() *common.KeywordInfo {
	if x != nil {
		return x.KeywordInfo
	}
	return nil
}

var File_google_ads_googleads_v11_resources_click_view_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v11_resources_click_view_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x07, 0x0a, 0x09,
	0x43, 0x6c, 0x69, 0x63, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x12, 0x4f, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2a, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x67, 0x63,
	0x6c, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00,
	0x52, 0x05, 0x67, 0x63, 0x6c, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x5d, 0x0a, 0x10, 0x61, 0x72,
	0x65, 0x61, 0x5f, 0x6f, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x61, 0x72, 0x65, 0x61, 0x4f,
	0x66, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x65, 0x0a, 0x14, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x29, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a, 0x0b, 0x61,
	0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2a, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x48, 0x02, 0x52, 0x09,
	0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x88, 0x01, 0x01, 0x12, 0x71, 0x0a, 0x18,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32,
	0xe0, 0x41, 0x03, 0xfa, 0x41, 0x2c, 0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x48, 0x03, 0x52, 0x16, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x4b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x04, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x07,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xe0,
	0x41, 0x03, 0xfa, 0x41, 0x2b, 0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x54, 0x0a, 0x0c, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x3a,
	0x5a, 0xea, 0x41, 0x57, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x12, 0x31, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x7b, 0x64, 0x61,
	0x74, 0x65, 0x7d, 0x7e, 0x7b, 0x67, 0x63, 0x6c, 0x69, 0x64, 0x7d, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x67, 0x63, 0x6c, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x61, 0x64, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x42, 0x80, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0e, 0x43, 0x6c, 0x69,
	0x63, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x31,
	0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v11_resources_click_view_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v11_resources_click_view_proto_rawDescData = file_google_ads_googleads_v11_resources_click_view_proto_rawDesc
)

func file_google_ads_googleads_v11_resources_click_view_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v11_resources_click_view_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v11_resources_click_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v11_resources_click_view_proto_rawDescData)
	})
	return file_google_ads_googleads_v11_resources_click_view_proto_rawDescData
}

var file_google_ads_googleads_v11_resources_click_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v11_resources_click_view_proto_goTypes = []interface{}{
	(*ClickView)(nil),            // 0: google.ads.googleads.v11.resources.ClickView
	(*common.ClickLocation)(nil), // 1: google.ads.googleads.v11.common.ClickLocation
	(*common.KeywordInfo)(nil),   // 2: google.ads.googleads.v11.common.KeywordInfo
}
var file_google_ads_googleads_v11_resources_click_view_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v11.resources.ClickView.area_of_interest:type_name -> google.ads.googleads.v11.common.ClickLocation
	1, // 1: google.ads.googleads.v11.resources.ClickView.location_of_presence:type_name -> google.ads.googleads.v11.common.ClickLocation
	2, // 2: google.ads.googleads.v11.resources.ClickView.keyword_info:type_name -> google.ads.googleads.v11.common.KeywordInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v11_resources_click_view_proto_init() }
func file_google_ads_googleads_v11_resources_click_view_proto_init() {
	if File_google_ads_googleads_v11_resources_click_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v11_resources_click_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickView); i {
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
	file_google_ads_googleads_v11_resources_click_view_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v11_resources_click_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v11_resources_click_view_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v11_resources_click_view_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v11_resources_click_view_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v11_resources_click_view_proto = out.File
	file_google_ads_googleads_v11_resources_click_view_proto_rawDesc = nil
	file_google_ads_googleads_v11_resources_click_view_proto_goTypes = nil
	file_google_ads_googleads_v11_resources_click_view_proto_depIdxs = nil
}
