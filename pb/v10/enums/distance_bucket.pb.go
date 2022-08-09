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
// source: google/ads/googleads/v10/enums/distance_bucket.proto

package enums

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

// The distance bucket for a user's distance from an advertiser's location
// extension.
type DistanceBucketEnum_DistanceBucket int32

const (
	// Not specified.
	DistanceBucketEnum_UNSPECIFIED DistanceBucketEnum_DistanceBucket = 0
	// Used for return value only. Represents value unknown in this version.
	DistanceBucketEnum_UNKNOWN DistanceBucketEnum_DistanceBucket = 1
	// User was within 700m of the location.
	DistanceBucketEnum_WITHIN_700M DistanceBucketEnum_DistanceBucket = 2
	// User was within 1KM of the location.
	DistanceBucketEnum_WITHIN_1KM DistanceBucketEnum_DistanceBucket = 3
	// User was within 5KM of the location.
	DistanceBucketEnum_WITHIN_5KM DistanceBucketEnum_DistanceBucket = 4
	// User was within 10KM of the location.
	DistanceBucketEnum_WITHIN_10KM DistanceBucketEnum_DistanceBucket = 5
	// User was within 15KM of the location.
	DistanceBucketEnum_WITHIN_15KM DistanceBucketEnum_DistanceBucket = 6
	// User was within 20KM of the location.
	DistanceBucketEnum_WITHIN_20KM DistanceBucketEnum_DistanceBucket = 7
	// User was within 25KM of the location.
	DistanceBucketEnum_WITHIN_25KM DistanceBucketEnum_DistanceBucket = 8
	// User was within 30KM of the location.
	DistanceBucketEnum_WITHIN_30KM DistanceBucketEnum_DistanceBucket = 9
	// User was within 35KM of the location.
	DistanceBucketEnum_WITHIN_35KM DistanceBucketEnum_DistanceBucket = 10
	// User was within 40KM of the location.
	DistanceBucketEnum_WITHIN_40KM DistanceBucketEnum_DistanceBucket = 11
	// User was within 45KM of the location.
	DistanceBucketEnum_WITHIN_45KM DistanceBucketEnum_DistanceBucket = 12
	// User was within 50KM of the location.
	DistanceBucketEnum_WITHIN_50KM DistanceBucketEnum_DistanceBucket = 13
	// User was within 55KM of the location.
	DistanceBucketEnum_WITHIN_55KM DistanceBucketEnum_DistanceBucket = 14
	// User was within 60KM of the location.
	DistanceBucketEnum_WITHIN_60KM DistanceBucketEnum_DistanceBucket = 15
	// User was within 65KM of the location.
	DistanceBucketEnum_WITHIN_65KM DistanceBucketEnum_DistanceBucket = 16
	// User was beyond 65KM of the location.
	DistanceBucketEnum_BEYOND_65KM DistanceBucketEnum_DistanceBucket = 17
	// User was within 0.7 miles of the location.
	DistanceBucketEnum_WITHIN_0_7MILES DistanceBucketEnum_DistanceBucket = 18
	// User was within 1 mile of the location.
	DistanceBucketEnum_WITHIN_1MILE DistanceBucketEnum_DistanceBucket = 19
	// User was within 5 miles of the location.
	DistanceBucketEnum_WITHIN_5MILES DistanceBucketEnum_DistanceBucket = 20
	// User was within 10 miles of the location.
	DistanceBucketEnum_WITHIN_10MILES DistanceBucketEnum_DistanceBucket = 21
	// User was within 15 miles of the location.
	DistanceBucketEnum_WITHIN_15MILES DistanceBucketEnum_DistanceBucket = 22
	// User was within 20 miles of the location.
	DistanceBucketEnum_WITHIN_20MILES DistanceBucketEnum_DistanceBucket = 23
	// User was within 25 miles of the location.
	DistanceBucketEnum_WITHIN_25MILES DistanceBucketEnum_DistanceBucket = 24
	// User was within 30 miles of the location.
	DistanceBucketEnum_WITHIN_30MILES DistanceBucketEnum_DistanceBucket = 25
	// User was within 35 miles of the location.
	DistanceBucketEnum_WITHIN_35MILES DistanceBucketEnum_DistanceBucket = 26
	// User was within 40 miles of the location.
	DistanceBucketEnum_WITHIN_40MILES DistanceBucketEnum_DistanceBucket = 27
	// User was beyond 40 miles of the location.
	DistanceBucketEnum_BEYOND_40MILES DistanceBucketEnum_DistanceBucket = 28
)

// Enum value maps for DistanceBucketEnum_DistanceBucket.
var (
	DistanceBucketEnum_DistanceBucket_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "WITHIN_700M",
		3:  "WITHIN_1KM",
		4:  "WITHIN_5KM",
		5:  "WITHIN_10KM",
		6:  "WITHIN_15KM",
		7:  "WITHIN_20KM",
		8:  "WITHIN_25KM",
		9:  "WITHIN_30KM",
		10: "WITHIN_35KM",
		11: "WITHIN_40KM",
		12: "WITHIN_45KM",
		13: "WITHIN_50KM",
		14: "WITHIN_55KM",
		15: "WITHIN_60KM",
		16: "WITHIN_65KM",
		17: "BEYOND_65KM",
		18: "WITHIN_0_7MILES",
		19: "WITHIN_1MILE",
		20: "WITHIN_5MILES",
		21: "WITHIN_10MILES",
		22: "WITHIN_15MILES",
		23: "WITHIN_20MILES",
		24: "WITHIN_25MILES",
		25: "WITHIN_30MILES",
		26: "WITHIN_35MILES",
		27: "WITHIN_40MILES",
		28: "BEYOND_40MILES",
	}
	DistanceBucketEnum_DistanceBucket_value = map[string]int32{
		"UNSPECIFIED":     0,
		"UNKNOWN":         1,
		"WITHIN_700M":     2,
		"WITHIN_1KM":      3,
		"WITHIN_5KM":      4,
		"WITHIN_10KM":     5,
		"WITHIN_15KM":     6,
		"WITHIN_20KM":     7,
		"WITHIN_25KM":     8,
		"WITHIN_30KM":     9,
		"WITHIN_35KM":     10,
		"WITHIN_40KM":     11,
		"WITHIN_45KM":     12,
		"WITHIN_50KM":     13,
		"WITHIN_55KM":     14,
		"WITHIN_60KM":     15,
		"WITHIN_65KM":     16,
		"BEYOND_65KM":     17,
		"WITHIN_0_7MILES": 18,
		"WITHIN_1MILE":    19,
		"WITHIN_5MILES":   20,
		"WITHIN_10MILES":  21,
		"WITHIN_15MILES":  22,
		"WITHIN_20MILES":  23,
		"WITHIN_25MILES":  24,
		"WITHIN_30MILES":  25,
		"WITHIN_35MILES":  26,
		"WITHIN_40MILES":  27,
		"BEYOND_40MILES":  28,
	}
)

func (x DistanceBucketEnum_DistanceBucket) Enum() *DistanceBucketEnum_DistanceBucket {
	p := new(DistanceBucketEnum_DistanceBucket)
	*p = x
	return p
}

func (x DistanceBucketEnum_DistanceBucket) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DistanceBucketEnum_DistanceBucket) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_enums_distance_bucket_proto_enumTypes[0].Descriptor()
}

func (DistanceBucketEnum_DistanceBucket) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_enums_distance_bucket_proto_enumTypes[0]
}

func (x DistanceBucketEnum_DistanceBucket) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DistanceBucketEnum_DistanceBucket.Descriptor instead.
func (DistanceBucketEnum_DistanceBucket) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDescGZIP(), []int{0, 0}
}

// Container for distance buckets of a user's distance from an advertiser's
// location extension.
type DistanceBucketEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DistanceBucketEnum) Reset() {
	*x = DistanceBucketEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_enums_distance_bucket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistanceBucketEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistanceBucketEnum) ProtoMessage() {}

func (x *DistanceBucketEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_enums_distance_bucket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistanceBucketEnum.ProtoReflect.Descriptor instead.
func (*DistanceBucketEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_enums_distance_bucket_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xad, 0x04, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x96, 0x04,
	0x0a, 0x0e, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x37, 0x30, 0x30, 0x4d, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x31, 0x4b, 0x4d, 0x10, 0x03, 0x12,
	0x0e, 0x0a, 0x0a, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x35, 0x4b, 0x4d, 0x10, 0x04, 0x12,
	0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x31, 0x30, 0x4b, 0x4d, 0x10, 0x05,
	0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x31, 0x35, 0x4b, 0x4d, 0x10,
	0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x32, 0x30, 0x4b, 0x4d,
	0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x32, 0x35, 0x4b,
	0x4d, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x33, 0x30,
	0x4b, 0x4d, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x33,
	0x35, 0x4b, 0x4d, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f,
	0x34, 0x30, 0x4b, 0x4d, 0x10, 0x0b, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e,
	0x5f, 0x34, 0x35, 0x4b, 0x4d, 0x10, 0x0c, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x54, 0x48, 0x49,
	0x4e, 0x5f, 0x35, 0x30, 0x4b, 0x4d, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x54, 0x48,
	0x49, 0x4e, 0x5f, 0x35, 0x35, 0x4b, 0x4d, 0x10, 0x0e, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49, 0x54,
	0x48, 0x49, 0x4e, 0x5f, 0x36, 0x30, 0x4b, 0x4d, 0x10, 0x0f, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x49,
	0x54, 0x48, 0x49, 0x4e, 0x5f, 0x36, 0x35, 0x4b, 0x4d, 0x10, 0x10, 0x12, 0x0f, 0x0a, 0x0b, 0x42,
	0x45, 0x59, 0x4f, 0x4e, 0x44, 0x5f, 0x36, 0x35, 0x4b, 0x4d, 0x10, 0x11, 0x12, 0x13, 0x0a, 0x0f,
	0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x30, 0x5f, 0x37, 0x4d, 0x49, 0x4c, 0x45, 0x53, 0x10,
	0x12, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x31, 0x4d, 0x49, 0x4c,
	0x45, 0x10, 0x13, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x35, 0x4d,
	0x49, 0x4c, 0x45, 0x53, 0x10, 0x14, 0x12, 0x12, 0x0a, 0x0e, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e,
	0x5f, 0x31, 0x30, 0x4d, 0x49, 0x4c, 0x45, 0x53, 0x10, 0x15, 0x12, 0x12, 0x0a, 0x0e, 0x57, 0x49,
	0x54, 0x48, 0x49, 0x4e, 0x5f, 0x31, 0x35, 0x4d, 0x49, 0x4c, 0x45, 0x53, 0x10, 0x16, 0x12, 0x12,
	0x0a, 0x0e, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x32, 0x30, 0x4d, 0x49, 0x4c, 0x45, 0x53,
	0x10, 0x17, 0x12, 0x12, 0x0a, 0x0e, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x32, 0x35, 0x4d,
	0x49, 0x4c, 0x45, 0x53, 0x10, 0x18, 0x12, 0x12, 0x0a, 0x0e, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e,
	0x5f, 0x33, 0x30, 0x4d, 0x49, 0x4c, 0x45, 0x53, 0x10, 0x19, 0x12, 0x12, 0x0a, 0x0e, 0x57, 0x49,
	0x54, 0x48, 0x49, 0x4e, 0x5f, 0x33, 0x35, 0x4d, 0x49, 0x4c, 0x45, 0x53, 0x10, 0x1a, 0x12, 0x12,
	0x0a, 0x0e, 0x57, 0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x34, 0x30, 0x4d, 0x49, 0x4c, 0x45, 0x53,
	0x10, 0x1b, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x45, 0x59, 0x4f, 0x4e, 0x44, 0x5f, 0x34, 0x30, 0x4d,
	0x49, 0x4c, 0x45, 0x53, 0x10, 0x1c, 0x42, 0xed, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x13, 0x44,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDescData = file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDesc
)

func file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDescData
}

var file_google_ads_googleads_v10_enums_distance_bucket_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_enums_distance_bucket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_enums_distance_bucket_proto_goTypes = []interface{}{
	(DistanceBucketEnum_DistanceBucket)(0), // 0: google.ads.googleads.v10.enums.DistanceBucketEnum.DistanceBucket
	(*DistanceBucketEnum)(nil),             // 1: google.ads.googleads.v10.enums.DistanceBucketEnum
}
var file_google_ads_googleads_v10_enums_distance_bucket_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_enums_distance_bucket_proto_init() }
func file_google_ads_googleads_v10_enums_distance_bucket_proto_init() {
	if File_google_ads_googleads_v10_enums_distance_bucket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_enums_distance_bucket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistanceBucketEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_enums_distance_bucket_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_enums_distance_bucket_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_enums_distance_bucket_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_enums_distance_bucket_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_enums_distance_bucket_proto = out.File
	file_google_ads_googleads_v10_enums_distance_bucket_proto_rawDesc = nil
	file_google_ads_googleads_v10_enums_distance_bucket_proto_goTypes = nil
	file_google_ads_googleads_v10_enums_distance_bucket_proto_depIdxs = nil
}
