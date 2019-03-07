package evaluator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
	"git.exsdev.ru/ExS/monkey/types"
)

type Testing interface {
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	FailNow()
}

func checkParserErrors(t Testing, p *parser.Parser) {
	if !assert.Empty(t, p.Errors()) {
		for _, e := range p.Errors() {
			t.Error(e)
		}

		t.FailNow()
	}
}

func testEval(t *testing.T, input string) types.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	env := GetBaseEnvironment()

	return Eval(program, env)
}
