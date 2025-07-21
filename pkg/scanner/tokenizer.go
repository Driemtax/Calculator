package scanner

import (
	"regexp"
)

// Tokenize breaks down a mathematical expression string into individual tokens.
// It recognizes numbers (integers), arithmetic operators (+, -, *, /),
// parentheses, and filters out whitespace and tabs.
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
	re := regexp.MustCompile(`\d+|[+\-*/()]|sin|cos|tan`)
	tokenList = re.FindAllString(expression, -1)

	// Replace constants for pi support

	return tokenList
}
