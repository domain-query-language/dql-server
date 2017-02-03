package repl

import (
	"testing"
	"github.com/domain-query-language/dql-server/src/server/adapter/repl"
	"os"
)

func TestStartingTheRepl(t *testing.T) {

	repl.Start(os.Stdin, os.Stdout);
}

func TestStartingStatementRep(t *testing.T) {

	repl.StartStatementRepl(os.Stdin, os.Stdout);
}
