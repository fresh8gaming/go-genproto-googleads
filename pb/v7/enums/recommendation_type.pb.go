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
// source: google/ads/googleads/v7/enums/recommendation_type.proto

package enums

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Types of recommendations.
type RecommendationTypeEnum_RecommendationType int32

const (
	// Not specified.
	RecommendationTypeEnum_UNSPECIFIED RecommendationTypeEnum_RecommendationType = 0
	// Used for return value only. Represents value unknown in this version.
	RecommendationTypeEnum_UNKNOWN RecommendationTypeEnum_RecommendationType = 1
	// Budget recommendation for campaigns that are currently budget-constrained
	// (as opposed to the FORECASTING_CAMPAIGN_BUDGET recommendation, which
	// applies to campaigns that are expected to become budget-constrained in
	// the future).
	RecommendationTypeEnum_CAMPAIGN_BUDGET RecommendationTypeEnum_RecommendationType = 2
	// Keyword recommendation.
	RecommendationTypeEnum_KEYWORD RecommendationTypeEnum_RecommendationType = 3
	// Recommendation to add a new text ad.
	RecommendationTypeEnum_TEXT_AD RecommendationTypeEnum_RecommendationType = 4
	// Recommendation to update a campaign to use a Target CPA bidding strategy.
	RecommendationTypeEnum_TARGET_CPA_OPT_IN RecommendationTypeEnum_RecommendationType = 5
	// Recommendation to update a campaign to use the Maximize Conversions
	// bidding strategy.
	RecommendationTypeEnum_MAXIMIZE_CONVERSIONS_OPT_IN RecommendationTypeEnum_RecommendationType = 6
	// Recommendation to enable Enhanced Cost Per Click for a campaign.
	RecommendationTypeEnum_ENHANCED_CPC_OPT_IN RecommendationTypeEnum_RecommendationType = 7
	// Recommendation to start showing your campaign's ads on Google Search
	// Partners Websites.
	RecommendationTypeEnum_SEARCH_PARTNERS_OPT_IN RecommendationTypeEnum_RecommendationType = 8
	// Recommendation to update a campaign to use a Maximize Clicks bidding
	// strategy.
	RecommendationTypeEnum_MAXIMIZE_CLICKS_OPT_IN RecommendationTypeEnum_RecommendationType = 9
	// Recommendation to start using the "Optimize" ad rotation setting for the
	// given ad group.
	RecommendationTypeEnum_OPTIMIZE_AD_ROTATION RecommendationTypeEnum_RecommendationType = 10
	// Recommendation to add callout extensions to a campaign.
	RecommendationTypeEnum_CALLOUT_EXTENSION RecommendationTypeEnum_RecommendationType = 11
	// Recommendation to add sitelink extensions to a campaign.
	RecommendationTypeEnum_SITELINK_EXTENSION RecommendationTypeEnum_RecommendationType = 12
	// Recommendation to add call extensions to a campaign.
	RecommendationTypeEnum_CALL_EXTENSION RecommendationTypeEnum_RecommendationType = 13
	// Recommendation to change an existing keyword from one match type to a
	// broader match type.
	RecommendationTypeEnum_KEYWORD_MATCH_TYPE RecommendationTypeEnum_RecommendationType = 14
	// Recommendation to move unused budget from one budget to a constrained
	// budget.
	RecommendationTypeEnum_MOVE_UNUSED_BUDGET RecommendationTypeEnum_RecommendationType = 15
	// Budget recommendation for campaigns that are expected to become
	// budget-constrained in the future (as opposed to the CAMPAIGN_BUDGET
	// recommendation, which applies to campaigns that are currently
	// budget-constrained).
	RecommendationTypeEnum_FORECASTING_CAMPAIGN_BUDGET RecommendationTypeEnum_RecommendationType = 16
	// Recommendation to update a campaign to use a Target ROAS bidding
	// strategy.
	RecommendationTypeEnum_TARGET_ROAS_OPT_IN RecommendationTypeEnum_RecommendationType = 17
	// Recommendation to add a new responsive search ad.
	RecommendationTypeEnum_RESPONSIVE_SEARCH_AD RecommendationTypeEnum_RecommendationType = 18
	// Budget recommendation for campaigns whose ROI is predicted to increase
	// with a budget adjustment.
	RecommendationTypeEnum_MARGINAL_ROI_CAMPAIGN_BUDGET RecommendationTypeEnum_RecommendationType = 19
)

// Enum value maps for RecommendationTypeEnum_RecommendationType.
var (
	RecommendationTypeEnum_RecommendationType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CAMPAIGN_BUDGET",
		3:  "KEYWORD",
		4:  "TEXT_AD",
		5:  "TARGET_CPA_OPT_IN",
		6:  "MAXIMIZE_CONVERSIONS_OPT_IN",
		7:  "ENHANCED_CPC_OPT_IN",
		8:  "SEARCH_PARTNERS_OPT_IN",
		9:  "MAXIMIZE_CLICKS_OPT_IN",
		10: "OPTIMIZE_AD_ROTATION",
		11: "CALLOUT_EXTENSION",
		12: "SITELINK_EXTENSION",
		13: "CALL_EXTENSION",
		14: "KEYWORD_MATCH_TYPE",
		15: "MOVE_UNUSED_BUDGET",
		16: "FORECASTING_CAMPAIGN_BUDGET",
		17: "TARGET_ROAS_OPT_IN",
		18: "RESPONSIVE_SEARCH_AD",
		19: "MARGINAL_ROI_CAMPAIGN_BUDGET",
	}
	RecommendationTypeEnum_RecommendationType_value = map[string]int32{
		"UNSPECIFIED":                  0,
		"UNKNOWN":                      1,
		"CAMPAIGN_BUDGET":              2,
		"KEYWORD":                      3,
		"TEXT_AD":                      4,
		"TARGET_CPA_OPT_IN":            5,
		"MAXIMIZE_CONVERSIONS_OPT_IN":  6,
		"ENHANCED_CPC_OPT_IN":          7,
		"SEARCH_PARTNERS_OPT_IN":       8,
		"MAXIMIZE_CLICKS_OPT_IN":       9,
		"OPTIMIZE_AD_ROTATION":         10,
		"CALLOUT_EXTENSION":            11,
		"SITELINK_EXTENSION":           12,
		"CALL_EXTENSION":               13,
		"KEYWORD_MATCH_TYPE":           14,
		"MOVE_UNUSED_BUDGET":           15,
		"FORECASTING_CAMPAIGN_BUDGET":  16,
		"TARGET_ROAS_OPT_IN":           17,
		"RESPONSIVE_SEARCH_AD":         18,
		"MARGINAL_ROI_CAMPAIGN_BUDGET": 19,
	}
)

func (x RecommendationTypeEnum_RecommendationType) Enum() *RecommendationTypeEnum_RecommendationType {
	p := new(RecommendationTypeEnum_RecommendationType)
	*p = x
	return p
}

func (x RecommendationTypeEnum_RecommendationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecommendationTypeEnum_RecommendationType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v7_enums_recommendation_type_proto_enumTypes[0].Descriptor()
}

func (RecommendationTypeEnum_RecommendationType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v7_enums_recommendation_type_proto_enumTypes[0]
}

func (x RecommendationTypeEnum_RecommendationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecommendationTypeEnum_RecommendationType.Descriptor instead.
func (RecommendationTypeEnum_RecommendationType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing types of recommendations.
type RecommendationTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RecommendationTypeEnum) Reset() {
	*x = RecommendationTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_enums_recommendation_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendationTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationTypeEnum) ProtoMessage() {}

func (x *RecommendationTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_enums_recommendation_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationTypeEnum.ProtoReflect.Descriptor instead.
func (*RecommendationTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v7_enums_recommendation_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x04, 0x0a, 0x16, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0xec, 0x03, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49,
	0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4b,
	0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x58, 0x54,
	0x5f, 0x41, 0x44, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f,
	0x43, 0x50, 0x41, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b,
	0x4d, 0x41, 0x58, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x06, 0x12, 0x17, 0x0a,
	0x13, 0x45, 0x4e, 0x48, 0x41, 0x4e, 0x43, 0x45, 0x44, 0x5f, 0x43, 0x50, 0x43, 0x5f, 0x4f, 0x50,
	0x54, 0x5f, 0x49, 0x4e, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48,
	0x5f, 0x50, 0x41, 0x52, 0x54, 0x4e, 0x45, 0x52, 0x53, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49, 0x4e,
	0x10, 0x08, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x41, 0x58, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x43,
	0x4c, 0x49, 0x43, 0x4b, 0x53, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x09, 0x12, 0x18,
	0x0a, 0x14, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x41, 0x44, 0x5f, 0x52, 0x4f,
	0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x41, 0x4c, 0x4c,
	0x4f, 0x55, 0x54, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x0b, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x49, 0x54, 0x45, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x45, 0x58, 0x54, 0x45,
	0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x0c, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x16, 0x0a, 0x12, 0x4b,
	0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x0e, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x55, 0x4e, 0x55, 0x53,
	0x45, 0x44, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x0f, 0x12, 0x1f, 0x0a, 0x1b, 0x46,
	0x4f, 0x52, 0x45, 0x43, 0x41, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41,
	0x49, 0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x10, 0x12, 0x16, 0x0a, 0x12,
	0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x4f, 0x41, 0x53, 0x5f, 0x4f, 0x50, 0x54, 0x5f,
	0x49, 0x4e, 0x10, 0x11, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x49,
	0x56, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x10, 0x12, 0x12, 0x20,
	0x0a, 0x1c, 0x4d, 0x41, 0x52, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x52, 0x4f, 0x49, 0x5f, 0x43,
	0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x13,
	0x42, 0xec, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x17, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x37, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x37, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x37, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDescData = file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDesc
)

func file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDescData
}

var file_google_ads_googleads_v7_enums_recommendation_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v7_enums_recommendation_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v7_enums_recommendation_type_proto_goTypes = []interface{}{
	(RecommendationTypeEnum_RecommendationType)(0), // 0: google.ads.googleads.v7.enums.RecommendationTypeEnum.RecommendationType
	(*RecommendationTypeEnum)(nil),                 // 1: google.ads.googleads.v7.enums.RecommendationTypeEnum
}
var file_google_ads_googleads_v7_enums_recommendation_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v7_enums_recommendation_type_proto_init() }
func file_google_ads_googleads_v7_enums_recommendation_type_proto_init() {
	if File_google_ads_googleads_v7_enums_recommendation_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v7_enums_recommendation_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendationTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v7_enums_recommendation_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v7_enums_recommendation_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v7_enums_recommendation_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v7_enums_recommendation_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v7_enums_recommendation_type_proto = out.File
	file_google_ads_googleads_v7_enums_recommendation_type_proto_rawDesc = nil
	file_google_ads_googleads_v7_enums_recommendation_type_proto_goTypes = nil
	file_google_ads_googleads_v7_enums_recommendation_type_proto_depIdxs = nil
}
