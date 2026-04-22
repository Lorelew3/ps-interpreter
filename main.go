package main

import "strings"

func main() {
	ip := NewInterpreter()

	// Example program
	program := `
	10 5 add
	3 mul
	`

	tokens := strings.Fields(program)

	ip.execute(tokens)

	// print final stack
	for len(ip.stack) > 0 {
		ip.PrintEq()
	}
}
