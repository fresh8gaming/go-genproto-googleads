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
// source: google/ads/googleads/v11/errors/ad_group_feed_error.proto

package errors

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

// Enum describing possible ad group feed errors.
type AdGroupFeedErrorEnum_AdGroupFeedError int32

const (
	// Enum unspecified.
	AdGroupFeedErrorEnum_UNSPECIFIED AdGroupFeedErrorEnum_AdGroupFeedError = 0
	// The received error code is not known in this version.
	AdGroupFeedErrorEnum_UNKNOWN AdGroupFeedErrorEnum_AdGroupFeedError = 1
	// An active feed already exists for this ad group and place holder type.
	AdGroupFeedErrorEnum_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE AdGroupFeedErrorEnum_AdGroupFeedError = 2
	// The specified feed is removed.
	AdGroupFeedErrorEnum_CANNOT_CREATE_FOR_REMOVED_FEED AdGroupFeedErrorEnum_AdGroupFeedError = 3
	// The AdGroupFeed already exists. UPDATE operation should be used to modify
	// the existing AdGroupFeed.
	AdGroupFeedErrorEnum_ADGROUP_FEED_ALREADY_EXISTS AdGroupFeedErrorEnum_AdGroupFeedError = 4
	// Cannot operate on removed AdGroupFeed.
	AdGroupFeedErrorEnum_CANNOT_OPERATE_ON_REMOVED_ADGROUP_FEED AdGroupFeedErrorEnum_AdGroupFeedError = 5
	// Invalid placeholder type.
	AdGroupFeedErrorEnum_INVALID_PLACEHOLDER_TYPE AdGroupFeedErrorEnum_AdGroupFeedError = 6
	// Feed mapping for this placeholder type does not exist.
	AdGroupFeedErrorEnum_MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE AdGroupFeedErrorEnum_AdGroupFeedError = 7
	// Location AdGroupFeeds cannot be created unless there is a location
	// CustomerFeed for the specified feed.
	AdGroupFeedErrorEnum_NO_EXISTING_LOCATION_CUSTOMER_FEED AdGroupFeedErrorEnum_AdGroupFeedError = 8
)

// Enum value maps for AdGroupFeedErrorEnum_AdGroupFeedError.
var (
	AdGroupFeedErrorEnum_AdGroupFeedError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
		3: "CANNOT_CREATE_FOR_REMOVED_FEED",
		4: "ADGROUP_FEED_ALREADY_EXISTS",
		5: "CANNOT_OPERATE_ON_REMOVED_ADGROUP_FEED",
		6: "INVALID_PLACEHOLDER_TYPE",
		7: "MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE",
		8: "NO_EXISTING_LOCATION_CUSTOMER_FEED",
	}
	AdGroupFeedErrorEnum_AdGroupFeedError_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE": 2,
		"CANNOT_CREATE_FOR_REMOVED_FEED":           3,
		"ADGROUP_FEED_ALREADY_EXISTS":              4,
		"CANNOT_OPERATE_ON_REMOVED_ADGROUP_FEED":   5,
		"INVALID_PLACEHOLDER_TYPE":                 6,
		"MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE": 7,
		"NO_EXISTING_LOCATION_CUSTOMER_FEED":       8,
	}
)

func (x AdGroupFeedErrorEnum_AdGroupFeedError) Enum() *AdGroupFeedErrorEnum_AdGroupFeedError {
	p := new(AdGroupFeedErrorEnum_AdGroupFeedError)
	*p = x
	return p
}

func (x AdGroupFeedErrorEnum_AdGroupFeedError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdGroupFeedErrorEnum_AdGroupFeedError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_enumTypes[0].Descriptor()
}

func (AdGroupFeedErrorEnum_AdGroupFeedError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_enumTypes[0]
}

func (x AdGroupFeedErrorEnum_AdGroupFeedError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdGroupFeedErrorEnum_AdGroupFeedError.Descriptor instead.
func (AdGroupFeedErrorEnum_AdGroupFeedError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible ad group feed errors.
type AdGroupFeedErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdGroupFeedErrorEnum) Reset() {
	*x = AdGroupFeedErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupFeedErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupFeedErrorEnum) ProtoMessage() {}

func (x *AdGroupFeedErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupFeedErrorEnum.ProtoReflect.Descriptor instead.
func (*AdGroupFeedErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v11_errors_ad_group_feed_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xdc, 0x02, 0x0a,
	0x14, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xc3, 0x02, 0x0a, 0x10, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x46, 0x65, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x2c, 0x0a, 0x28, 0x46, 0x45, 0x45, 0x44,
	0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x5f,
	0x46, 0x4f, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x4d, 0x4f,
	0x56, 0x45, 0x44, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x44,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41,
	0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x04, 0x12, 0x2a, 0x0a, 0x26, 0x43,
	0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e,
	0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x10, 0x06, 0x12, 0x2c, 0x0a, 0x28, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47,
	0x5f, 0x46, 0x45, 0x45, 0x44, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x07, 0x12, 0x26, 0x0a, 0x22, 0x4e, 0x4f, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x55, 0x53, 0x54,
	0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x08, 0x42, 0xf5, 0x01, 0x0a, 0x23,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x42, 0x15, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDescData = file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDesc
)

func file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDescData
}

var file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_goTypes = []interface{}{
	(AdGroupFeedErrorEnum_AdGroupFeedError)(0), // 0: google.ads.googleads.v11.errors.AdGroupFeedErrorEnum.AdGroupFeedError
	(*AdGroupFeedErrorEnum)(nil),               // 1: google.ads.googleads.v11.errors.AdGroupFeedErrorEnum
}
var file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_init() }
func file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_init() {
	if File_google_ads_googleads_v11_errors_ad_group_feed_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupFeedErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v11_errors_ad_group_feed_error_proto = out.File
	file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_rawDesc = nil
	file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_goTypes = nil
	file_google_ads_googleads_v11_errors_ad_group_feed_error_proto_depIdxs = nil
}
