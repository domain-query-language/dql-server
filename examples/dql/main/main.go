package main

import (
	"net/http"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure/adapter"
	"github.com/domain-query-language/dql-server/examples/dql/application/query/list-databases"
	"github.com/domain-query-language/dql-server/examples/dql/application"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure/application/projection"
	"strings"
	"encoding/json"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database/command"
)

func schema(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	statements := strings.TrimSpace(
		r.FormValue("statements"),
	)

	adapter := adapter.NewMockAdapter(statements)

	handleable, err := adapter.Next()

	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if(handleable.Typ == "query") {

		result, handle_err := application.QueryHandler.Handle(
			handleable.Query,
		)

		if handle_err != nil {
			w.Header().Add("error", handle_err.Error())
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		response, _ := result.MarshalJSON()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

		return
	}

	if(handleable.Typ == "command") {

		events, handle_err := application.CommandHandler.Handle(
			handleable.Command,
		)

		if handle_err != nil {
			w.Header().Add("error", handle_err.Error())
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		json_events, _ := json.Marshal(events)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json_events)

		return
	}
}

func main() {

	application.ProjectionsRepository.Save(
		projection.ListDatabasesProjection,
	)

	application.QueryHandler.Add(
		projection.ListDatabasesProjectionID,
		list_databases.Handler,
	)

	application.AggregatesRepository.Add(
		database.Identifier,
		database.Aggregate,
	)

	application.CommandHandler.Handle(
		vm.NewCommand(
			vm.NewAggregateIdentifier(
				uuid.NewV4(),
				database.Identifier,
			),
			command.Create{
				Name: "dql",
			},
		),
	)

	/*
	http.HandleFunc("/schema", schema)

	err := http.ListenAndServe(":4242", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	*/
}
