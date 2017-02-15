package store

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/store"
)

type MemoryLog struct {

	aggregates map[vm.Identifier][]vm.Event

	events_index map[vm.Identifier]int
	events []vm.Event
}

func (self *MemoryLog) Reset() {

	self.aggregates = map[vm.Identifier][]vm.Event{}

	self.events_index = map[vm.Identifier]int{}
	self.events = []vm.Event{}
}

func (self *MemoryLog) Append(events []vm.Event) {

	for _, event := range events {

		self.events = append(self.events, event)
		self.events_index[event.Id()] = len(self.events) - 1
		self.aggregates[event.AggregateId().TypeId] = append(self.aggregates[event.AggregateId().TypeId], event)
	}
}

func (self *MemoryLog) Stream() store.Stream {
	return NewMemoryStream(self)
}

func NewMemoryLog() *MemoryLog {

	return &MemoryLog {
		aggregates: map[vm.Identifier][]vm.Event{},
		events_index: map[vm.Identifier]int{},
		events: []vm.Event{},
	}
}
