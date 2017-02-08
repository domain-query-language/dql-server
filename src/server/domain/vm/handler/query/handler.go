package query

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"errors"
)

type QueryHandler func(query vm.Query, queryable Queryable) Result

var (
	QUERY_HANDLER_NOT_EXISTS = errors.New("The query handler does not exist.")
)

type Handler interface {

	Handle(query vm.Query) (result Result, err error)
}

type Handler_ struct {

	handlers map[vm.Identifier]QueryHandler
	repository_projection projection.Repository
}

func (self *Handler_) Add(id vm.Identifier, query_handler QueryHandler) {
	self.handlers[id] = query_handler
}

func (self *Handler_) Handle(query vm.Query) (result Result, err error) {

	handler, ok := self.handlers[query.Id()]

	if !ok {
		return nil, QUERY_HANDLER_NOT_EXISTS
	}

	projection, err := self.repository_projection.Get(query.Id())

	return handler(query, projection), nil
}

func CreateHandler(repository projection.Repository) *Handler_ {

	return &Handler_ {
		handlers: map[vm.Identifier]QueryHandler{},
		repository_projection: repository,
	}
}
