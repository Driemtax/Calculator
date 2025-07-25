package scanner

import (
	"fmt"
	"math"
	"regexp"
)

// Tokenize breaks down a mathematical expression string into individual tokens.
// It recognizes numbers (integers), arithmetic operators (+, -, *, /),
// parentheses, functions, pi and filters out whitespace and tabs.
// It also replaces constants like pi with its acutal value.
//
// The function returns a slice of strings where each string represents
// a single token from the input expression.
//
// Example:
//
//	tokens := Tokenize("2 + 3 * (4 - 1)")
//	// Returns: ["2", "+", "3", "*", "(", "4", "-", "1", ")"]
func Tokenize(expression string) []string {
	var tokenList = []string{}
	re := regexp.MustCompile(`\d+|[+\-*/()]|sin|cos|tan|pi`)
	tokenList = re.FindAllString(expression, -1)

	// Replace constants for pi support
	for i, token := range tokenList {
		if token == "pi" {
			tokenList[i] = fmt.Sprintf("%f", math.Pi)
		}
	}
	return tokenList
}
