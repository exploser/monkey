package evaluator_test

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/evaltests"
	"git.exsdev.ru/ExS/monkey/evaluator"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
	"git.exsdev.ru/ExS/monkey/types"
)

func e(t testing.TB, input string) types.Object {
	t.Helper()

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return evaluator.Eval(program, evaluator.GetBaseEnvironment())
}

func TestEverything(t *testing.T) {
	evaltests.RunAll(t, e)
}

func BenchmarkEverything(b *testing.B) {
	evaltests.BenchmarkAll(b, e)
}
