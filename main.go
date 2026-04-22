package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ip := NewInterpreter()

	fmt.Println("PostScript Interpreter (type 'exit' to quit)")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("PS> ")

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		if line == "exit" {
			break
		}

		tokens := strings.Fields(line)

		for i := 0; i < len(tokens); i++ {
			ip.Step(tokens, &i)

			// 🔥 SHOW STACK AFTER EACH STEP
			fmt.Println("Stack:", ip.stack)
		}
	}
}
