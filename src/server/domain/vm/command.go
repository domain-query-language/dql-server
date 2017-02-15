package vm

import (
	"github.com/satori/go.uuid"
	"time"
)

type Command interface {

	Id() Identifier

	TypeId() Identifier

	AggregateId() *AggregateIdentifier

	Payload() Payload
}

type Command_  struct {

	id Identifier
	typeId Identifier
	aggregateId *AggregateIdentifier

	occurredAt time.Time

	payload Payload
}


func (self *Command_) Id() Identifier {
	return self.id
}

func (self *Command_) TypeId() Identifier {
	return self.typeId
}

func (self *Command_) AggregateId() *AggregateIdentifier {
	return self.aggregateId
}

func (self *Command_) Payload() Payload {
	return self.payload
}

func NewCommand(aggregate_id *AggregateIdentifier, payload Payload) *Command_ {

	return &Command_ {
		id: uuid.NewV4(),
		typeId: payload.TypeId(),
		aggregateId: aggregate_id,
		occurredAt: time.Now(),
		payload: payload,
	}
}
