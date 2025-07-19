// Package parser implements a recursive descent parser for mathematical expressions.
//
// The parser uses the following EBNF grammar:
//
//	exp     -> product (('+' | '-') product)*
//	product -> factor (('*' | '/') factor)*
//	factor  -> '(' exp ')' | NUMBER
//
// Parse parses a slice of tokens representing a mathematical expression and returns
// the evaluated result as a float64. It ensures all tokens are consumed during parsing.
//
// parseExpression handles addition and subtraction operations with left associativity.
//
// parseProduct handles multiplication and division operations with left associativity.
//
// parseFactor handles parenthesized expressions and numeric literals.
package parser

import (
	"errors"
	"slices"
	"strconv"

	"github.com/driemtax/Calculator/pkg/arithmetics"
)

// Parse parses a slice of tokens representing a mathematical expression and returns
// the evaluated result as a float64. It ensures all tokens are consumed during parsing.
func Parse(tokens []string) (float64, error) {
	result, tokenRest, err := parseExpression(tokens)
	if len(tokenRest) != 0 {
		err = errors.Join(err, errors.New("still tokens left"))
	}

	return result, err
}

func parseExpression(tokens []string) (float64, []string, error) {
	operators := []string{"+", "-"}
	result, tokenRest, err := parseProduct(tokens)
	// check that at least 2 tokens are left and the first is an operator. If there would only be a single token left
	// and it is an operator, then there is an error
	for len(tokenRest) > 1 && slices.Contains(operators, tokenRest[0]) {
		operator := tokenRest[0]
		var arg = 0.0
		arg, tokenRest, err = parseProduct(tokenRest[1:])

		if operator == "+" {
			result = arithmetics.Add(result, arg)
		} else { // operator must be '-'
			result = arithmetics.Subtract(result, arg)
		}
	}
	return result, tokenRest, err
}

func parseProduct(tokens []string) (float64, []string, error) {
	operators := []string{"*", "/"}
	result, tokenRest, err := parseFactor(tokens)

	for len(tokenRest) > 1 && slices.Contains(operators, tokenRest[0]) {
		operator := tokenRest[0]
		var arg = 0.0
		arg, tokenRest, err = parseFactor(tokenRest[1:])

		if operator == "*" {
			result = arithmetics.Multiply(result, arg)
		} else {
			result, err = arithmetics.Divide(result, arg)
		}
	}
	return result, tokenRest, err
}

func parseFactor(tokens []string) (float64, []string, error) {
	result := 0.0
	var err error = nil
	num := false
	tokenRest := tokens

	if tokenRest[0] == "(" {
		result, tokenRest, err = parseExpression(tokens[1:])
		if !(tokenRest[0] == ")") {
			errorMessage := "')' expexted, got " + tokenRest[0]
			err = errors.New(errorMessage)
			return result, tokenRest, err
		}
	} else {
		result, num = isNumber(tokenRest[0])
		if !(num) {
			errorMessage := "Number expected, got " + tokenRest[0]
			err = errors.New(errorMessage)
			return result, tokenRest, err
		}
	}

	return result, tokenRest[1:], err
}

func isNumber(s string) (float64, bool) {
	result, err := strconv.ParseFloat(s, 64)

	return result, err == nil
}
