package repl

import (
	"bufio"
	"fmt"
	"io"

	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
)

const prompt = "=> "

func Start(in io.Reader, out io.Writer) {
	for {
		fmt.Print(prompt)
		scanner := bufio.NewScanner(in)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		fmt.Println(p.ParseProgram())
	}
}
