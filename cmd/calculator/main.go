package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/driemtax/Calculator/pkg/calculator"
)

func main() {
	fmt.Print("$ Welcome! To exit just type exit\n")
	for {
		// Getting input from user
		fmt.Print("$ ")
		input, inputErr := bufio.NewReader(os.Stdin).ReadString('\n')
		if inputErr != nil {
			fmt.Println("Error reading input:", inputErr)
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			os.Exit(0)
		}
		// calculate and display the result
		var result, err = calculator.Evaluate(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("$ Result:", result)
		}
	}
}
