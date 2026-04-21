package main

// execution engine
func (ip *Interpreter) execute(tokens []string) {
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		switch t {

		case "add":
			ip.Add()
		case "sub":
			ip.Sub()
		case "mul":
			ip.Mul()
		case "div":
			ip.Div()
		case "mod":
			ip.Mod()
		case "abs":
			ip.Abs()
		case "neg":
			ip.Neg()

		case "floor":
			ip.Floor()
		case "ceiling":
			ip.Ceiling()
		case "round":
			ip.Round()
		case "sqrt":
			ip.Sqrt()
		case "idiv":
			ip.Idiv()
		case "copy":
			ip.Copy()

		case "dictLength":
			ip.DictLength()
		case "maxlength":
			ip.DictMaxLength()

		case "dup":
			ip.Dup()
		case "quit":
			ip.Quit()

		case "exch":
			ip.Exch()
		case "pop":
			ip.PopOp()
		case "clear":
			ip.Clear()
		case "count":
			ip.Count()
		case "for":
			ip.For()

		case "length":
			ip.StrLength()
		case "get":
			ip.Get()
		case "getinterval":
			ip.GetInterval()
		case "putinterval":
			ip.PutInterval()

		case "eq":
			ip.Eq()
		case "ne":
			ip.Ne()
		case "lt":
			ip.Lt()
		case "gt":
			ip.Gt()
		case "ge":
			ip.Ge()
		case "le":
			ip.Le()

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

		case "dict":
			ip.Dict()
		case "begin":
			ip.BeginDict()
		case "end":
			ip.EndDict()
		case "def":
			ip.DefOp()

		case "=":
			ip.PrintEq()
		case "print":
			ip.Print()
		case "==":
			ip.PrintPP()

		default:
			if val, ok := ip.Lookup(t); ok {
				switch v := val.(type) {
				case Procedure:
					ip.Call(v)
				default:
					ip.Push(v)
				}
			}
		}
	}
}
