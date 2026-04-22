package main

import "strconv"

// helper function
func extractProcedure(tokens []string, start int) (Procedure, int) {
	depth := 0
	var body []string

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
				return Procedure{tokens: body}, i
			}
			body = append(body, t)
			continue
		}

		body = append(body, t)
	}

	panic("unclosed procedure")
}

func (ip *Interpreter) execute(tokens []string) {
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		// ===================== PROCEDURE HANDLING =====================
		if t == "{" {
			proc, end := extractProcedure(tokens, i)
			ip.Push(proc)
			i = end
			continue
		}

		// ===================== NAME LITERAL =====================
		if len(t) > 1 && t[0] == '/' {
			ip.Push(t[1:])
			continue
		}

		// ===================== STRING LITERAL =====================
		if len(t) >= 2 && t[0] == '(' && t[len(t)-1] == ')' {
			ip.Push(t[1 : len(t)-1])
			continue
		}

		// ===================== COMMANDS =====================
		switch t {

		// STACK
		case "dup":
			ip.Dup()
			continue
		case "exch":
			ip.Exch()
			continue
		case "pop":
			ip.PopOp()
			continue
		case "clear":
			ip.Clear()
			continue
		case "count":
			ip.Count()
			continue
		case "copy":
			ip.Copy()
			continue

		// ARITHMETIC
		case "add":
			ip.Add()
			continue
		case "sub":
			ip.Sub()
			continue
		case "mul":
			ip.Mul()
			continue
		case "div":
			ip.Div()
			continue
		case "idiv":
			ip.Idiv()
			continue
		case "mod":
			ip.Mod()
			continue
		case "abs":
			ip.Abs()
			continue
		case "neg":
			ip.Neg()
			continue
		case "ceiling":
			ip.Ceiling()
			continue
		case "floor":
			ip.Floor()
			continue
		case "round":
			ip.Round()
			continue
		case "sqrt":
			ip.Sqrt()
			continue

		// DICTIONARY
		case "dict":
			ip.Dict()
			continue
		case "length":
			ip.Length()
			continue
		case "maxlength":
			ip.DictMaxLength()
			continue
		case "begin":
			ip.BeginDict()
			continue
		case "end":
			ip.EndDict()
			continue
		case "def":
			ip.DefOp()
			continue

		// STRING
		case "get":
			ip.Get()
			continue
		case "getinterval":
			ip.GetInterval()
			continue
		case "putinterval":
			ip.PutInterval()
			continue

		// BOOLEAN
		case "eq":
			ip.Eq()
			continue
		case "ne":
			ip.Ne()
			continue
		case "lt":
			ip.Lt()
			continue
		case "gt":
			ip.Gt()
			continue
		case "le":
			ip.Le()
			continue
		case "ge":
			ip.Ge()
			continue
		case "and":
			ip.And()
			continue
		case "or":
			ip.Or()
			continue
		case "not":
			ip.Not()
			continue
		case "true":
			ip.Push(true)
			continue
		case "false":
			ip.Push(false)
			continue

		// FLOW
		case "if":
			p := ip.Pop()
			proc, ok := p.(Procedure)
			if !ok {
				panic("if expects procedure")
			}
			ip.If(proc)
			continue

		case "ifelse":
			p2 := ip.Pop()
			p1 := ip.Pop()

			proc1, ok1 := p1.(Procedure)
			proc2, ok2 := p2.(Procedure)

			if !ok1 || !ok2 {
				panic("ifelse expects procedures")
			}

			ip.IfElse(proc1, proc2)
			continue

		case "for":
			ip.For()
			continue

		case "repeat":
			p := ip.Pop()
			n := ip.Pop()

			proc, ok1 := p.(Procedure)
			count, ok2 := n.(int)

			if !ok1 || !ok2 {
				panic("repeat expects (int procedure)")
			}

			ip.Repeat(proc, count)
			continue

		case "quit":
			panic("quit")

		// IO
		case "print":
			ip.Print()
			continue
		case "=":
			ip.PrintEq()
			continue
		case "==":
			ip.PrintPP()
			continue
		}

		// ===================== LITERALS =====================

		// int
		if i, err := strconv.Atoi(t); err == nil {
			ip.Push(i)
			continue
		}

		// float
		if f, err := strconv.ParseFloat(t, 64); err == nil {
			ip.Push(f)
			continue
		}

		// ===================== LOOKUP =====================
		val, ok := ip.Lookup(t)
		if !ok {
			panic("unknown token: " + t)
		}

		switch v := val.(type) {
		case Procedure:
			ip.Call(v)
		default:
			ip.Push(v)
		}
	}
}
