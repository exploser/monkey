package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/vasilevp/monkey/evaluator"

	"github.com/vasilevp/monkey/lexer"
	"github.com/vasilevp/monkey/parser"
)

const prompt = "=> "

func Start(in io.Reader, out io.Writer) {
	env := evaluator.GetBaseEnvironment()
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

		prog := p.ParseProgram()
		for _, e := range p.Errors() {
			fmt.Println(e)
		}

		evaluated := evaluator.Eval(prog, env)
		if evaluated == nil {
			fmt.Println("eval: (nil)")
			continue
		}

		fmt.Println(evaluated)
		fmt.Println("input: ", prog.String())
	}
}
