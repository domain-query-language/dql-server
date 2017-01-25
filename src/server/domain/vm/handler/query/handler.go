package query

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"errors"
)

type QueryHandler func(query Query, projection projection.Projection) Result

var (
	QUERY_HANDLER_NOT_EXISTS = errors.New("Query handler does not exist.")
)

type Handler interface {

	Handle(query Query) (result Result, err error)
}

type SimpleHandler struct {

	handlers map[vm.Identifier]QueryHandler
	repository_projection projection.Repository
}

func (self *SimpleHandler) Add(id vm.Identifier, query_handler QueryHandler) {
	self.handlers[id] = query_handler
}

func (self *SimpleHandler) Handle(query Query) (result Result, err error) {

	handler, ok := self.handlers[query.Id()]

	if !ok {
		return nil, QUERY_HANDLER_NOT_EXISTS
	}

	projection, err := self.repository_projection.Get(query.Id())

	return handler(query, projection), nil
}

func CreateHandler(repository projection.Repository) Handler {

	return &SimpleHandler {
		handlers: map[vm.Identifier]QueryHandler{},
		repository_projection: repository,
	}
}
