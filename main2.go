package main

import (
	"os"

	"git.exsdev.ru/ExS/monkey/repl"
)

//go:generate go build tools/cmd/stringer/stringer.go
//go:generate go build

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
