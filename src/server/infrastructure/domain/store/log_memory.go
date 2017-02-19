package store

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/store"
)

type MemoryLog struct {

	aggregates map[vm.AggregateIdentifier][]vm.Loggable

	events_index map[vm.Identifier]int
	events []vm.Loggable
}

func (self *MemoryLog) Reset() {

	self.aggregates = map[vm.AggregateIdentifier][]vm.Loggable{}

	self.events_index = map[vm.Identifier]int{}
	self.events = []vm.Loggable{}
}

func (self *MemoryLog) Append(loggable vm.Loggable) {

	aggregate_id := *loggable.AggregateId()

	self.events = append(self.events, loggable)
	self.events_index[loggable.Id()] = len(self.events) - 1
	self.aggregates[aggregate_id] = append(self.aggregates[aggregate_id], loggable)
}

func (self *MemoryLog) Stream() store.Stream {
	return NewMemoryStream(self)
}

func (self *MemoryLog) AggregateStream(id *vm.AggregateIdentifier) store.AggregateStream {
	return NewMemoryAggregateStream(id, self)
}

func NewMemoryLog() *MemoryLog {

	return &MemoryLog {
		aggregates: map[vm.AggregateIdentifier][]vm.Loggable{},
		events_index: map[vm.Identifier]int{},
		events: []vm.Loggable{},
	}
}
