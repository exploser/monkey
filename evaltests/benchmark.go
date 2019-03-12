package evaltests

import (
	"testing"
)

func benchmarkFibonacci(b *testing.B, e Evaluator) {
	input := "fib := fn(x) { if (x < 2) { return x; } return fib(x-1)+fib(x-2); }; fib(10);"

	for n := 0; n < b.N; n++ {
		e(b, input)
	}
}
