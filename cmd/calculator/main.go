package main

import (
	"fmt"

	"github.com/driemtax/Calculator/pkg/calculator"
	"github.com/driemtax/Calculator/pkg/scanner"
)

func main() {
	var input = "(1 + 2) * 3"
	var result, err = calculator.Evaluate(input)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}

	var tokenList = scanner.Tokenize(input)

	for i, token := range tokenList {
		fmt.Println("Token ", i, ": ", token)
	}
}
