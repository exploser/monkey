package evaltests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vasilevp/monkey/types"
)

type Evaluator func(t testing.TB, input string) types.Object

type iparser interface {
	Errors() []error
}

func CheckParserErrors(t *testing.T, p iparser) {
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

var benchmarks = map[string]func(t *testing.B, e Evaluator){
	"fibonacci": benchmarkFibonacci,
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
		if tt, ok := tests[v]; ok {
			tt(t, e)
		} else {
			t.Errorf("test %q not found", v)
		}
	}
}

func BenchmarkAll(b *testing.B, e Evaluator) {
	b.Helper()

	for _, v := range benchmarks {
		v(b, e)
	}
}

func Benchmark(b *testing.B, e Evaluator, usertests []string) {
	b.Helper()

	for _, v := range usertests {
		if tt, ok := benchmarks[v]; ok {
			tt(b, e)
		} else {
			b.Errorf("benchmark %q not found", v)
		}
	}
}
