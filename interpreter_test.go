package main

import (
	"strings"
	"testing"
)

func run(t *testing.T, prog string) *Interpreter {
	t.Helper()
	ip := NewInterpreter()
	ip.execute(strings.Fields(prog))
	return ip
}

func top(t *testing.T, ip *Interpreter, want interface{}) {
	t.Helper()

	got := ip.stack[len(ip.stack)-1]
	if got != want {
		t.Fatalf("expected %v got %v", want, got)
	}
}

// ---------------- STACK ----------------

func TestStack(t *testing.T) {
	top(t, run(t, "1 2 add"), 3)
	top(t, run(t, "5 2 sub"), 3)
	top(t, run(t, "2 3 mul"), 6)
}

// ---------------- SCOPING ----------------

func TestScoping(t *testing.T) {
	prog := `
/x 10 def
/foo { x } def
10 dict begin
/x 20 def
foo
`

	ip := NewInterpreter()
	ip.lexical = false
	ip.execute(strings.Fields(prog))
	if ip.Pop().(int) != 20 {
		t.Fatal("dynamic failed")
	}

	ip = NewInterpreter()
	ip.lexical = true
	ip.execute(strings.Fields(prog))
	if ip.Pop().(int) != 10 {
		t.Fatal("lexical failed")
	}
}
