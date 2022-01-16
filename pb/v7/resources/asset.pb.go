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
// source: google/ads/googleads/v7/resources/asset.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/dictav/go-genproto-googleads/pb/v7/common"
	enums "github.com/dictav/go-genproto-googleads/pb/v7/enums"
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

// Asset is a part of an ad which can be shared across multiple ads.
// It can be an image (ImageAsset), a video (YoutubeVideoAsset), etc.
// Assets are immutable and cannot be removed. To stop an asset from serving,
// remove the asset from the entity that is using it.
type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the asset.
	// Asset resource names have the form:
	//
	// `customers/{customer_id}/assets/{asset_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the asset.
	Id *int64 `protobuf:"varint,11,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Optional name of the asset.
	Name *string `protobuf:"bytes,12,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Output only. Type of the asset.
	Type enums.AssetTypeEnum_AssetType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v7.enums.AssetTypeEnum_AssetType" json:"type,omitempty"`
	// A list of possible final URLs after all cross domain redirects.
	FinalUrls []string `protobuf:"bytes,14,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// A list of possible final mobile URLs after all cross domain redirects.
	FinalMobileUrls []string `protobuf:"bytes,16,rep,name=final_mobile_urls,json=finalMobileUrls,proto3" json:"final_mobile_urls,omitempty"`
	// URL template for constructing a tracking URL.
	TrackingUrlTemplate *string `protobuf:"bytes,17,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3,oneof" json:"tracking_url_template,omitempty"`
	// A list of mappings to be used for substituting URL custom parameter tags in
	// the tracking_url_template, final_urls, and/or final_mobile_urls.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,18,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// URL template for appending params to landing page URLs served with parallel
	// tracking.
	FinalUrlSuffix *string `protobuf:"bytes,19,opt,name=final_url_suffix,json=finalUrlSuffix,proto3,oneof" json:"final_url_suffix,omitempty"`
	// Output only. Policy information for the asset.
	PolicySummary *AssetPolicySummary `protobuf:"bytes,13,opt,name=policy_summary,json=policySummary,proto3" json:"policy_summary,omitempty"`
	// The specific type of the asset.
	//
	// Types that are assignable to AssetData:
	//	*Asset_YoutubeVideoAsset
	//	*Asset_MediaBundleAsset
	//	*Asset_ImageAsset
	//	*Asset_TextAsset
	//	*Asset_LeadFormAsset
	//	*Asset_BookOnGoogleAsset
	//	*Asset_PromotionAsset
	//	*Asset_CalloutAsset
	//	*Asset_StructuredSnippetAsset
	//	*Asset_SitelinkAsset
	AssetData isAsset_AssetData `protobuf_oneof:"asset_data"`
}

func (x *Asset) Reset() {
	*x = Asset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_resources_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_resources_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_resources_asset_proto_rawDescGZIP(), []int{0}
}

func (x *Asset) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *Asset) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Asset) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Asset) GetType() enums.AssetTypeEnum_AssetType {
	if x != nil {
		return x.Type
	}
	return enums.AssetTypeEnum_UNSPECIFIED
}

func (x *Asset) GetFinalUrls() []string {
	if x != nil {
		return x.FinalUrls
	}
	return nil
}

func (x *Asset) GetFinalMobileUrls() []string {
	if x != nil {
		return x.FinalMobileUrls
	}
	return nil
}

func (x *Asset) GetTrackingUrlTemplate() string {
	if x != nil && x.TrackingUrlTemplate != nil {
		return *x.TrackingUrlTemplate
	}
	return ""
}

func (x *Asset) GetUrlCustomParameters() []*common.CustomParameter {
	if x != nil {
		return x.UrlCustomParameters
	}
	return nil
}

func (x *Asset) GetFinalUrlSuffix() string {
	if x != nil && x.FinalUrlSuffix != nil {
		return *x.FinalUrlSuffix
	}
	return ""
}

func (x *Asset) GetPolicySummary() *AssetPolicySummary {
	if x != nil {
		return x.PolicySummary
	}
	return nil
}

func (m *Asset) GetAssetData() isAsset_AssetData {
	if m != nil {
		return m.AssetData
	}
	return nil
}

func (x *Asset) GetYoutubeVideoAsset() *common.YoutubeVideoAsset {
	if x, ok := x.GetAssetData().(*Asset_YoutubeVideoAsset); ok {
		return x.YoutubeVideoAsset
	}
	return nil
}

func (x *Asset) GetMediaBundleAsset() *common.MediaBundleAsset {
	if x, ok := x.GetAssetData().(*Asset_MediaBundleAsset); ok {
		return x.MediaBundleAsset
	}
	return nil
}

func (x *Asset) GetImageAsset() *common.ImageAsset {
	if x, ok := x.GetAssetData().(*Asset_ImageAsset); ok {
		return x.ImageAsset
	}
	return nil
}

func (x *Asset) GetTextAsset() *common.TextAsset {
	if x, ok := x.GetAssetData().(*Asset_TextAsset); ok {
		return x.TextAsset
	}
	return nil
}

func (x *Asset) GetLeadFormAsset() *common.LeadFormAsset {
	if x, ok := x.GetAssetData().(*Asset_LeadFormAsset); ok {
		return x.LeadFormAsset
	}
	return nil
}

func (x *Asset) GetBookOnGoogleAsset() *common.BookOnGoogleAsset {
	if x, ok := x.GetAssetData().(*Asset_BookOnGoogleAsset); ok {
		return x.BookOnGoogleAsset
	}
	return nil
}

func (x *Asset) GetPromotionAsset() *common.PromotionAsset {
	if x, ok := x.GetAssetData().(*Asset_PromotionAsset); ok {
		return x.PromotionAsset
	}
	return nil
}

func (x *Asset) GetCalloutAsset() *common.CalloutAsset {
	if x, ok := x.GetAssetData().(*Asset_CalloutAsset); ok {
		return x.CalloutAsset
	}
	return nil
}

func (x *Asset) GetStructuredSnippetAsset() *common.StructuredSnippetAsset {
	if x, ok := x.GetAssetData().(*Asset_StructuredSnippetAsset); ok {
		return x.StructuredSnippetAsset
	}
	return nil
}

func (x *Asset) GetSitelinkAsset() *common.SitelinkAsset {
	if x, ok := x.GetAssetData().(*Asset_SitelinkAsset); ok {
		return x.SitelinkAsset
	}
	return nil
}

type isAsset_AssetData interface {
	isAsset_AssetData()
}

type Asset_YoutubeVideoAsset struct {
	// Immutable. A YouTube video asset.
	YoutubeVideoAsset *common.YoutubeVideoAsset `protobuf:"bytes,5,opt,name=youtube_video_asset,json=youtubeVideoAsset,proto3,oneof"`
}

type Asset_MediaBundleAsset struct {
	// Immutable. A media bundle asset.
	MediaBundleAsset *common.MediaBundleAsset `protobuf:"bytes,6,opt,name=media_bundle_asset,json=mediaBundleAsset,proto3,oneof"`
}

type Asset_ImageAsset struct {
	// Output only. An image asset.
	ImageAsset *common.ImageAsset `protobuf:"bytes,7,opt,name=image_asset,json=imageAsset,proto3,oneof"`
}

type Asset_TextAsset struct {
	// Output only. A text asset.
	TextAsset *common.TextAsset `protobuf:"bytes,8,opt,name=text_asset,json=textAsset,proto3,oneof"`
}

type Asset_LeadFormAsset struct {
	// A lead form asset.
	LeadFormAsset *common.LeadFormAsset `protobuf:"bytes,9,opt,name=lead_form_asset,json=leadFormAsset,proto3,oneof"`
}

type Asset_BookOnGoogleAsset struct {
	// A book on google asset.
	BookOnGoogleAsset *common.BookOnGoogleAsset `protobuf:"bytes,10,opt,name=book_on_google_asset,json=bookOnGoogleAsset,proto3,oneof"`
}

type Asset_PromotionAsset struct {
	// A promotion asset.
	PromotionAsset *common.PromotionAsset `protobuf:"bytes,15,opt,name=promotion_asset,json=promotionAsset,proto3,oneof"`
}

type Asset_CalloutAsset struct {
	// A callout asset.
	CalloutAsset *common.CalloutAsset `protobuf:"bytes,20,opt,name=callout_asset,json=calloutAsset,proto3,oneof"`
}

type Asset_StructuredSnippetAsset struct {
	// A structured snippet asset.
	StructuredSnippetAsset *common.StructuredSnippetAsset `protobuf:"bytes,21,opt,name=structured_snippet_asset,json=structuredSnippetAsset,proto3,oneof"`
}

type Asset_SitelinkAsset struct {
	// A sitelink asset.
	SitelinkAsset *common.SitelinkAsset `protobuf:"bytes,22,opt,name=sitelink_asset,json=sitelinkAsset,proto3,oneof"`
}

func (*Asset_YoutubeVideoAsset) isAsset_AssetData() {}

func (*Asset_MediaBundleAsset) isAsset_AssetData() {}

func (*Asset_ImageAsset) isAsset_AssetData() {}

func (*Asset_TextAsset) isAsset_AssetData() {}

func (*Asset_LeadFormAsset) isAsset_AssetData() {}

func (*Asset_BookOnGoogleAsset) isAsset_AssetData() {}

func (*Asset_PromotionAsset) isAsset_AssetData() {}

func (*Asset_CalloutAsset) isAsset_AssetData() {}

func (*Asset_StructuredSnippetAsset) isAsset_AssetData() {}

func (*Asset_SitelinkAsset) isAsset_AssetData() {}

// Contains policy information for an asset.
type AssetPolicySummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The list of policy findings for this asset.
	PolicyTopicEntries []*common.PolicyTopicEntry `protobuf:"bytes,1,rep,name=policy_topic_entries,json=policyTopicEntries,proto3" json:"policy_topic_entries,omitempty"`
	// Output only. Where in the review process this asset is.
	ReviewStatus enums.PolicyReviewStatusEnum_PolicyReviewStatus `protobuf:"varint,2,opt,name=review_status,json=reviewStatus,proto3,enum=google.ads.googleads.v7.enums.PolicyReviewStatusEnum_PolicyReviewStatus" json:"review_status,omitempty"`
	// Output only. The overall approval status of this asset, calculated based on the status
	// of its individual policy topic entries.
	ApprovalStatus enums.PolicyApprovalStatusEnum_PolicyApprovalStatus `protobuf:"varint,3,opt,name=approval_status,json=approvalStatus,proto3,enum=google.ads.googleads.v7.enums.PolicyApprovalStatusEnum_PolicyApprovalStatus" json:"approval_status,omitempty"`
}

func (x *AssetPolicySummary) Reset() {
	*x = AssetPolicySummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_resources_asset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetPolicySummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetPolicySummary) ProtoMessage() {}

func (x *AssetPolicySummary) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_resources_asset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetPolicySummary.ProtoReflect.Descriptor instead.
func (*AssetPolicySummary) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_resources_asset_proto_rawDescGZIP(), []int{1}
}

func (x *AssetPolicySummary) GetPolicyTopicEntries() []*common.PolicyTopicEntry {
	if x != nil {
		return x.PolicyTopicEntries
	}
	return nil
}

func (x *AssetPolicySummary) GetReviewStatus() enums.PolicyReviewStatusEnum_PolicyReviewStatus {
	if x != nil {
		return x.ReviewStatus
	}
	return enums.PolicyReviewStatusEnum_UNSPECIFIED
}

func (x *AssetPolicySummary) GetApprovalStatus() enums.PolicyApprovalStatusEnum_PolicyApprovalStatus {
	if x != nil {
		return x.ApprovalStatus
	}
	return enums.PolicyApprovalStatusEnum_UNSPECIFIED
}

var File_google_ads_googleads_v7_resources_asset_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v7_resources_asset_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x37, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x0d, 0x0a, 0x05, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x12, 0x4b, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xe0, 0x41, 0x05, 0xfa, 0x41,
	0x20, 0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x4f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c,
	0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72,
	0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x37,
	0x0a, 0x15, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x13, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x63, 0x0a, 0x15, 0x75, 0x72, 0x6c, 0x5f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x13, 0x75, 0x72, 0x6c, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x10,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x55,
	0x72, 0x6c, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x12, 0x61, 0x0a, 0x0e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x68,
	0x0a, 0x13, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x59, 0x6f, 0x75,
	0x74, 0x75, 0x62, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x03,
	0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x11, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x65, 0x0a, 0x12, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x10, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12,
	0x52, 0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x12, 0x4f, 0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x09, 0x74, 0x65, 0x78, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x12, 0x57, 0x0a, 0x0f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c,
	0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0d,
	0x6c, 0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x64, 0x0a,
	0x14, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6f, 0x6e, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x4f, 0x6e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x11, 0x62, 0x6f, 0x6f, 0x6b, 0x4f, 0x6e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x12, 0x59, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0e,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x53,
	0x0a, 0x0d, 0x63, 0x61, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x12, 0x72, 0x0a, 0x18, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x64, 0x5f, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x64, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52,
	0x16, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x53, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x56, 0x0a, 0x0e, 0x73, 0x69, 0x74, 0x65, 0x6c,
	0x69, 0x6e, 0x6b, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x53, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x0d, 0x73, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x73, 0x73, 0x65, 0x74, 0x3a,
	0x4e, 0xea, 0x41, 0x4b, 0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x12, 0x29, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42,
	0x0c, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x18, 0x0a,
	0x16, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x22, 0xed, 0x02, 0x0a,
	0x12, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x67, 0x0a, 0x14, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x72, 0x0a, 0x0d,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x7a, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0xf7, 0x01, 0x0a,
	0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0a, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x37,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x37, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x37, 0x3a, 0x3a, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v7_resources_asset_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v7_resources_asset_proto_rawDescData = file_google_ads_googleads_v7_resources_asset_proto_rawDesc
)

func file_google_ads_googleads_v7_resources_asset_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v7_resources_asset_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v7_resources_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v7_resources_asset_proto_rawDescData)
	})
	return file_google_ads_googleads_v7_resources_asset_proto_rawDescData
}

var file_google_ads_googleads_v7_resources_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v7_resources_asset_proto_goTypes = []interface{}{
	(*Asset)(nil),                                            // 0: google.ads.googleads.v7.resources.Asset
	(*AssetPolicySummary)(nil),                               // 1: google.ads.googleads.v7.resources.AssetPolicySummary
	(enums.AssetTypeEnum_AssetType)(0),                       // 2: google.ads.googleads.v7.enums.AssetTypeEnum.AssetType
	(*common.CustomParameter)(nil),                           // 3: google.ads.googleads.v7.common.CustomParameter
	(*common.YoutubeVideoAsset)(nil),                         // 4: google.ads.googleads.v7.common.YoutubeVideoAsset
	(*common.MediaBundleAsset)(nil),                          // 5: google.ads.googleads.v7.common.MediaBundleAsset
	(*common.ImageAsset)(nil),                                // 6: google.ads.googleads.v7.common.ImageAsset
	(*common.TextAsset)(nil),                                 // 7: google.ads.googleads.v7.common.TextAsset
	(*common.LeadFormAsset)(nil),                             // 8: google.ads.googleads.v7.common.LeadFormAsset
	(*common.BookOnGoogleAsset)(nil),                         // 9: google.ads.googleads.v7.common.BookOnGoogleAsset
	(*common.PromotionAsset)(nil),                            // 10: google.ads.googleads.v7.common.PromotionAsset
	(*common.CalloutAsset)(nil),                              // 11: google.ads.googleads.v7.common.CalloutAsset
	(*common.StructuredSnippetAsset)(nil),                    // 12: google.ads.googleads.v7.common.StructuredSnippetAsset
	(*common.SitelinkAsset)(nil),                             // 13: google.ads.googleads.v7.common.SitelinkAsset
	(*common.PolicyTopicEntry)(nil),                          // 14: google.ads.googleads.v7.common.PolicyTopicEntry
	(enums.PolicyReviewStatusEnum_PolicyReviewStatus)(0),     // 15: google.ads.googleads.v7.enums.PolicyReviewStatusEnum.PolicyReviewStatus
	(enums.PolicyApprovalStatusEnum_PolicyApprovalStatus)(0), // 16: google.ads.googleads.v7.enums.PolicyApprovalStatusEnum.PolicyApprovalStatus
}
var file_google_ads_googleads_v7_resources_asset_proto_depIdxs = []int32{
	2,  // 0: google.ads.googleads.v7.resources.Asset.type:type_name -> google.ads.googleads.v7.enums.AssetTypeEnum.AssetType
	3,  // 1: google.ads.googleads.v7.resources.Asset.url_custom_parameters:type_name -> google.ads.googleads.v7.common.CustomParameter
	1,  // 2: google.ads.googleads.v7.resources.Asset.policy_summary:type_name -> google.ads.googleads.v7.resources.AssetPolicySummary
	4,  // 3: google.ads.googleads.v7.resources.Asset.youtube_video_asset:type_name -> google.ads.googleads.v7.common.YoutubeVideoAsset
	5,  // 4: google.ads.googleads.v7.resources.Asset.media_bundle_asset:type_name -> google.ads.googleads.v7.common.MediaBundleAsset
	6,  // 5: google.ads.googleads.v7.resources.Asset.image_asset:type_name -> google.ads.googleads.v7.common.ImageAsset
	7,  // 6: google.ads.googleads.v7.resources.Asset.text_asset:type_name -> google.ads.googleads.v7.common.TextAsset
	8,  // 7: google.ads.googleads.v7.resources.Asset.lead_form_asset:type_name -> google.ads.googleads.v7.common.LeadFormAsset
	9,  // 8: google.ads.googleads.v7.resources.Asset.book_on_google_asset:type_name -> google.ads.googleads.v7.common.BookOnGoogleAsset
	10, // 9: google.ads.googleads.v7.resources.Asset.promotion_asset:type_name -> google.ads.googleads.v7.common.PromotionAsset
	11, // 10: google.ads.googleads.v7.resources.Asset.callout_asset:type_name -> google.ads.googleads.v7.common.CalloutAsset
	12, // 11: google.ads.googleads.v7.resources.Asset.structured_snippet_asset:type_name -> google.ads.googleads.v7.common.StructuredSnippetAsset
	13, // 12: google.ads.googleads.v7.resources.Asset.sitelink_asset:type_name -> google.ads.googleads.v7.common.SitelinkAsset
	14, // 13: google.ads.googleads.v7.resources.AssetPolicySummary.policy_topic_entries:type_name -> google.ads.googleads.v7.common.PolicyTopicEntry
	15, // 14: google.ads.googleads.v7.resources.AssetPolicySummary.review_status:type_name -> google.ads.googleads.v7.enums.PolicyReviewStatusEnum.PolicyReviewStatus
	16, // 15: google.ads.googleads.v7.resources.AssetPolicySummary.approval_status:type_name -> google.ads.googleads.v7.enums.PolicyApprovalStatusEnum.PolicyApprovalStatus
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v7_resources_asset_proto_init() }
func file_google_ads_googleads_v7_resources_asset_proto_init() {
	if File_google_ads_googleads_v7_resources_asset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v7_resources_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asset); i {
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
		file_google_ads_googleads_v7_resources_asset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetPolicySummary); i {
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
	file_google_ads_googleads_v7_resources_asset_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Asset_YoutubeVideoAsset)(nil),
		(*Asset_MediaBundleAsset)(nil),
		(*Asset_ImageAsset)(nil),
		(*Asset_TextAsset)(nil),
		(*Asset_LeadFormAsset)(nil),
		(*Asset_BookOnGoogleAsset)(nil),
		(*Asset_PromotionAsset)(nil),
		(*Asset_CalloutAsset)(nil),
		(*Asset_StructuredSnippetAsset)(nil),
		(*Asset_SitelinkAsset)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v7_resources_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v7_resources_asset_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v7_resources_asset_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v7_resources_asset_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v7_resources_asset_proto = out.File
	file_google_ads_googleads_v7_resources_asset_proto_rawDesc = nil
	file_google_ads_googleads_v7_resources_asset_proto_goTypes = nil
	file_google_ads_googleads_v7_resources_asset_proto_depIdxs = nil
}
