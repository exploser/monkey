package main

import (
	"os"

	"github.com/vasilevp/monkey/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
