package vm

import (
	"github.com/satori/go.uuid"
	"time"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/aggregate"
)

type Command interface {

	Id() Identifier

	TypeId() Identifier

	AggregateId() Identifier

	AggregateTypeId() Identifier
}

type Command_  struct {

	id Identifier
	typeId Identifier
	aggregateId Identifier
	aggregateTypeId Identifier

	occurredAt time.Time

	Payload Payload
}


func (self *Command_) Id() Identifier {
	return self.id
}

func (self *Command_) TypeId() Identifier {
	return self.typeId
}

func (self *Command_) AggregateId() Identifier {
	return self.aggregateId
}

func (self *Command_) AggregateTypeId() Identifier {
	return self.aggregateTypeId
}

func NewCommand(aggregate_id aggregate.Identifier, payload Payload) *Command_ {

	return &Command_ {
		id: uuid.NewV4(),
		typeId: payload.TypeId(),
		aggregateId: aggregate_id.Id(),
		aggregateTypeId: aggregate_id.TypeId(),
		occurredAt: time.Now(),
		Payload: payload,
	}
}
