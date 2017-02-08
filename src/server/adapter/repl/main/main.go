package main

import (
	"github.com/domain-query-language/dql-server/src/server/adapter/repl"
	"os"
)

func main() {

	if (len(os.Args) > 1 && os.Args[1] == "statement") {

		repl.StartStatementRepl(os.Stdin, os.Stdout)
	} else {

		repl.Start(os.Stdin, os.Stdout)
	}
}
