// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: api.proto

package api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TodoType int32

const (
	TodoType_task  TodoType = 0
	TodoType_event TodoType = 1
)

// Enum value maps for TodoType.
var (
	TodoType_name = map[int32]string{
		0: "task",
		1: "event",
	}
	TodoType_value = map[string]int32{
		"task":  0,
		"event": 1,
	}
)

func (x TodoType) Enum() *TodoType {
	p := new(TodoType)
	*p = x
	return p
}

func (x TodoType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TodoType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (TodoType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x TodoType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TodoType.Descriptor instead.
func (TodoType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

// CreateTodo
type CreateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Type TodoType `protobuf:"varint,2,opt,name=Type,json=type,proto3,enum=grpc_server.TodoType" json:"Type,omitempty"`
}

func (x *CreateTodoRequest) Reset() {
	*x = CreateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest) ProtoMessage() {}

func (x *CreateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTodoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTodoRequest) GetType() TodoType {
	if x != nil {
		return x.Type
	}
	return TodoType_task
}

// CreateTodoV2
type CreateTodoRequestV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *CreateTodoRequest `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *CreateTodoRequestV2) Reset() {
	*x = CreateTodoRequestV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoRequestV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequestV2) ProtoMessage() {}

func (x *CreateTodoRequestV2) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequestV2.ProtoReflect.Descriptor instead.
func (*CreateTodoRequestV2) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTodoRequestV2) GetTodo() *CreateTodoRequest {
	if x != nil {
		return x.Todo
	}
	return nil
}

// GetAllTodo
type TodoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *TodoList) Reset() {
	*x = TodoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoList) ProtoMessage() {}

func (x *TodoList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoList.ProtoReflect.Descriptor instead.
func (*TodoList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *TodoList) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

// GetTodoById
type GetTodoByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTodoByIDRequest) Reset() {
	*x = GetTodoByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoByIDRequest) ProtoMessage() {}

func (x *GetTodoByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoByIDRequest.ProtoReflect.Descriptor instead.
func (*GetTodoByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetTodoByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetTodoByIDRequestWithQueryParams
type GetTodoByIDRequestWithQueryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetTodoByIDRequestWithQueryParams) Reset() {
	*x = GetTodoByIDRequestWithQueryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoByIDRequestWithQueryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoByIDRequestWithQueryParams) ProtoMessage() {}

func (x *GetTodoByIDRequestWithQueryParams) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoByIDRequestWithQueryParams.ProtoReflect.Descriptor instead.
func (*GetTodoByIDRequestWithQueryParams) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetTodoByIDRequestWithQueryParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetTodoByIDRequestWithQueryParams) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// Todo
type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string                 `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Type      TodoType               `protobuf:"varint,3,opt,name=Type,json=type,proto3,enum=grpc_server.TodoType" json:"Type,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *Todo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Todo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Todo) GetType() TodoType {
	if x != nil {
		return x.Type
	}
	return TodoType_task
}

func (x *Todo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x12, 0x92, 0x41, 0x06, 0x32, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x01, 0x18, 0x1e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x09, 0x92, 0x41, 0x06, 0x32, 0x04, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x3a, 0x13, 0x92, 0x41, 0x10, 0x0a, 0x0e, 0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2,
	0x01, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x32, 0x12, 0x32, 0x0a,
	0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x74, 0x6f, 0x64,
	0x6f, 0x22, 0x33, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a,
	0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64,
	0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x21,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x17,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0x92, 0x41, 0x04, 0x32,
	0x02, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x92, 0x41, 0x06, 0x32, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09, 0x92, 0x41, 0x06,
	0x32, 0x04, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x12, 0x92, 0x41, 0x0f,
	0x32, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x1f, 0x0a, 0x08, 0x54, 0x6f,
	0x64, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x32, 0xf4, 0x03, 0x0a, 0x0b,
	0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x5f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x64, 0x6f, 0x56, 0x32, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x32, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x3a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x32, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x51, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x6f, 0x64, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x5c, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x79, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x64, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x57, 0x69, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x42, 0x9c, 0x01, 0x92, 0x41, 0x59, 0x12, 0x21, 0x0a, 0x0c, 0x47, 0x52, 0x50, 0x43,
	0x2d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0c, 0x47, 0x52, 0x50, 0x43, 0x20, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x32, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x72, 0x0d, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x6c,
	0x61, 0x6e, 0x63, 0x6b, 0x31, 0x38, 0x35, 0x38, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x31,
	0x30, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_goTypes = []interface{}{
	(TodoType)(0),                             // 0: grpc_server.TodoType
	(*CreateTodoRequest)(nil),                 // 1: grpc_server.CreateTodoRequest
	(*CreateTodoRequestV2)(nil),               // 2: grpc_server.CreateTodoRequestV2
	(*TodoList)(nil),                          // 3: grpc_server.TodoList
	(*GetTodoByIDRequest)(nil),                // 4: grpc_server.GetTodoByIDRequest
	(*GetTodoByIDRequestWithQueryParams)(nil), // 5: grpc_server.GetTodoByIDRequestWithQueryParams
	(*Todo)(nil),                              // 6: grpc_server.Todo
	(*timestamppb.Timestamp)(nil),             // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                     // 8: google.protobuf.Empty
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: grpc_server.CreateTodoRequest.Type:type_name -> grpc_server.TodoType
	1,  // 1: grpc_server.CreateTodoRequestV2.todo:type_name -> grpc_server.CreateTodoRequest
	6,  // 2: grpc_server.TodoList.todos:type_name -> grpc_server.Todo
	0,  // 3: grpc_server.Todo.Type:type_name -> grpc_server.TodoType
	7,  // 4: grpc_server.Todo.createdAt:type_name -> google.protobuf.Timestamp
	1,  // 5: grpc_server.GrpcService.CreateTodo:input_type -> grpc_server.CreateTodoRequest
	2,  // 6: grpc_server.GrpcService.CreateTodoV2:input_type -> grpc_server.CreateTodoRequestV2
	8,  // 7: grpc_server.GrpcService.GetAllTodo:input_type -> google.protobuf.Empty
	4,  // 8: grpc_server.GrpcService.GetTodoByID:input_type -> grpc_server.GetTodoByIDRequest
	5,  // 9: grpc_server.GrpcService.GetTodoByIDWithQueryParams:input_type -> grpc_server.GetTodoByIDRequestWithQueryParams
	6,  // 10: grpc_server.GrpcService.CreateTodo:output_type -> grpc_server.Todo
	6,  // 11: grpc_server.GrpcService.CreateTodoV2:output_type -> grpc_server.Todo
	3,  // 12: grpc_server.GrpcService.GetAllTodo:output_type -> grpc_server.TodoList
	6,  // 13: grpc_server.GrpcService.GetTodoByID:output_type -> grpc_server.Todo
	3,  // 14: grpc_server.GrpcService.GetTodoByIDWithQueryParams:output_type -> grpc_server.TodoList
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoRequestV2); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoList); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoByIDRequest); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoByIDRequestWithQueryParams); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
