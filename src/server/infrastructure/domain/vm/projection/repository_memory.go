package projection

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
)

type MemoryRepository struct {

	projections map[vm.Identifier]projection.Projection
}

func (self *MemoryRepository) Add(projection projection.Projection) {
	self.projections[projection.Id()] = projection
}

func (self *MemoryRepository) Get(id vm.Identifier) (projection projection.Projection, err error) {
	return self.projections[id], nil
}

func (self *MemoryRepository) Save(projection projection.Projection) error {
	self.projections[projection.Id()] = projection

	return nil
}

func CreateMemoryRepository() *MemoryRepository {
	return &MemoryRepository {
		projections: map[vm.Identifier]projection.Projection{},
	}
}
