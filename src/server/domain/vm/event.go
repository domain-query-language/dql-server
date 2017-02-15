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
	CommandId() Identifier
	TypeId() Identifier
	AggregateTypeId() Identifier

	OccurredAt() time.Time

	Payload() Payload
}

type Event_  struct {

	id Identifier
	type_id Identifier
	aggregate_id Identifier
	command_id Identifier
	aggregate_type_id Identifier

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

func (self *Event_) AggregateTypeId() Identifier {
	return self.aggregate_type_id
}

func (self *Event_) OccurredAt() time.Time {
	return self.occurred_at
}

func (self *Event_) Payload() Payload {
	return self.payload
}

func NewEvent(id *AggregateIdentifier, payload Payload) *Event_ {

	return &Event_ {
		id: uuid.NewV4(),
		command_id: id.Id,
		type_id: payload.TypeId(),
		aggregate_type_id: id.TypeId,
		occurred_at: time.Now(),
		payload: payload,
	}
}

func FromRawEvent(
	id Identifier,
	type_id Identifier,
	command_id Identifier,
	aggregate_id Identifier,
	aggregate_type_id Identifier,
	occurred_at time.Time,
	payload Payload,
) *Event_ {

	return &Event_ {
		id: id,
		type_id: type_id,
		command_id: command_id,
		aggregate_id: aggregate_id,
		aggregate_type_id: aggregate_type_id,
		occurred_at: occurred_at,
		payload: payload,
	}
}
