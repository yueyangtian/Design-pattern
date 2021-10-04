package Command_Pattern

import "testing"

func TestCommandPattern(t *testing.T) {
	l := &Light{}
	tv := &TV{}
	s := Switcher{}
	s.ExecandStore(l)
	s.ExecandStore(tv)
}
