package vm

import (
	"time"
	"github.com/satori/go.uuid"
)

type Payload interface {

}

type Event interface {

	Id() Identifier
	CommandId() Identifier
	TypeId() Identifier
	ContextId() Identifier

	OccurredAt() time.Time

	Payload() struct{}
}

type Event_  struct {

	id Identifier
	type_id Identifier
	aggregate_id Identifier
	command_id Identifier
	context_id Identifier

	occurred_at time.Time

	payload Payload
}


func NewEvent(type_id Identifier, command_id Identifier, aggregate_id Identifier, context_id Identifier, payload Payload) *Event_ {

	return &Event_ {
		id: uuid.NewV4(),
		command_id: command_id,
		type_id: type_id,
		context_id: context_id,
		occurred_at: time.Now(),
		payload: payload,
	}
}
