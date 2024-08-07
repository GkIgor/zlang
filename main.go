package main

import (
	"flag"
	"fmt"

	"github.com/GkIgor/zlang/lexer"
)

var Target string = "./default.etq"

func ReadFlags() {
	target := flag.String("o", "default.etq", "Output file")
	file := flag.String("", "", "Input file")
	flag.Parse()

	if target != nil {
		fmt.Println("Output file:", *target)
		ToggleTarget(target)
	}

	if file != nil {
		fmt.Println("Input file:", *file)
	}
}

func ToggleTarget(target *string) {
	Target = *target
}

func main() {
	ReadFlags()
	input := `image:"example.jpg", text:"hello world", attr:"width" "100px", attr:"height" "200px";
	image:"another.jpg", text:"another text";`
	tokens := lexer.Lex(input)

	for _, token := range tokens {
		fmt.Printf("Type: %s, Value: %s\n", token.Type, token.Value)
	}

}
