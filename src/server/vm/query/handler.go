package query

import (
	"github.com/domain-query-language/dql-server/src/server/vm/projector"
)

type QueryHandler func(query Query, projection projector.Projection) Result

type Handler interface {

	Handle(query Query) (result Result, error)
}
