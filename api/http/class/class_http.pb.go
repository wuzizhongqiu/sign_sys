// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.0
// source: http/class/class_http.proto

package class

import (
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

// 创建班级
type CreateClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateClassRequest) Reset() {
	*x = CreateClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassRequest) ProtoMessage() {}

func (x *CreateClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassRequest.ProtoReflect.Descriptor instead.
func (*CreateClassRequest) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClassRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateClassReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateClassReply) Reset() {
	*x = CreateClassReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClassReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassReply) ProtoMessage() {}

func (x *CreateClassReply) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassReply.ProtoReflect.Descriptor instead.
func (*CreateClassReply) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClassReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 获取班级信息
type GetClassByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetClassByNameRequest) Reset() {
	*x = GetClassByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassByNameRequest) ProtoMessage() {}

func (x *GetClassByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassByNameRequest.ProtoReflect.Descriptor instead.
func (*GetClassByNameRequest) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{2}
}

func (x *GetClassByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetClassByNameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Student             string `protobuf:"bytes,2,opt,name=student,proto3" json:"student,omitempty"`
	Teacher             string `protobuf:"bytes,3,opt,name=teacher,proto3" json:"teacher,omitempty"`
	StudentAbsence      string `protobuf:"bytes,4,opt,name=student_absence,json=studentAbsence,proto3" json:"student_absence,omitempty"`
	StudentAbsenceCount int64  `protobuf:"varint,5,opt,name=student_absence_count,json=studentAbsenceCount,proto3" json:"student_absence_count,omitempty"`
}

func (x *GetClassByNameReply) Reset() {
	*x = GetClassByNameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClassByNameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassByNameReply) ProtoMessage() {}

func (x *GetClassByNameReply) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassByNameReply.ProtoReflect.Descriptor instead.
func (*GetClassByNameReply) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{3}
}

func (x *GetClassByNameReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetClassByNameReply) GetStudent() string {
	if x != nil {
		return x.Student
	}
	return ""
}

func (x *GetClassByNameReply) GetTeacher() string {
	if x != nil {
		return x.Teacher
	}
	return ""
}

func (x *GetClassByNameReply) GetStudentAbsence() string {
	if x != nil {
		return x.StudentAbsence
	}
	return ""
}

func (x *GetClassByNameReply) GetStudentAbsenceCount() int64 {
	if x != nil {
		return x.StudentAbsenceCount
	}
	return 0
}

// 获取所有班级列表
type GetAllClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllClassRequest) Reset() {
	*x = GetAllClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllClassRequest) ProtoMessage() {}

func (x *GetAllClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllClassRequest.ProtoReflect.Descriptor instead.
func (*GetAllClassRequest) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{4}
}

type GetAllClassReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name []string `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAllClassReply) Reset() {
	*x = GetAllClassReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllClassReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllClassReply) ProtoMessage() {}

func (x *GetAllClassReply) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllClassReply.ProtoReflect.Descriptor instead.
func (*GetAllClassReply) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllClassReply) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

// 加入班级
type JoinClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *JoinClassRequest) Reset() {
	*x = JoinClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinClassRequest) ProtoMessage() {}

func (x *JoinClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinClassRequest.ProtoReflect.Descriptor instead.
func (*JoinClassRequest) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{6}
}

func (x *JoinClassRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type JoinClassReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *JoinClassReply) Reset() {
	*x = JoinClassReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinClassReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinClassReply) ProtoMessage() {}

func (x *JoinClassReply) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinClassReply.ProtoReflect.Descriptor instead.
func (*JoinClassReply) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{7}
}

func (x *JoinClassReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 设置考勤
type SetAttendanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time       int32  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	AttendName string `protobuf:"bytes,2,opt,name=attend_name,json=attendName,proto3" json:"attend_name,omitempty"`
}

func (x *SetAttendanceRequest) Reset() {
	*x = SetAttendanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAttendanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAttendanceRequest) ProtoMessage() {}

func (x *SetAttendanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAttendanceRequest.ProtoReflect.Descriptor instead.
func (*SetAttendanceRequest) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{8}
}

func (x *SetAttendanceRequest) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *SetAttendanceRequest) GetAttendName() string {
	if x != nil {
		return x.AttendName
	}
	return ""
}

type SetAttendanceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SetAttendanceReply) Reset() {
	*x = SetAttendanceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAttendanceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAttendanceReply) ProtoMessage() {}

func (x *SetAttendanceReply) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAttendanceReply.ProtoReflect.Descriptor instead.
func (*SetAttendanceReply) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{9}
}

func (x *SetAttendanceReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 学生签到
type StudentSignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttendName string `protobuf:"bytes,1,opt,name=attend_name,json=attendName,proto3" json:"attend_name,omitempty"`
}

func (x *StudentSignInRequest) Reset() {
	*x = StudentSignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentSignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentSignInRequest) ProtoMessage() {}

func (x *StudentSignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentSignInRequest.ProtoReflect.Descriptor instead.
func (*StudentSignInRequest) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{10}
}

func (x *StudentSignInRequest) GetAttendName() string {
	if x != nil {
		return x.AttendName
	}
	return ""
}

type StudentSignInReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StudentSignInReply) Reset() {
	*x = StudentSignInReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_class_class_http_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentSignInReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentSignInReply) ProtoMessage() {}

func (x *StudentSignInReply) ProtoReflect() protoreflect.Message {
	mi := &file_http_class_class_http_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentSignInReply.ProtoReflect.Descriptor instead.
func (*StudentSignInReply) Descriptor() ([]byte, []int) {
	return file_http_class_class_http_proto_rawDescGZIP(), []int{11}
}

func (x *StudentSignInReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_http_class_class_http_proto protoreflect.FileDescriptor

var file_http_class_class_http_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x62,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x41, 0x74,
	0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x14, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xa7, 0x05, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x68, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1e, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x76, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x69, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1e,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x12, 0x60, 0x0a, 0x09, 0x4a, 0x6f,
	0x69, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1c, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x76, 0x0a, 0x0d,
	0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x74,
	0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x74,
	0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x0d, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x20, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a,
	0x01, 0x2a, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x42, 0x20, 0x5a,
	0x1e, 0x77, 0x75, 0x7a, 0x69, 0x67, 0x6f, 0x77, 0x65, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_http_class_class_http_proto_rawDescOnce sync.Once
	file_http_class_class_http_proto_rawDescData = file_http_class_class_http_proto_rawDesc
)

func file_http_class_class_http_proto_rawDescGZIP() []byte {
	file_http_class_class_http_proto_rawDescOnce.Do(func() {
		file_http_class_class_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_class_class_http_proto_rawDescData)
	})
	return file_http_class_class_http_proto_rawDescData
}

var file_http_class_class_http_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_http_class_class_http_proto_goTypes = []interface{}{
	(*CreateClassRequest)(nil),    // 0: http.class.CreateClassRequest
	(*CreateClassReply)(nil),      // 1: http.class.CreateClassReply
	(*GetClassByNameRequest)(nil), // 2: http.class.GetClassByNameRequest
	(*GetClassByNameReply)(nil),   // 3: http.class.GetClassByNameReply
	(*GetAllClassRequest)(nil),    // 4: http.class.GetAllClassRequest
	(*GetAllClassReply)(nil),      // 5: http.class.GetAllClassReply
	(*JoinClassRequest)(nil),      // 6: http.class.JoinClassRequest
	(*JoinClassReply)(nil),        // 7: http.class.JoinClassReply
	(*SetAttendanceRequest)(nil),  // 8: http.class.SetAttendanceRequest
	(*SetAttendanceReply)(nil),    // 9: http.class.SetAttendanceReply
	(*StudentSignInRequest)(nil),  // 10: http.class.StudentSignInRequest
	(*StudentSignInReply)(nil),    // 11: http.class.StudentSignInReply
}
var file_http_class_class_http_proto_depIdxs = []int32{
	0,  // 0: http.class.Class.CreateClass:input_type -> http.class.CreateClassRequest
	2,  // 1: http.class.Class.GetClassByName:input_type -> http.class.GetClassByNameRequest
	4,  // 2: http.class.Class.GetAllClass:input_type -> http.class.GetAllClassRequest
	6,  // 3: http.class.Class.JoinClass:input_type -> http.class.JoinClassRequest
	8,  // 4: http.class.Class.SetAttendance:input_type -> http.class.SetAttendanceRequest
	10, // 5: http.class.Class.StudentSignIn:input_type -> http.class.StudentSignInRequest
	1,  // 6: http.class.Class.CreateClass:output_type -> http.class.CreateClassReply
	3,  // 7: http.class.Class.GetClassByName:output_type -> http.class.GetClassByNameReply
	5,  // 8: http.class.Class.GetAllClass:output_type -> http.class.GetAllClassReply
	7,  // 9: http.class.Class.JoinClass:output_type -> http.class.JoinClassReply
	9,  // 10: http.class.Class.SetAttendance:output_type -> http.class.SetAttendanceReply
	11, // 11: http.class.Class.StudentSignIn:output_type -> http.class.StudentSignInReply
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_http_class_class_http_proto_init() }
func file_http_class_class_http_proto_init() {
	if File_http_class_class_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_class_class_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClassRequest); i {
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
		file_http_class_class_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClassReply); i {
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
		file_http_class_class_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassByNameRequest); i {
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
		file_http_class_class_http_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClassByNameReply); i {
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
		file_http_class_class_http_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllClassRequest); i {
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
		file_http_class_class_http_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllClassReply); i {
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
		file_http_class_class_http_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinClassRequest); i {
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
		file_http_class_class_http_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinClassReply); i {
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
		file_http_class_class_http_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAttendanceRequest); i {
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
		file_http_class_class_http_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAttendanceReply); i {
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
		file_http_class_class_http_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentSignInRequest); i {
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
		file_http_class_class_http_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentSignInReply); i {
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
			RawDescriptor: file_http_class_class_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_http_class_class_http_proto_goTypes,
		DependencyIndexes: file_http_class_class_http_proto_depIdxs,
		MessageInfos:      file_http_class_class_http_proto_msgTypes,
	}.Build()
	File_http_class_class_http_proto = out.File
	file_http_class_class_http_proto_rawDesc = nil
	file_http_class_class_http_proto_goTypes = nil
	file_http_class_class_http_proto_depIdxs = nil
}
