package workflow

import (
	"errors"
	"github.com/domain-query-language/dql-server/src/server/vm"
	"github.com/domain-query-language/dql-server/src/server/vm/handler"
)

var (
	WORKFLOW_HANDLER_NOT_EXISTS = errors.New("Projector handler does not exist.")
)

type WorkflowHandler func(handler handler.Handler, event vm.Event) []vm.Event

type Workflow interface {

	Reset()

	Apply(event vm.Event) error

	Version() int
}

type SimpleWorkflow struct {

	id vm.Identifier
	last_id vm.Identifier

}
