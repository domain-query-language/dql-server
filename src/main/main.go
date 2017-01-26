package main

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/query"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/application"
	"fmt"
)

func main() {

	repository_projection := projection.CreateRepository()
	repository_projection.Save(application.ListDatabasesProjection)

	query_handler := query.CreateHandler(repository_projection)
	query_handler.(query.SimpleHandler).Add(
		application.ListDatabasesQueryID,
		application.ListDatabasesQueryHandler,
	)

	result, _ := query_handler.Handle(application.ListDatabasesQuery)

	result_string, _ := result.MarshalJSON()

	fmt.Printf(string(result_string))
}
