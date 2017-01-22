package aggregate

import (
	"errors"
	"github.com/domain-query-language/dql-server/src/server/vm/projector"
	"github.com/domain-query-language/dql-server/src/server/vm"
)

var (
	AGGREGATE_HANDLER_NOT_EXISTS = errors.New("Aggregate handler does not exist.")
)

type AggregateHandler func(projector projector.Projector, command vm.Command) ([]vm.Event, error)

type Aggregate interface {

	Reset()

	Flush()

	Commands() []vm.Command

	Events() []vm.Event

	Snapshot() projector.Snapshot

	Handle(command vm.Command) ([]vm.Event, error)
}

/**

	Implementation of Aggregate

 */

type Aggregate_ struct {

	handlers map[Identifier]AggregateHandler

	projector projector.Projector

	commands []vm.Command
	events []vm.Event
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

func (self *Aggregate_) Snapshot() projector.Snapshot {

	return self.projector.Snapshot()
}

func (self *Aggregate_) Handle(command vm.Command) ([]vm.Event, error) {

	handler, handler_exists := self.handlers[command.TypeId()]

	if(!handler_exists) {
		return nil, AGGREGATE_HANDLER_NOT_EXISTS
	}

	events, handling_err := handler(self.projector, command)

	if(handling_err) {
		return nil, handling_err
	}

	self.commands = append(self.commands, command)
	self.events = append(self.events, events)

	return events, nil
}
