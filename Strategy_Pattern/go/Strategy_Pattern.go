package Strategy_Pattern

type strategy interface {
	Operation(int, int) int
}

type AddOperation struct{}

func (*AddOperation) Operation(a, b int) int {
	return a + b
}

type MutilOperation struct{}

func (*MutilOperation) Operation(a, b int) int {
	return a * b
}

type Context struct {
	Option strategy
}

func (this *Context) Exec(a, b int) int {
	return this.Option.Operation(a, b)
}
