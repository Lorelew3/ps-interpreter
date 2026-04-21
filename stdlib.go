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

	start := len(ip.stack) - n
	copied := make([]interface{}, n)
	copy(copied, ip.stack[start:])
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
	ip.Push(ip.Pop().(int))
}

func (ip *Interpreter) Ceiling() {
	ip.Push(ip.Pop().(int))
}

func (ip *Interpreter) Round() {
	ip.Push(ip.Pop().(int))
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

// ===================== COMPARISONS =====================

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
	b := ip.Pop().(int)
	a := ip.Pop().(int)
	ip.Push(a < b)
}

func (ip *Interpreter) Gt() {
	b := ip.Pop().(int)
	a := ip.Pop().(int)
	ip.Push(a > b)
}

func (ip *Interpreter) Ge() {
	b := ip.Pop().(int)
	a := ip.Pop().(int)
	ip.Push(a >= b)
}

func (ip *Interpreter) Le() {
	b := ip.Pop().(int)
	a := ip.Pop().(int)
	ip.Push(a <= b)
}

// ===================== BOOLEAN =====================

func (ip *Interpreter) And() {
	b := ip.Pop().(bool)
	a := ip.Pop().(bool)
	ip.Push(a && b)
}

func (ip *Interpreter) Or() {
	b := ip.Pop().(bool)
	a := ip.Pop().(bool)
	ip.Push(a || b)
}

func (ip *Interpreter) Not() {
	a := ip.Pop().(bool)
	ip.Push(!a)
}

func (ip *Interpreter) True() {
	ip.Push(true)
}

func (ip *Interpreter) False() {
	ip.Push(false)
}

// ===================== DICTIONARY =====================

func (ip *Interpreter) Dict() {
	ip.Pop() // ignore size
	ip.Push(make(map[string]interface{}))
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
	ip.Def(name, val)
}

// ===================== OUTPUT =====================

func (ip *Interpreter) PrintEq() {
	fmt.Println(ip.Pop())
}

func (ip *Interpreter) Print() {
	fmt.Print(ip.Pop())
}

func (ip *Interpreter) PrintPP() {
	fmt.Printf("%#v\n", ip.Pop())
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
