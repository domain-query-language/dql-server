package main

import (
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure"
	"github.com/domain-query-language/dql-server/examples/dql/application/http"
)

func main() {

	infrastructure.Boot()

	http.StartHttpServer()
}
