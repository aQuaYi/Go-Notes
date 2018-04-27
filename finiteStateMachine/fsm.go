package main

import (
	"fmt"
	"sync"
)

type fsmState string
type fsmEvent string
type fsmHandler func() fsmState

type fsm struct {
	mu       sync.Mutex
	state    fsmState
	handlers map[fsmState]map[fsmEvent]fsmHandler
}

func (f *fsm) getState() fsmState {
	return f.state
}
func (f *fsm) setState(newState fsmState) {
	f.state = newState
}

func (f *fsm) addHandler(state fsmState, event fsmEvent, handler fsmHandler) {
	if _, ok := f.handlers[state]; !ok {
		f.handlers[state] = make(map[fsmEvent]fsmHandler)
	}
	if _, ok := f.handlers[state][event]; ok {
		fmt.Printf("[警告] 状态(%s)的事件(%s)的处理方法的旧方法被覆盖\n", state, event)
	}
	f.handlers[state][event] = handler
	// 为什么要返回 f
	// return f
}

func (f *fsm) call(event fsmEvent) fsmState {
	f.mu.Lock()
	defer f.mu.Unlock()

	events, ok := f.handlers[f.getState()]
	if !ok || events == nil {
		return f.getState()
	}

	if fn, ok := events[event]; ok {
		oldState := f.getState()
		f.setState(fn())
		newState := f.getState()
		fmt.Printf("\t状态从 [%s] 变成 [%s]\n", oldState, newState)
	}

	return f.getState()
}

func newFSM(initState fsmState) *fsm {
	f := &fsm{
		state:    initState,
		handlers: make(map[fsmState]map[fsmEvent]fsmHandler),
	}
	fmt.Printf("\t状态机生成，初始状态 [%s]\n", f.getState())
	return f
}
