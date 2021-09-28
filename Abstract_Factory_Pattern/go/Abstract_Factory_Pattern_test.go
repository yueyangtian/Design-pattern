package Abstract_Factory_Pattern

import (
	"testing"
)

func TestAbstract(t *testing.T) {
	var f1 Factory
	f1 = new(BlackRectangleFactory)
	s1 := f1.CreateShape()
	c1 := f1.CreateColor()
	t.Log(s1.Draw())
	t.Log(c1.Tint())
}
