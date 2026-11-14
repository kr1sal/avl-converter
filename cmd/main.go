package main

import (
	"fmt"

	"kr1sal.com/avl-converter/internal/converter"
)

func main() {
	records, err := converter.ReadCSVFiles("./data/")
	if err != nil {
		panic(err)
	}
	err = converter.WriteJSONFile("./data/program-result.json", records)
	if err != nil {
		panic(err)
	}
	fmt.Print(records)
}
