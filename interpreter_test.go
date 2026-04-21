package main

import "testing"

// stack tests
func TestDup(t *testing.T) {
	ip := NewInterpreter()
	ip.Push(5)
	ip.Dup()

	a := ip.Pop()
	b := ip.Pop()

	if a != 5 || b != 5 {
		t.Errorf("expected 5,5 got %v,%v", a, b)
	}
}

func TestExch(t *testing.T) {
	ip := NewInterpreter()
	ip.Push(1)
	ip.Push(2)
	ip.Exch()

	if ip.Pop() != 1 || ip.Pop() != 2 {
		t.Fail()
	}
}

func TestCount(t *testing.T) {
	ip := NewInterpreter()
	ip.Push(1)
	ip.Push(2)
	ip.Count()

	if ip.Pop() != 2 {
		t.Fail()
	}
}

// arithmetic tests
func TestAdd(t *testing.T) {
	ip := NewInterpreter()
	ip.Push(10)
	ip.Push(20)
	ip.Add()

	if ip.Pop() != 30 {
		t.Fail()
	}
}

func TestSubMulDivMod(t *testing.T) {
	ip := NewInterpreter()

	ip.Push(10)
	ip.Push(2)
	ip.Sub()
	if ip.Pop() != 8 {
		t.Fail()
	}

	ip.Push(3)
	ip.Push(4)
	ip.Mul()
	if ip.Pop() != 12 {
		t.Fail()
	}

	ip.Push(20)
	ip.Push(5)
	ip.Div()
	if ip.Pop() != 4 {
		t.Fail()
	}

	ip.Push(10)
	ip.Push(3)
	ip.Mod()
	if ip.Pop() != 1 {
		t.Fail()
	}
}

// boolean and comparison tests
func TestBoolOps(t *testing.T) {
	ip := NewInterpreter()

	ip.Push(true)
	ip.Push(false)
	ip.And()
	if ip.Pop() != false {
		t.Fail()
	}

	ip.Push(true)
	ip.Push(false)
	ip.Or()
	if ip.Pop() != true {
		t.Fail()
	}

	ip.Push(true)
	ip.Not()
	if ip.Pop() != false {
		t.Fail()
	}
}

func TestCompare(t *testing.T) {
	ip := NewInterpreter()

	ip.Push(5)
	ip.Push(5)
	ip.Eq()
	if ip.Pop() != true {
		t.Fail()
	}

	ip.Push(3)
	ip.Push(5)
	ip.Lt()
	if ip.Pop() != true {
		t.Fail()
	}

	ip.Push(5)
	ip.Push(3)
	ip.Gt()
	if ip.Pop() != true {
		t.Fail()
	}
}

// dictionary tests
func TestDefLookup(t *testing.T) {
	ip := NewInterpreter()

	ip.Push("x")
	ip.Push(10)
	ip.DefOp()

	val, ok := ip.Lookup("x")
	if !ok || val != 10 {
		t.Fail()
	}
}

// scoping test
func TestScopingDynamic(t *testing.T) {
	ip := NewInterpreter()
	ip.lexical = false // dynamic

	// x = 10
	ip.Push("x")
	ip.Push(10)
	ip.DefOp()

	val, _ := ip.Lookup("x")
	if val != 10 {
		t.Fail()
	}
}

func TestScopingLexicalToggle(t *testing.T) {
	ip := NewInterpreter()
	ip.lexical = true // lexical mode ON

	ip.Push("x")
	ip.Push(20)
	ip.DefOp()

	val, ok := ip.Lookup("x")
	if !ok || val != 20 {
		t.Fail()
	}
}
