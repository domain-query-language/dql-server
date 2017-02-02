package repl

import (
	"testing"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/repl"
	"os"
)

func TestStartingTheRepl(t *testing.T){

	repl.Start(os.Stdin, os.Stdout);

}
