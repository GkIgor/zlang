package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/GkIgor/zlang/file"
	"github.com/GkIgor/zlang/flags"
	"github.com/GkIgor/zlang/lexer"
)

func main() {
	flags.ReadFlags()
	fileName := os.Args[1]

	if len(fileName) <= 0 {
		panic("No input file specified")
	}

	lines := file.ReadFiles(fileName)

	t := lexer.Lex(strings.Join(lines, ""))

	fmt.Println(t)

}
