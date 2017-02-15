package aggregate

import (
	"errors"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"time"
)

var (
	AGGREGATE_HANDLER_NOT_EXISTS = errors.New("Aggregate command does not exist.")
)

type AggregateHandler func(aggregate Aggregate, command vm.Command) error

type Aggregate interface {

	Id() *vm.AggregateIdentifier

	ContextId() vm.Identifier

	Reset()

	Apply(event vm.Event) error

	Flush()

	Commands() []vm.Command

	Events() []vm.Event

	Projector() projection.Projector

	Snapshot() Snapshot

	Handle(command vm.Command) ([]vm.Event, error)

	Copy(id vm.Identifier) Aggregate
}

/**

	Implementation of Aggregate

 */

type Aggregate_ struct {

	id *vm.AggregateIdentifier
	context_id vm.Identifier

	handlers *map[vm.Identifier]AggregateHandler

	projector projection.Projector

	commands []vm.Command
	events []vm.Event
}

func (self *Aggregate_) Id() *vm.AggregateIdentifier {
	return self.id
}

func (self *Aggregate_) ContextId() vm.Identifier {
	return self.context_id
}

func (self *Aggregate_) Reset() {

	self.projector.Reset()

	self.commands = []vm.Command{}
	self.events = []vm.Event{}
}

func (self *Aggregate_) Apply(event vm.Event) error {

	error := self.projector.Apply(event)

	if error != nil {
		return error
	}

	self.events = append(self.events, event)

	return nil
}

func (self *Aggregate_) Flush() {

	self.commands = []vm.Command{}
	self.events = []vm.Event{}
}

func (self *Aggregate_) Commands() []vm.Command {
	return self.commands
}

func (self *Aggregate_) Events() []vm.Event {
	return self.events
}

func (self *Aggregate_) Projector() projection.Projector {
	return self.projector
}

func (self *Aggregate_) Snapshot() Snapshot {

	return Snapshot {
		Id: self.id,
		OccurredAt: time.Now(),
		Version: self.projector.Version(),
		//Payload: self.projector.Projection().GobDecode(),
	}
}

func (self *Aggregate_) Handle(command vm.Command) ([]vm.Event, error) {

	handler, handler_exists := (*self.handlers)[command.TypeId()]

	if(!handler_exists) {
		return nil, AGGREGATE_HANDLER_NOT_EXISTS
	}

	handling_err := handler(self, command)

	if handling_err != nil {
		return nil, handling_err
	}

	self.commands = append(self.commands, command)

	return self.events, nil
}

func (self *Aggregate_) Copy(id vm.Identifier) Aggregate {

	aggregate := *self
	aggregate.id.Id = id
	aggregate.projector = aggregate.projector.Copy()

	return &aggregate
}

func NewAggregate(id vm.Identifier, context_id vm.Identifier, projector projection.Projector, handlers *map[vm.Identifier]AggregateHandler) *Aggregate_ {

	return &Aggregate_ {
		id: vm.NewAggregateIdentifier(
			nil,
			id,
		),
		context_id: context_id,
		projector: projector,
		handlers: handlers,
	}
}
