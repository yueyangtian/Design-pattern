package Mediator_Pattern

import "testing"

func TestMediatorPattern(t *testing.T) {
	u1 := User{Name: "Lili"}
	u2 := User{Name: "Tom"}
	u1.SendMessage("hello")
	u2.SendMessage("world")
}
