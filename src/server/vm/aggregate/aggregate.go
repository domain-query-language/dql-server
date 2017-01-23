package aggregate

import (
	"errors"
	"github.com/domain-query-language/dql-server/src/server/vm/projector"
	"github.com/domain-query-language/dql-server/src/server/vm"
	"time"
	"github.com/domain-query-language/dql-server/src/server/vm/handler/command"
)

var (
	AGGREGATE_HANDLER_NOT_EXISTS = errors.New("Aggregate command does not exist.")
)

type AggregateHandler func(projector projector.Projector, command command.Command) ([]vm.Event, error)

type Aggregate interface {
 d
	Reset()

	Flush()

	Commands() []command.Command

	Events() []vm.Event

	Snapshot() Snapshot

	Handle(command command.Command) ([]vm.Event, error)
}

/**

	Implementation of Aggregate

 */

type Aggregate_ struct {

	id Identifier

	handlers map[vm.Identifier]AggregateHandler

	projector projector.Projector

	commands []command.Command
	events []vm.Event
}

func (self *Aggregate_) Reset() {

	self.projector.Reset()

	self.commands = []command.Command{}
	self.events = []vm.Event{}
}

func (self *Aggregate_) Flush() {

	self.commands = []command.Command{}
	self.events = []vm.Event{}
}

func (self *Aggregate_) Commands() []command.Command {
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
		Payload: self.projector.Projection().GobDecode(),
	}
}

func (self *Aggregate_) Handle(command command.Command) ([]vm.Event, error) {

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
