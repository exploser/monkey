package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/vasilevp/monkey/bytecode"
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
			fmt.Println("parser:", e)
			continue
		}

		fmt.Println("input:", prog.String())

		evaluated := evaluator.Eval(prog, env)
		if evaluated == nil {
			fmt.Println("eval: (nil)")
			continue
		}

		fmt.Println("eval:", evaluated)

		compiler := bytecode.New()
		err := compiler.Compile(prog)
		if err != nil {
			fmt.Println("compiler:", err)
			continue
		}

		fmt.Println("compiler:", compiler.Bytecode)

	}
}
