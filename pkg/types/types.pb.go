// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/types/types.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	pkg/types/types.proto

It has these top-level messages:
	Workflow
	WorkflowSpec
	WorkflowStatus
	WorkflowInvocation
	WorkflowInvocationSpec
	WorkflowInvocationStatus
	DependencyConfig
	Task
	TaskSpec
	TaskStatus
	TaskDependencyParameters
	TaskInvocation
	TaskInvocationSpec
	TaskInvocationStatus
	ObjectMetadata
	TypedValue
	Error
	FnRef
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowStatus_Status int32

const (
	WorkflowStatus_PENDING WorkflowStatus_Status = 0
	//        PARSING = 1; // During validation/parsing
	WorkflowStatus_READY   WorkflowStatus_Status = 2
	WorkflowStatus_FAILED  WorkflowStatus_Status = 3
	WorkflowStatus_DELETED WorkflowStatus_Status = 4
)

var WorkflowStatus_Status_name = map[int32]string{
	0: "PENDING",
	2: "READY",
	3: "FAILED",
	4: "DELETED",
}
var WorkflowStatus_Status_value = map[string]int32{
	"PENDING": 0,
	"READY":   2,
	"FAILED":  3,
	"DELETED": 4,
}

func (x WorkflowStatus_Status) String() string {
	return proto.EnumName(WorkflowStatus_Status_name, int32(x))
}
func (WorkflowStatus_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type WorkflowInvocationStatus_Status int32

const (
	WorkflowInvocationStatus_UNKNOWN     WorkflowInvocationStatus_Status = 0
	WorkflowInvocationStatus_SCHEDULED   WorkflowInvocationStatus_Status = 1
	WorkflowInvocationStatus_IN_PROGRESS WorkflowInvocationStatus_Status = 2
	WorkflowInvocationStatus_SUCCEEDED   WorkflowInvocationStatus_Status = 3
	WorkflowInvocationStatus_FAILED      WorkflowInvocationStatus_Status = 4
	WorkflowInvocationStatus_ABORTED     WorkflowInvocationStatus_Status = 5
)

var WorkflowInvocationStatus_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SCHEDULED",
	2: "IN_PROGRESS",
	3: "SUCCEEDED",
	4: "FAILED",
	5: "ABORTED",
}
var WorkflowInvocationStatus_Status_value = map[string]int32{
	"UNKNOWN":     0,
	"SCHEDULED":   1,
	"IN_PROGRESS": 2,
	"SUCCEEDED":   3,
	"FAILED":      4,
	"ABORTED":     5,
}

func (x WorkflowInvocationStatus_Status) String() string {
	return proto.EnumName(WorkflowInvocationStatus_Status_name, int32(x))
}
func (WorkflowInvocationStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

type TaskStatus_Status int32

const (
	TaskStatus_STARTED TaskStatus_Status = 0
	TaskStatus_READY   TaskStatus_Status = 1
	TaskStatus_FAILED  TaskStatus_Status = 2
)

var TaskStatus_Status_name = map[int32]string{
	0: "STARTED",
	1: "READY",
	2: "FAILED",
}
var TaskStatus_Status_value = map[string]int32{
	"STARTED": 0,
	"READY":   1,
	"FAILED":  2,
}

func (x TaskStatus_Status) String() string {
	return proto.EnumName(TaskStatus_Status_name, int32(x))
}
func (TaskStatus_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

type TaskDependencyParameters_DependencyType int32

const (
	TaskDependencyParameters_DATA           TaskDependencyParameters_DependencyType = 0
	TaskDependencyParameters_CONTROL        TaskDependencyParameters_DependencyType = 1
	TaskDependencyParameters_DYNAMIC_OUTPUT TaskDependencyParameters_DependencyType = 2
)

var TaskDependencyParameters_DependencyType_name = map[int32]string{
	0: "DATA",
	1: "CONTROL",
	2: "DYNAMIC_OUTPUT",
}
var TaskDependencyParameters_DependencyType_value = map[string]int32{
	"DATA":           0,
	"CONTROL":        1,
	"DYNAMIC_OUTPUT": 2,
}

func (x TaskDependencyParameters_DependencyType) String() string {
	return proto.EnumName(TaskDependencyParameters_DependencyType_name, int32(x))
}
func (TaskDependencyParameters_DependencyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

type TaskInvocationStatus_Status int32

const (
	TaskInvocationStatus_UNKNOWN     TaskInvocationStatus_Status = 0
	TaskInvocationStatus_SCHEDULED   TaskInvocationStatus_Status = 1
	TaskInvocationStatus_IN_PROGRESS TaskInvocationStatus_Status = 2
	TaskInvocationStatus_SUCCEEDED   TaskInvocationStatus_Status = 3
	TaskInvocationStatus_FAILED      TaskInvocationStatus_Status = 4
	TaskInvocationStatus_ABORTED     TaskInvocationStatus_Status = 5
	TaskInvocationStatus_SKIPPED     TaskInvocationStatus_Status = 6
)

var TaskInvocationStatus_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SCHEDULED",
	2: "IN_PROGRESS",
	3: "SUCCEEDED",
	4: "FAILED",
	5: "ABORTED",
	6: "SKIPPED",
}
var TaskInvocationStatus_Status_value = map[string]int32{
	"UNKNOWN":     0,
	"SCHEDULED":   1,
	"IN_PROGRESS": 2,
	"SUCCEEDED":   3,
	"FAILED":      4,
	"ABORTED":     5,
	"SKIPPED":     6,
}

func (x TaskInvocationStatus_Status) String() string {
	return proto.EnumName(TaskInvocationStatus_Status_name, int32(x))
}
func (TaskInvocationStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{13, 0}
}

//
// Workflow Model
//
type Workflow struct {
	Metadata *ObjectMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *WorkflowSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *WorkflowStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *Workflow) Reset()                    { *m = Workflow{} }
func (m *Workflow) String() string            { return proto.CompactTextString(m) }
func (*Workflow) ProtoMessage()               {}
func (*Workflow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Workflow) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Workflow) GetSpec() *WorkflowSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Workflow) GetStatus() *WorkflowStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Workflow Definition
//
// The workflowDefinition contains the definition of a workflow.
//
// Ideally the source code (json, yaml) can be converted directly to this message.
// Naming, triggers and versioning of the workflow itself is out of the scope of this data structure, which is delegated
// to the user/system upon the creation of a workflow.
type WorkflowSpec struct {
	// apiVersion describes what version is of the workflow definition.
	// By default the workflow engine will assume the latest version to be used.
	ApiVersion string `protobuf:"bytes,1,opt,name=apiVersion" json:"apiVersion,omitempty"`
	// Tasks contains the specs of the tasks, with the key being the task id.
	//
	// Note: Dependency graph is build into the tasks.
	Tasks map[string]*TaskSpec `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// From which task should the workflow return the output? Future: multiple? Implicit?
	OutputTask  string `protobuf:"bytes,3,opt,name=outputTask" json:"outputTask,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// The UID that the workflow should have. Only use this in case you want to force a specific UID.
	ForceId string `protobuf:"bytes,5,opt,name=forceId" json:"forceId,omitempty"`
	// Name is solely for human-readablity
	Name string `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	// Internal indicates whether is a workflow should be visible to a human (default) or not.
	//
	Internal bool `protobuf:"varint,7,opt,name=internal" json:"internal,omitempty"`
}

func (m *WorkflowSpec) Reset()                    { *m = WorkflowSpec{} }
func (m *WorkflowSpec) String() string            { return proto.CompactTextString(m) }
func (*WorkflowSpec) ProtoMessage()               {}
func (*WorkflowSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WorkflowSpec) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *WorkflowSpec) GetTasks() map[string]*TaskSpec {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *WorkflowSpec) GetOutputTask() string {
	if m != nil {
		return m.OutputTask
	}
	return ""
}

func (m *WorkflowSpec) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WorkflowSpec) GetForceId() string {
	if m != nil {
		return m.ForceId
	}
	return ""
}

func (m *WorkflowSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowSpec) GetInternal() bool {
	if m != nil {
		return m.Internal
	}
	return false
}

type WorkflowStatus struct {
	Status    WorkflowStatus_Status      `protobuf:"varint,1,opt,name=status,enum=WorkflowStatus_Status" json:"status,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=updatedAt" json:"updatedAt,omitempty"`
	// Tasks contains the status of the tasks, with the key being the task id.
	Tasks map[string]*TaskStatus `protobuf:"bytes,3,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Error *Error                 `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *WorkflowStatus) Reset()                    { *m = WorkflowStatus{} }
func (m *WorkflowStatus) String() string            { return proto.CompactTextString(m) }
func (*WorkflowStatus) ProtoMessage()               {}
func (*WorkflowStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WorkflowStatus) GetStatus() WorkflowStatus_Status {
	if m != nil {
		return m.Status
	}
	return WorkflowStatus_PENDING
}

func (m *WorkflowStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *WorkflowStatus) GetTasks() map[string]*TaskStatus {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *WorkflowStatus) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

//
// Workflow Invocation Model
//
type WorkflowInvocation struct {
	Metadata *ObjectMetadata           `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *WorkflowInvocationSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *WorkflowInvocationStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *WorkflowInvocation) Reset()                    { *m = WorkflowInvocation{} }
func (m *WorkflowInvocation) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocation) ProtoMessage()               {}
func (*WorkflowInvocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WorkflowInvocation) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *WorkflowInvocation) GetSpec() *WorkflowInvocationSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *WorkflowInvocation) GetStatus() *WorkflowInvocationStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Workflow Invocation Model
type WorkflowInvocationSpec struct {
	WorkflowId string                 `protobuf:"bytes,1,opt,name=workflowId" json:"workflowId,omitempty"`
	Inputs     map[string]*TypedValue `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// ParentId contains the id of the encapsulating workflow invocation.
	//
	// This used within the workflow engine; for user-provided workflow invocations the parentId is ignored.
	ParentId string `protobuf:"bytes,3,opt,name=parentId" json:"parentId,omitempty"`
}

func (m *WorkflowInvocationSpec) Reset()                    { *m = WorkflowInvocationSpec{} }
func (m *WorkflowInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationSpec) ProtoMessage()               {}
func (*WorkflowInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WorkflowInvocationSpec) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *WorkflowInvocationSpec) GetInputs() map[string]*TypedValue {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *WorkflowInvocationSpec) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

type WorkflowInvocationStatus struct {
	Status    WorkflowInvocationStatus_Status `protobuf:"varint,1,opt,name=status,enum=WorkflowInvocationStatus_Status" json:"status,omitempty"`
	UpdatedAt *google_protobuf.Timestamp      `protobuf:"bytes,2,opt,name=updatedAt" json:"updatedAt,omitempty"`
	Tasks     map[string]*TaskInvocation      `protobuf:"bytes,3,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Output    *TypedValue                     `protobuf:"bytes,4,opt,name=output" json:"output,omitempty"`
	// In case the task ID also exists in the workflow spec, the dynamic task will be
	// used as an overlay over the static task.
	DynamicTasks map[string]*Task `protobuf:"bytes,5,rep,name=dynamicTasks" json:"dynamicTasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Error        *Error           `protobuf:"bytes,6,opt,name=error" json:"error,omitempty"`
}

func (m *WorkflowInvocationStatus) Reset()                    { *m = WorkflowInvocationStatus{} }
func (m *WorkflowInvocationStatus) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationStatus) ProtoMessage()               {}
func (*WorkflowInvocationStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WorkflowInvocationStatus) GetStatus() WorkflowInvocationStatus_Status {
	if m != nil {
		return m.Status
	}
	return WorkflowInvocationStatus_UNKNOWN
}

func (m *WorkflowInvocationStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *WorkflowInvocationStatus) GetTasks() map[string]*TaskInvocation {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *WorkflowInvocationStatus) GetOutput() *TypedValue {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *WorkflowInvocationStatus) GetDynamicTasks() map[string]*Task {
	if m != nil {
		return m.DynamicTasks
	}
	return nil
}

func (m *WorkflowInvocationStatus) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type DependencyConfig struct {
	// Dependencies for this task to execute
	Requires map[string]*TaskDependencyParameters `protobuf:"bytes,1,rep,name=requires" json:"requires,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Number of dependencies to wait for
	Await int32 `protobuf:"varint,2,opt,name=await" json:"await,omitempty"`
}

func (m *DependencyConfig) Reset()                    { *m = DependencyConfig{} }
func (m *DependencyConfig) String() string            { return proto.CompactTextString(m) }
func (*DependencyConfig) ProtoMessage()               {}
func (*DependencyConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DependencyConfig) GetRequires() map[string]*TaskDependencyParameters {
	if m != nil {
		return m.Requires
	}
	return nil
}

func (m *DependencyConfig) GetAwait() int32 {
	if m != nil {
		return m.Await
	}
	return 0
}

//
// Task Model
//
type Task struct {
	Metadata *ObjectMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *TaskSpec       `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *TaskStatus     `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Task) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Task) GetSpec() *TaskSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Task) GetStatus() *TaskStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// A task is the primitive unit of a workflow, representing an action that needs to be performed in order to continue.
//
// A task as a number of inputs and exactly two outputs
// Id is specified outside of TaskSpec
type TaskSpec struct {
	// Name/identifier of the function
	FunctionRef string                 `protobuf:"bytes,1,opt,name=functionRef" json:"functionRef,omitempty"`
	Inputs      map[string]*TypedValue `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Dependencies for this task to execute
	Requires map[string]*TaskDependencyParameters `protobuf:"bytes,3,rep,name=requires" json:"requires,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Number of dependencies to wait for
	Await int32 `protobuf:"varint,4,opt,name=await" json:"await,omitempty"`
	// Transform the output, or override the output with a literal
	Output *TypedValue `protobuf:"bytes,5,opt,name=output" json:"output,omitempty"`
}

func (m *TaskSpec) Reset()                    { *m = TaskSpec{} }
func (m *TaskSpec) String() string            { return proto.CompactTextString(m) }
func (*TaskSpec) ProtoMessage()               {}
func (*TaskSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TaskSpec) GetFunctionRef() string {
	if m != nil {
		return m.FunctionRef
	}
	return ""
}

func (m *TaskSpec) GetInputs() map[string]*TypedValue {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TaskSpec) GetRequires() map[string]*TaskDependencyParameters {
	if m != nil {
		return m.Requires
	}
	return nil
}

func (m *TaskSpec) GetAwait() int32 {
	if m != nil {
		return m.Await
	}
	return 0
}

func (m *TaskSpec) GetOutput() *TypedValue {
	if m != nil {
		return m.Output
	}
	return nil
}

type TaskStatus struct {
	Status    TaskStatus_Status          `protobuf:"varint,1,opt,name=status,enum=TaskStatus_Status" json:"status,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=updatedAt" json:"updatedAt,omitempty"`
	FnRef     *FnRef                     `protobuf:"bytes,3,opt,name=fnRef" json:"fnRef,omitempty"`
	Error     *Error                     `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *TaskStatus) Reset()                    { *m = TaskStatus{} }
func (m *TaskStatus) String() string            { return proto.CompactTextString(m) }
func (*TaskStatus) ProtoMessage()               {}
func (*TaskStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TaskStatus) GetStatus() TaskStatus_Status {
	if m != nil {
		return m.Status
	}
	return TaskStatus_STARTED
}

func (m *TaskStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *TaskStatus) GetFnRef() *FnRef {
	if m != nil {
		return m.FnRef
	}
	return nil
}

func (m *TaskStatus) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type TaskDependencyParameters struct {
	Type  TaskDependencyParameters_DependencyType `protobuf:"varint,1,opt,name=type,enum=TaskDependencyParameters_DependencyType" json:"type,omitempty"`
	Alias string                                  `protobuf:"bytes,2,opt,name=alias" json:"alias,omitempty"`
}

func (m *TaskDependencyParameters) Reset()                    { *m = TaskDependencyParameters{} }
func (m *TaskDependencyParameters) String() string            { return proto.CompactTextString(m) }
func (*TaskDependencyParameters) ProtoMessage()               {}
func (*TaskDependencyParameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TaskDependencyParameters) GetType() TaskDependencyParameters_DependencyType {
	if m != nil {
		return m.Type
	}
	return TaskDependencyParameters_DATA
}

func (m *TaskDependencyParameters) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

//
// Task Invocation Model
//
type TaskInvocation struct {
	Metadata *ObjectMetadata       `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *TaskInvocationSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *TaskInvocationStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *TaskInvocation) Reset()                    { *m = TaskInvocation{} }
func (m *TaskInvocation) String() string            { return proto.CompactTextString(m) }
func (*TaskInvocation) ProtoMessage()               {}
func (*TaskInvocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *TaskInvocation) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TaskInvocation) GetSpec() *TaskInvocationSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *TaskInvocation) GetStatus() *TaskInvocationStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type TaskInvocationSpec struct {
	// Id of the task to be invoked (no ambiguatity at this point
	FnRef  *FnRef                 `protobuf:"bytes,1,opt,name=fnRef" json:"fnRef,omitempty"`
	TaskId string                 `protobuf:"bytes,2,opt,name=taskId" json:"taskId,omitempty"`
	Inputs map[string]*TypedValue `protobuf:"bytes,3,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TaskInvocationSpec) Reset()                    { *m = TaskInvocationSpec{} }
func (m *TaskInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*TaskInvocationSpec) ProtoMessage()               {}
func (*TaskInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *TaskInvocationSpec) GetFnRef() *FnRef {
	if m != nil {
		return m.FnRef
	}
	return nil
}

func (m *TaskInvocationSpec) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TaskInvocationSpec) GetInputs() map[string]*TypedValue {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type TaskInvocationStatus struct {
	Status    TaskInvocationStatus_Status `protobuf:"varint,1,opt,name=status,enum=TaskInvocationStatus_Status" json:"status,omitempty"`
	UpdatedAt *google_protobuf.Timestamp  `protobuf:"bytes,2,opt,name=updatedAt" json:"updatedAt,omitempty"`
	Output    *TypedValue                 `protobuf:"bytes,3,opt,name=output" json:"output,omitempty"`
	Error     *Error                      `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *TaskInvocationStatus) Reset()                    { *m = TaskInvocationStatus{} }
func (m *TaskInvocationStatus) String() string            { return proto.CompactTextString(m) }
func (*TaskInvocationStatus) ProtoMessage()               {}
func (*TaskInvocationStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *TaskInvocationStatus) GetStatus() TaskInvocationStatus_Status {
	if m != nil {
		return m.Status
	}
	return TaskInvocationStatus_UNKNOWN
}

func (m *TaskInvocationStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *TaskInvocationStatus) GetOutput() *TypedValue {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *TaskInvocationStatus) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

//
// Common
//
type ObjectMetadata struct {
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=createdAt" json:"createdAt,omitempty"`
}

func (m *ObjectMetadata) Reset()                    { *m = ObjectMetadata{} }
func (m *ObjectMetadata) String() string            { return proto.CompactTextString(m) }
func (*ObjectMetadata) ProtoMessage()               {}
func (*ObjectMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ObjectMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ObjectMetadata) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// Copy of protobuf's Any, to avoid protobuf requirement of a protobuf-based type.
type TypedValue struct {
	Type  string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *TypedValue) Reset()                    { *m = TypedValue{} }
func (m *TypedValue) String() string            { return proto.CompactTextString(m) }
func (*TypedValue) ProtoMessage()               {}
func (*TypedValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *TypedValue) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TypedValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Error struct {
	Code    string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// FnRef is an immutable, unique reference to a function on a specific function runtime environment.
//
// The string representation (via String or Format): runtime://runtimeId
type FnRef struct {
	// Runtime is the Function Runtime environment (fnenv) that was used to resolve the function.
	Runtime string `protobuf:"bytes,2,opt,name=runtime" json:"runtime,omitempty"`
	// RuntimeId is the runtime-specific identifier of the function.
	RuntimeId string `protobuf:"bytes,3,opt,name=runtimeId" json:"runtimeId,omitempty"`
}

func (m *FnRef) Reset()                    { *m = FnRef{} }
func (m *FnRef) String() string            { return proto.CompactTextString(m) }
func (*FnRef) ProtoMessage()               {}
func (*FnRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *FnRef) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *FnRef) GetRuntimeId() string {
	if m != nil {
		return m.RuntimeId
	}
	return ""
}

func init() {
	proto.RegisterType((*Workflow)(nil), "Workflow")
	proto.RegisterType((*WorkflowSpec)(nil), "WorkflowSpec")
	proto.RegisterType((*WorkflowStatus)(nil), "WorkflowStatus")
	proto.RegisterType((*WorkflowInvocation)(nil), "WorkflowInvocation")
	proto.RegisterType((*WorkflowInvocationSpec)(nil), "WorkflowInvocationSpec")
	proto.RegisterType((*WorkflowInvocationStatus)(nil), "WorkflowInvocationStatus")
	proto.RegisterType((*DependencyConfig)(nil), "DependencyConfig")
	proto.RegisterType((*Task)(nil), "Task")
	proto.RegisterType((*TaskSpec)(nil), "TaskSpec")
	proto.RegisterType((*TaskStatus)(nil), "TaskStatus")
	proto.RegisterType((*TaskDependencyParameters)(nil), "TaskDependencyParameters")
	proto.RegisterType((*TaskInvocation)(nil), "TaskInvocation")
	proto.RegisterType((*TaskInvocationSpec)(nil), "TaskInvocationSpec")
	proto.RegisterType((*TaskInvocationStatus)(nil), "TaskInvocationStatus")
	proto.RegisterType((*ObjectMetadata)(nil), "ObjectMetadata")
	proto.RegisterType((*TypedValue)(nil), "TypedValue")
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterType((*FnRef)(nil), "FnRef")
	proto.RegisterEnum("WorkflowStatus_Status", WorkflowStatus_Status_name, WorkflowStatus_Status_value)
	proto.RegisterEnum("WorkflowInvocationStatus_Status", WorkflowInvocationStatus_Status_name, WorkflowInvocationStatus_Status_value)
	proto.RegisterEnum("TaskStatus_Status", TaskStatus_Status_name, TaskStatus_Status_value)
	proto.RegisterEnum("TaskDependencyParameters_DependencyType", TaskDependencyParameters_DependencyType_name, TaskDependencyParameters_DependencyType_value)
	proto.RegisterEnum("TaskInvocationStatus_Status", TaskInvocationStatus_Status_name, TaskInvocationStatus_Status_value)
}

func init() { proto.RegisterFile("pkg/types/types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0x29, 0x51, 0x96, 0x46, 0x89, 0xc2, 0x6e, 0xf3, 0xc3, 0xa8, 0x6e, 0xed, 0xd0, 0x2d,
	0x62, 0xd4, 0x0d, 0xdd, 0x3a, 0xfd, 0x71, 0xe3, 0x02, 0x85, 0x22, 0xd2, 0xa9, 0x10, 0x47, 0x12,
	0x68, 0xd9, 0x41, 0x02, 0x14, 0x01, 0x2d, 0xae, 0x0c, 0xc6, 0x16, 0xc9, 0x92, 0x54, 0x0c, 0x9d,
	0xfb, 0x0c, 0x7d, 0x80, 0x9e, 0x7b, 0xe8, 0xa9, 0xc7, 0x3e, 0x41, 0x7b, 0xe8, 0xb5, 0x40, 0xdf,
	0xa0, 0x2f, 0x51, 0xec, 0x72, 0x49, 0x2e, 0x29, 0x29, 0xae, 0x01, 0xe7, 0x22, 0x70, 0x66, 0x67,
	0x67, 0x76, 0xbe, 0x99, 0x6f, 0x56, 0x0b, 0x37, 0xfd, 0x93, 0xe3, 0xcd, 0x68, 0xea, 0xe3, 0x30,
	0xfe, 0xd5, 0xfc, 0xc0, 0x8b, 0xbc, 0xe6, 0xca, 0xb1, 0xe7, 0x1d, 0x9f, 0xe2, 0x4d, 0x2a, 0x1d,
	0x4d, 0x46, 0x9b, 0x91, 0x33, 0xc6, 0x61, 0x64, 0x8d, 0xfd, 0xd8, 0x40, 0xfd, 0x51, 0x80, 0xea,
	0x33, 0x2f, 0x38, 0x19, 0x9d, 0x7a, 0x67, 0x68, 0x03, 0xaa, 0x63, 0x1c, 0x59, 0xb6, 0x15, 0x59,
	0x8a, 0xb0, 0x2a, 0xac, 0xd7, 0xb7, 0xae, 0x6b, 0xbd, 0xa3, 0x57, 0x78, 0x18, 0x3d, 0x65, 0x6a,
	0x33, 0x35, 0x40, 0x77, 0xa1, 0x1c, 0xfa, 0x78, 0xa8, 0x88, 0xd4, 0xf0, 0x9a, 0x96, 0x78, 0xd9,
	0xf7, 0xf1, 0xd0, 0xa4, 0x4b, 0xe8, 0x1e, 0x54, 0xc2, 0xc8, 0x8a, 0x26, 0xa1, 0x52, 0x62, 0xde,
	0x52, 0x23, 0xaa, 0x36, 0xd9, 0xb2, 0xfa, 0xab, 0x08, 0x57, 0xf9, 0xfd, 0xe8, 0x03, 0x00, 0xcb,
	0x77, 0x0e, 0x71, 0x10, 0x3a, 0x9e, 0x4b, 0xcf, 0x52, 0x33, 0x39, 0x0d, 0xd2, 0x40, 0x8a, 0xac,
	0xf0, 0x24, 0x54, 0xc4, 0xd5, 0xd2, 0x7a, 0x7d, 0x4b, 0xc9, 0x45, 0xd7, 0x06, 0x64, 0xc9, 0x70,
	0xa3, 0x60, 0x6a, 0xc6, 0x66, 0xc4, 0x9f, 0x37, 0x89, 0xfc, 0x49, 0x44, 0x96, 0xe8, 0x69, 0x6a,
	0x26, 0xa7, 0x41, 0xab, 0x50, 0xb7, 0x71, 0x38, 0x0c, 0x1c, 0x3f, 0x22, 0x01, 0xcb, 0xd4, 0x80,
	0x57, 0x21, 0x05, 0x96, 0x46, 0x5e, 0x30, 0xc4, 0x1d, 0x5b, 0x91, 0xe8, 0x6a, 0x22, 0x22, 0x04,
	0x65, 0xd7, 0x1a, 0x63, 0xa5, 0x42, 0xd5, 0xf4, 0x1b, 0x35, 0xa1, 0xea, 0xb8, 0x11, 0x0e, 0x5c,
	0xeb, 0x54, 0x59, 0x5a, 0x15, 0xd6, 0xab, 0x66, 0x2a, 0x37, 0xdb, 0x00, 0xd9, 0x01, 0x91, 0x0c,
	0xa5, 0x13, 0x3c, 0x65, 0x29, 0x92, 0x4f, 0xb4, 0x02, 0xd2, 0x6b, 0xeb, 0x74, 0x82, 0x19, 0xb2,
	0x35, 0x9a, 0x0e, 0x45, 0x35, 0xd6, 0x3f, 0x14, 0xb7, 0x05, 0xf5, 0x4f, 0x11, 0x1a, 0x79, 0x30,
	0x91, 0x96, 0xa2, 0x4d, 0x9c, 0x35, 0xb6, 0x6e, 0x15, 0xd0, 0xd6, 0xf2, 0xa0, 0xa3, 0x6d, 0xa8,
	0x4d, 0x7c, 0xdb, 0x8a, 0xb0, 0xdd, 0x8a, 0x58, 0xac, 0xa6, 0x16, 0xf7, 0x8b, 0x96, 0xf4, 0x8b,
	0x36, 0x48, 0xfa, 0xc5, 0xcc, 0x8c, 0xd1, 0xa7, 0x09, 0xfa, 0x25, 0x8a, 0x7e, 0xb3, 0x18, 0x68,
	0x16, 0xff, 0x65, 0x90, 0x70, 0x10, 0x78, 0x01, 0x45, 0xb6, 0xbe, 0x55, 0xd1, 0x0c, 0x22, 0x99,
	0xb1, 0xb2, 0x69, 0x9c, 0x83, 0xc8, 0xdd, 0x3c, 0x22, 0xf5, 0x18, 0x91, 0x38, 0x1b, 0x0e, 0x93,
	0xaf, 0xa1, 0xc2, 0xa0, 0xa8, 0xc3, 0x52, 0xdf, 0xe8, 0xea, 0x9d, 0xee, 0x63, 0xf9, 0x0a, 0xaa,
	0x81, 0x64, 0x1a, 0x2d, 0xfd, 0xb9, 0x2c, 0x22, 0x80, 0xca, 0x6e, 0xab, 0xb3, 0x67, 0xe8, 0x72,
	0x89, 0xd8, 0xe8, 0xc6, 0x9e, 0x31, 0x30, 0x74, 0xb9, 0xac, 0xfe, 0x2c, 0x00, 0x4a, 0x92, 0xe8,
	0xb8, 0xaf, 0xbd, 0xa1, 0x45, 0x8b, 0x7e, 0x21, 0x42, 0x6c, 0xe4, 0x08, 0x71, 0x5b, 0x9b, 0xf5,
	0xc7, 0x51, 0xe3, 0xb3, 0x02, 0x35, 0xee, 0xcc, 0x33, 0xcf, 0x93, 0xe4, 0x6f, 0x01, 0x6e, 0xcd,
	0xf7, 0x49, 0xda, 0xfb, 0x2c, 0x59, 0xb1, 0x13, 0xba, 0x64, 0x1a, 0xb4, 0x03, 0x15, 0xc7, 0xf5,
	0x27, 0x51, 0xc2, 0x97, 0xb5, 0x05, 0x87, 0xd3, 0x3a, 0xd4, 0x2a, 0x2e, 0x1d, 0xdb, 0x42, 0x7a,
	0xd9, 0xb7, 0x02, 0xec, 0x46, 0x1d, 0x9b, 0x31, 0x27, 0x95, 0x9b, 0xbb, 0x50, 0xe7, 0xb6, 0xfc,
	0xaf, 0xd2, 0x4d, 0x7d, 0x6c, 0x1f, 0x12, 0x15, 0x5f, 0xba, 0xbf, 0xca, 0xa0, 0x2c, 0x02, 0x00,
	0x6d, 0x17, 0x1a, 0x7b, 0x75, 0x21, 0x56, 0x97, 0xd7, 0xe2, 0x0f, 0xf3, 0x2d, 0xfe, 0xe1, 0xe2,
	0x90, 0xb3, 0xcd, 0xbe, 0x06, 0x95, 0x78, 0xb4, 0xb0, 0x6e, 0xcf, 0x25, 0xcd, 0x96, 0x50, 0x0f,
	0xae, 0xda, 0x53, 0xd7, 0x1a, 0x3b, 0x43, 0xea, 0x40, 0x91, 0x68, 0x9c, 0x8d, 0xc5, 0x71, 0x74,
	0xce, 0x3a, 0x0e, 0x97, 0x73, 0x90, 0x51, 0xac, 0x32, 0x8f, 0x62, 0x9d, 0x73, 0x28, 0xf6, 0x51,
	0xbe, 0x4e, 0xd7, 0x69, 0x5a, 0xd9, 0x19, 0xb8, 0x5a, 0x35, 0x77, 0xe1, 0x9d, 0x99, 0xb3, 0xcc,
	0xf1, 0xf8, 0x5e, 0xde, 0xa3, 0x44, 0x3d, 0xf2, 0x35, 0xff, 0x9e, 0xa7, 0xeb, 0x41, 0xf7, 0x49,
	0xb7, 0xf7, 0xac, 0x2b, 0x5f, 0x41, 0xd7, 0xa0, 0xb6, 0xdf, 0xfe, 0xce, 0xd0, 0x0f, 0x08, 0x4d,
	0x05, 0x74, 0x1d, 0xea, 0x9d, 0xee, 0xcb, 0xbe, 0xd9, 0x7b, 0x6c, 0x1a, 0xfb, 0xfb, 0xb2, 0x48,
	0xd7, 0x0f, 0xda, 0x6d, 0xc3, 0xd0, 0x29, 0x8d, 0x33, 0x4a, 0x97, 0x89, 0x9f, 0xd6, 0xa3, 0x9e,
	0x49, 0x28, 0x2d, 0xa9, 0xbf, 0x0b, 0x20, 0xeb, 0xd8, 0xc7, 0xae, 0x8d, 0xdd, 0xe1, 0xb4, 0xed,
	0xb9, 0x23, 0xe7, 0x18, 0xed, 0x40, 0x35, 0xc0, 0x3f, 0x4c, 0x9c, 0x00, 0x93, 0x66, 0x22, 0x88,
	0xaf, 0x68, 0x45, 0x23, 0xcd, 0x64, 0x16, 0x31, 0xca, 0xe9, 0x06, 0x74, 0x03, 0x24, 0xeb, 0xcc,
	0x72, 0xe2, 0x4e, 0x92, 0xcc, 0x58, 0x68, 0x1e, 0xc2, 0xb5, 0xdc, 0x86, 0x39, 0x50, 0x6c, 0xe6,
	0xa1, 0xb8, 0x43, 0xa1, 0xc8, 0xc2, 0xf6, 0xad, 0xc0, 0x1a, 0xe3, 0x08, 0x07, 0xb9, 0x69, 0x76,
	0x06, 0x65, 0x7a, 0x35, 0x5d, 0x68, 0x06, 0xbd, 0x9f, 0x9b, 0x41, 0xdc, 0xd5, 0x11, 0x4f, 0x9d,
	0xb5, 0xc2, 0xd4, 0xc9, 0x4d, 0xd2, 0x64, 0xce, 0xfc, 0x2b, 0x42, 0x35, 0xd9, 0x47, 0x2e, 0xc6,
	0xd1, 0xc4, 0x1d, 0xd2, 0x1e, 0xc0, 0x23, 0x96, 0x14, 0xaf, 0x42, 0xf7, 0x0b, 0xb3, 0xe5, 0x66,
	0x1a, 0x74, 0xee, 0x34, 0x79, 0xc0, 0x55, 0x20, 0xe6, 0xd6, 0xed, 0x6c, 0xc3, 0xb9, 0xc8, 0x97,
	0x39, 0xe4, 0x39, 0x9e, 0x49, 0x0b, 0x79, 0x76, 0x59, 0x13, 0xea, 0xad, 0x95, 0xf9, 0x1f, 0x21,
	0x66, 0x26, 0xa3, 0xc2, 0xc7, 0x85, 0x59, 0x87, 0xb8, 0x0a, 0x5d, 0xde, 0x74, 0x5b, 0x06, 0x69,
	0x44, 0xeb, 0x59, 0x62, 0xb3, 0x62, 0x97, 0x48, 0x66, 0xac, 0x7c, 0xf3, 0x65, 0xad, 0x7e, 0xc2,
	0xd3, 0x76, 0x7f, 0xd0, 0xa2, 0x74, 0xe3, 0x6e, 0x59, 0x81, 0xa3, 0xa4, 0xa8, 0xfe, 0x22, 0x80,
	0xb2, 0x08, 0x06, 0xf4, 0x0d, 0x94, 0xc9, 0x9f, 0x55, 0x96, 0xea, 0xfa, 0x42, 0xbc, 0x38, 0x8a,
	0x92, 0xe2, 0x98, 0x74, 0x17, 0x6d, 0x8a, 0x53, 0xc7, 0x0a, 0x69, 0xea, 0x35, 0x33, 0x16, 0xd4,
	0x1d, 0x68, 0xe4, 0xad, 0x51, 0x15, 0xca, 0x7a, 0x6b, 0xd0, 0x92, 0xaf, 0x90, 0x03, 0xb7, 0x7b,
	0xdd, 0x81, 0xd9, 0xdb, 0x93, 0x05, 0x84, 0xa0, 0xa1, 0x3f, 0xef, 0xb6, 0x9e, 0x76, 0xda, 0x2f,
	0x7b, 0x07, 0x83, 0xfe, 0xc1, 0x40, 0x16, 0xd5, 0x9f, 0x04, 0x68, 0xe4, 0x07, 0xdf, 0xc5, 0xe8,
	0x77, 0x2f, 0x47, 0xbf, 0x77, 0x0b, 0x43, 0x94, 0x23, 0xe2, 0xfd, 0x02, 0x11, 0x6f, 0x16, 0x4d,
	0xf3, 0x94, 0xfc, 0x43, 0x00, 0x34, 0xeb, 0x2b, 0x2b, 0xa3, 0x30, 0xaf, 0x8c, 0xb7, 0xa0, 0x42,
	0xee, 0xa3, 0x8e, 0xcd, 0x00, 0x62, 0x12, 0xfa, 0x2a, 0x25, 0x6c, 0x89, 0x4d, 0xc0, 0x59, 0xd7,
	0xf3, 0xa8, 0x7b, 0x69, 0x97, 0xfd, 0x6f, 0x22, 0xdc, 0x98, 0x97, 0x2e, 0xfa, 0xbc, 0xd0, 0xfc,
	0xcb, 0x73, 0x51, 0xb9, 0x3c, 0x1a, 0x64, 0x03, 0xa4, 0xb4, 0xf8, 0xa2, 0x7e, 0x33, 0x1b, 0x5e,
	0xbd, 0xd5, 0x4b, 0x8c, 0x52, 0xec, 0x49, 0xa7, 0xdf, 0x37, 0x74, 0xb9, 0xa2, 0xbe, 0x80, 0x46,
	0xbe, 0xf3, 0x50, 0x03, 0x44, 0x27, 0xf9, 0xbf, 0x27, 0x3a, 0x36, 0x81, 0x62, 0x18, 0x60, 0x06,
	0x45, 0xe9, 0x7c, 0x28, 0x52, 0x63, 0xf5, 0x4b, 0x80, 0x2c, 0x77, 0xf2, 0xa4, 0x49, 0x89, 0x59,
	0xcb, 0xe8, 0x96, 0x15, 0xf7, 0x2a, 0xab, 0xa7, 0xfa, 0x05, 0x48, 0x14, 0x0f, 0xb2, 0x65, 0xe8,
	0xd9, 0xe9, 0x16, 0xf2, 0x4d, 0xde, 0x4c, 0x63, 0x1c, 0x86, 0xd6, 0x31, 0x66, 0x2d, 0x98, 0x88,
	0xea, 0xb7, 0x20, 0xd1, 0x5e, 0x25, 0x26, 0xc1, 0xc4, 0x25, 0xaf, 0xd2, 0xc4, 0x84, 0x89, 0x68,
	0x19, 0x6a, 0xec, 0x33, 0xfd, 0xdf, 0x99, 0x29, 0x1e, 0x2d, 0xbd, 0x90, 0xe8, 0x3b, 0xf7, 0xa8,
	0x42, 0xf3, 0x7a, 0xf0, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x69, 0xe2, 0x1e, 0x01, 0x0f,
	0x00, 0x00,
}
