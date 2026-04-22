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

// ===================== INIT =====================

func NewInterpreter() *Interpreter {
	ip := &Interpreter{
		stack:     []interface{}{},
		dictStack: []map[string]interface{}{},
		lexical:   false,
	}

	ip.dictStack = append(ip.dictStack, make(map[string]interface{}))
	return ip
}

func (ip *Interpreter) NewProcedure(tokens []string) Procedure {
	return Procedure{
		tokens: append([]string{}, tokens...),
		env: []map[string]interface{}{
			ip.dictStack[len(ip.dictStack)-1],
		},
	}
}

// ===================== STACK =====================

func (ip *Interpreter) Push(v interface{}) {
	ip.stack = append(ip.stack, v)
}

func (ip *Interpreter) Pop() interface{} {
	if len(ip.stack) == 0 {
		panic("stack underflow")
	}
	v := ip.stack[len(ip.stack)-1]
	ip.stack = ip.stack[:len(ip.stack)-1]
	return v
}

func cloneDictStack(src []map[string]interface{}) []map[string]interface{} {
	dst := make([]map[string]interface{}, len(src))
	for i, d := range src {
		nd := make(map[string]interface{}, len(d))
		for k, v := range d {
			nd[k] = v
		}
		dst[i] = nd
	}
	return dst
}

// ===================== DICTIONARY =====================

func (ip *Interpreter) Def(name string, val interface{}) {
	top := ip.dictStack[len(ip.dictStack)-1]
	top[name] = val
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
	if len(ip.dictStack) <= 1 {
		panic("cannot pop global dict")
	}
	ip.dictStack = ip.dictStack[:len(ip.dictStack)-1]
}

// ===================== PROCEDURES =====================
func (ip *Interpreter) Call(proc Procedure) {
	oldLen := len(ip.dictStack)

	// restore even if panic happens
	defer func() {
		ip.dictStack = ip.dictStack[:oldLen]
	}()

	// IMPORTANT: shallow env is unsafe → clone
	for _, d := range proc.env {
		copyMap := make(map[string]interface{}, len(d))
		for k, v := range d {
			copyMap[k] = v
		}
		ip.dictStack = append(ip.dictStack, copyMap)
	}

	ip.execute(proc.tokens)
}
