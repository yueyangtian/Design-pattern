package Mediator_Pattern

import (
	"fmt"
)

type User struct {
	Name string
}

func (this *User) SendMessage(mess string) {
	cr := ChartRoom{No: "01"}
	cr.ShowMessage(fmt.Sprintf("[%s]:%s", this.Name, mess))
}

type ChartRoom struct {
	No string
}

func (this *ChartRoom) ShowMessage(mess string) {
	fmt.Printf("%s %s\n", this.No, mess)
}
