package main

import (
	"flag"
	"fmt"
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

}
