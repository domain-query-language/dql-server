package main

import (
	"net/http"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure/adapter"
	"strings"
	"encoding/json"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database/command"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure"
	"github.com/davecgh/go-spew/spew"
	"github.com/domain-query-language/dql-server/examples/dql/application/projection/list-databases"
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

		result, handle_err := infrastructure.QueryHandler.Handle(
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

		events, handle_err := infrastructure.CommandHandler.Handle(
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

	infrastructure.Boot()

	aggregate_id := uuid.NewV4()

	infrastructure.CommandHandler.Handle(
		vm.NewCommand(
			vm.NewAggregateIdentifier(
				aggregate_id,
				database.Identifier,
			),
			command.Create {
				Name: "dql",
			},
		),
	)

	infrastructure.CommandHandler.Handle(
		vm.NewCommand(
			vm.NewAggregateIdentifier(
				aggregate_id,
				database.Identifier,
			),
			command.Rename {
				Name: "dql-lol",
			},
		),
	)

	infrastructure.CommandHandler.Handle(
		vm.NewCommand(
			vm.NewAggregateIdentifier(
				aggregate_id,
				database.Identifier,
			),
			command.Rename {
				Name: "dql-rofl",
			},
		),
	)

	result, _ := infrastructure.QueryHandler.Handle(
		vm.NewQuery(
			list_databases.Identifier,
			list_databases.Query {

			},
		),
	)

	spew.Dump(result)

	/*
	http.HandleFunc("/schema", schema)

	err := http.ListenAndServe(":4242", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	*/
}
