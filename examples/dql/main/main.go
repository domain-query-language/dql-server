package main

import (
	"net/http"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser"
	infraParser "github.com/domain-query-language/dql-server/src/server/infrastructure/adapter/parser"
	"strings"
	"encoding/json"
	"log"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure"
)

func schema(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	statements := strings.TrimSpace(
		r.FormValue("statements"),
	)

	adapter := parser.New(infraParser.NewUuidGenerator(), statements)

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

	http.HandleFunc("/schema", schema)

	err := http.ListenAndServe(":4242", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
