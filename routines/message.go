package routines

// Message - implementation of routineMsg
type Message struct {
	id   string
	name string
	msg  interface{}
}

// NewMessage - create a new reoutines.Message instance implements container.RoutineMsg
func NewMessage(id string, name string, msg interface{}) Message {
	m := Message{}
	m.id = id
	m.name = name
	m.msg = msg
	return m
}

// GetId - container.RoutineMsg
func (msg *Message) GetId() *string {
	return &msg.id
}

// GetName - container.RoutineMsg
func (msg *Message) GetName() *string {
	return &msg.name
}

// GetMsg - container.RoutineMsg
func (msg *Message) GetMsg() interface{} {
	return msg.msg
}
