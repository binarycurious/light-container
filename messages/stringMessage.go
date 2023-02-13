package messages

import (
	"fmt"
	"time"

	"github.com/binarycurious/light-container/container"
)

// StringMessage - @impl of container.RoutineMsg for simple string messages
type StringMessage struct {
	id   string
	name string
	msg  string
}

// NewStringMessage - create an http request message for the channel pipeline
func NewStringMessage(name string, msg string) container.RoutineMsg {
	return &StringMessage{
		id:   fmt.Sprintf("%v", time.Now().UnixNano()),
		name: name,
		msg:  msg,
	}
}

// GetId -
func (r *StringMessage) GetId() *string {
	return &r.id
}

// GetName -
func (r *StringMessage) GetName() *string {
	return &r.name
}

// GetMsg -
func (r *StringMessage) GetMsg() interface{} {
	return r.msg
}
