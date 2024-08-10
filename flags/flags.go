package flags

import (
	"flag"

	"github.com/GkIgor/zlang/target"
)

func ReadFlags() {
	output := flag.String("o", "default.etq", "Output file")
	file := flag.String("", "", "Input file")
	flag.Parse()

	if output != nil {
		// fmt.Println("Output file:", *output)
		target.SetTarget(*output)
	}

	if file != nil {
		// fmt.Println("Input file:", *file)
	}
}
