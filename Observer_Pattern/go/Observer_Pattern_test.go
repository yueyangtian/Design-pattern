package Observer_Pattern

import "testing"

func TestBoserverPattern(t *testing.T) {
	obj := &Object1{}

	ob1 := &Observer1{}
	ob2 := &Observer2{}
	ob3 := &Observer2{}

	obj.Attach(ob1)
	obj.Attach(ob2)
	obj.Attach(ob3)

	obj.SetStat(1)
}
