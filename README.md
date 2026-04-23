# PostScript Interpreter in Go

## Overview
This project implements a simplified PostScript interpreter in Go. It supports core PostScript functionality including stack operations, arithmetic, control flow, dictionaries, procedures, and both dynamic and lexical scoping.

---

## Build Instructions

Ensure Go is installed.

Build the project:
go build

This produces an executable:
- Windows: ps-interpreter.exe
- macOS/Linux: ps-interpreter

---

## Run Instructions

Start the interpreter:
go run .

or after building:
./ps-interpreter

Exit with:
exit

---

## Scoping Behavior (Required Feature)

This interpreter supports two scoping modes:

### Dynamic Scoping (default)
- Variable lookup occurs at runtime using the current dictionary stack.
- Inner scopes override outer scopes.

Enable with:
dynamic

### Lexical Scoping
- Procedures capture the dictionary environment at definition time.
- Variable lookup uses the stored snapshot environment.

Enable with:
lexical

### Example
/x 10 def
/foo { x = } def

10 dict begin
/x 20 def
foo
end

Dynamic output:
20

Lexical output:
10

---

## Unimplemented / Modified Commands (Section 3.2 Compliance)

This section documents commands that are missing or implemented differently, with technical justification.

---

### dict (MODIFIED)

Specification:
- Takes an integer capacity and creates a dictionary with that size.

Implementation:
- Capacity argument is ignored.
- A Go map[string]interface{} is created instead.

Justification:
- Go maps do not expose or support manual capacity semantics in a way that affects correctness.
- Maps automatically resize and do not behave like fixed-capacity PostScript dictionaries.
- Therefore, the argument is accepted but intentionally unused.

---

### % Comments (PARTIALLY IMPLEMENTED)

Specification:
- Everything after % on a line is treated as a comment and ignored by the interpreter.

Implementation:
- Comments are stripped at the line level before tokenization using strings.Index.

Justification:
- The tokenizer uses strings.Fields, which does not preserve lexical structure or token boundaries required for full inline comment parsing.
- Full PostScript comment handling would require a character-level lexer, which is outside the simplified interpreter design.

---

### Floating-Point Numbers (PARTIAL SUPPORT)

Specification:
- PostScript supports integers and real numbers seamlessly.

Implementation:
- Uses Go float64 and strconv.ParseFloat for numeric parsing.
- Some behavior (especially rounding edge cases) follows Go’s math library rather than PostScript specification.

Justification:
- Go’s math package defines rounding and floating-point behavior.
- Exact PostScript numeric semantics are not fully replicated to avoid implementing a custom decimal arithmetic system.

---

### File Execution (NOT IMPLEMENTED)

Specification:
- PostScript supports loading and executing external files.

Implementation:
- No file loading or execution of external scripts is supported.

Justification:
- The interpreter is REPL-based only.
- Adding file execution would require a file parser and module system beyond the assignment scope.

---

### Tokenization (SIMPLIFIED IMPLEMENTATION)

Specification:
- Full PostScript parsing supports nested structures, multi-line procedures, and complex token rules.

Implementation:
- Input is split using strings.Fields (whitespace-based tokenization only).

Justification:
- A full lexer/parser is not implemented.
- This simplifies parsing but restricts multi-line and advanced syntax handling.

---

### Error Handling (DESIGN CHOICE)

Specification:
- PostScript typically uses runtime errors handled by the interpreter environment.

Implementation:
- Errors use panic() for:
  - stack underflow
  - unknown tokens
  - invalid operations

Justification:
- Go does not use exception-based control flow in the same way.
- panic() is used to simplify error propagation in a project-focused interpreter.

---

## Supported Commands

Stack:
dup, exch, pop, clear, count, copy

Arithmetic:
add, sub, mul, div, idiv, mod, abs, neg, ceiling, floor, round, sqrt

Boolean / Comparison:
eq, ne, lt, gt, le, ge, and, or, not, true, false

Dictionary:
dict, begin, end, def, length, maxlength

String:
get, getinterval, putinterval, length

Control Flow:
if, ifelse, repeat, for, quit

I/O:
print, =, ==

Mode Switching:
lexical, dynamic

---

## Notes

- This is a stack-based interpreter.
- Values are only printed when explicitly requested using = or ==.

Example:
5 5 add =

---

## Summary

This interpreter demonstrates:
- Stack-based execution model
- Procedure handling and closures
- Dictionary-based environments
- Dynamic vs lexical scoping
- Core PostScript execution semantics

---

## Future Improvements

- Full lexer/parser (character-level)
- Multi-line procedure support
- File execution system
- Improved error handling without panic
- Debugging tools (stack tracing)