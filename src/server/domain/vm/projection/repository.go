package projection

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Repository interface {

	Get(id vm.Identifier) (projection Projection, err error)

	Save(projection Projection) error
}
