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
	"fmt"
	"math"
	"slices"
	"strconv"

	"github.com/driemtax/Calculator/pkg/arithmetics"
)

// Parse parses a slice of tokens representing a mathematical expression and returns
// the evaluated result as a float64. It ensures all tokens are consumed during parsing.
func Parse(tokens []string) (float64, error) {
	result, tokenRest, err := parseFunction(tokens)
	if len(tokenRest) != 0 {
		err = errors.Join(err, errors.New("still tokens left"))
	}

	return result, err
}

func parseFunction(tokens []string) (float64, []string, error) {
	operators := []string{"sin", "cos", "tan"}
	var result float64
	var tokenRest []string
	var err error

	if slices.Contains(operators, tokens[0]) {
		if tokens[1] != "(" {
			err = errors.New("did you miss a '('?")
			return result, tokenRest, err
		}
		fmt.Printf("Before iterating with: %q\n", tokens)
		result, tokenRest, err = parseExpression(tokens[1:])

		fmt.Printf("Result: %2f, TokenRest: %q\n", result, tokenRest)
		// now calculate the actual sin,cos, or tan of the result
		switch tokens[0] {
		case "sin":
			result = math.Sin(result)
		case "cos":
			result = math.Cos(result)
		case "tan":
			result = math.Tan(result)
		}
	} else {
		fmt.Printf("Function again: %q\n", tokens)
		result, tokenRest, err = parseExpression(tokens)
	}

	return result, tokenRest, err
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
	operators := []string{"sin", "cos", "tan"}
	result := 0.0
	var err error = nil
	num := false
	tokenRest := tokens

	if tokenRest[0] == "(" {
		fmt.Printf("Factor: %q\n", tokenRest)
		result, tokenRest, err = parseFunction(tokens[1:])
		fmt.Printf("After Function: %2f, %q\n", result, tokenRest)
		if len(tokenRest) == 0 {
			return result, tokenRest, err
		}
		if !(tokenRest[0] == ")") {
			errorMessage := "')' expexted, got " + tokenRest[0]
			err = errors.New(errorMessage)
			return result, tokenRest, err
		}
	} else if slices.Contains(operators, tokenRest[0]) {
		result, tokenRest, err = parseFunction(tokens)
		if len(tokenRest) == 0 {
			return result, tokenRest, err
		}
	} else {
		fmt.Printf("Number: %s\n", tokenRest[0])
		result, num = isNumber(tokenRest[0])
		if !(num) {
			errorMessage := "Number expected, got " + tokenRest[0]
			err = errors.New(errorMessage)
			return result, tokenRest, err
		}
	}

	fmt.Printf("Returning: %2f, %q\n", result, tokenRest)
	return result, tokenRest[1:], err
}

func isNumber(s string) (float64, bool) {
	result, err := strconv.ParseFloat(s, 64)

	return result, err == nil
}
