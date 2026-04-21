package main

// for I/O
import "fmt"

// stack ops
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

// arithmetic (full set)
func (ip *Interpreter) Add() { b := ip.Pop().(int); a := ip.Pop().(int); ip.Push(a + b) }
func (ip *Interpreter) Sub() { b := ip.Pop().(int); a := ip.Pop().(int); ip.Push(a - b) }
func (ip *Interpreter) Mul() { b := ip.Pop().(int); a := ip.Pop().(int); ip.Push(a * b) }
func (ip *Interpreter) Div() { b := ip.Pop().(int); a := ip.Pop().(int); ip.Push(a / b) }
func (ip *Interpreter) Mod() { b := ip.Pop().(int); a := ip.Pop().(int); ip.Push(a % b) }

// comparisons
func (ip *Interpreter) Eq() { b := ip.Pop(); a := ip.Pop(); ip.Push(a == b) }
func (ip *Interpreter) Ne() { b := ip.Pop(); a := ip.Pop(); ip.Push(a != b) }
func (ip *Interpreter) Lt() { b := ip.Pop().(int); a := ip.Pop().(int); ip.Push(a < b) }
func (ip *Interpreter) Gt() { b := ip.Pop().(int); a := ip.Pop().(int); ip.Push(a > b) }

// boolean
func (ip *Interpreter) And() { b := ip.Pop().(bool); a := ip.Pop().(bool); ip.Push(a && b) }
func (ip *Interpreter) Or()  { b := ip.Pop().(bool); a := ip.Pop().(bool); ip.Push(a || b) }
func (ip *Interpreter) Not() { a := ip.Pop().(bool); ip.Push(!a) }

// dictionary ops
func (ip *Interpreter) Dict() {
	ip.Pop()
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

// I/O
func (ip *Interpreter) PrintEq() {
	fmt.Println(ip.Pop())
}

// scoping
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
