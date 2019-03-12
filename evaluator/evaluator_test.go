package evaluator_test

import (
	"testing"

	"github.com/vasilevp/monkey/evaltests"
	"github.com/vasilevp/monkey/evaluator"
	"github.com/vasilevp/monkey/lexer"
	"github.com/vasilevp/monkey/parser"
	"github.com/vasilevp/monkey/types"
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
