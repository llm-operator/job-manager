// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api/v1/job_manager_service.proto

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

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt int64      `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Error     *Job_Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// The name of the fine-tuned model that is being created. The value will be null if the fine-tuning job is still running.
	FineTunedModel  string               `protobuf:"bytes,4,opt,name=fine_tuned_model,json=fineTunedModel,proto3" json:"fine_tuned_model,omitempty"`
	FinishedAt      int64                `protobuf:"varint,5,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	Hyperparameters *Job_Hyperparameters `protobuf:"bytes,6,opt,name=hyperparameters,proto3" json:"hyperparameters,omitempty"`
	// The base model that is being fine-tuned.
	Model          string   `protobuf:"bytes,7,opt,name=model,proto3" json:"model,omitempty"`
	Object         string   `protobuf:"bytes,8,opt,name=object,proto3" json:"object,omitempty"`
	OrganizationId string   `protobuf:"bytes,9,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	ResultFiles    []string `protobuf:"bytes,10,rep,name=result_files,json=resultFiles,proto3" json:"result_files,omitempty"`
	// The current status of the fine-tuning job, which can be either validating_files, queued, running, succeeded, failed, or cancelled.
	Status         string `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	TrainedTokens  int32  `protobuf:"varint,12,opt,name=trained_tokens,json=trainedTokens,proto3" json:"trained_tokens,omitempty"`
	TrainingFile   string `protobuf:"bytes,13,opt,name=training_file,json=trainingFile,proto3" json:"training_file,omitempty"`
	ValidationFile string `protobuf:"bytes,14,opt,name=validation_file,json=validationFile,proto3" json:"validation_file,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_manager_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_manager_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_api_v1_job_manager_service_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Job) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Job) GetError() *Job_Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Job) GetFineTunedModel() string {
	if x != nil {
		return x.FineTunedModel
	}
	return ""
}

func (x *Job) GetFinishedAt() int64 {
	if x != nil {
		return x.FinishedAt
	}
	return 0
}

func (x *Job) GetHyperparameters() *Job_Hyperparameters {
	if x != nil {
		return x.Hyperparameters
	}
	return nil
}

func (x *Job) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Job) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *Job) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Job) GetResultFiles() []string {
	if x != nil {
		return x.ResultFiles
	}
	return nil
}

func (x *Job) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Job) GetTrainedTokens() int32 {
	if x != nil {
		return x.TrainedTokens
	}
	return 0
}

func (x *Job) GetTrainingFile() string {
	if x != nil {
		return x.TrainingFile
	}
	return ""
}

func (x *Job) GetValidationFile() string {
	if x != nil {
		return x.ValidationFile
	}
	return ""
}

type CreateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model           string                            `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	TrainingFile    string                            `protobuf:"bytes,2,opt,name=training_file,json=trainingFile,proto3" json:"training_file,omitempty"`
	Hyperparameters *CreateJobRequest_Hyperparameters `protobuf:"bytes,3,opt,name=hyperparameters,proto3" json:"hyperparameters,omitempty"`
	// A string of up to 18 characters that will be added to your fine-tuned model name.
	//
	// For example, a suffix of "custom-model-name" would produce a
	// model name like
	// ft:gpt-3.5-turbo:openai:custom-model-name:7p4lURel.
	Suffix         string `protobuf:"bytes,4,opt,name=suffix,proto3" json:"suffix,omitempty"`
	ValidationFile string `protobuf:"bytes,5,opt,name=validation_file,json=validationFile,proto3" json:"validation_file,omitempty"`
}

func (x *CreateJobRequest) Reset() {
	*x = CreateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_manager_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobRequest) ProtoMessage() {}

func (x *CreateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_manager_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobRequest.ProtoReflect.Descriptor instead.
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_job_manager_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CreateJobRequest) GetTrainingFile() string {
	if x != nil {
		return x.TrainingFile
	}
	return ""
}

func (x *CreateJobRequest) GetHyperparameters() *CreateJobRequest_Hyperparameters {
	if x != nil {
		return x.Hyperparameters
	}
	return nil
}

func (x *CreateJobRequest) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

func (x *CreateJobRequest) GetValidationFile() string {
	if x != nil {
		return x.ValidationFile
	}
	return ""
}

type ListJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// after is the identifier for the last job from the previous pagination request.
	After string `protobuf:"bytes,1,opt,name=after,proto3" json:"after,omitempty"`
	// limit is the number of fine-tuning jobs to retrieve. Defaults to 20.
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListJobsRequest) Reset() {
	*x = ListJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_manager_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsRequest) ProtoMessage() {}

func (x *ListJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_manager_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsRequest.ProtoReflect.Descriptor instead.
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_job_manager_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListJobsRequest) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *ListJobsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object  string `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Data    []*Job `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	HasMore bool   `protobuf:"varint,3,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
}

func (x *ListJobsResponse) Reset() {
	*x = ListJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_manager_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsResponse) ProtoMessage() {}

func (x *ListJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_manager_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsResponse.ProtoReflect.Descriptor instead.
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_job_manager_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListJobsResponse) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ListJobsResponse) GetData() []*Job {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListJobsResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

type GetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_manager_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_manager_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_job_manager_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelJobRequest) Reset() {
	*x = CancelJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_manager_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelJobRequest) ProtoMessage() {}

func (x *CancelJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_manager_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelJobRequest.ProtoReflect.Descriptor instead.
func (*CancelJobRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_job_manager_service_proto_rawDescGZIP(), []int{5}
}

func (x *CancelJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Job_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Param   string `protobuf:"bytes,3,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *Job_Error) Reset() {
	*x = Job_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_manager_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job_Error) ProtoMessage() {}

func (x *Job_Error) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_manager_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job_Error.ProtoReflect.Descriptor instead.
func (*Job_Error) Descriptor() ([]byte, []int) {
	return file_api_v1_job_manager_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Job_Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Job_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Job_Error) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

type Job_Hyperparameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Note: OpenAI API supports string or interger.
	BatchSize int32 `protobuf:"varint,1,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	// Note: OpenAI API supports string or number.
	LearningRateMultiplier float64 `protobuf:"fixed64,2,opt,name=learning_rate_multiplier,json=learningRateMultiplier,proto3" json:"learning_rate_multiplier,omitempty"`
	// Note: OpenAI API supports string or interger.
	NEpochs int32 `protobuf:"varint,3,opt,name=n_epochs,json=nEpochs,proto3" json:"n_epochs,omitempty"`
}

func (x *Job_Hyperparameters) Reset() {
	*x = Job_Hyperparameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_manager_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job_Hyperparameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job_Hyperparameters) ProtoMessage() {}

func (x *Job_Hyperparameters) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_manager_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job_Hyperparameters.ProtoReflect.Descriptor instead.
func (*Job_Hyperparameters) Descriptor() ([]byte, []int) {
	return file_api_v1_job_manager_service_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Job_Hyperparameters) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *Job_Hyperparameters) GetLearningRateMultiplier() float64 {
	if x != nil {
		return x.LearningRateMultiplier
	}
	return 0
}

func (x *Job_Hyperparameters) GetNEpochs() int32 {
	if x != nil {
		return x.NEpochs
	}
	return 0
}

type CreateJobRequest_Hyperparameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Note: OpenAI API supports string or interger.
	BatchSize int32 `protobuf:"varint,1,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	// Note: OpenAI API supports string or number.
	LearningRateMultiplier float64 `protobuf:"fixed64,2,opt,name=learning_rate_multiplier,json=learningRateMultiplier,proto3" json:"learning_rate_multiplier,omitempty"`
	// Note: OpenAI API supports string or interger.
	NEpochs int32 `protobuf:"varint,3,opt,name=n_epochs,json=nEpochs,proto3" json:"n_epochs,omitempty"`
}

func (x *CreateJobRequest_Hyperparameters) Reset() {
	*x = CreateJobRequest_Hyperparameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_manager_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobRequest_Hyperparameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobRequest_Hyperparameters) ProtoMessage() {}

func (x *CreateJobRequest_Hyperparameters) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_manager_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobRequest_Hyperparameters.ProtoReflect.Descriptor instead.
func (*CreateJobRequest_Hyperparameters) Descriptor() ([]byte, []int) {
	return file_api_v1_job_manager_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CreateJobRequest_Hyperparameters) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *CreateJobRequest_Hyperparameters) GetLearningRateMultiplier() float64 {
	if x != nil {
		return x.LearningRateMultiplier
	}
	return 0
}

func (x *CreateJobRequest_Hyperparameters) GetNEpochs() int32 {
	if x != nil {
		return x.NEpochs
	}
	return 0
}

var File_api_v1_job_manager_service_proto protoreflect.FileDescriptor

var file_api_v1_job_manager_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x81, 0x06, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f,
	0x62, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x28,
	0x0a, 0x10, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x6e, 0x65, 0x54, 0x75,
	0x6e, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x60, 0x0a, 0x0f, 0x68, 0x79, 0x70,
	0x65, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x48, 0x79, 0x70, 0x65, 0x72,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0f, 0x68, 0x79, 0x70, 0x65,
	0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69,
	0x6c, 0x65, 0x1a, 0x4b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x1a,
	0x85, 0x01, 0x0a, 0x0f, 0x48, 0x79, 0x70, 0x65, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x16, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x61,
	0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08,
	0x6e, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x73, 0x22, 0x85, 0x03, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x6d, 0x0a, 0x0f, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x43, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x66,
	0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x79, 0x70, 0x65, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x27,
	0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x85, 0x01, 0x0a, 0x0f, 0x48, 0x79, 0x70, 0x65,
	0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x16, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x73, 0x22,
	0x3d, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x81,
	0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f,
	0x62, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d,
	0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f,
	0x72, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xd0, 0x04, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x65,
	0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x33, 0x2e, 0x6c, 0x6c,
	0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74,
	0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x66,
	0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e,
	0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x91, 0x01, 0x0a, 0x08, 0x4c, 0x69,
	0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x32, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a,
	0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6c, 0x6c, 0x6d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6e,
	0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x85, 0x01,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x30, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6c, 0x6c, 0x6d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a,
	0x6f, 0x62, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x92, 0x01, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x4a, 0x6f, 0x62, 0x12, 0x33, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62,
	0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69,
	0x6e, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6c, 0x6d, 0x2d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6a, 0x6f, 0x62, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_job_manager_service_proto_rawDescOnce sync.Once
	file_api_v1_job_manager_service_proto_rawDescData = file_api_v1_job_manager_service_proto_rawDesc
)

func file_api_v1_job_manager_service_proto_rawDescGZIP() []byte {
	file_api_v1_job_manager_service_proto_rawDescOnce.Do(func() {
		file_api_v1_job_manager_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_job_manager_service_proto_rawDescData)
	})
	return file_api_v1_job_manager_service_proto_rawDescData
}

var file_api_v1_job_manager_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_v1_job_manager_service_proto_goTypes = []interface{}{
	(*Job)(nil),                              // 0: llmoperator.fine_tuning.server.v1.Job
	(*CreateJobRequest)(nil),                 // 1: llmoperator.fine_tuning.server.v1.CreateJobRequest
	(*ListJobsRequest)(nil),                  // 2: llmoperator.fine_tuning.server.v1.ListJobsRequest
	(*ListJobsResponse)(nil),                 // 3: llmoperator.fine_tuning.server.v1.ListJobsResponse
	(*GetJobRequest)(nil),                    // 4: llmoperator.fine_tuning.server.v1.GetJobRequest
	(*CancelJobRequest)(nil),                 // 5: llmoperator.fine_tuning.server.v1.CancelJobRequest
	(*Job_Error)(nil),                        // 6: llmoperator.fine_tuning.server.v1.Job.Error
	(*Job_Hyperparameters)(nil),              // 7: llmoperator.fine_tuning.server.v1.Job.Hyperparameters
	(*CreateJobRequest_Hyperparameters)(nil), // 8: llmoperator.fine_tuning.server.v1.CreateJobRequest.Hyperparameters
}
var file_api_v1_job_manager_service_proto_depIdxs = []int32{
	6, // 0: llmoperator.fine_tuning.server.v1.Job.error:type_name -> llmoperator.fine_tuning.server.v1.Job.Error
	7, // 1: llmoperator.fine_tuning.server.v1.Job.hyperparameters:type_name -> llmoperator.fine_tuning.server.v1.Job.Hyperparameters
	8, // 2: llmoperator.fine_tuning.server.v1.CreateJobRequest.hyperparameters:type_name -> llmoperator.fine_tuning.server.v1.CreateJobRequest.Hyperparameters
	0, // 3: llmoperator.fine_tuning.server.v1.ListJobsResponse.data:type_name -> llmoperator.fine_tuning.server.v1.Job
	1, // 4: llmoperator.fine_tuning.server.v1.FineTuningService.CreateJob:input_type -> llmoperator.fine_tuning.server.v1.CreateJobRequest
	2, // 5: llmoperator.fine_tuning.server.v1.FineTuningService.ListJobs:input_type -> llmoperator.fine_tuning.server.v1.ListJobsRequest
	4, // 6: llmoperator.fine_tuning.server.v1.FineTuningService.GetJob:input_type -> llmoperator.fine_tuning.server.v1.GetJobRequest
	5, // 7: llmoperator.fine_tuning.server.v1.FineTuningService.CancelJob:input_type -> llmoperator.fine_tuning.server.v1.CancelJobRequest
	0, // 8: llmoperator.fine_tuning.server.v1.FineTuningService.CreateJob:output_type -> llmoperator.fine_tuning.server.v1.Job
	3, // 9: llmoperator.fine_tuning.server.v1.FineTuningService.ListJobs:output_type -> llmoperator.fine_tuning.server.v1.ListJobsResponse
	0, // 10: llmoperator.fine_tuning.server.v1.FineTuningService.GetJob:output_type -> llmoperator.fine_tuning.server.v1.Job
	0, // 11: llmoperator.fine_tuning.server.v1.FineTuningService.CancelJob:output_type -> llmoperator.fine_tuning.server.v1.Job
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_job_manager_service_proto_init() }
func file_api_v1_job_manager_service_proto_init() {
	if File_api_v1_job_manager_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_job_manager_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_api_v1_job_manager_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobRequest); i {
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
		file_api_v1_job_manager_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsRequest); i {
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
		file_api_v1_job_manager_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsResponse); i {
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
		file_api_v1_job_manager_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobRequest); i {
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
		file_api_v1_job_manager_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelJobRequest); i {
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
		file_api_v1_job_manager_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job_Error); i {
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
		file_api_v1_job_manager_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job_Hyperparameters); i {
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
		file_api_v1_job_manager_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobRequest_Hyperparameters); i {
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
			RawDescriptor: file_api_v1_job_manager_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_job_manager_service_proto_goTypes,
		DependencyIndexes: file_api_v1_job_manager_service_proto_depIdxs,
		MessageInfos:      file_api_v1_job_manager_service_proto_msgTypes,
	}.Build()
	File_api_v1_job_manager_service_proto = out.File
	file_api_v1_job_manager_service_proto_rawDesc = nil
	file_api_v1_job_manager_service_proto_goTypes = nil
	file_api_v1_job_manager_service_proto_depIdxs = nil
}
