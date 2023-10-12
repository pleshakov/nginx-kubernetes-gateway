// Code generated by counterfeiter. DO NOT EDIT.
package eventsfakes

import (
	"context"
	"sync"

	"github.com/go-logr/logr"
	"github.com/nginxinc/nginx-gateway-fabric/internal/framework/events"
)

type FakeEventHandler struct {
	HandleEventBatchStub        func(context.Context, logr.Logger, events.EventBatch)
	handleEventBatchMutex       sync.RWMutex
	handleEventBatchArgsForCall []struct {
		arg1 context.Context
		arg2 logr.Logger
		arg3 events.EventBatch
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventHandler) HandleEventBatch(arg1 context.Context, arg2 logr.Logger, arg3 events.EventBatch) {
	fake.handleEventBatchMutex.Lock()
	fake.handleEventBatchArgsForCall = append(fake.handleEventBatchArgsForCall, struct {
		arg1 context.Context
		arg2 logr.Logger
		arg3 events.EventBatch
	}{arg1, arg2, arg3})
	stub := fake.HandleEventBatchStub
	fake.recordInvocation("HandleEventBatch", []interface{}{arg1, arg2, arg3})
	fake.handleEventBatchMutex.Unlock()
	if stub != nil {
		fake.HandleEventBatchStub(arg1, arg2, arg3)
	}
}

func (fake *FakeEventHandler) HandleEventBatchCallCount() int {
	fake.handleEventBatchMutex.RLock()
	defer fake.handleEventBatchMutex.RUnlock()
	return len(fake.handleEventBatchArgsForCall)
}

func (fake *FakeEventHandler) HandleEventBatchCalls(stub func(context.Context, logr.Logger, events.EventBatch)) {
	fake.handleEventBatchMutex.Lock()
	defer fake.handleEventBatchMutex.Unlock()
	fake.HandleEventBatchStub = stub
}

func (fake *FakeEventHandler) HandleEventBatchArgsForCall(i int) (context.Context, logr.Logger, events.EventBatch) {
	fake.handleEventBatchMutex.RLock()
	defer fake.handleEventBatchMutex.RUnlock()
	argsForCall := fake.handleEventBatchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeEventHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleEventBatchMutex.RLock()
	defer fake.handleEventBatchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventHandler) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ events.EventHandler = new(FakeEventHandler)
