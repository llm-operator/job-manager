// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: api/v1/batch_service.proto

package v1

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

type BatchJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt  int64               `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	FinishedAt int64               `protobuf:"varint,3,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	Error      *BatchJob_Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Status     string              `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Image      string              `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Command    string              `protobuf:"bytes,7,opt,name=command,proto3" json:"command,omitempty"`
	Resources  *BatchJob_Resources `protobuf:"bytes,8,opt,name=resources,proto3" json:"resources,omitempty"`
	Envs       map[string]string   `protobuf:"bytes,9,rep,name=envs,proto3" json:"envs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DataFiles  []string            `protobuf:"bytes,10,rep,name=data_files,json=dataFiles,proto3" json:"data_files,omitempty"`
}

func (x *BatchJob) Reset() {
	*x = BatchJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_batch_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchJob) ProtoMessage() {}

func (x *BatchJob) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_batch_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchJob.ProtoReflect.Descriptor instead.
func (*BatchJob) Descriptor() ([]byte, []int) {
	return file_api_v1_batch_service_proto_rawDescGZIP(), []int{0}
}

func (x *BatchJob) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BatchJob) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *BatchJob) GetFinishedAt() int64 {
	if x != nil {
		return x.FinishedAt
	}
	return 0
}

func (x *BatchJob) GetError() *BatchJob_Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *BatchJob) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BatchJob) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *BatchJob) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *BatchJob) GetResources() *BatchJob_Resources {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *BatchJob) GetEnvs() map[string]string {
	if x != nil {
		return x.Envs
	}
	return nil
}

func (x *BatchJob) GetDataFiles() []string {
	if x != nil {
		return x.DataFiles
	}
	return nil
}

type CreateBatchJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image   string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	// scripts is a map of script names to script contents.
	// The total size of the scripts should not exceed 1MB.
	Scripts   map[string][]byte   `protobuf:"bytes,3,rep,name=scripts,proto3" json:"scripts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Resources *BatchJob_Resources `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
	Envs      map[string]string   `protobuf:"bytes,5,rep,name=envs,proto3" json:"envs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// data_files is a list of file IDs that will be downloaded to the container.
	DataFiles []string `protobuf:"bytes,6,rep,name=data_files,json=dataFiles,proto3" json:"data_files,omitempty"`
}

func (x *CreateBatchJobRequest) Reset() {
	*x = CreateBatchJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_batch_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBatchJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchJobRequest) ProtoMessage() {}

func (x *CreateBatchJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_batch_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchJobRequest.ProtoReflect.Descriptor instead.
func (*CreateBatchJobRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_batch_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBatchJobRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateBatchJobRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CreateBatchJobRequest) GetScripts() map[string][]byte {
	if x != nil {
		return x.Scripts
	}
	return nil
}

func (x *CreateBatchJobRequest) GetResources() *BatchJob_Resources {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *CreateBatchJobRequest) GetEnvs() map[string]string {
	if x != nil {
		return x.Envs
	}
	return nil
}

func (x *CreateBatchJobRequest) GetDataFiles() []string {
	if x != nil {
		return x.DataFiles
	}
	return nil
}

type ListBatchJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// after is the identifier for the last batch job from the previous pagination request.
	After string `protobuf:"bytes,1,opt,name=after,proto3" json:"after,omitempty"`
	// limit is the number of batch jobs to retrieve. Defaults to 20.
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListBatchJobsRequest) Reset() {
	*x = ListBatchJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_batch_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBatchJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBatchJobsRequest) ProtoMessage() {}

func (x *ListBatchJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_batch_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBatchJobsRequest.ProtoReflect.Descriptor instead.
func (*ListBatchJobsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_batch_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListBatchJobsRequest) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *ListBatchJobsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListBatchJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs    []*BatchJob `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	HasMore bool        `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
}

func (x *ListBatchJobsResponse) Reset() {
	*x = ListBatchJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_batch_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBatchJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBatchJobsResponse) ProtoMessage() {}

func (x *ListBatchJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_batch_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBatchJobsResponse.ProtoReflect.Descriptor instead.
func (*ListBatchJobsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_batch_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListBatchJobsResponse) GetJobs() []*BatchJob {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *ListBatchJobsResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

type GetBatchJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBatchJobRequest) Reset() {
	*x = GetBatchJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_batch_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchJobRequest) ProtoMessage() {}

func (x *GetBatchJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_batch_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchJobRequest.ProtoReflect.Descriptor instead.
func (*GetBatchJobRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_batch_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetBatchJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelBatchJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelBatchJobRequest) Reset() {
	*x = CancelBatchJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_batch_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelBatchJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBatchJobRequest) ProtoMessage() {}

func (x *CancelBatchJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_batch_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBatchJobRequest.ProtoReflect.Descriptor instead.
func (*CancelBatchJobRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_batch_service_proto_rawDescGZIP(), []int{5}
}

func (x *CancelBatchJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BatchJob_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BatchJob_Error) Reset() {
	*x = BatchJob_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_batch_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchJob_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchJob_Error) ProtoMessage() {}

func (x *BatchJob_Error) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_batch_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchJob_Error.ProtoReflect.Descriptor instead.
func (*BatchJob_Error) Descriptor() ([]byte, []int) {
	return file_api_v1_batch_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BatchJob_Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BatchJob_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BatchJob_Resources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GpuCount int32 `protobuf:"varint,1,opt,name=gpu_count,json=gpuCount,proto3" json:"gpu_count,omitempty"`
}

func (x *BatchJob_Resources) Reset() {
	*x = BatchJob_Resources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_batch_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchJob_Resources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchJob_Resources) ProtoMessage() {}

func (x *BatchJob_Resources) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_batch_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchJob_Resources.ProtoReflect.Descriptor instead.
func (*BatchJob_Resources) Descriptor() ([]byte, []int) {
	return file_api_v1_batch_service_proto_rawDescGZIP(), []int{0, 1}
}

func (x *BatchJob_Resources) GetGpuCount() int32 {
	if x != nil {
		return x.GpuCount
	}
	return 0
}

var File_api_v1_batch_service_proto protoreflect.FileDescriptor

var file_api_v1_batch_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6c, 0x6c,
	0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x04, 0x0a, 0x08, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x41, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x4d, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x43,
	0x0a, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6c,
	0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4a, 0x6f, 0x62, 0x2e, 0x45, 0x6e, 0x76, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x65,
	0x6e, 0x76, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x1a, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x28, 0x0a, 0x09, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x70, 0x75, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x70, 0x75, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x1a, 0x37, 0x0a, 0x09, 0x45, 0x6e, 0x76, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd7, 0x03, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x59, 0x0a, 0x07, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x73, 0x12, 0x4d, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x50, 0x0a, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x65, 0x6e,
	0x76, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a,
	0x09, 0x45, 0x6e, 0x76, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x42, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6d, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x27, 0x0a, 0x15, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xbf, 0x04, 0x0a, 0x0c, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x12, 0x32, 0x2e, 0x6c,
	0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a,
	0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x6a, 0x6f,
	0x62, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4a, 0x6f, 0x62, 0x73, 0x12, 0x31, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a,
	0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x6a,
	0x6f, 0x62, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4a, 0x6f, 0x62, 0x12, 0x2f, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x22, 0x1b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x6a,
	0x6f, 0x62, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x8f, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x12, 0x32, 0x2e, 0x6c, 0x6c,
	0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x1a,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6c, 0x6d, 0x2d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6a, 0x6f, 0x62, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_batch_service_proto_rawDescOnce sync.Once
	file_api_v1_batch_service_proto_rawDescData = file_api_v1_batch_service_proto_rawDesc
)

func file_api_v1_batch_service_proto_rawDescGZIP() []byte {
	file_api_v1_batch_service_proto_rawDescOnce.Do(func() {
		file_api_v1_batch_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_batch_service_proto_rawDescData)
	})
	return file_api_v1_batch_service_proto_rawDescData
}

var file_api_v1_batch_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_v1_batch_service_proto_goTypes = []interface{}{
	(*BatchJob)(nil),              // 0: llmoperator.batch.server.v1.BatchJob
	(*CreateBatchJobRequest)(nil), // 1: llmoperator.batch.server.v1.CreateBatchJobRequest
	(*ListBatchJobsRequest)(nil),  // 2: llmoperator.batch.server.v1.ListBatchJobsRequest
	(*ListBatchJobsResponse)(nil), // 3: llmoperator.batch.server.v1.ListBatchJobsResponse
	(*GetBatchJobRequest)(nil),    // 4: llmoperator.batch.server.v1.GetBatchJobRequest
	(*CancelBatchJobRequest)(nil), // 5: llmoperator.batch.server.v1.CancelBatchJobRequest
	(*BatchJob_Error)(nil),        // 6: llmoperator.batch.server.v1.BatchJob.Error
	(*BatchJob_Resources)(nil),    // 7: llmoperator.batch.server.v1.BatchJob.Resources
	nil,                           // 8: llmoperator.batch.server.v1.BatchJob.EnvsEntry
	nil,                           // 9: llmoperator.batch.server.v1.CreateBatchJobRequest.ScriptsEntry
	nil,                           // 10: llmoperator.batch.server.v1.CreateBatchJobRequest.EnvsEntry
}
var file_api_v1_batch_service_proto_depIdxs = []int32{
	6,  // 0: llmoperator.batch.server.v1.BatchJob.error:type_name -> llmoperator.batch.server.v1.BatchJob.Error
	7,  // 1: llmoperator.batch.server.v1.BatchJob.resources:type_name -> llmoperator.batch.server.v1.BatchJob.Resources
	8,  // 2: llmoperator.batch.server.v1.BatchJob.envs:type_name -> llmoperator.batch.server.v1.BatchJob.EnvsEntry
	9,  // 3: llmoperator.batch.server.v1.CreateBatchJobRequest.scripts:type_name -> llmoperator.batch.server.v1.CreateBatchJobRequest.ScriptsEntry
	7,  // 4: llmoperator.batch.server.v1.CreateBatchJobRequest.resources:type_name -> llmoperator.batch.server.v1.BatchJob.Resources
	10, // 5: llmoperator.batch.server.v1.CreateBatchJobRequest.envs:type_name -> llmoperator.batch.server.v1.CreateBatchJobRequest.EnvsEntry
	0,  // 6: llmoperator.batch.server.v1.ListBatchJobsResponse.jobs:type_name -> llmoperator.batch.server.v1.BatchJob
	1,  // 7: llmoperator.batch.server.v1.BatchService.CreateBatchJob:input_type -> llmoperator.batch.server.v1.CreateBatchJobRequest
	2,  // 8: llmoperator.batch.server.v1.BatchService.ListBatchJobs:input_type -> llmoperator.batch.server.v1.ListBatchJobsRequest
	4,  // 9: llmoperator.batch.server.v1.BatchService.GetBatchJob:input_type -> llmoperator.batch.server.v1.GetBatchJobRequest
	5,  // 10: llmoperator.batch.server.v1.BatchService.CancelBatchJob:input_type -> llmoperator.batch.server.v1.CancelBatchJobRequest
	0,  // 11: llmoperator.batch.server.v1.BatchService.CreateBatchJob:output_type -> llmoperator.batch.server.v1.BatchJob
	3,  // 12: llmoperator.batch.server.v1.BatchService.ListBatchJobs:output_type -> llmoperator.batch.server.v1.ListBatchJobsResponse
	0,  // 13: llmoperator.batch.server.v1.BatchService.GetBatchJob:output_type -> llmoperator.batch.server.v1.BatchJob
	0,  // 14: llmoperator.batch.server.v1.BatchService.CancelBatchJob:output_type -> llmoperator.batch.server.v1.BatchJob
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_v1_batch_service_proto_init() }
func file_api_v1_batch_service_proto_init() {
	if File_api_v1_batch_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_batch_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchJob); i {
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
		file_api_v1_batch_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBatchJobRequest); i {
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
		file_api_v1_batch_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBatchJobsRequest); i {
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
		file_api_v1_batch_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBatchJobsResponse); i {
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
		file_api_v1_batch_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchJobRequest); i {
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
		file_api_v1_batch_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelBatchJobRequest); i {
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
		file_api_v1_batch_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchJob_Error); i {
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
		file_api_v1_batch_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchJob_Resources); i {
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
			RawDescriptor: file_api_v1_batch_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_batch_service_proto_goTypes,
		DependencyIndexes: file_api_v1_batch_service_proto_depIdxs,
		MessageInfos:      file_api_v1_batch_service_proto_msgTypes,
	}.Build()
	File_api_v1_batch_service_proto = out.File
	file_api_v1_batch_service_proto_rawDesc = nil
	file_api_v1_batch_service_proto_goTypes = nil
	file_api_v1_batch_service_proto_depIdxs = nil
}
