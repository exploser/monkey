package main

import (
	"os"

	"git.exsdev.ru/ExS/gop/repl"
)

func main() {

	repl.Start(os.Stdin, os.Stdout)
}
