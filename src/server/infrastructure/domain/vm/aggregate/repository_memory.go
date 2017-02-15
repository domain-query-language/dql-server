package aggregate

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/aggregate"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"errors"
)

type MemoryRepository struct {

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

	return nil
}

func CreateMemoryRepository() *MemoryRepository {

	return &MemoryRepository {
		aggregates: map[vm.Identifier]aggregate.Aggregate{},
		aggregate_instances: map[*vm.AggregateIdentifier]aggregate.Aggregate{},
	}
}
