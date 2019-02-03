// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/operation_access_denied_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible operation access denied errors.
type OperationAccessDeniedErrorEnum_OperationAccessDeniedError int32

const (
	// Enum unspecified.
	OperationAccessDeniedErrorEnum_UNSPECIFIED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 0
	// The received error code is not known in this version.
	OperationAccessDeniedErrorEnum_UNKNOWN OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 1
	// Unauthorized invocation of a service's method (get, mutate, etc.)
	OperationAccessDeniedErrorEnum_ACTION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 2
	// Unauthorized CREATE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_CREATE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 3
	// Unauthorized REMOVE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_REMOVE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 4
	// Unauthorized UPDATE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_UPDATE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 5
	// A mutate action is not allowed on this campaign, from this client.
	OperationAccessDeniedErrorEnum_MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 6
	// This operation is not permitted on this campaign type
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 7
	// A CREATE operation may not set status to REMOVED.
	OperationAccessDeniedErrorEnum_CREATE_AS_REMOVED_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 8
	// This operation is not allowed because the campaign or adgroup is removed.
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 9
	// This operation is not permitted on this ad group type.
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 10
	// The mutate is not allowed for this customer.
	OperationAccessDeniedErrorEnum_MUTATE_NOT_PERMITTED_FOR_CUSTOMER OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 11
)

var OperationAccessDeniedErrorEnum_OperationAccessDeniedError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "ACTION_NOT_PERMITTED",
	3:  "CREATE_OPERATION_NOT_PERMITTED",
	4:  "REMOVE_OPERATION_NOT_PERMITTED",
	5:  "UPDATE_OPERATION_NOT_PERMITTED",
	6:  "MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT",
	7:  "OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE",
	8:  "CREATE_AS_REMOVED_NOT_PERMITTED",
	9:  "OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE",
	10: "OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE",
	11: "MUTATE_NOT_PERMITTED_FOR_CUSTOMER",
}
var OperationAccessDeniedErrorEnum_OperationAccessDeniedError_value = map[string]int32{
	"UNSPECIFIED":                                  0,
	"UNKNOWN":                                      1,
	"ACTION_NOT_PERMITTED":                         2,
	"CREATE_OPERATION_NOT_PERMITTED":               3,
	"REMOVE_OPERATION_NOT_PERMITTED":               4,
	"UPDATE_OPERATION_NOT_PERMITTED":               5,
	"MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT":       6,
	"OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE":    7,
	"CREATE_AS_REMOVED_NOT_PERMITTED":              8,
	"OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE": 9,
	"OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE":    10,
	"MUTATE_NOT_PERMITTED_FOR_CUSTOMER":            11,
}

func (x OperationAccessDeniedErrorEnum_OperationAccessDeniedError) String() string {
	return proto.EnumName(OperationAccessDeniedErrorEnum_OperationAccessDeniedError_name, int32(x))
}
func (OperationAccessDeniedErrorEnum_OperationAccessDeniedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_operation_access_denied_error_03fdde94fbc45f34, []int{0, 0}
}

// Container for enum describing possible operation access denied errors.
type OperationAccessDeniedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationAccessDeniedErrorEnum) Reset()         { *m = OperationAccessDeniedErrorEnum{} }
func (m *OperationAccessDeniedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*OperationAccessDeniedErrorEnum) ProtoMessage()    {}
func (*OperationAccessDeniedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_operation_access_denied_error_03fdde94fbc45f34, []int{0}
}
func (m *OperationAccessDeniedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationAccessDeniedErrorEnum.Unmarshal(m, b)
}
func (m *OperationAccessDeniedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationAccessDeniedErrorEnum.Marshal(b, m, deterministic)
}
func (dst *OperationAccessDeniedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationAccessDeniedErrorEnum.Merge(dst, src)
}
func (m *OperationAccessDeniedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_OperationAccessDeniedErrorEnum.Size(m)
}
func (m *OperationAccessDeniedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationAccessDeniedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_OperationAccessDeniedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OperationAccessDeniedErrorEnum)(nil), "google.ads.googleads.v0.errors.OperationAccessDeniedErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.OperationAccessDeniedErrorEnum_OperationAccessDeniedError", OperationAccessDeniedErrorEnum_OperationAccessDeniedError_name, OperationAccessDeniedErrorEnum_OperationAccessDeniedError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/operation_access_denied_error.proto", fileDescriptor_operation_access_denied_error_03fdde94fbc45f34)
}

var fileDescriptor_operation_access_denied_error_03fdde94fbc45f34 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0xaa, 0xd3, 0x30,
	0x1c, 0xc6, 0xdd, 0xa6, 0xe7, 0x68, 0x76, 0x61, 0x09, 0x5e, 0x88, 0x17, 0x3b, 0x58, 0x51, 0x50,
	0x34, 0x2d, 0x78, 0x17, 0xaf, 0xb2, 0x36, 0x67, 0x14, 0x6d, 0x12, 0xb2, 0x74, 0xa2, 0x0c, 0x42,
	0x5d, 0x4b, 0x19, 0x9c, 0xd3, 0x8c, 0xe6, 0x78, 0x1e, 0xc8, 0x4b, 0x9f, 0x44, 0x7c, 0x14, 0x7d,
	0x03, 0xaf, 0xa4, 0xcd, 0xba, 0x0b, 0x59, 0xe7, 0x55, 0xff, 0x34, 0xbf, 0xfc, 0xf2, 0x85, 0x7c,
	0x60, 0x5e, 0x19, 0x53, 0x5d, 0x95, 0x41, 0x5e, 0xd8, 0xc0, 0x8d, 0xed, 0x74, 0x1b, 0x06, 0x65,
	0xd3, 0x98, 0xc6, 0x06, 0x66, 0x57, 0x36, 0xf9, 0xcd, 0xd6, 0xd4, 0x3a, 0xdf, 0x6c, 0x4a, 0x6b,
	0x75, 0x51, 0xd6, 0xdb, 0xb2, 0xd0, 0xdd, 0x32, 0xda, 0x35, 0xe6, 0xc6, 0xc0, 0x99, 0xdb, 0x88,
	0xf2, 0xc2, 0xa2, 0x83, 0x03, 0xdd, 0x86, 0xc8, 0x39, 0xfc, 0xdf, 0x13, 0x30, 0xe3, 0xbd, 0x87,
	0x74, 0x9a, 0xb8, 0xb3, 0xd0, 0x76, 0x9d, 0xd6, 0x5f, 0xaf, 0xfd, 0x1f, 0x13, 0xf0, 0x64, 0x18,
	0x81, 0x0f, 0xc1, 0x34, 0x63, 0x4b, 0x41, 0xa3, 0xe4, 0x32, 0xa1, 0xb1, 0x77, 0x07, 0x4e, 0xc1,
	0x79, 0xc6, 0xde, 0x33, 0xfe, 0x91, 0x79, 0x23, 0xf8, 0x18, 0x3c, 0x22, 0x91, 0x4a, 0x38, 0xd3,
	0x8c, 0x2b, 0x2d, 0xa8, 0x4c, 0x13, 0xa5, 0x68, 0xec, 0x8d, 0xa1, 0x0f, 0x66, 0x91, 0xa4, 0x44,
	0x51, 0xcd, 0x05, 0x95, 0xe4, 0x08, 0x33, 0x69, 0x19, 0x49, 0x53, 0xbe, 0x1a, 0x66, 0xee, 0xb6,
	0x4c, 0x26, 0xe2, 0x53, 0x9e, 0x7b, 0xf0, 0x15, 0x78, 0x91, 0x66, 0xaa, 0x65, 0x8e, 0x85, 0xd1,
	0x97, 0x5c, 0xea, 0xe8, 0x43, 0x42, 0x99, 0xf2, 0xce, 0xe0, 0x1b, 0xf0, 0x72, 0x40, 0xe4, 0x38,
	0x92, 0x0a, 0x92, 0x2c, 0x98, 0x56, 0x9f, 0x04, 0xf5, 0xce, 0xe1, 0x33, 0x70, 0xb1, 0xbf, 0x06,
	0x59, 0x6a, 0x17, 0x36, 0xfe, 0xe7, 0xfc, 0xfb, 0x30, 0x04, 0xaf, 0x4f, 0x39, 0xfb, 0x6d, 0x92,
	0x2e, 0x79, 0x26, 0x23, 0xea, 0x3d, 0xf8, 0x5f, 0x0a, 0x12, 0xeb, 0x85, 0xe4, 0x99, 0x70, 0x29,
	0x00, 0x7c, 0x0e, 0x9e, 0xee, 0x2f, 0x78, 0x24, 0x71, 0xb6, 0x54, 0x3c, 0xa5, 0xd2, 0x9b, 0xce,
	0xff, 0x8c, 0x80, 0xbf, 0x31, 0xd7, 0xe8, 0x74, 0x29, 0xe6, 0x17, 0xc3, 0xcf, 0x2d, 0xda, 0x56,
	0x89, 0xd1, 0xe7, 0x78, 0xaf, 0xa8, 0xcc, 0x55, 0x5e, 0x57, 0xc8, 0x34, 0x55, 0x50, 0x95, 0x75,
	0xd7, 0xb9, 0xbe, 0xab, 0xbb, 0xad, 0x1d, 0xaa, 0xee, 0x3b, 0xf7, 0xf9, 0x36, 0x9e, 0x2c, 0x08,
	0xf9, 0x3e, 0x9e, 0x2d, 0x9c, 0x8c, 0x14, 0x16, 0xb9, 0xb1, 0x9d, 0x56, 0x21, 0xea, 0x8e, 0xb4,
	0x3f, 0x7b, 0x60, 0x4d, 0x0a, 0xbb, 0x3e, 0x00, 0xeb, 0x55, 0xb8, 0x76, 0xc0, 0xaf, 0xb1, 0xef,
	0xfe, 0x62, 0x4c, 0x0a, 0x8b, 0xf1, 0x01, 0xc1, 0x78, 0x15, 0x62, 0xec, 0xa0, 0x2f, 0x67, 0x5d,
	0xba, 0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x16, 0x16, 0x0c, 0x83, 0x57, 0x03, 0x00, 0x00,
}