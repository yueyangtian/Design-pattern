package Command_Pattern

import "fmt"

type Command interface {
	Execute()
}

type Light struct{}

func (this *Light) TrunOn() {
	fmt.Println("Light turn on")
}

func (this *Light) Execute() {
	this.TrunOn()
}

type TV struct{}

func (this *TV) TrunOn() {
	fmt.Println("tv turn on")
}

func (this *TV) Execute() {
	this.TrunOn()
}

type Switcher struct {
	CommandList []Command
}

func (this *Switcher) ExecandStore(c Command) {
	this.CommandList = append(this.CommandList, c)
	c.Execute()
}
