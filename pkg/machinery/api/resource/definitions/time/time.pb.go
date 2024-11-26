// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: resource/definitions/time/time.proto

package time

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AdjtimeStatusSpec describes Linux internal adjtime state.
type AdjtimeStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset                   *durationpb.Duration `protobuf:"bytes,1,opt,name=offset,proto3" json:"offset,omitempty"`
	FrequencyAdjustmentRatio float64              `protobuf:"fixed64,2,opt,name=frequency_adjustment_ratio,json=frequencyAdjustmentRatio,proto3" json:"frequency_adjustment_ratio,omitempty"`
	MaxError                 *durationpb.Duration `protobuf:"bytes,3,opt,name=max_error,json=maxError,proto3" json:"max_error,omitempty"`
	EstError                 *durationpb.Duration `protobuf:"bytes,4,opt,name=est_error,json=estError,proto3" json:"est_error,omitempty"`
	Status                   string               `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Constant                 int64                `protobuf:"varint,6,opt,name=constant,proto3" json:"constant,omitempty"`
	SyncStatus               bool                 `protobuf:"varint,7,opt,name=sync_status,json=syncStatus,proto3" json:"sync_status,omitempty"`
	State                    string               `protobuf:"bytes,8,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *AdjtimeStatusSpec) Reset() {
	*x = AdjtimeStatusSpec{}
	mi := &file_resource_definitions_time_time_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdjtimeStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjtimeStatusSpec) ProtoMessage() {}

func (x *AdjtimeStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_time_time_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdjtimeStatusSpec.ProtoReflect.Descriptor instead.
func (*AdjtimeStatusSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_time_time_proto_rawDescGZIP(), []int{0}
}

func (x *AdjtimeStatusSpec) GetOffset() *durationpb.Duration {
	if x != nil {
		return x.Offset
	}
	return nil
}

func (x *AdjtimeStatusSpec) GetFrequencyAdjustmentRatio() float64 {
	if x != nil {
		return x.FrequencyAdjustmentRatio
	}
	return 0
}

func (x *AdjtimeStatusSpec) GetMaxError() *durationpb.Duration {
	if x != nil {
		return x.MaxError
	}
	return nil
}

func (x *AdjtimeStatusSpec) GetEstError() *durationpb.Duration {
	if x != nil {
		return x.EstError
	}
	return nil
}

func (x *AdjtimeStatusSpec) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AdjtimeStatusSpec) GetConstant() int64 {
	if x != nil {
		return x.Constant
	}
	return 0
}

func (x *AdjtimeStatusSpec) GetSyncStatus() bool {
	if x != nil {
		return x.SyncStatus
	}
	return false
}

func (x *AdjtimeStatusSpec) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

// StatusSpec describes time sync state.
type StatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Synced       bool  `protobuf:"varint,1,opt,name=synced,proto3" json:"synced,omitempty"`
	Epoch        int64 `protobuf:"varint,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
	SyncDisabled bool  `protobuf:"varint,3,opt,name=sync_disabled,json=syncDisabled,proto3" json:"sync_disabled,omitempty"`
}

func (x *StatusSpec) Reset() {
	*x = StatusSpec{}
	mi := &file_resource_definitions_time_time_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusSpec) ProtoMessage() {}

func (x *StatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_time_time_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusSpec.ProtoReflect.Descriptor instead.
func (*StatusSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_time_time_proto_rawDescGZIP(), []int{1}
}

func (x *StatusSpec) GetSynced() bool {
	if x != nil {
		return x.Synced
	}
	return false
}

func (x *StatusSpec) GetEpoch() int64 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *StatusSpec) GetSyncDisabled() bool {
	if x != nil {
		return x.SyncDisabled
	}
	return false
}

var File_resource_definitions_time_time_proto protoreflect.FileDescriptor

var file_resource_definitions_time_time_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x02, 0x0a, 0x11, 0x41, 0x64, 0x6a, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x31, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x3c, 0x0a, 0x1a, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x61, 0x64,
	0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x18, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x41,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x36,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x36, 0x0a, 0x09, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x65, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x5f, 0x0a, 0x0a, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x79,
	0x6e, 0x63, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x72, 0x0a, 0x27, 0x64, 0x65,
	0x76, 0x2e, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x74, 0x61, 0x6c,
	0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_definitions_time_time_proto_rawDescOnce sync.Once
	file_resource_definitions_time_time_proto_rawDescData = file_resource_definitions_time_time_proto_rawDesc
)

func file_resource_definitions_time_time_proto_rawDescGZIP() []byte {
	file_resource_definitions_time_time_proto_rawDescOnce.Do(func() {
		file_resource_definitions_time_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_definitions_time_time_proto_rawDescData)
	})
	return file_resource_definitions_time_time_proto_rawDescData
}

var file_resource_definitions_time_time_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resource_definitions_time_time_proto_goTypes = []any{
	(*AdjtimeStatusSpec)(nil),   // 0: talos.resource.definitions.time.AdjtimeStatusSpec
	(*StatusSpec)(nil),          // 1: talos.resource.definitions.time.StatusSpec
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
}
var file_resource_definitions_time_time_proto_depIdxs = []int32{
	2, // 0: talos.resource.definitions.time.AdjtimeStatusSpec.offset:type_name -> google.protobuf.Duration
	2, // 1: talos.resource.definitions.time.AdjtimeStatusSpec.max_error:type_name -> google.protobuf.Duration
	2, // 2: talos.resource.definitions.time.AdjtimeStatusSpec.est_error:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_resource_definitions_time_time_proto_init() }
func file_resource_definitions_time_time_proto_init() {
	if File_resource_definitions_time_time_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resource_definitions_time_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_definitions_time_time_proto_goTypes,
		DependencyIndexes: file_resource_definitions_time_time_proto_depIdxs,
		MessageInfos:      file_resource_definitions_time_time_proto_msgTypes,
	}.Build()
	File_resource_definitions_time_time_proto = out.File
	file_resource_definitions_time_time_proto_rawDesc = nil
	file_resource_definitions_time_time_proto_goTypes = nil
	file_resource_definitions_time_time_proto_depIdxs = nil
}
