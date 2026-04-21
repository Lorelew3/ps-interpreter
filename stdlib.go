package main

import (
	"fmt"
	"math"
)

// ===================== STACK OPS =====================

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
	n := ip.Pop().(int)

	if n > len(ip.stack) {
		panic("stack underflow")
	}
	if n < 0 {
		panic("copy expects non-negative integer")
	}

	start := len(ip.stack) - n
	copied := append([]interface{}{}, ip.stack[start:]...)
	ip.stack = append(ip.stack, copied...)
}

// ===================== ARITHMETIC =====================

func (ip *Interpreter) Add() {
	b := ip.Pop().(int)
	a := ip.Pop().(int)
	ip.Push(a + b)
}

func (ip *Interpreter) Sub() {
	b := ip.Pop().(int)
	a := ip.Pop().(int)
	ip.Push(a - b)
}

func (ip *Interpreter) Mul() {
	b := ip.Pop().(int)
	a := ip.Pop().(int)
	ip.Push(a * b)
}

func (ip *Interpreter) Div() {
	b := ip.Pop().(int)
	a := ip.Pop().(int)
	ip.Push(a / b)
}

func (ip *Interpreter) Mod() {
	b := ip.Pop().(int)
	a := ip.Pop().(int)
	ip.Push(a % b)
}

func (ip *Interpreter) Abs() {
	a := ip.Pop().(int)
	if a < 0 {
		a = -a
	}
	ip.Push(a)
}

func (ip *Interpreter) Neg() {
	a := ip.Pop().(int)
	ip.Push(-a)
}

func (ip *Interpreter) Floor() {
	a := ip.Pop().(int)
	ip.Push(int(math.Floor(float64(a))))
}

func (ip *Interpreter) Ceiling() {
	a := ip.Pop().(int)
	ip.Push(int(math.Ceil(float64(a))))
}

func (ip *Interpreter) Round() {
	a := ip.Pop().(int)
	ip.Push(int(math.Round(float64(a))))
}

func (ip *Interpreter) Sqrt() {
	a := ip.Pop().(int)
	ip.Push(int(math.Sqrt(float64(a))))
}

func (ip *Interpreter) Idiv() {
	b := ip.Pop().(int)
	a := ip.Pop().(int)
	ip.Push(a / b)
}

// ===================== boolean/bibtwise =====================

func (ip *Interpreter) Eq() {
	b := ip.Pop()
	a := ip.Pop()

	switch x := a.(type) {
	case int:
		ip.Push(x == b.(int))
	case bool:
		ip.Push(x == b.(bool))
	case string:
		ip.Push(x == b.(string))
	default:
		ip.Push(false) // PostScript usually returns false, not crash
	}
}

func (ip *Interpreter) Ne() {
	b := ip.Pop()
	a := ip.Pop()

	switch x := a.(type) {
	case int:
		ip.Push(x != b.(int))
	case bool:
		ip.Push(x != b.(bool))
	case string:
		ip.Push(x != b.(string))
	default:
		ip.Push(true)
	}
}

func (ip *Interpreter) Lt() {
	b := ip.Pop()
	a := ip.Pop()

	switch x := a.(type) {
	case int:
		ip.Push(x < b.(int))
	case string:
		ip.Push(x < b.(string))
	default:
		panic("lt expects int or string")
	}
}

func (ip *Interpreter) Gt() {
	b := ip.Pop()
	a := ip.Pop()

	switch x := a.(type) {
	case int:
		ip.Push(x > b.(int))
	case string:
		ip.Push(x > b.(string))
	default:
		panic("gt expects int or string")
	}
}

func (ip *Interpreter) Ge() {
	b := ip.Pop()
	a := ip.Pop()

	switch x := a.(type) {
	case int:
		ip.Push(x >= b.(int))
	case string:
		ip.Push(x >= b.(string))
	default:
		panic("ge expects int or string")
	}
}

func (ip *Interpreter) Le() {
	b := ip.Pop()
	a := ip.Pop()

	switch x := a.(type) {
	case int:
		ip.Push(x <= b.(int))
	case string:
		ip.Push(x <= b.(string))
	default:
		panic("le expects int or string")
	}
}

func (ip *Interpreter) And() {
	b := ip.Pop()
	a := ip.Pop()

	switch x := a.(type) {
	case bool:
		ip.Push(x && b.(bool))
	case int:
		ip.Push(x & b.(int))
	default:
		panic("and expects bool or int")
	}
}

func (ip *Interpreter) Or() {
	b := ip.Pop()
	a := ip.Pop()

	switch x := a.(type) {
	case bool:
		ip.Push(x || b.(bool))
	case int:
		ip.Push(x | b.(int))
	default:
		panic("or expects bool or int")
	}
}

func (ip *Interpreter) Not() {
	a := ip.Pop()

	switch x := a.(type) {
	case bool:
		ip.Push(!x)
	case int:
		ip.Push(^x)
	default:
		panic("not expects bool or int")
	}
}

func (ip *Interpreter) True() {
	ip.Push(true)
}

func (ip *Interpreter) False() {
	ip.Push(false)
}

// ===================== DICTIONARY =====================

func (ip *Interpreter) Dict() {
	size := ip.Pop().(int)
	ip.Push(make(map[string]interface{}, size))
}

func (ip *Interpreter) BeginDict() {
	d := ip.Pop().(map[string]interface{})
	ip.Begin(d)
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

func (ip *Interpreter) DictMaxLength() {
	d := ip.Pop().(map[string]interface{})
	ip.Push(len(d))
}

// ===================== OUTPUT =====================

func (ip *Interpreter) PrintEq() {
	fmt.Println(ip.Pop())
}

func (ip *Interpreter) Print() {
	s := ip.Pop().(string)
	fmt.Print(s)
}

func (ip *Interpreter) PrintPP() {
	v := ip.Pop()

	switch x := v.(type) {
	case string:
		fmt.Printf("(%s)\n", x)
	default:
		fmt.Printf("%v\n", x)
	}
}

// ===================== FLOW CONTROL =====================

func (ip *Interpreter) If(proc Procedure) {
	cond := ip.Pop().(bool)
	if cond {
		ip.Call(proc)
	}
}

func (ip *Interpreter) IfElse(p1, p2 Procedure) {
	cond := ip.Pop().(bool)
	if cond {
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
	end := ip.Pop().(int)
	step := ip.Pop().(int)
	start := ip.Pop().(int)

	if step == 0 {
		panic("for loop step cannot be 0")
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

// ===================== SCOPING =====================

func (ip *Interpreter) Call(proc Procedure) {
	if ip.lexical {
		old := ip.dictStack
		ip.dictStack = proc.env
		ip.execute(proc.tokens)
		ip.dictStack = old
	} else {
		ip.execute(proc.tokens)
	}
}

// strings
func (ip *Interpreter) Get() {
	index := ip.Pop().(int)
	s := ip.Pop().(string)

	r := []rune(s)

	if index < 0 || index >= len(r) {
		panic("index out of bounds")
	}

	ip.Push(int(r[index]))
}

func (ip *Interpreter) GetInterval() {
	count := ip.Pop().(int)
	start := ip.Pop().(int)
	s := ip.Pop().(string)

	r := []rune(s)

	if start < 0 || start+count > len(r) {
		panic("substring out of bounds")
	}

	ip.Push(string(r[start : start+count]))
}

func (ip *Interpreter) PutInterval() {
	replacement := ip.Pop().(string)
	start := ip.Pop().(int)
	target := ip.Pop().(string)

	r := []rune(target)
	rep := []rune(replacement)

	if start < 0 || start+len(rep) > len(r) {
		panic("putinterval out of bounds")
	}

	copy(r[start:], rep)
	ip.Push(string(r))
}

// string and dictionary length
func (ip *Interpreter) Length() {
	v := ip.Pop()

	switch x := v.(type) {
	case string:
		ip.Push(len([]rune(x))) // safer for unicode

	case map[string]interface{}:
		ip.Push(len(x))

	case []interface{}: // <-- ADD THIS
		ip.Push(len(x))

	default:
		panic("length expects string, array, or dictionary")
	}
}
