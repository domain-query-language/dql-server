package http

import (
	adapter "github.com/domain-query-language/dql-server/src/server/adapter/parser"
	infraParser "github.com/domain-query-language/dql-server/src/server/infrastructure/adapter/parser"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure"
	"net/http"
	"io/ioutil"
)

func Schema(w http.ResponseWriter, r *http.Request) {

	statements, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if (len(statements) == 0) {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	adptr := adapter.New(infraParser.NewUuidGenerator(), string(statements))

	handleable, err := adptr.Next()

	if (err != nil) {
		panic(err.Error())
	}

	if(handleable.Typ == "query") {

		result, err := infrastructure.QueryHandler.Handle(
			handleable.Query,
		)

		if err != nil {

			w.Header().Add("error", err.Error())
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		respondWithJson(w, result)

		return
	}

	if(handleable.Typ == "command") {

		events, err := infrastructure.CommandHandler.Handle(
			handleable.Command,
		)

		if err != nil {
			w.Header().Add("error", err.Error())
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		respondWithJson(w, events)

		return
	}

}
