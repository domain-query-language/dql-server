package aggregate

import (
	"errors"
	"github.com/domain-query-language/dql-server/src/server/vm/projector"
	"github.com/domain-query-language/dql-server/src/server/vm"
)

var (
	AGGREGATE_HANDLER_NOT_EXISTS = errors.New("Aggregate handler does not exist.")
)

type AggregateHandler func(projector projector.Projector) []vm.Event

type Aggregate interface {

	Handle(command vm.Command) ([]vm.Event, error)

	Reset()

	Flush()

	State() projector.Projection

	Changes() []vm.Event
}

/**

	Implementation of SimpleAggregate

 */

type SimpleAggregate struct {

	id Identifier

	handlers map[Identifier]AggregateHandler

	projector projector.Projector

	changes []vm.Event
}

func (self *SimpleAggregate) Handle(command vm.Command) ([]vm.Event, error) {

	handler, ok := self.handlers[command.TypeId()]

	if(!ok) {
		return []vm.Event{}, AGGREGATE_HANDLER_NOT_EXISTS
	}

	events := handler(self.projector)

	self.changes = append(self.changes, events)

	return events, nil
}

func (self *SimpleAggregate) Reset() {
	self.projector.Reset()
	self.changes = []vm.Event{}
}

func (self *SimpleAggregate) Flush() {
	self.changes = []vm.Event{}
}

func (self *SimpleAggregate) State() projector.Projection {
	return self.projector.Projection()
}

func (self *SimpleAggregate) Changes() []vm.Event {
	return self.changes
}
