package calculator

import (
	"github.com/driemtax/Calculator/pkg/parser"
	"github.com/driemtax/Calculator/pkg/scanner"
)

func Evaluate(expression string) (float64, error) {
	tokenList := scanner.Tokenize(expression)
	result, err := parser.Parse(tokenList)

	return result, err
}
