// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: errs/errors.proto

package errs

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	// 未知错误
	ErrorReason_UNKNOWN ErrorReason = 0
	// 内部错误
	ErrorReason_INTERNAL ErrorReason = 500
	// 参数错误
	ErrorReason_PARAMETER ErrorReason = 400
	// 业务错误
	ErrorReason_BUSINESS ErrorReason = 600
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:   "UNKNOWN",
		500: "INTERNAL",
		400: "PARAMETER",
		600: "BUSINESS",
	}
	ErrorReason_value = map[string]int32{
		"UNKNOWN":   0,
		"INTERNAL":  500,
		"PARAMETER": 400,
		"BUSINESS":  600,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_errs_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_errs_errors_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_errs_errors_proto_rawDescGZIP(), []int{0}
}

var File_errs_errors_proto protoreflect.FileDescriptor

var file_errs_errors_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x72, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x72, 0x72, 0x73, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x62,
	0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x08, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0xf4, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12,
	0x0e, 0x0a, 0x09, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x90, 0x03, 0x12,
	0x13, 0x0a, 0x08, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45, 0x53, 0x53, 0x10, 0xd8, 0x04, 0x1a, 0x04,
	0xa8, 0x45, 0x95, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xc8, 0x01, 0x22, 0x06, 0x08, 0xc8, 0x01, 0x10,
	0xc8, 0x01, 0x42, 0x24, 0x0a, 0x04, 0x65, 0x72, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x1a, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errs_errors_proto_rawDescOnce sync.Once
	file_errs_errors_proto_rawDescData = file_errs_errors_proto_rawDesc
)

func file_errs_errors_proto_rawDescGZIP() []byte {
	file_errs_errors_proto_rawDescOnce.Do(func() {
		file_errs_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_errs_errors_proto_rawDescData)
	})
	return file_errs_errors_proto_rawDescData
}

var file_errs_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errs_errors_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: errs.ErrorReason
}
var file_errs_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errs_errors_proto_init() }
func file_errs_errors_proto_init() {
	if File_errs_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errs_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errs_errors_proto_goTypes,
		DependencyIndexes: file_errs_errors_proto_depIdxs,
		EnumInfos:         file_errs_errors_proto_enumTypes,
	}.Build()
	File_errs_errors_proto = out.File
	file_errs_errors_proto_rawDesc = nil
	file_errs_errors_proto_goTypes = nil
	file_errs_errors_proto_depIdxs = nil
}
