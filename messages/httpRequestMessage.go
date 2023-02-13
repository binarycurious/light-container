package messages

import (
	"fmt"
	"time"

	"github.com/binarycurious/light-container/container"
)

// HttpRequestMessage @imp container.RoutineMsg
type HttpRequestMessage struct {
	id     string
	url    string
	method string
	body   string
}

// NewHttpRequestMessage - create an http request message for the channel pipeline
func NewHttpRequestMessage(url string, method string, body string) container.RoutineMsg {
	return &HttpRequestMessage{
		id:     fmt.Sprintf("%s_%s_%v", url, method, time.Now().UnixMicro()),
		url:    url,
		method: method,
		body:   body,
	}
}

// GetId -
func (r *HttpRequestMessage) GetId() *string {
	return &r.id
}

// GetName -
func (r *HttpRequestMessage) GetName() *string {
	name := fmt.Sprintf("%s %s", r.method, r.url)
	return &name
}

// GetMsg -
func (r *HttpRequestMessage) GetMsg() interface{} {
	return r
}
