// The MIT License (MIT)
// 
// Copyright (c) 2017 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by thriftrw-plugin-yarpc
// @generated

package historyserviceclient

import (
	context "context"
	history "github.com/uber/cadence/.gen/go/history"
	shared "github.com/uber/cadence/.gen/go/shared"
	wire "go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	reflect "reflect"
)

// Interface is a client for the HistoryService service.
type Interface interface {
	DescribeHistoryHost(
		ctx context.Context,
		Request *shared.DescribeHistoryHostRequest,
		opts ...yarpc.CallOption,
	) (*shared.DescribeHistoryHostResponse, error)

	DescribeMutableState(
		ctx context.Context,
		Request *history.DescribeMutableStateRequest,
		opts ...yarpc.CallOption,
	) (*history.DescribeMutableStateResponse, error)

	DescribeWorkflowExecution(
		ctx context.Context,
		DescribeRequest *history.DescribeWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) (*shared.DescribeWorkflowExecutionResponse, error)

	GetMutableState(
		ctx context.Context,
		GetRequest *history.GetMutableStateRequest,
		opts ...yarpc.CallOption,
	) (*history.GetMutableStateResponse, error)

	RecordActivityTaskHeartbeat(
		ctx context.Context,
		HeartbeatRequest *history.RecordActivityTaskHeartbeatRequest,
		opts ...yarpc.CallOption,
	) (*shared.RecordActivityTaskHeartbeatResponse, error)

	RecordActivityTaskStarted(
		ctx context.Context,
		AddRequest *history.RecordActivityTaskStartedRequest,
		opts ...yarpc.CallOption,
	) (*history.RecordActivityTaskStartedResponse, error)

	RecordChildExecutionCompleted(
		ctx context.Context,
		CompletionRequest *history.RecordChildExecutionCompletedRequest,
		opts ...yarpc.CallOption,
	) error

	RecordDecisionTaskStarted(
		ctx context.Context,
		AddRequest *history.RecordDecisionTaskStartedRequest,
		opts ...yarpc.CallOption,
	) (*history.RecordDecisionTaskStartedResponse, error)

	RemoveSignalMutableState(
		ctx context.Context,
		RemoveRequest *history.RemoveSignalMutableStateRequest,
		opts ...yarpc.CallOption,
	) error

	ReplicateEvents(
		ctx context.Context,
		ReplicateRequest *history.ReplicateEventsRequest,
		opts ...yarpc.CallOption,
	) error

	ReplicateRawEvents(
		ctx context.Context,
		ReplicateRequest *history.ReplicateRawEventsRequest,
		opts ...yarpc.CallOption,
	) error

	RequestCancelWorkflowExecution(
		ctx context.Context,
		CancelRequest *history.RequestCancelWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) error

	ResetStickyTaskList(
		ctx context.Context,
		ResetRequest *history.ResetStickyTaskListRequest,
		opts ...yarpc.CallOption,
	) (*history.ResetStickyTaskListResponse, error)

	ResetWorkflowExecution(
		ctx context.Context,
		ResetRequest *history.ResetWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) (*shared.ResetWorkflowExecutionResponse, error)

	RespondActivityTaskCanceled(
		ctx context.Context,
		CanceledRequest *history.RespondActivityTaskCanceledRequest,
		opts ...yarpc.CallOption,
	) error

	RespondActivityTaskCompleted(
		ctx context.Context,
		CompleteRequest *history.RespondActivityTaskCompletedRequest,
		opts ...yarpc.CallOption,
	) error

	RespondActivityTaskFailed(
		ctx context.Context,
		FailRequest *history.RespondActivityTaskFailedRequest,
		opts ...yarpc.CallOption,
	) error

	RespondDecisionTaskCompleted(
		ctx context.Context,
		CompleteRequest *history.RespondDecisionTaskCompletedRequest,
		opts ...yarpc.CallOption,
	) (*history.RespondDecisionTaskCompletedResponse, error)

	RespondDecisionTaskFailed(
		ctx context.Context,
		FailedRequest *history.RespondDecisionTaskFailedRequest,
		opts ...yarpc.CallOption,
	) error

	ScheduleDecisionTask(
		ctx context.Context,
		ScheduleRequest *history.ScheduleDecisionTaskRequest,
		opts ...yarpc.CallOption,
	) error

	SignalWithStartWorkflowExecution(
		ctx context.Context,
		SignalWithStartRequest *history.SignalWithStartWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) (*shared.StartWorkflowExecutionResponse, error)

	SignalWorkflowExecution(
		ctx context.Context,
		SignalRequest *history.SignalWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) error

	StartWorkflowExecution(
		ctx context.Context,
		StartRequest *history.StartWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) (*shared.StartWorkflowExecutionResponse, error)

	SyncActivity(
		ctx context.Context,
		SyncActivityRequest *history.SyncActivityRequest,
		opts ...yarpc.CallOption,
	) error

	SyncShardStatus(
		ctx context.Context,
		SyncShardStatusRequest *history.SyncShardStatusRequest,
		opts ...yarpc.CallOption,
	) error

	TerminateWorkflowExecution(
		ctx context.Context,
		TerminateRequest *history.TerminateWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) error
}

// New builds a new client for the HistoryService service.
//
// 	client := historyserviceclient.New(dispatcher.ClientConfig("historyservice"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "HistoryService",
			ClientConfig: c,
		}, opts...),
	}
}

func init() {
	yarpc.RegisterClientBuilder(
		func(c transport.ClientConfig, f reflect.StructField) Interface {
			return New(c, thrift.ClientBuilderOptions(c, f)...)
		},
	)
}

type client struct {
	c thrift.Client
}

func (c client) DescribeHistoryHost(
	ctx context.Context,
	_Request *shared.DescribeHistoryHostRequest,
	opts ...yarpc.CallOption,
) (success *shared.DescribeHistoryHostResponse, err error) {

	args := history.HistoryService_DescribeHistoryHost_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_DescribeHistoryHost_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_DescribeHistoryHost_Helper.UnwrapResponse(&result)
	return
}

func (c client) DescribeMutableState(
	ctx context.Context,
	_Request *history.DescribeMutableStateRequest,
	opts ...yarpc.CallOption,
) (success *history.DescribeMutableStateResponse, err error) {

	args := history.HistoryService_DescribeMutableState_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_DescribeMutableState_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_DescribeMutableState_Helper.UnwrapResponse(&result)
	return
}

func (c client) DescribeWorkflowExecution(
	ctx context.Context,
	_DescribeRequest *history.DescribeWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.DescribeWorkflowExecutionResponse, err error) {

	args := history.HistoryService_DescribeWorkflowExecution_Helper.Args(_DescribeRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_DescribeWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_DescribeWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) GetMutableState(
	ctx context.Context,
	_GetRequest *history.GetMutableStateRequest,
	opts ...yarpc.CallOption,
) (success *history.GetMutableStateResponse, err error) {

	args := history.HistoryService_GetMutableState_Helper.Args(_GetRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_GetMutableState_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_GetMutableState_Helper.UnwrapResponse(&result)
	return
}

func (c client) RecordActivityTaskHeartbeat(
	ctx context.Context,
	_HeartbeatRequest *history.RecordActivityTaskHeartbeatRequest,
	opts ...yarpc.CallOption,
) (success *shared.RecordActivityTaskHeartbeatResponse, err error) {

	args := history.HistoryService_RecordActivityTaskHeartbeat_Helper.Args(_HeartbeatRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RecordActivityTaskHeartbeat_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_RecordActivityTaskHeartbeat_Helper.UnwrapResponse(&result)
	return
}

func (c client) RecordActivityTaskStarted(
	ctx context.Context,
	_AddRequest *history.RecordActivityTaskStartedRequest,
	opts ...yarpc.CallOption,
) (success *history.RecordActivityTaskStartedResponse, err error) {

	args := history.HistoryService_RecordActivityTaskStarted_Helper.Args(_AddRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RecordActivityTaskStarted_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_RecordActivityTaskStarted_Helper.UnwrapResponse(&result)
	return
}

func (c client) RecordChildExecutionCompleted(
	ctx context.Context,
	_CompletionRequest *history.RecordChildExecutionCompletedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RecordChildExecutionCompleted_Helper.Args(_CompletionRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RecordChildExecutionCompleted_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RecordChildExecutionCompleted_Helper.UnwrapResponse(&result)
	return
}

func (c client) RecordDecisionTaskStarted(
	ctx context.Context,
	_AddRequest *history.RecordDecisionTaskStartedRequest,
	opts ...yarpc.CallOption,
) (success *history.RecordDecisionTaskStartedResponse, err error) {

	args := history.HistoryService_RecordDecisionTaskStarted_Helper.Args(_AddRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RecordDecisionTaskStarted_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_RecordDecisionTaskStarted_Helper.UnwrapResponse(&result)
	return
}

func (c client) RemoveSignalMutableState(
	ctx context.Context,
	_RemoveRequest *history.RemoveSignalMutableStateRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RemoveSignalMutableState_Helper.Args(_RemoveRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RemoveSignalMutableState_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RemoveSignalMutableState_Helper.UnwrapResponse(&result)
	return
}

func (c client) ReplicateEvents(
	ctx context.Context,
	_ReplicateRequest *history.ReplicateEventsRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_ReplicateEvents_Helper.Args(_ReplicateRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ReplicateEvents_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_ReplicateEvents_Helper.UnwrapResponse(&result)
	return
}

func (c client) ReplicateRawEvents(
	ctx context.Context,
	_ReplicateRequest *history.ReplicateRawEventsRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_ReplicateRawEvents_Helper.Args(_ReplicateRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ReplicateRawEvents_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_ReplicateRawEvents_Helper.UnwrapResponse(&result)
	return
}

func (c client) RequestCancelWorkflowExecution(
	ctx context.Context,
	_CancelRequest *history.RequestCancelWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RequestCancelWorkflowExecution_Helper.Args(_CancelRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RequestCancelWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RequestCancelWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) ResetStickyTaskList(
	ctx context.Context,
	_ResetRequest *history.ResetStickyTaskListRequest,
	opts ...yarpc.CallOption,
) (success *history.ResetStickyTaskListResponse, err error) {

	args := history.HistoryService_ResetStickyTaskList_Helper.Args(_ResetRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ResetStickyTaskList_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_ResetStickyTaskList_Helper.UnwrapResponse(&result)
	return
}

func (c client) ResetWorkflowExecution(
	ctx context.Context,
	_ResetRequest *history.ResetWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.ResetWorkflowExecutionResponse, err error) {

	args := history.HistoryService_ResetWorkflowExecution_Helper.Args(_ResetRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ResetWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_ResetWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) RespondActivityTaskCanceled(
	ctx context.Context,
	_CanceledRequest *history.RespondActivityTaskCanceledRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RespondActivityTaskCanceled_Helper.Args(_CanceledRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RespondActivityTaskCanceled_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RespondActivityTaskCanceled_Helper.UnwrapResponse(&result)
	return
}

func (c client) RespondActivityTaskCompleted(
	ctx context.Context,
	_CompleteRequest *history.RespondActivityTaskCompletedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RespondActivityTaskCompleted_Helper.Args(_CompleteRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RespondActivityTaskCompleted_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RespondActivityTaskCompleted_Helper.UnwrapResponse(&result)
	return
}

func (c client) RespondActivityTaskFailed(
	ctx context.Context,
	_FailRequest *history.RespondActivityTaskFailedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RespondActivityTaskFailed_Helper.Args(_FailRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RespondActivityTaskFailed_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RespondActivityTaskFailed_Helper.UnwrapResponse(&result)
	return
}

func (c client) RespondDecisionTaskCompleted(
	ctx context.Context,
	_CompleteRequest *history.RespondDecisionTaskCompletedRequest,
	opts ...yarpc.CallOption,
) (success *history.RespondDecisionTaskCompletedResponse, err error) {

	args := history.HistoryService_RespondDecisionTaskCompleted_Helper.Args(_CompleteRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RespondDecisionTaskCompleted_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_RespondDecisionTaskCompleted_Helper.UnwrapResponse(&result)
	return
}

func (c client) RespondDecisionTaskFailed(
	ctx context.Context,
	_FailedRequest *history.RespondDecisionTaskFailedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RespondDecisionTaskFailed_Helper.Args(_FailedRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RespondDecisionTaskFailed_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RespondDecisionTaskFailed_Helper.UnwrapResponse(&result)
	return
}

func (c client) ScheduleDecisionTask(
	ctx context.Context,
	_ScheduleRequest *history.ScheduleDecisionTaskRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_ScheduleDecisionTask_Helper.Args(_ScheduleRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ScheduleDecisionTask_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_ScheduleDecisionTask_Helper.UnwrapResponse(&result)
	return
}

func (c client) SignalWithStartWorkflowExecution(
	ctx context.Context,
	_SignalWithStartRequest *history.SignalWithStartWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.StartWorkflowExecutionResponse, err error) {

	args := history.HistoryService_SignalWithStartWorkflowExecution_Helper.Args(_SignalWithStartRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_SignalWithStartWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_SignalWithStartWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) SignalWorkflowExecution(
	ctx context.Context,
	_SignalRequest *history.SignalWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_SignalWorkflowExecution_Helper.Args(_SignalRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_SignalWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_SignalWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) StartWorkflowExecution(
	ctx context.Context,
	_StartRequest *history.StartWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.StartWorkflowExecutionResponse, err error) {

	args := history.HistoryService_StartWorkflowExecution_Helper.Args(_StartRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_StartWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_StartWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) SyncActivity(
	ctx context.Context,
	_SyncActivityRequest *history.SyncActivityRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_SyncActivity_Helper.Args(_SyncActivityRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_SyncActivity_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_SyncActivity_Helper.UnwrapResponse(&result)
	return
}

func (c client) SyncShardStatus(
	ctx context.Context,
	_SyncShardStatusRequest *history.SyncShardStatusRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_SyncShardStatus_Helper.Args(_SyncShardStatusRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_SyncShardStatus_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_SyncShardStatus_Helper.UnwrapResponse(&result)
	return
}

func (c client) TerminateWorkflowExecution(
	ctx context.Context,
	_TerminateRequest *history.TerminateWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_TerminateWorkflowExecution_Helper.Args(_TerminateRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_TerminateWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_TerminateWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}
