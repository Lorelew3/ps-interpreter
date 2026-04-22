package main

import "strconv"

// ===================== PROCEDURE PARSER =====================

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

// ===================== STEP EXECUTION =====================

func (ip *Interpreter) Step(tokens []string, i *int) {
	t := tokens[*i]

	// PROCEDURE
	if t == "{" {
		proc, end := extractProcedure(tokens, *i)
		ip.Push(proc)
		*i = end
		return
	}

	// NAME LITERAL
	if len(t) > 1 && t[0] == '/' {
		ip.Push(t[1:])
		return
	}

	// STRING
	if len(t) >= 2 && t[0] == '(' && t[len(t)-1] == ')' {
		ip.Push(t[1 : len(t)-1])
		return
	}

	// COMMANDS
	switch t {

	// STACK
	case "dup":
		ip.Dup()
		return
	case "exch":
		ip.Exch()
		return
	case "pop":
		ip.PopOp()
		return
	case "clear":
		ip.Clear()
		return
	case "count":
		ip.Count()
		return
	case "copy":
		ip.Copy()
		return

	// ARITHMETIC
	case "add":
		ip.Add()
		return
	case "sub":
		ip.Sub()
		return
	case "mul":
		ip.Mul()
		return
	case "div":
		ip.Div()
		return
	case "idiv":
		ip.Idiv()
		return
	case "mod":
		ip.Mod()
		return
	case "abs":
		ip.Abs()
		return
	case "neg":
		ip.Neg()
		return
	case "ceiling":
		ip.Ceiling()
		return
	case "floor":
		ip.Floor()
		return
	case "round":
		ip.Round()
		return
	case "sqrt":
		ip.Sqrt()
		return

	// DICTIONARY
	case "dict":
		ip.Dict()
		return
	case "length":
		ip.Length()
		return
	case "maxlength":
		ip.DictMaxLength()
		return
	case "begin":
		ip.BeginDict()
		return
	case "end":
		ip.EndDict()
		return
	case "def":
		ip.DefOp()
		return

	// STRING
	case "get":
		ip.Get()
		return
	case "getinterval":
		ip.GetInterval()
		return
	case "putinterval":
		ip.PutInterval()
		return

	// BOOLEAN
	case "eq":
		ip.Eq()
		return
	case "ne":
		ip.Ne()
		return
	case "lt":
		ip.Lt()
		return
	case "gt":
		ip.Gt()
		return
	case "le":
		ip.Le()
		return
	case "ge":
		ip.Ge()
		return
	case "and":
		ip.And()
		return
	case "or":
		ip.Or()
		return
	case "not":
		ip.Not()
		return
	case "true":
		ip.Push(true)
		return
	case "false":
		ip.Push(false)
		return

	// FLOW
	case "if":
		p := ip.Pop().(Procedure)
		ip.If(p)
		return

	case "ifelse":
		p2 := ip.Pop().(Procedure)
		p1 := ip.Pop().(Procedure)
		ip.IfElse(p1, p2)
		return

	case "for":
		ip.For()
		return

	case "repeat":
		proc := ip.Pop().(Procedure)
		count := asInt(ip.Pop())
		ip.Repeat(proc, count)
		return

	case "quit":
		panic("quit")

	// IO
	case "print":
		ip.Print()
		return
	case "=":
		ip.PrintEq()
		return
	case "==":
		ip.PrintPP()
		return
	}

	// INT
	if v, err := strconv.Atoi(t); err == nil {
		ip.Push(v)
		return
	}

	// FLOAT
	if v, err := strconv.ParseFloat(t, 64); err == nil {
		ip.Push(v)
		return
	}

	// LOOKUP
	val, ok := ip.Lookup(t)
	if !ok {
		panic("unknown token: " + t)
	}

	if proc, ok := val.(Procedure); ok {
		ip.Call(proc)
	} else {
		ip.Push(val)
	}
}

// ===================== ORIGINAL EXECUTE =====================

func (ip *Interpreter) execute(tokens []string) {
	for i := 0; i < len(tokens); i++ {
		ip.Step(tokens, &i)
	}
}
