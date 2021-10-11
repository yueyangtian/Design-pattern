package Interpreter_Pattern

import (
	"strconv"
	"strings"
)

type interpreter interface {
	Interpreter() int
}

type MutilInterpreter struct {
	First  interpreter
	Second interpreter
}

func (this *MutilInterpreter) Interpreter() int {
	return this.First.Interpreter() * this.Second.Interpreter()
}

type AddInterpreter struct {
	First  interpreter
	Second interpreter
}

func (this *AddInterpreter) Interpreter() int {
	return this.First.Interpreter() + this.Second.Interpreter()
}

type NumberInterpreter struct {
	Num string
}

func (this *NumberInterpreter) Interpreter() int {
	if n, err := strconv.Atoi(this.Num); err == nil {
		return n
	}
	return 0
}

type Context struct {
	Interpreter []interpreter
}

func (this *Context) Pop() interpreter {
	if len(this.Interpreter) == 0 {
		return nil
	}
	I := this.Interpreter[len(this.Interpreter)-1]
	this.Interpreter = this.Interpreter[0 : len(this.Interpreter)-1]
	return I
}
func (this *Context) Push(I interpreter) {
	this.Interpreter = append(this.Interpreter, I)
}

func Parser(s string) int {
	ctx := Context{}
	ss := strings.Split(s, " ")
	return Option(ss, &ctx)
}

func Option(ss []string, ctx *Context) int {
	if len(ss) == 0 && len(ctx.Interpreter) == 1 {
		return ctx.Pop().Interpreter()
	}
	if len(ss) == 0 {
		return 0
	}
	n := ss[0]
	if n == "*" {
		second, frist := ctx.Pop(), ctx.Pop()
		if second == nil || frist == nil {
			return 0
		}
		i := MutilInterpreter{First: frist, Second: second}
		ctx.Push(&i)
	} else if n == "+" {
		second, frist := ctx.Pop(), ctx.Pop()
		if second == nil || frist == nil {
			return 0
		}
		i := AddInterpreter{First: frist, Second: second}
		ctx.Push(&i)
	} else if IsNumber(n) {
		i := NumberInterpreter{Num: n}
		ctx.Push(&i)
	} else {
		panic("unkonw synxt ")
	}
	return Option(ss[1:], ctx)
}
func IsNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, n := range s {
		if n > '9' || n < '0' {
			return false
		}
	}
	return true
}
