package vm

import (
	"time"
	"github.com/satori/go.uuid"
)

type Payload interface {
	TypeId() Identifier
}

type Event interface {

	Id() Identifier
	TypeId() Identifier

	CommandId() Identifier

	AggregateId() *AggregateIdentifier

	OccurredAt() time.Time

	Payload() Payload
}

type Event_  struct {

	id Identifier
	type_id Identifier

	command_id Identifier

	aggregate_id *AggregateIdentifier

	occurred_at time.Time

	payload Payload
}

func (self *Event_) Id() Identifier {
	return self.id
}

func (self *Event_) CommandId() Identifier {
	return self.command_id
}

func (self *Event_) TypeId() Identifier {
	return self.type_id
}

func (self *Event_) AggregateId() *AggregateIdentifier {
	return self.aggregate_id
}

func (self *Event_) OccurredAt() time.Time {
	return self.occurred_at
}

func (self *Event_) Payload() Payload {
	return self.payload
}

func NewEvent(aggregate_id *AggregateIdentifier, command_id Identifier, payload Payload) *Event_ {

	return &Event_ {
		id: uuid.NewV4(),
		command_id: command_id,
		type_id: payload.TypeId(),
		aggregate_id: aggregate_id,
		occurred_at: time.Now(),
		payload: payload,
	}
}

