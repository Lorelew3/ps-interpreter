package main

import (
	"strconv"
)

// ===================== PROCEDURE =====================

func extractProcedure(ip *Interpreter, tokens []string, start int) (Procedure, int) {
	depth := 0
	body := []string{}

	for i := start; i < len(tokens); i++ {
		t := tokens[i]

		if t == "{" {
			depth++
			if depth > 1 {
				body = append(body, t)
			}
			continue
		}

		if t == "}" {
			depth--
			if depth == 0 {
				return ip.NewProcedure(body), i
			}
			body = append(body, t)
			continue
		}

		body = append(body, t)
	}

	panic("unclosed procedure")
}

// ===================== STEP =====================

func (ip *Interpreter) Step(tokens []string, i *int) {
	t := tokens[*i]

	// procedure
	if t == "{" {
		p, end := extractProcedure(ip, tokens, *i)
		ip.Push(p)
		*i = end
		return
	}

	// literal name
	if len(t) > 0 && t[0] == '/' {
		ip.Push(t[1:])
		return
	}

	// string
	if len(t) >= 2 && t[0] == '(' && t[len(t)-1] == ')' {
		ip.Push(t[1 : len(t)-1])
		return
	}

	// ===== ALL COMMANDS FROM stdlib.go =====
	switch t {

	// STACK
	case "dup":
		ip.Dup()
	case "exch":
		ip.Exch()
	case "pop":
		ip.PopOp()
	case "clear":
		ip.Clear()
	case "count":
		ip.Count()
	case "copy":
		ip.Copy()

	// ARITH
	case "add":
		ip.Add()
	case "sub":
		ip.Sub()
	case "mul":
		ip.Mul()
	case "div":
		ip.Div()
	case "idiv":
		ip.Idiv()
	case "mod":
		ip.Mod()
	case "abs":
		ip.Abs()
	case "neg":
		ip.Neg()
	case "ceiling":
		ip.Ceiling()
	case "floor":
		ip.Floor()
	case "round":
		ip.Round()
	case "sqrt":
		ip.Sqrt()

	// DICT
	case "dict":
		ip.Dict()
	case "length":
		ip.Length()
	case "maxlength":
		ip.DictMaxLength()
	case "begin":
		ip.BeginDict()
	case "end":
		ip.EndDict()
	case "def":
		ip.DefOp()

	// STRING
	case "get":
		ip.Get()
	case "getinterval":
		ip.GetInterval()
	case "putinterval":
		ip.PutInterval()

	// BOOL
	case "eq":
		ip.Eq()
	case "ne":
		ip.Ne()
	case "lt":
		ip.Lt()
	case "gt":
		ip.Gt()
	case "le":
		ip.Le()
	case "ge":
		ip.Ge()
	case "and":
		ip.And()
	case "or":
		ip.Or()
	case "not":
		ip.Not()
	case "true":
		ip.True()
	case "false":
		ip.False()

	// FLOW
	case "if":
		p := ip.Pop().(Procedure)
		if ip.Pop().(bool) {
			ip.Call(p)
		}

	case "ifelse":
		p2 := ip.Pop().(Procedure)
		p1 := ip.Pop().(Procedure)
		if ip.Pop().(bool) {
			ip.Call(p1)
		} else {
			ip.Call(p2)
		}

	case "repeat":
		p := ip.Pop().(Procedure)
		n := ip.Pop().(int)
		for i := 0; i < n; i++ {
			ip.Call(p)
		}

	case "for":
		ip.For()

	case "quit":
		ip.Quit()

	// IO
	case "print":
		ip.Print()
	case "=":
		ip.PrintEq()
	case "==":
		ip.PrintPP()

	// MODE SWITCH
	case "lexical":
		ip.lexical = true
	case "dynamic":
		ip.lexical = false

	// NUMBER
	default:
		if v, err := strconv.Atoi(t); err == nil {
			ip.Push(v)
			return
		}

		val, ok := ip.Lookup(t)
		if !ok {
			panic("unknown token: " + t)
		}

		if p, ok := val.(Procedure); ok {
			ip.Call(p)
		} else {
			ip.Push(val)
		}
	}
}
