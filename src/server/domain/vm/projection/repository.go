package projection

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Repository interface {

	Get(id vm.Identifier) (projection Projection, err error)

	Save(projection Projection) error
}

type Repository_ struct {

	projections map[vm.Identifier]Projection
}

func (self *Repository_) Get(id vm.Identifier) (projection Projection, err error) {
	return self.projections[id], nil
}

func (self *Repository_) Save(projection Projection) error {
	self.projections[projection.Id()] = projection

	return nil
}

func CreateRepository() Repository {
	return &Repository_{
		projections: map[vm.Identifier]Projection{},
	}
}
