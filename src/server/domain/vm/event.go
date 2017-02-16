package vm

import (
	"time"
	"github.com/satori/go.uuid"
	"github.com/pquerna/ffjson/ffjson"
)

type Payload interface {

	TypeId() Identifier
}

type Event interface {

	Id() Identifier
	TypeId() Identifier

	CommandId() Identifier

	AggregateId() AggregateIdentifier

	OccurredAt() time.Time

	Payload() Payload
}

type Event_  struct {

	id Identifier `json:"id"`
	typeId Identifier `json:"type_id"`

	commandId Identifier `json:"command_id"`

	aggregateId AggregateIdentifier `json:"aggregate"`

	occurredAt time.Time `json:"occurred_at"`

	payload Payload `json:"payload"`
}

func (self *Event_) Id() Identifier {
	return self.id
}

func (self *Event_) TypeId() Identifier {
	return self.typeId
}

func (self *Event_) CommandId() Identifier {
	return self.commandId
}

func (self *Event_) AggregateId() AggregateIdentifier {
	return self.aggregateId
}

func (self *Event_) OccurredAt() time.Time {
	return self.occurredAt
}

func (self *Event_) Payload() Payload {
	return self.payload
}

func (self *Event_) MarshalJSON() ([]byte, error) {

	return ffjson.Marshal(
			struct {
				Id Identifier `json:"id"`
				TypeId Identifier `json:"type_id"`

				CommandId Identifier `json:"command_id"`

				AggregateId AggregateIdentifier `json:"aggregate"`

				OccurredAt time.Time `json:"occurred_at"`

				Payload Payload `json:"payload"`
			}{
				self.id,
				self.typeId,
				self.commandId,
				self.aggregateId,
				self.occurredAt,
				self.payload,
			},
	)
}

func NewEvent(aggregateId AggregateIdentifier, commandId Identifier, payload Payload) *Event_ {

	return &Event_ {
		id: uuid.NewV4(),
		typeId: payload.TypeId(),

		commandId: commandId,

		aggregateId: aggregateId,
		occurredAt: time.Now(),
		payload: payload,
	}
}
