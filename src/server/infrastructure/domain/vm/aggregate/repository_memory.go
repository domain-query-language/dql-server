package aggregate

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/aggregate"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"errors"
	"github.com/domain-query-language/dql-server/src/server/domain/store"
)

type MemoryRepository struct {

	event_log store.Log
	aggregates map[vm.Identifier]aggregate.Aggregate
	aggregate_instances map[*vm.AggregateIdentifier]aggregate.Aggregate
}

func (self *MemoryRepository) Add(id vm.Identifier, aggregate aggregate.Aggregate) {
	self.aggregates[id] = aggregate
}

func (self *MemoryRepository) Get(id *vm.AggregateIdentifier) (aggregate.Aggregate, error) {

	aggregate, ok := self.aggregate_instances[id]

	if !ok {
		aggregate, aggregate_found := self.aggregates[id.TypeId]

		if !aggregate_found {
			return nil, errors.New("The aggregate type does not exist.")
		}

		return aggregate.Copy(id.Id), nil
	}

	return aggregate, nil
}

func (self *MemoryRepository) Save(aggregate aggregate.Aggregate) error {

	self.aggregate_instances[aggregate.Id()] = aggregate

	self.event_log.Append(aggregate.Events())

	return nil
}

func CreateMemoryRepository(event_log store.Log) *MemoryRepository {

	return &MemoryRepository {
		event_log: event_log,
		aggregates: map[vm.Identifier]aggregate.Aggregate{},
		aggregate_instances: map[*vm.AggregateIdentifier]aggregate.Aggregate{},
	}
}
