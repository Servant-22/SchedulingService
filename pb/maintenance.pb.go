// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: pb/maintenance.proto

package SchedulingService_pb

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

type AppointmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskDescription string `protobuf:"bytes,2,opt,name=task_description,json=taskDescription,proto3" json:"task_description,omitempty"`
	PreferredTime   string `protobuf:"bytes,3,opt,name=preferred_time,json=preferredTime,proto3" json:"preferred_time,omitempty"`
}

func (x *AppointmentRequest) Reset() {
	*x = AppointmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_maintenance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppointmentRequest) ProtoMessage() {}

func (x *AppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_maintenance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppointmentRequest.ProtoReflect.Descriptor instead.
func (*AppointmentRequest) Descriptor() ([]byte, []int) {
	return file_pb_maintenance_proto_rawDescGZIP(), []int{0}
}

func (x *AppointmentRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AppointmentRequest) GetTaskDescription() string {
	if x != nil {
		return x.TaskDescription
	}
	return ""
}

func (x *AppointmentRequest) GetPreferredTime() string {
	if x != nil {
		return x.PreferredTime
	}
	return ""
}

type AppointmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfirmationId string `protobuf:"bytes,1,opt,name=confirmation_id,json=confirmationId,proto3" json:"confirmation_id,omitempty"`
	ScheduledTime  string `protobuf:"bytes,2,opt,name=scheduled_time,json=scheduledTime,proto3" json:"scheduled_time,omitempty"`
}

func (x *AppointmentResponse) Reset() {
	*x = AppointmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_maintenance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppointmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppointmentResponse) ProtoMessage() {}

func (x *AppointmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_maintenance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppointmentResponse.ProtoReflect.Descriptor instead.
func (*AppointmentResponse) Descriptor() ([]byte, []int) {
	return file_pb_maintenance_proto_rawDescGZIP(), []int{1}
}

func (x *AppointmentResponse) GetConfirmationId() string {
	if x != nil {
		return x.ConfirmationId
	}
	return ""
}

func (x *AppointmentResponse) GetScheduledTime() string {
	if x != nil {
		return x.ScheduledTime
	}
	return ""
}

var File_pb_maintenance_proto protoreflect.FileDescriptor

var file_pb_maintenance_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x7f, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x61,
	0x73, 0x6b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x6e, 0x0a, 0x12, 0x4d,
	0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x58, 0x0a, 0x13, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x61, 0x6e,
	0x74, 0x2d, 0x32, 0x32, 0x2f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_maintenance_proto_rawDescOnce sync.Once
	file_pb_maintenance_proto_rawDescData = file_pb_maintenance_proto_rawDesc
)

func file_pb_maintenance_proto_rawDescGZIP() []byte {
	file_pb_maintenance_proto_rawDescOnce.Do(func() {
		file_pb_maintenance_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_maintenance_proto_rawDescData)
	})
	return file_pb_maintenance_proto_rawDescData
}

var file_pb_maintenance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_maintenance_proto_goTypes = []interface{}{
	(*AppointmentRequest)(nil),  // 0: maintenance.AppointmentRequest
	(*AppointmentResponse)(nil), // 1: maintenance.AppointmentResponse
}
var file_pb_maintenance_proto_depIdxs = []int32{
	0, // 0: maintenance.MaintenanceService.ScheduleAppointment:input_type -> maintenance.AppointmentRequest
	1, // 1: maintenance.MaintenanceService.ScheduleAppointment:output_type -> maintenance.AppointmentResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_maintenance_proto_init() }
func file_pb_maintenance_proto_init() {
	if File_pb_maintenance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_maintenance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppointmentRequest); i {
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
		file_pb_maintenance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppointmentResponse); i {
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
			RawDescriptor: file_pb_maintenance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_maintenance_proto_goTypes,
		DependencyIndexes: file_pb_maintenance_proto_depIdxs,
		MessageInfos:      file_pb_maintenance_proto_msgTypes,
	}.Build()
	File_pb_maintenance_proto = out.File
	file_pb_maintenance_proto_rawDesc = nil
	file_pb_maintenance_proto_goTypes = nil
	file_pb_maintenance_proto_depIdxs = nil
}