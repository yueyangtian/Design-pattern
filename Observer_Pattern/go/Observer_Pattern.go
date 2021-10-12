package Observer_Pattern

import "fmt"

type Object interface {
	SetStat(int)
	NotifyAll()
	Attach(Observer)
}

type Observer interface {
	Notify(string)
}

type Object1 struct {
	List []Observer
	stat int
}

type Observer1 struct{}

func (this *Object1) SetStat(s int) {
	this.stat = s
	this.NotifyAll()
}

func (this *Object1) NotifyAll() {
	s := fmt.Sprintf("from object 1 stat %d notified", this.stat)
	for _, n := range this.List {
		n.Notify(s)
	}
}

func (this *Object1) Attach(ob Observer) {
	this.List = append(this.List, ob)
}

func (this *Observer1) Notify(s string) {
	fmt.Printf("ob1 %s\n", s)
}

type Observer2 struct{}

func (this *Observer2) Notify(s string) {
	fmt.Printf("ob2 %s\n", s)
}
