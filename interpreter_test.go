package main

import (
	"strings"
	"testing"
)

// ===================== HELPERS =====================

func run(t *testing.T, prog string) *Interpreter {
	t.Helper()
	ip := NewInterpreter()
	ip.execute(strings.Fields(prog))
	return ip
}

func top(t *testing.T, ip *Interpreter, want interface{}) {
	t.Helper()

	if len(ip.stack) == 0 {
		t.Fatalf("empty stack, expected %v", want)
	}

	got := ip.stack[len(ip.stack)-1]

	if got != want {
		t.Fatalf("expected %v got %v", want, got)
	}
}

// ===================== STACK OPS (1–6) =====================

func TestStackOps(t *testing.T) {
	ip := run(t, "1 2 dup")
	top(t, ip, 2)

	ip = run(t, "1 2 exch")
	top(t, ip, 1)

	ip = run(t, "1 2 pop")
	top(t, ip, 1)

	ip = run(t, "1 2 3 count")
	top(t, ip, 3)

	ip = run(t, "1 2 3 2 copy")
	if len(ip.stack) != 5 {
		t.Fatalf("copy failed")
	}

	ip = run(t, "1 2 3 clear")
	if len(ip.stack) != 0 {
		t.Fatalf("clear failed")
	}
}

// ===================== ARITHMETIC (7–18) =====================

func TestArithmetic(t *testing.T) {
	tests := []struct {
		p    string
		want interface{}
	}{
		{"2 3 add", 5},
		{"5 3 sub", 2},
		{"2 3 mul", 6},
		{"6 3 div", 2.0},
		{"7 3 idiv", 2},
		{"7 3 mod", 1},
		{"-5 abs", 5},
		{"5 neg", -5},
		{"2.9 floor", 2},
		{"2.1 ceiling", 3},
		{"2.6 round", 3},
		{"9 sqrt", 3.0},
	}

	for _, tt := range tests {
		ip := run(t, tt.p)
		top(t, ip, tt.want)
	}
}

// ===================== DICTIONARY (19–24) =====================

func TestDict(t *testing.T) {
	ip := NewInterpreter()

	ip.execute(strings.Fields("/x 10 def /y 20 def"))
	v, ok := ip.Lookup("x")

	if !ok || v.(int) != 10 {
		t.Fatal("def/lookup failed")
	}

	ip = run(t, "10 dict")
	if _, ok := ip.Pop().(map[string]interface{}); !ok {
		t.Fatal("dict failed")
	}
}

// ===================== STRINGS (25–28) =====================

func TestStrings(t *testing.T) {
	ip := run(t, "(hello) 1 get")
	top(t, ip, "e")

	ip = run(t, "(hello) 1 3 getinterval")
	top(t, ip, "ell")
}

// ===================== BOOLEAN + BITWISE (29–39) =====================

func TestBoolean(t *testing.T) {
	tests := []struct {
		p    string
		want interface{}
	}{
		{"true false and", false},
		{"true false or", true},
		{"true not", false},
		{"3 3 eq", true},
		{"3 4 ne", true},
		{"3 5 lt", true},
		{"5 3 gt", true},
		{"3 3 ge", true},
		{"3 3 le", true},
	}

	for _, tt := range tests {
		ip := run(t, tt.p)
		top(t, ip, tt.want)
	}
}

// ===================== FLOW (40–43) =====================

func TestFlow(t *testing.T) {
	ip := NewInterpreter()

	ip.Push(true)
	ip.If(ip.NewProcedure([]string{"1"}))
	top(t, ip, 1)

	ip = NewInterpreter()
	ip.Push(false)
	ip.IfElse(
		ip.NewProcedure([]string{"1"}),
		ip.NewProcedure([]string{"2"}),
	)
	top(t, ip, 2)

	ip = NewInterpreter()
	ip.execute(strings.Fields("1 3 1 { dup } repeat"))
	if len(ip.stack) != 3 {
		t.Fatal("repeat failed")
	}
}

// ===================== IO (45–47) =====================

func TestIO(t *testing.T) {
	ip := run(t, "42 =")
	ip = run(t, "42 ==")
	_ = ip // just ensure no crash
}

// ===================== EDGE CASES =====================

func TestErrors(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()

	ip := NewInterpreter()
	ip.Pop() // underflow
}
