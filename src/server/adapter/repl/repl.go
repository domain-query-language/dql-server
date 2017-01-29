package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/domain-query-language/dql-server/src/server/adapter/tokenizer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {

	io.WriteString(out, LOGO)

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

		t := tokenizer.NewTokenizer(line)

		tokens := t.Tokens()

		for _, tok := range tokens {
			io.WriteString(out, tok.String()+"\n")
		}
		io.WriteString(out, "\n")
	}
}

const LOGO = `    ____  ____    __
   / __ \/ __ \  / /
  / / / / / / / / /
 / /_/ / /_/ / / /___
/_____/\___\_\/_____/

`

