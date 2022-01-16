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
// source: google/ads/googleads/v8/enums/asset_performance_label.proto

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

// Enum describing the possible performance labels of an asset, usually
// computed in the context of a linkage.
type AssetPerformanceLabelEnum_AssetPerformanceLabel int32

const (
	// Not specified.
	AssetPerformanceLabelEnum_UNSPECIFIED AssetPerformanceLabelEnum_AssetPerformanceLabel = 0
	// Used for return value only. Represents value unknown in this version.
	AssetPerformanceLabelEnum_UNKNOWN AssetPerformanceLabelEnum_AssetPerformanceLabel = 1
	// This asset does not yet have any performance informantion. This may be
	// because it is still under review.
	AssetPerformanceLabelEnum_PENDING AssetPerformanceLabelEnum_AssetPerformanceLabel = 2
	// The asset has started getting impressions but the stats are not
	// statistically significant enough to get an asset performance label.
	AssetPerformanceLabelEnum_LEARNING AssetPerformanceLabelEnum_AssetPerformanceLabel = 3
	// Worst performing assets.
	AssetPerformanceLabelEnum_LOW AssetPerformanceLabelEnum_AssetPerformanceLabel = 4
	// Good performing assets.
	AssetPerformanceLabelEnum_GOOD AssetPerformanceLabelEnum_AssetPerformanceLabel = 5
	// Best performing assets.
	AssetPerformanceLabelEnum_BEST AssetPerformanceLabelEnum_AssetPerformanceLabel = 6
)

// Enum value maps for AssetPerformanceLabelEnum_AssetPerformanceLabel.
var (
	AssetPerformanceLabelEnum_AssetPerformanceLabel_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "PENDING",
		3: "LEARNING",
		4: "LOW",
		5: "GOOD",
		6: "BEST",
	}
	AssetPerformanceLabelEnum_AssetPerformanceLabel_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"PENDING":     2,
		"LEARNING":    3,
		"LOW":         4,
		"GOOD":        5,
		"BEST":        6,
	}
)

func (x AssetPerformanceLabelEnum_AssetPerformanceLabel) Enum() *AssetPerformanceLabelEnum_AssetPerformanceLabel {
	p := new(AssetPerformanceLabelEnum_AssetPerformanceLabel)
	*p = x
	return p
}

func (x AssetPerformanceLabelEnum_AssetPerformanceLabel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetPerformanceLabelEnum_AssetPerformanceLabel) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v8_enums_asset_performance_label_proto_enumTypes[0].Descriptor()
}

func (AssetPerformanceLabelEnum_AssetPerformanceLabel) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v8_enums_asset_performance_label_proto_enumTypes[0]
}

func (x AssetPerformanceLabelEnum_AssetPerformanceLabel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetPerformanceLabelEnum_AssetPerformanceLabel.Descriptor instead.
func (AssetPerformanceLabelEnum_AssetPerformanceLabel) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the performance label of an asset.
type AssetPerformanceLabelEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetPerformanceLabelEnum) Reset() {
	*x = AssetPerformanceLabelEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_enums_asset_performance_label_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetPerformanceLabelEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetPerformanceLabelEnum) ProtoMessage() {}

func (x *AssetPerformanceLabelEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_enums_asset_performance_label_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetPerformanceLabelEnum.ProtoReflect.Descriptor instead.
func (*AssetPerformanceLabelEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v8_enums_asset_performance_label_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x19, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x6d, 0x0a, 0x15, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x4c, 0x45, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f,
	0x57, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x4f, 0x4f, 0x44, 0x10, 0x05, 0x12, 0x08, 0x0a,
	0x04, 0x42, 0x45, 0x53, 0x54, 0x10, 0x06, 0x42, 0xef, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1a, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x38, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDescData = file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDesc
)

func file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDescData
}

var file_google_ads_googleads_v8_enums_asset_performance_label_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v8_enums_asset_performance_label_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v8_enums_asset_performance_label_proto_goTypes = []interface{}{
	(AssetPerformanceLabelEnum_AssetPerformanceLabel)(0), // 0: google.ads.googleads.v8.enums.AssetPerformanceLabelEnum.AssetPerformanceLabel
	(*AssetPerformanceLabelEnum)(nil),                    // 1: google.ads.googleads.v8.enums.AssetPerformanceLabelEnum
}
var file_google_ads_googleads_v8_enums_asset_performance_label_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_enums_asset_performance_label_proto_init() }
func file_google_ads_googleads_v8_enums_asset_performance_label_proto_init() {
	if File_google_ads_googleads_v8_enums_asset_performance_label_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_enums_asset_performance_label_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetPerformanceLabelEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_enums_asset_performance_label_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_enums_asset_performance_label_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v8_enums_asset_performance_label_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v8_enums_asset_performance_label_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_enums_asset_performance_label_proto = out.File
	file_google_ads_googleads_v8_enums_asset_performance_label_proto_rawDesc = nil
	file_google_ads_googleads_v8_enums_asset_performance_label_proto_goTypes = nil
	file_google_ads_googleads_v8_enums_asset_performance_label_proto_depIdxs = nil
}
