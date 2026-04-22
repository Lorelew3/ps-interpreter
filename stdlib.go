package main

import (
	"fmt"
	"math"
)

// ===================== HELPERS =====================

func asInt(v interface{}) int {
	if x, ok := v.(int); ok {
		return x
	}
	panic("expected int")
}

func asBool(v interface{}) bool {
	if x, ok := v.(bool); ok {
		return x
	}
	panic("expected bool")
}

func asFloat(v interface{}) float64 {
	switch x := v.(type) {
	case int:
		return float64(x)
	case float64:
		return x
	default:
		panic("expected number")
	}
}

// ===================== STACK =====================

func (ip *Interpreter) Dup() {
	v := ip.Pop()
	ip.Push(v)
	ip.Push(v)
}

func (ip *Interpreter) Exch() {
	b := ip.Pop()
	a := ip.Pop()
	ip.Push(b)
	ip.Push(a)
}

func (ip *Interpreter) PopOp() {
	ip.Pop()
}

func (ip *Interpreter) Clear() {
	ip.stack = []interface{}{}
}

func (ip *Interpreter) Count() {
	ip.Push(len(ip.stack))
}

func (ip *Interpreter) Copy() {
	n := asInt(ip.Pop())

	if n < 0 || n > len(ip.stack) {
		panic("copy out of bounds")
	}

	start := len(ip.stack) - n
	slice := append([]interface{}{}, ip.stack[start:]...)

	ip.stack = append(ip.stack, slice...)
}

// ===================== ARITHMETIC =====================

func (ip *Interpreter) Add() {
	b := asInt(ip.Pop())
	a := asInt(ip.Pop())
	ip.Push(a + b)
}

func (ip *Interpreter) Sub() {
	b := asInt(ip.Pop())
	a := asInt(ip.Pop())
	ip.Push(a - b)
}

func (ip *Interpreter) Mul() {
	b := asInt(ip.Pop())
	a := asInt(ip.Pop())
	ip.Push(a * b)
}

func (ip *Interpreter) Div() {
	b := asFloat(ip.Pop())
	a := asFloat(ip.Pop())

	if b == 0 {
		panic("division by zero")
	}

	ip.Push(a / b)
}

func (ip *Interpreter) Idiv() { ip.Div() }

func (ip *Interpreter) Mod() {
	b := asInt(ip.Pop())
	a := asInt(ip.Pop())

	if b == 0 {
		panic("mod by zero")
	}

	ip.Push(a % b)
}

func (ip *Interpreter) Abs() {
	a := asInt(ip.Pop())
	if a < 0 {
		a = -a
	}
	ip.Push(a)
}

func (ip *Interpreter) Neg() {
	ip.Push(-asInt(ip.Pop()))
}

func (ip *Interpreter) Ceiling() {
	ip.Push(int(math.Ceil(asFloat(ip.Pop()))))
}

func (ip *Interpreter) Floor() {
	ip.Push(int(math.Floor(asFloat(ip.Pop()))))
}

func (ip *Interpreter) Round() {
	ip.Push(int(math.Round(asFloat(ip.Pop()))))
}

func (ip *Interpreter) Sqrt() {
	ip.Push(math.Sqrt(asFloat(ip.Pop())))
}

// ===================== DICTIONARY =====================

func (ip *Interpreter) Dict() {
	_ = asInt(ip.Pop()) // capacity ignored (PostScript compatibility)
	ip.Push(make(map[string]interface{}))
}

func (ip *Interpreter) Length() {
	v := ip.Pop()

	switch x := v.(type) {
	case string:
		ip.Push(len([]rune(x)))
	case map[string]interface{}:
		ip.Push(len(x))
	default:
		panic("length expects string or dict")
	}
}

func (ip *Interpreter) DictMaxLength() {
	d := ip.Pop().(map[string]interface{})
	ip.Push(len(d))
}

func (ip *Interpreter) BeginDict() {
	d := ip.Pop()

	m, ok := d.(map[string]interface{})
	if !ok {
		panic("begin expects dict")
	}

	ip.Begin(m)
}

func (ip *Interpreter) EndDict() {
	ip.End()
}

func (ip *Interpreter) DefOp() {
	val := ip.Pop()
	name := ip.Pop().(string)

	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}

	ip.Def(name, val)
}

// ===================== STRING =====================

func (ip *Interpreter) Get() {
	index := asInt(ip.Pop())
	s := []rune(ip.Pop().(string))

	if index < 0 || index >= len(s) {
		panic("out of bounds")
	}

	ip.Push(string(s[index]))
}

func (ip *Interpreter) GetInterval() {
	count := asInt(ip.Pop())
	start := asInt(ip.Pop())
	s := []rune(ip.Pop().(string))

	if start < 0 || start+count > len(s) {
		panic("out of bounds")
	}

	ip.Push(string(s[start : start+count]))
}

func (ip *Interpreter) PutInterval() {
	repl := []rune(ip.Pop().(string))
	start := asInt(ip.Pop())
	target := []rune(ip.Pop().(string))

	if start < 0 || start+len(repl) > len(target) {
		panic("out of bounds")
	}

	for i := 0; i < len(repl); i++ {
		target[start+i] = repl[i]
	}

	ip.Push(string(target))
}

// ===================== BOOLEAN / BITWISE =====================

func (ip *Interpreter) Eq() {
	b := ip.Pop()
	a := ip.Pop()

	ip.Push(a == b)
}

func (ip *Interpreter) Ne() {
	b := ip.Pop()
	a := ip.Pop()

	ip.Push(a != b)
}

func (ip *Interpreter) Lt() {
	b := asInt(ip.Pop())
	a := asInt(ip.Pop())
	ip.Push(a < b)
}

func (ip *Interpreter) Gt() {
	b := asInt(ip.Pop())
	a := asInt(ip.Pop())
	ip.Push(a > b)
}

func (ip *Interpreter) Le() {
	b := asInt(ip.Pop())
	a := asInt(ip.Pop())
	ip.Push(a <= b)
}

func (ip *Interpreter) Ge() {
	b := asInt(ip.Pop())
	a := asInt(ip.Pop())
	ip.Push(a >= b)
}

func (ip *Interpreter) And() {
	b := ip.Pop()
	a := ip.Pop()

	ab, okA := a.(bool)
	bb, okB := b.(bool)

	if okA && okB {
		ip.Push(ab && bb)
		return
	}

	ip.Push(asInt(a) & asInt(b))
}

func (ip *Interpreter) Or() {
	b := ip.Pop()
	a := ip.Pop()

	ab, okA := a.(bool)
	bb, okB := b.(bool)

	if okA && okB {
		ip.Push(ab || bb)
		return
	}

	ip.Push(asInt(a) | asInt(b))
}

func (ip *Interpreter) Not() {
	v := ip.Pop()

	if b, ok := v.(bool); ok {
		ip.Push(!b)
		return
	}

	ip.Push(^asInt(v))
}

func (ip *Interpreter) True()  { ip.Push(true) }
func (ip *Interpreter) False() { ip.Push(false) }

// ===================== FLOW =====================

func (ip *Interpreter) If(proc Procedure) {
	if asBool(ip.Pop()) {
		ip.Call(proc)
	}
}

func (ip *Interpreter) IfElse(p1, p2 Procedure) {
	if asBool(ip.Pop()) {
		ip.Call(p1)
	} else {
		ip.Call(p2)
	}
}

func (ip *Interpreter) Repeat(proc Procedure, n int) {
	for i := 0; i < n; i++ {
		ip.Call(proc)
	}
}

func (ip *Interpreter) For() {
	proc := ip.Pop().(Procedure)
	end := asInt(ip.Pop())
	step := asInt(ip.Pop())
	start := asInt(ip.Pop())

	if step == 0 {
		panic("step 0")
	}

	if step > 0 {
		for i := start; i <= end; i += step {
			ip.Push(i)
			ip.Call(proc)
		}
	} else {
		for i := start; i >= end; i += step {
			ip.Push(i)
			ip.Call(proc)
		}
	}
}

func (ip *Interpreter) Quit() {
	panic("quit")
}

// ===================== IO =====================
func (ip *Interpreter) Print() {
	fmt.Print(ip.Pop())
}

func (ip *Interpreter) PrintEq() {
	fmt.Println(ip.Pop())
}

func (ip *Interpreter) PrintPP() {
	fmt.Printf("%#v\n", ip.Pop())
}
