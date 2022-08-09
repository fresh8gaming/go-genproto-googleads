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
// source: google/ads/googleads/v10/errors/conversion_value_rule_error.proto

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

// Enum describing possible conversion value rule errors.
type ConversionValueRuleErrorEnum_ConversionValueRuleError int32

const (
	// Enum unspecified.
	ConversionValueRuleErrorEnum_UNSPECIFIED ConversionValueRuleErrorEnum_ConversionValueRuleError = 0
	// The received error code is not known in this version.
	ConversionValueRuleErrorEnum_UNKNOWN ConversionValueRuleErrorEnum_ConversionValueRuleError = 1
	// The value rule's geo location condition contains invalid geo target
	// constant(s), i.e. there's no matching geo target.
	ConversionValueRuleErrorEnum_INVALID_GEO_TARGET_CONSTANT ConversionValueRuleErrorEnum_ConversionValueRuleError = 2
	// The value rule's geo location condition contains conflicting included and
	// excluded geo targets. Specifically, some of the excluded geo target(s)
	// are the same as or contain some of the included geo target(s). For
	// example, the geo location condition includes California but excludes U.S.
	ConversionValueRuleErrorEnum_CONFLICTING_INCLUDED_AND_EXCLUDED_GEO_TARGET ConversionValueRuleErrorEnum_ConversionValueRuleError = 3
	// User specified conflicting conditions for two value rules in the same
	// value rule set.
	ConversionValueRuleErrorEnum_CONFLICTING_CONDITIONS ConversionValueRuleErrorEnum_ConversionValueRuleError = 4
	// The value rule cannot be removed because it's still included in some
	// value rule set.
	ConversionValueRuleErrorEnum_CANNOT_REMOVE_IF_INCLUDED_IN_VALUE_RULE_SET ConversionValueRuleErrorEnum_ConversionValueRuleError = 5
	// The value rule contains a condition that's not allowed by the value rule
	// set including this value rule.
	ConversionValueRuleErrorEnum_CONDITION_NOT_ALLOWED ConversionValueRuleErrorEnum_ConversionValueRuleError = 6
	// The value rule contains a field that should be unset.
	ConversionValueRuleErrorEnum_FIELD_MUST_BE_UNSET ConversionValueRuleErrorEnum_ConversionValueRuleError = 7
	// Pausing the value rule requires pausing the value rule set because the
	// value rule is (one of) the last enabled in the value rule set.
	ConversionValueRuleErrorEnum_CANNOT_PAUSE_UNLESS_VALUE_RULE_SET_IS_PAUSED ConversionValueRuleErrorEnum_ConversionValueRuleError = 8
	// The value rule's geo location condition contains untargetable geo target
	// constant(s).
	ConversionValueRuleErrorEnum_UNTARGETABLE_GEO_TARGET ConversionValueRuleErrorEnum_ConversionValueRuleError = 9
	// The value rule's audience condition contains invalid user list(s). In
	// another word, there's no matching user list.
	ConversionValueRuleErrorEnum_INVALID_AUDIENCE_USER_LIST ConversionValueRuleErrorEnum_ConversionValueRuleError = 10
	// The value rule's audience condition contains inaccessible user list(s).
	ConversionValueRuleErrorEnum_INACCESSIBLE_USER_LIST ConversionValueRuleErrorEnum_ConversionValueRuleError = 11
	// The value rule's audience condition contains invalid user_interest(s).
	// This might be because there is no matching user interest, or the user
	// interest is not visible.
	ConversionValueRuleErrorEnum_INVALID_AUDIENCE_USER_INTEREST ConversionValueRuleErrorEnum_ConversionValueRuleError = 12
	// When a value rule is created, it shouldn't have REMOVED status.
	ConversionValueRuleErrorEnum_CANNOT_ADD_RULE_WITH_STATUS_REMOVED ConversionValueRuleErrorEnum_ConversionValueRuleError = 13
)

// Enum value maps for ConversionValueRuleErrorEnum_ConversionValueRuleError.
var (
	ConversionValueRuleErrorEnum_ConversionValueRuleError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "INVALID_GEO_TARGET_CONSTANT",
		3:  "CONFLICTING_INCLUDED_AND_EXCLUDED_GEO_TARGET",
		4:  "CONFLICTING_CONDITIONS",
		5:  "CANNOT_REMOVE_IF_INCLUDED_IN_VALUE_RULE_SET",
		6:  "CONDITION_NOT_ALLOWED",
		7:  "FIELD_MUST_BE_UNSET",
		8:  "CANNOT_PAUSE_UNLESS_VALUE_RULE_SET_IS_PAUSED",
		9:  "UNTARGETABLE_GEO_TARGET",
		10: "INVALID_AUDIENCE_USER_LIST",
		11: "INACCESSIBLE_USER_LIST",
		12: "INVALID_AUDIENCE_USER_INTEREST",
		13: "CANNOT_ADD_RULE_WITH_STATUS_REMOVED",
	}
	ConversionValueRuleErrorEnum_ConversionValueRuleError_value = map[string]int32{
		"UNSPECIFIED":                 0,
		"UNKNOWN":                     1,
		"INVALID_GEO_TARGET_CONSTANT": 2,
		"CONFLICTING_INCLUDED_AND_EXCLUDED_GEO_TARGET": 3,
		"CONFLICTING_CONDITIONS":                       4,
		"CANNOT_REMOVE_IF_INCLUDED_IN_VALUE_RULE_SET":  5,
		"CONDITION_NOT_ALLOWED":                        6,
		"FIELD_MUST_BE_UNSET":                          7,
		"CANNOT_PAUSE_UNLESS_VALUE_RULE_SET_IS_PAUSED": 8,
		"UNTARGETABLE_GEO_TARGET":                      9,
		"INVALID_AUDIENCE_USER_LIST":                   10,
		"INACCESSIBLE_USER_LIST":                       11,
		"INVALID_AUDIENCE_USER_INTEREST":               12,
		"CANNOT_ADD_RULE_WITH_STATUS_REMOVED":          13,
	}
)

func (x ConversionValueRuleErrorEnum_ConversionValueRuleError) Enum() *ConversionValueRuleErrorEnum_ConversionValueRuleError {
	p := new(ConversionValueRuleErrorEnum_ConversionValueRuleError)
	*p = x
	return p
}

func (x ConversionValueRuleErrorEnum_ConversionValueRuleError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionValueRuleErrorEnum_ConversionValueRuleError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_enumTypes[0].Descriptor()
}

func (ConversionValueRuleErrorEnum_ConversionValueRuleError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_enumTypes[0]
}

func (x ConversionValueRuleErrorEnum_ConversionValueRuleError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionValueRuleErrorEnum_ConversionValueRuleError.Descriptor instead.
func (ConversionValueRuleErrorEnum_ConversionValueRuleError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible conversion value rule errors.
type ConversionValueRuleErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConversionValueRuleErrorEnum) Reset() {
	*x = ConversionValueRuleErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionValueRuleErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRuleErrorEnum) ProtoMessage() {}

func (x *ConversionValueRuleErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRuleErrorEnum.ProtoReflect.Descriptor instead.
func (*ConversionValueRuleErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_errors_conversion_value_rule_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x22, 0x85, 0x04, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xe4, 0x03, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x47, 0x45, 0x4f, 0x5f,
	0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x10,
	0x02, 0x12, 0x30, 0x0a, 0x2c, 0x43, 0x4f, 0x4e, 0x46, 0x4c, 0x49, 0x43, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x45, 0x58,
	0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x47, 0x45, 0x4f, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x46, 0x4c, 0x49, 0x43, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x04, 0x12,
	0x2f, 0x0a, 0x2b, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45,
	0x5f, 0x49, 0x46, 0x5f, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x05,
	0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x46,
	0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x45, 0x54, 0x10, 0x07, 0x12, 0x30, 0x0a, 0x2c, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x50,
	0x41, 0x55, 0x53, 0x45, 0x5f, 0x55, 0x4e, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x50, 0x41,
	0x55, 0x53, 0x45, 0x44, 0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x4e, 0x54, 0x41, 0x52, 0x47,
	0x45, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x47, 0x45, 0x4f, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41,
	0x55, 0x44, 0x49, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53,
	0x54, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x49,
	0x42, 0x4c, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0b, 0x12,
	0x22, 0x0a, 0x1e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x45,
	0x4e, 0x43, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x45, 0x53,
	0x54, 0x10, 0x0c, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x44,
	0x44, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x0d, 0x42, 0xfd, 0x01, 0x0a,
	0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x42, 0x1d, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64,
	0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDescData = file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDesc
)

func file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDescData
}

var file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_goTypes = []interface{}{
	(ConversionValueRuleErrorEnum_ConversionValueRuleError)(0), // 0: google.ads.googleads.v10.errors.ConversionValueRuleErrorEnum.ConversionValueRuleError
	(*ConversionValueRuleErrorEnum)(nil),                       // 1: google.ads.googleads.v10.errors.ConversionValueRuleErrorEnum
}
var file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_init() }
func file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_init() {
	if File_google_ads_googleads_v10_errors_conversion_value_rule_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionValueRuleErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_errors_conversion_value_rule_error_proto = out.File
	file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_rawDesc = nil
	file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_goTypes = nil
	file_google_ads_googleads_v10_errors_conversion_value_rule_error_proto_depIdxs = nil
}
