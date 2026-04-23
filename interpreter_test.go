package main

import (
	"testing"
)

// ===================== HELPERS =====================

func newIP() *Interpreter {
	return NewInterpreter()
}

func mustInt(t *testing.T, v interface{}) int {
	t.Helper()
	x, ok := v.(int)
	if !ok {
		t.Fatalf("expected int, got %T (%v)", v, v)
	}
	return x
}

func mustBool(t *testing.T, v interface{}) bool {
	t.Helper()
	x, ok := v.(bool)
	if !ok {
		t.Fatalf("expected bool, got %T (%v)", v, v)
	}
	return x
}

// ===================== STACK OPS =====================

func TestStackOps(t *testing.T) {
	ip := newIP()

	ip.Push(1)
	ip.Push(2)
	ip.Dup()

	if len(ip.stack) != 3 {
		t.Fatalf("expected stack size 3, got %d", len(ip.stack))
	}

	ip.Exch()
	top := mustInt(t, ip.Pop())
	if top != 2 {
		t.Fatalf("expected 2, got %d", top)
	}
}

func TestClearCount(t *testing.T) {
	ip := newIP()

	ip.Push(1)
	ip.Push(2)
	ip.Push(3)

	ip.Count()
	cnt := mustInt(t, ip.Pop())
	if cnt != 3 {
		t.Fatalf("expected count 3, got %d", cnt)
	}

	ip.Clear()
	if len(ip.stack) != 0 {
		t.Fatalf("expected empty stack")
	}
}

func TestCopy(t *testing.T) {
	ip := newIP()

	ip.Push(1)
	ip.Push(2)
	ip.Push(3)

	ip.Push(2)

	ip.Copy()

	if len(ip.stack) != 5 {
		t.Fatalf("expected 5 items, got %d", len(ip.stack))
	}

	if mustInt(t, ip.stack[len(ip.stack)-1]) != 2 {
		t.Fatalf("copy failed")
	}
}

// ===================== ARITHMETIC =====================

func TestArithmetic(t *testing.T) {
	ip := newIP()

	ip.Push(10)
	ip.Push(5)
	ip.Add()
	if mustInt(t, ip.Pop()) != 15 {
		t.Fatal("add failed")
	}

	ip.Push(10)
	ip.Push(3)
	ip.Sub()
	if mustInt(t, ip.Pop()) != 7 {
		t.Fatal("sub failed")
	}

	ip.Push(6)
	ip.Push(2)
	ip.Mul()
	if mustInt(t, ip.Pop()) != 12 {
		t.Fatal("mul failed")
	}

	ip.Push(10)
	ip.Push(2)
	ip.Div()
	if mustInt(t, ip.Pop()) != 5 {
		t.Fatal("div failed")
	}

	ip.Push(10)
	ip.Push(3)
	ip.Mod()
	if mustInt(t, ip.Pop()) != 1 {
		t.Fatal("mod failed")
	}
}

// ===================== DICTIONARY =====================

func TestDefAndLookup(t *testing.T) {
	ip := newIP()

	ip.Push("x")
	ip.Push(42)
	ip.DefOp()

	v, ok := ip.Lookup("x")
	if !ok {
		t.Fatal("lookup failed")
	}
	if mustInt(t, v) != 42 {
		t.Fatal("def/lookup mismatch")
	}
}

func TestBeginEndDict(t *testing.T) {
	ip := newIP()

	d := make(map[string]interface{})
	d["a"] = 1

	ip.Push(d)
	ip.BeginDict()

	ip.Push(99)
	ip.Def("b", 99)

	v, ok := ip.Lookup("b")
	if !ok || mustInt(t, v) != 99 {
		t.Fatal("begin/end dict failed")
	}

	ip.EndDict()
}

// ===================== STRING OPS =====================

func TestStringLength(t *testing.T) {
	ip := newIP()

	ip.Push("hello")
	ip.Length()

	if mustInt(t, ip.Pop()) != 5 {
		t.Fatal("string length failed")
	}
}

// ===================== BOOLEAN =====================

func TestBoolOps(t *testing.T) {
	ip := newIP()

	ip.Push(3)
	ip.Push(5)
	ip.Lt()
	if mustBool(t, ip.Pop()) != true {
		t.Fatal("lt failed")
	}

	ip.Push(true)
	ip.Push(false)
	ip.And()
	if mustBool(t, ip.Pop()) != false {
		t.Fatal("and failed")
	}
}

// ===================== PROCEDURE =====================

func TestProcedureSimple(t *testing.T) {
	ip := newIP()

	// { 1 2 add }
	tokens := []string{"{", "1", "2", "add", "}"}

	p, _ := extractProcedure(ip, tokens, 0)
	ip.Push(p)
	proc := ip.Pop().(Procedure)

	ip.Call(proc)

	if mustInt(t, ip.Pop()) != 3 {
		t.Fatal("procedure execution failed")
	}
}

// ===================== FLOW CONTROL =====================

func TestRepeat(t *testing.T) {
	ip := newIP()

	// { 1 } 3 repeat => 1 1 1
	tokens := []string{"3", "{", "1", "}", "repeat"}
	// simulate step parsing
	ip.execute(tokens)

	if len(ip.stack) != 3 {
		t.Fatalf("expected 3 items, got %d", len(ip.stack))
	}
}

// ===================== LOOKUP + DYNAMIC CALL =====================

func TestDynamicProcedureCall(t *testing.T) {
	ip := newIP()

	ip.Push(Procedure{tokens: []string{"2", "3", "add"}})
	ip.Def("addTwo", ip.Pop())

	tokens := []string{"addTwo"}
	ip.execute(tokens)

	if mustInt(t, ip.Pop()) != 5 {
		t.Fatal("dynamic call failed")
	}
}
