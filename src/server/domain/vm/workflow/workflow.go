package workflow

import (
	"errors"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/command"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/query"
)

var (
	WORKFLOW_HANDLER_NOT_EXISTS = errors.New("Projector command does not exist.")
)

type WorkflowHandler func(commandHandler command.Handler, queryHandler query.Handler, event vm.Event)

type Workflow_ struct {

	commandHandler command.Handler
	queryHandler query.Handler

	handlers *map[vm.Identifier]WorkflowHandler
}

func (self *Workflow_) Reset() {
	// Do Nothing
}

func (self *Workflow_) Apply(event vm.Event) error {

	handler, ok := (*self.handlers)[event.TypeId()]

	if !ok {
		return WORKFLOW_HANDLER_NOT_EXISTS
	}

	handler(self.commandHandler, self.queryHandler, event)

	return nil
}

func NewWorkflow(commandHandler command.Handler, queryHandler query.Handler, handlers *map[vm.Identifier]WorkflowHandler) *Workflow_ {
	return &Workflow_ {
		commandHandler: commandHandler,
		queryHandler: queryHandler,
		handlers: handlers,
	}
}
