package main

import "testing"

func TestAdd(t *testing.T) {
	ip := NewInterpreter()
	ip.Push(10)
	ip.Push(20)
	ip.Add()

	if ip.Pop() != 30 {
		t.Fail()
	}
}

func TestDup(t *testing.T) {
	ip := NewInterpreter()
	ip.Push(5)
	ip.Dup()

	a := ip.Pop()
	b := ip.Pop()

	if a != 5 || b != 5 {
		t.Fail()
	}
}
