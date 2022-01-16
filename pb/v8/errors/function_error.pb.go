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
// source: google/ads/googleads/v8/errors/function_error.proto

package errors

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

// Enum describing possible function errors.
type FunctionErrorEnum_FunctionError int32

const (
	// Enum unspecified.
	FunctionErrorEnum_UNSPECIFIED FunctionErrorEnum_FunctionError = 0
	// The received error code is not known in this version.
	FunctionErrorEnum_UNKNOWN FunctionErrorEnum_FunctionError = 1
	// The format of the function is not recognized as a supported function
	// format.
	FunctionErrorEnum_INVALID_FUNCTION_FORMAT FunctionErrorEnum_FunctionError = 2
	// Operand data types do not match.
	FunctionErrorEnum_DATA_TYPE_MISMATCH FunctionErrorEnum_FunctionError = 3
	// The operands cannot be used together in a conjunction.
	FunctionErrorEnum_INVALID_CONJUNCTION_OPERANDS FunctionErrorEnum_FunctionError = 4
	// Invalid numer of Operands.
	FunctionErrorEnum_INVALID_NUMBER_OF_OPERANDS FunctionErrorEnum_FunctionError = 5
	// Operand Type not supported.
	FunctionErrorEnum_INVALID_OPERAND_TYPE FunctionErrorEnum_FunctionError = 6
	// Operator not supported.
	FunctionErrorEnum_INVALID_OPERATOR FunctionErrorEnum_FunctionError = 7
	// Request context type not supported.
	FunctionErrorEnum_INVALID_REQUEST_CONTEXT_TYPE FunctionErrorEnum_FunctionError = 8
	// The matching function is not allowed for call placeholders
	FunctionErrorEnum_INVALID_FUNCTION_FOR_CALL_PLACEHOLDER FunctionErrorEnum_FunctionError = 9
	// The matching function is not allowed for the specified placeholder
	FunctionErrorEnum_INVALID_FUNCTION_FOR_PLACEHOLDER FunctionErrorEnum_FunctionError = 10
	// Invalid operand.
	FunctionErrorEnum_INVALID_OPERAND FunctionErrorEnum_FunctionError = 11
	// Missing value for the constant operand.
	FunctionErrorEnum_MISSING_CONSTANT_OPERAND_VALUE FunctionErrorEnum_FunctionError = 12
	// The value of the constant operand is invalid.
	FunctionErrorEnum_INVALID_CONSTANT_OPERAND_VALUE FunctionErrorEnum_FunctionError = 13
	// Invalid function nesting.
	FunctionErrorEnum_INVALID_NESTING FunctionErrorEnum_FunctionError = 14
	// The Feed ID was different from another Feed ID in the same function.
	FunctionErrorEnum_MULTIPLE_FEED_IDS_NOT_SUPPORTED FunctionErrorEnum_FunctionError = 15
	// The matching function is invalid for use with a feed with a fixed schema.
	FunctionErrorEnum_INVALID_FUNCTION_FOR_FEED_WITH_FIXED_SCHEMA FunctionErrorEnum_FunctionError = 16
	// Invalid attribute name.
	FunctionErrorEnum_INVALID_ATTRIBUTE_NAME FunctionErrorEnum_FunctionError = 17
)

// Enum value maps for FunctionErrorEnum_FunctionError.
var (
	FunctionErrorEnum_FunctionError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "INVALID_FUNCTION_FORMAT",
		3:  "DATA_TYPE_MISMATCH",
		4:  "INVALID_CONJUNCTION_OPERANDS",
		5:  "INVALID_NUMBER_OF_OPERANDS",
		6:  "INVALID_OPERAND_TYPE",
		7:  "INVALID_OPERATOR",
		8:  "INVALID_REQUEST_CONTEXT_TYPE",
		9:  "INVALID_FUNCTION_FOR_CALL_PLACEHOLDER",
		10: "INVALID_FUNCTION_FOR_PLACEHOLDER",
		11: "INVALID_OPERAND",
		12: "MISSING_CONSTANT_OPERAND_VALUE",
		13: "INVALID_CONSTANT_OPERAND_VALUE",
		14: "INVALID_NESTING",
		15: "MULTIPLE_FEED_IDS_NOT_SUPPORTED",
		16: "INVALID_FUNCTION_FOR_FEED_WITH_FIXED_SCHEMA",
		17: "INVALID_ATTRIBUTE_NAME",
	}
	FunctionErrorEnum_FunctionError_value = map[string]int32{
		"UNSPECIFIED":                                 0,
		"UNKNOWN":                                     1,
		"INVALID_FUNCTION_FORMAT":                     2,
		"DATA_TYPE_MISMATCH":                          3,
		"INVALID_CONJUNCTION_OPERANDS":                4,
		"INVALID_NUMBER_OF_OPERANDS":                  5,
		"INVALID_OPERAND_TYPE":                        6,
		"INVALID_OPERATOR":                            7,
		"INVALID_REQUEST_CONTEXT_TYPE":                8,
		"INVALID_FUNCTION_FOR_CALL_PLACEHOLDER":       9,
		"INVALID_FUNCTION_FOR_PLACEHOLDER":            10,
		"INVALID_OPERAND":                             11,
		"MISSING_CONSTANT_OPERAND_VALUE":              12,
		"INVALID_CONSTANT_OPERAND_VALUE":              13,
		"INVALID_NESTING":                             14,
		"MULTIPLE_FEED_IDS_NOT_SUPPORTED":             15,
		"INVALID_FUNCTION_FOR_FEED_WITH_FIXED_SCHEMA": 16,
		"INVALID_ATTRIBUTE_NAME":                      17,
	}
)

func (x FunctionErrorEnum_FunctionError) Enum() *FunctionErrorEnum_FunctionError {
	p := new(FunctionErrorEnum_FunctionError)
	*p = x
	return p
}

func (x FunctionErrorEnum_FunctionError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FunctionErrorEnum_FunctionError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v8_errors_function_error_proto_enumTypes[0].Descriptor()
}

func (FunctionErrorEnum_FunctionError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v8_errors_function_error_proto_enumTypes[0]
}

func (x FunctionErrorEnum_FunctionError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FunctionErrorEnum_FunctionError.Descriptor instead.
func (FunctionErrorEnum_FunctionError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_errors_function_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible function errors.
type FunctionErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FunctionErrorEnum) Reset() {
	*x = FunctionErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_errors_function_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionErrorEnum) ProtoMessage() {}

func (x *FunctionErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_errors_function_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionErrorEnum.ProtoReflect.Descriptor instead.
func (*FunctionErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_errors_function_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v8_errors_function_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_errors_function_error_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x04, 0x0a, 0x11, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xab, 0x04, 0x0a, 0x0d, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x4f,
	0x52, 0x4d, 0x41, 0x54, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x03, 0x12, 0x20,
	0x0a, 0x1c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x4a, 0x55, 0x4e,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x53, 0x10, 0x04,
	0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4e, 0x55, 0x4d, 0x42,
	0x45, 0x52, 0x5f, 0x4f, 0x46, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x53, 0x10, 0x05,
	0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52,
	0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x07,
	0x12, 0x20, 0x0a, 0x1c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x08, 0x12, 0x29, 0x0a, 0x25, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x55,
	0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x09, 0x12, 0x24, 0x0a,
	0x20, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45,
	0x52, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x10, 0x0b, 0x12, 0x22, 0x0a, 0x1e, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x5f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x4e, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x0c, 0x12, 0x22, 0x0a, 0x1e,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x54,
	0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x0d,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4e, 0x45, 0x53, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x0e, 0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c,
	0x45, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x49, 0x44, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x0f, 0x12, 0x2f, 0x0a, 0x2b, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46,
	0x4f, 0x52, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x46, 0x49, 0x58,
	0x45, 0x44, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x10, 0x10, 0x12, 0x1a, 0x0a, 0x16, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x11, 0x42, 0xed, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x12,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x38, 0x3a,
	0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_errors_function_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_errors_function_error_proto_rawDescData = file_google_ads_googleads_v8_errors_function_error_proto_rawDesc
)

func file_google_ads_googleads_v8_errors_function_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_errors_function_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_errors_function_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_errors_function_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_errors_function_error_proto_rawDescData
}

var file_google_ads_googleads_v8_errors_function_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v8_errors_function_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v8_errors_function_error_proto_goTypes = []interface{}{
	(FunctionErrorEnum_FunctionError)(0), // 0: google.ads.googleads.v8.errors.FunctionErrorEnum.FunctionError
	(*FunctionErrorEnum)(nil),            // 1: google.ads.googleads.v8.errors.FunctionErrorEnum
}
var file_google_ads_googleads_v8_errors_function_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_errors_function_error_proto_init() }
func file_google_ads_googleads_v8_errors_function_error_proto_init() {
	if File_google_ads_googleads_v8_errors_function_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_errors_function_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v8_errors_function_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_errors_function_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_errors_function_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v8_errors_function_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v8_errors_function_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_errors_function_error_proto = out.File
	file_google_ads_googleads_v8_errors_function_error_proto_rawDesc = nil
	file_google_ads_googleads_v8_errors_function_error_proto_goTypes = nil
	file_google_ads_googleads_v8_errors_function_error_proto_depIdxs = nil
}
