package DesginPattern

import (
	"testing"
)

func TestLazyModel(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()
	t.Logf("%p:%p", s1, s2)

}
func TestHungryModel(t *testing.T) {
	s1 := GetInstanceHungry()
	s2 := GetInstanceHungry()
	t.Logf("%p:%p", s1, s2)
}
