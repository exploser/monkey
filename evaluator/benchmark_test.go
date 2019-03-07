package evaluator

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
)

func BenchmarkFibonacci(b *testing.B) {
	input := "fib := fn(x) { if (x < 2) { return x; } return fib(x-1)+fib(x-2); }; fib(10);"
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	checkParserErrors(b, p)

	for n := 0; n < b.N; n++ {
		env := GetBaseEnvironment()
		Eval(program, env)
	}
}
