package vm

import (
	"github.com/satori/go.uuid"
	"time"
)

type Command interface {

	Id() Identifier

	TypeId() Identifier

	AggregateId() Identifier

	ContextId() Identifier
}

type Command_  struct {

	id Identifier
	typeId Identifier
	aggregateId Identifier
	contextId Identifier

	occurredAt time.Time

	Payload Payload
}


func NewCommand(type_id Identifier, aggregate_id Identifier, context_id Identifier, payload Payload) *Command_ {

	return &Command_ {
		id: uuid.NewV4(),
		typeId: type_id,
		aggregateId: aggregate_id,
		contextId: context_id,
		occurredAt: time.Now(),
		Payload: payload,
	}
}
