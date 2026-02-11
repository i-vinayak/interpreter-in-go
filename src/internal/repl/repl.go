package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/i-vinayak/interpreter-in-go/src/internal/lexer"
	"github.com/i-vinayak/interpreter-in-go/src/internal/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	sc := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := sc.Scan()
		if !scanned {
			return
		}

		line := sc.Text()
		l := lexer.New(line)

		var tok *token.Token
		for tok = l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
