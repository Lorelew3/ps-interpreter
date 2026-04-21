package main

type Procedure struct {
	tokens []string
	env    []map[string]interface{}
}

type Interpreter struct {
	stack     []interface{}
	dictStack []map[string]interface{}
	lexical   bool
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		stack:     []interface{}{},
		dictStack: []map[string]interface{}{{}},
		lexical:   false,
	}
}

// ---------- STACK ----------
func (ip *Interpreter) Push(v interface{}) {
	ip.stack = append(ip.stack, v)
}

func (ip *Interpreter) Pop() interface{} {
	v := ip.stack[len(ip.stack)-1]
	ip.stack = ip.stack[:len(ip.stack)-1]
	return v
}

// ---------- DICTIONARY ----------
func (ip *Interpreter) Def(name string, val interface{}) {
	ip.dictStack[len(ip.dictStack)-1][name] = val
}

func (ip *Interpreter) Lookup(name string) (interface{}, bool) {
	for i := len(ip.dictStack) - 1; i >= 0; i-- {
		if v, ok := ip.dictStack[i][name]; ok {
			return v, true
		}
	}
	return nil, false
}

func (ip *Interpreter) Begin(d map[string]interface{}) {
	ip.dictStack = append(ip.dictStack, d)
}

func (ip *Interpreter) End() {
	ip.dictStack = ip.dictStack[:len(ip.dictStack)-1]
}
