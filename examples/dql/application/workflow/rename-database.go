package workflow

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/workflow"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/aggregate/database/event"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/command"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/query"
	"github.com/domain-query-language/dql-server/examples/dql/application/projection/list-databases"
)

var RenameDatabaseHandlers = &map[vm.Identifier]workflow.WorkflowHandler {

	event.TypeCreated: func(commandHandler command.Handler, queryHandler query.Handler, event vm.Event) {

		database_names_result, err := queryHandler.Handle(
			vm.NewQuery(
				list_databases.Identifier,
				list_databases.Query{

				},
			),
		)

		name_list := database_names_result.(query.Result_).Data


	},
}

var RenameDatabaseWorkflow = workflow.NewWorkflow(

)
