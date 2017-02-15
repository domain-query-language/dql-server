package main

import (
	"log"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure"
	"net/http"
	controllers "github.com/domain-query-language/dql-server/examples/dql/application/http"
)

func main() {

	infrastructure.Boot()

	http.HandleFunc("/schema", controllers.Schema)

	err := http.ListenAndServe(":4242", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
