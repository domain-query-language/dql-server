package aggregate

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/aggregate"
	"errors"
)

type MemoryRepository struct {

	aggregates map[[]byte]aggregate.Aggregate
}

func (self *MemoryRepository) Get(id aggregate.Identifier) (aggregate.Aggregate, error) {

	aggregate, ok := self.aggregates[id.Bytes()]

	if !ok {
		return nil, errors.New("Aggregate does not exist.")
	}

	return aggregate, nil
}

func (self *MemoryRepository) Save(aggregate aggregate.Aggregate) error {

	self.aggregates[aggregate.Id().Bytes()] = aggregate

	return nil
}

func CreateMemoryRepository() *MemoryRepository {

	return &MemoryRepository {
		map[[]byte]aggregate.Aggregate {

		},
	}
}
