package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ip := NewInterpreter()

	fmt.Println("PostScript Interpreter (type 'exit')")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("PS> ")

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		if idx := strings.Index(line, "%"); idx != -1 {
			line = line[:idx]
		}

		if line == "exit" {
			break
		}

		tokens := strings.Fields(line)
		ip.execute(tokens)
	}
}
