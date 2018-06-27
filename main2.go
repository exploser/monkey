package main

import (
	"os"

	"git.exsdev.ru/ExS/monkey/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
