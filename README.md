# Calculator

A simple mathematical expression calculator built with Go as a learning project for implementing recursive descent parsers.

## What it does

Evaluates mathematical expressions with proper operator precedence:
- Basic operations: `+`, `-`, `*`, `/`
- Parentheses for grouping: `(1 + 2) * 3`
- Handles operator precedence correctly: `1 + 2 * 3` = `7`

## Grammar
I decided to implement the following EBNF grammar, since it elimates left-recursion in a more efficent way then context-free grammars do.
```
exp     -> product (('+' | '-') product)*
product -> factor (('*' | '/') factor)*
factor  -> '(' exp ')' | NUMBER
```

## Usage

```bash
git clone https://github.com/driemtax/Calculator.git
cd Calculator
go run cmd/calculator/main.go
```

Or use as a library:

```go
import "github.com/driemtax/Calculator/pkg/calculator"

result, err := calculator.Evaluate("(1 + 2) * 3")
// result: 9.0
```

## Project Structure

```
Calculator/
├── cmd/calculator/main.go       # CLI entry point
├── pkg/
│   ├── calculator/              # Main API
│   ├── scanner/                 # Tokenization
│   ├── parser/                  # Recursive descent parser
│   └── arithmetics/             # Basic math operations
```

## Learning Goals

This project demonstrates:
- **Lexical Analysis**: Breaking strings into tokens
- **Recursive Descent Parsing**: Implementing a grammar with proper precedence
- **Go Project Structure**: Clean package organization
- **Error Handling**: Meaningful error messages for invalid input

Built purely with Go's standard library - no external dependencies.
