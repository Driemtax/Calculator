# Calculator

A recursive descent parser-based mathematical expression calculator written in Go.

## Features

- **Expression Evaluation**: Parses and evaluates mathematical expressions with proper operator precedence
- **Supported Operations**: Addition (`+`), Subtraction (`-`), Multiplication (`*`), Division (`/`)
- **Parentheses Support**: Handles nested parenthetical expressions
- **Error Handling**: Comprehensive error reporting for invalid expressions and division by zero
- **Clean Architecture**: Modular design with separate packages for scanning, parsing, and arithmetic operations

## Grammar

The calculator implements the following EBNF grammar:

```
exp     -> product (('+' | '-') product)*
product -> factor (('*' | '/') factor)*
factor  -> '(' exp ')' | NUMBER
```

This ensures proper operator precedence where multiplication and division have higher precedence than addition and subtraction.

## Installation

```bash
git clone https://github.com/driemtax/Calculator.git
cd Calculator
go mod download
```

## Usage

### Command Line

```bash
go run cmd/calculator/main.go
```

By default, it evaluates the expression `"1 + 2 * 3"`. To use with different expressions, modify the `input` variable in `main.go`.

### As a Library

```go
package main

import (
    "fmt"
    "github.com/driemtax/Calculator/pkg/calculator"
)

func main() {
    result, err := calculator.Evaluate("(1 + 2) * 3")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result) // Output: Result: 9
    }
}
```

## Examples

```go
calculator.Evaluate("2 + 3")           // 5
calculator.Evaluate("2 * 3 + 4")       // 10
calculator.Evaluate("(2 + 3) * 4")     // 20
calculator.Evaluate("10 / 2 - 3")      // 2
calculator.Evaluate("1 + 2 * 3")       // 7 (not 9, due to operator precedence)
```

## Project Structure

```
Calculator/
├── cmd/
│   └── calculator/
│       └── main.go          # CLI application entry point
├── pkg/
│   ├── calculator/
│   │   └── calculator.go    # Main API interface
│   ├── scanner/
│   │   └── tokenizer.go     # Lexical analysis (tokenization)
│   ├── parser/
│   │   └── parser.go        # Recursive descent parser
│   └── arithmetics/
│       └── arithmetics.go   # Basic arithmetic operations
├── go.mod
├── go.sum
└── README.md
```

## Architecture

### Scanner (`pkg/scanner`)
- **Tokenizer**: Breaks input strings into tokens (numbers, operators, parentheses)
- Uses regular expressions to identify valid tokens
- Automatically filters out whitespace

### Parser (`pkg/parser`)
- **Recursive Descent Parser**: Implements the mathematical grammar
- Handles operator precedence correctly
- Provides detailed error messages for malformed expressions
- Returns both the result and any remaining unparsed tokens

### Arithmetics (`pkg/arithmetics`)
- **Basic Operations**: Addition, subtraction, multiplication, division
- **Error Handling**: Detects and reports division by zero
- Pure functions with no side effects

### Calculator (`pkg/calculator`)
- **Main API**: Simple interface combining scanner and parser
- Single entry point for expression evaluation

## Error Handling

The calculator provides comprehensive error reporting:

```go
calculator.Evaluate("1 +")           // Error: Number expected, got end of expression
calculator.Evaluate("(1 + 2")        // Error: ')' expected, got end of expression
calculator.Evaluate("5 / 0")         // Error: division by zero
calculator.Evaluate("1 + 2 3")       // Error: still tokens left
```

## Testing

Run the tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Create a Pull Request

## Future Enhancements

- [ ] Support for floating-point numbers
- [ ] Additional mathematical functions (sin, cos, sqrt, etc.)
- [ ] Variable support
- [ ] Interactive REPL mode
- [ ] More comprehensive test suite
- [ ] Performance optimizations

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Technical Details

- **Language**: Go 1.21+
- **Dependencies**: Only standard library
- **Architecture**: Clean architecture with dependency inversion
- **Parser Type**: Recursive descent parser
- **Error Strategy**: Panic-free error handling with detailed messages

## Performance

The calculator is optimized for correctness and clarity rather than raw performance. For typical mathematical expressions, performance is excellent:

- **Time Complexity**: O(n) where n is the number of tokens
- **Space Complexity**: O(d) where d is the maximum nesting depth of parentheses
- **Memory Allocation**: Minimal allocations, mostly for token slices
