package main

import (
	"log"
	"net/http"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure/adapter"
	"github.com/domain-query-language/dql-server/examples/dql/application/query/list-databases"
	"github.com/domain-query-language/dql-server/examples/dql/application"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure/application/projection"
	"strings"
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
}

func main() {

	application.ProjectionsRepository.Save(
		projection.ListDatabasesProjection,
	)

	application.QueryHandler.Add(
		projection.ListDatabasesProjectionID,
		list_databases.Handler,
	)

	http.HandleFunc("/schema", schema)

	err := http.ListenAndServe(":4242", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
