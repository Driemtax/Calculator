package main

import (
	"fmt"

	"github.com/driemtax/Calculator/pkg/calculator"
)

func main() {
	var input = "1 + 2 * 3"
	var result, err = calculator.Evaluate(input)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}
}
