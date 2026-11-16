package main

import (
	"flag"
	"fmt"

	"kr1sal.com/avl-converter/internal/converter"
)

func main() {
	input := flag.String("input", "./data", "Input csv or directory")
	output := flag.String("output", "./data/result.json", "Output json file")
	flag.Parse()

	if input == nil || *input == "" {
		panic(fmt.Errorf("input is undefined!"))
	}
	if output == nil || *output == "" {
		panic(fmt.Errorf("output is undefined!"))
	}

	records, err := converter.ReadCSVFiles(*input)
	if err != nil {
		panic(err)
	}
	err = converter.WriteJSONFile(*output, records)
	if err != nil {
		panic(err)
	}
	fmt.Print(records)
}
