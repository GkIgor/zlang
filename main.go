package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/GkIgor/zlang/file"
	"github.com/GkIgor/zlang/lexer"
)

var Target string = "./default.etq"

func ReadFlags() {
	target := flag.String("o", "default.etq", "Output file")
	file := flag.String("", "", "Input file")
	flag.Parse()

	if target != nil {
		// fmt.Println("Output file:", *target)
		ToggleTarget(target)
	}

	if file != nil {
		// fmt.Println("Input file:", *file)
	}
}

func ToggleTarget(target *string) {
	Target = *target
}

func main() {
	ReadFlags()
	fileName := os.Args[1]

	if len(fileName) <= 0 {
		panic("No input file specified")
	}

	lines := file.ReadFiles(fileName)

	t := lexer.Lex(strings.Join(lines, ""))

	fmt.Println(t)

}
