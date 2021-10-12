package State_Pattern

import "fmt"

type Vote interface {
	DoAction()
}

type Voted struct{}

func (*Voted) DoAction() {
	fmt.Println("already voted")
}

type UnVote struct{}

func (*UnVote) DoAction() {
	fmt.Println("vote success")
}

type Votor struct {
	isDone bool
	doDove Vote
}

func (this *Votor) DoVote() {
	if !this.isDone {
		this.doDove = &UnVote{}
	}
	this.doDove.DoAction()
	this.isDone = true
	this.doDove = &Voted{}
}
