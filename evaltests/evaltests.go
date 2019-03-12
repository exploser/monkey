package evaltests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"git.exsdev.ru/ExS/monkey/types"
)

type Evaluator func(t testing.TB, input string) types.Object

var e Evaluator

type iparser interface {
	Errors() []error
}

func CheckParserErrors(t testing.TB, p iparser) {
	t.Helper()

	if !assert.Empty(t, p.Errors()) {
		for _, e := range p.Errors() {
			t.Error(e)
		}

		t.FailNow()
	}
}

var tests = map[string]func(t *testing.T, e Evaluator){
	"array":                 testArray,
	"bang":                  testBang,
	"boolean":               testBoolean,
	"declare":               testDeclare,
	"error":                 testError,
	"evalIntegerExpression": testEvalIntegerExpression,
	"function":              testFunction,
	"functionCall":          testFunctionCall,
	"if":                    testIf,
	"infix":                 testInfix,
	"len":                   testLen,
	"let":                   testLet,
	"return":                testReturn,
	"string":                testString,
	"stringConcat":          testStringConcat,
}

func RunAll(t *testing.T, e Evaluator) {
	t.Helper()

	for _, v := range tests {
		v(t, e)
	}
}

func Run(t *testing.T, e Evaluator, usertests []string) {
	t.Helper()

	for _, v := range usertests {
		tests[v](t, e)
	}
}
