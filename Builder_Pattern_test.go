package DesginPattern

import (
	"testing"
)

func TestBuilderPattrn(t *testing.T) {
	builder := NewCarBuilder()
	builder.Look("yellow")
	builder.Speed(200)
	b1 := builder.Build()
	t.Log("b1 look:", b1.Look())
	t.Log("b1 driver:", b1.Drive())
	builder.Look("red")
	builder.Speed(180)
	b2 := builder.Build()
	t.Log("b1 look:", b2.Look())
	t.Log("b1 driver:", b2.Drive())
}
