package main

import "fmt"

func main() {
	ip := NewInterpreter()

	ip.Push(10)
	ip.Push(20)
	ip.Add()

	fmt.Println(ip.Pop()) // should print 30
}
