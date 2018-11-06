package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/token"	
	"io"
	"os"
)
//Start reads lines in and writes tokens out
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("\ndigijan $ " + os.Getenv("USER") + " ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		tokenizer := lexer.New(line)

		for tokenized := tokenizer.NextToken(); tokenized.Type != token.EOF; tokenized = tokenizer.NextToken() {
			fmt.Printf("%+v\n", tokenized)
		}
	}
}
