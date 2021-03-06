package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/davecgh/go-spew/spew"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser"
)

func StartStatementRepl(in io.Reader, out io.Writer) {

	io.WriteString(out, STATEMENT_LOGO)

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

		p := parser.NewStatement(line)

		blkStmt, err := p.ParseBlockStatement()

		if (err != nil) {
			io.WriteString(out, "\n")
			io.WriteString(out, "Error: "+err.Error());
			io.WriteString(out, "\n")
		} else {
			spew.Config.DisableMethods = true;
			spew.Config.DisablePointerAddresses = true;
			io.WriteString(out, "\n")
			io.WriteString(out, "Formatted: "+ blkStmt.String()+"\n")
			io.WriteString(out, "\n")
			io.WriteString(out, spew.Sdump(blkStmt))

		}

		io.WriteString(out, "\n")
	}
}

const STATEMENT_LOGO = `    ____  ____    __
   / __ \/ __ \  / /
  / / / / / / / / /
 / /_/ / /_/ / / /___
/_____/\___\_\/_____/  StatementParser

`
