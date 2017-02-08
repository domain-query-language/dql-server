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

type AggregateHandler func(projector projection.Projector, command vm.Command) ([]vm.Event, error)

type Aggregate interface {

	Id() Identifier

	Reset()

	Flush()

	Commands() []vm.Command

	Events() []vm.Event

	Snapshot() Snapshot

	Handle(command vm.Command) ([]vm.Event, error)
}

/**

	Implementation of Aggregate

 */

type Aggregate_ struct {

	id Identifier

	handlers map[vm.Identifier]AggregateHandler

	projector projection.Projector

	commands []vm.Command
	events []vm.Event
}

func (self *Aggregate_) Id() Identifier {
	return self.id
}

func (self *Aggregate_) Reset() {

	self.projector.Reset()

	self.commands = []vm.Command{}
	self.events = []vm.Event{}
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

func (self *Aggregate_) Snapshot() Snapshot {

	return Snapshot {
		Id: self.id,
		OccurredAt: time.Now(),
		Version: self.projector.Version(),
		//Payload: self.projector.Projection().GobDecode(),
	}
}

func (self *Aggregate_) Handle(command vm.Command) ([]vm.Event, error) {

	handler, handler_exists := self.handlers[command.TypeId()]

	if(!handler_exists) {
		return nil, AGGREGATE_HANDLER_NOT_EXISTS
	}

	events, handling_err := handler(self.projector, command)

	if handling_err != nil {
		return nil, handling_err
	}

	self.commands = append(self.commands, command)
	self.events = append(self.events, events...)

	return events, nil
}

func CreateAggregate(id vm.Identifier, projector projection.Projector, handlers map[vm.Identifier]AggregateHandler) *Aggregate_ {

	return &Aggregate_ {
		id: id,
		projector: projector,
		handlers: handlers,
	}
}
