package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {

	io.WriteString(out, QUERY_LOGO)

	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if (line == "exit") {
			io.WriteString(out, "See ya!\n")
			break;
		}

		p := parser.New(nil, line)

		handleable, err := p.Next()

		if (err != nil) {
			io.WriteString(out, "Error: "+err.Error());
		} else {
			io.WriteString(out, handleable.String())

		}

		io.WriteString(out, "\n")
	}
}

const QUERY_LOGO = `    ____  ____    __
   / __ \/ __ \  / /
  / / / / / / / / /
 / /_/ / /_/ / / /___
/_____/\___\_\/_____/  QueryRepl

`

