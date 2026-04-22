package main

type Procedure struct {
	tokens     []string
	lexicalEnv []map[string]interface{}
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

// ===================== DICT =====================

func (ip *Interpreter) Def(name string, val interface{}) {
	ip.dictStack[len(ip.dictStack)-1][name] = val
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

// ===================== LOOKUP =====================

func (ip *Interpreter) Lookup(name string) (interface{}, bool) {
	for i := len(ip.dictStack) - 1; i >= 0; i-- {
		if v, ok := ip.dictStack[i][name]; ok {
			return v, true
		}
	}
	return nil, false
}

// ===================== PROCEDURES =====================

func (ip *Interpreter) NewProcedure(tokens []string) Procedure {
	envCopy := make([]map[string]interface{}, len(ip.dictStack))

	for i := range ip.dictStack {
		envCopy[i] = make(map[string]interface{})
		for k, v := range ip.dictStack[i] {
			envCopy[i][k] = v
		}
	}

	return Procedure{
		tokens:     append([]string{}, tokens...),
		lexicalEnv: envCopy,
	}
}

func (ip *Interpreter) Call(p Procedure) {
	old := ip.dictStack
	defer func() { ip.dictStack = old }()

	// 🔥 IMPORTANT: install lexical snapshot if needed
	if ip.lexical {
		ip.dictStack = make([]map[string]interface{}, len(p.lexicalEnv))
		for i := range p.lexicalEnv {
			ip.dictStack[i] = make(map[string]interface{})
			for k, v := range p.lexicalEnv[i] {
				ip.dictStack[i][k] = v
			}
		}
	}

	ip.execute(p.tokens)
}

// ===================== EXEC =====================

func (ip *Interpreter) execute(tokens []string) {
	for i := 0; i < len(tokens); i++ {
		ip.Step(tokens, &i)
	}
}
