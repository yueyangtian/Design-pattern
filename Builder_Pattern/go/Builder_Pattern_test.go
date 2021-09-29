package Builder_Pattern

import (
	"testing"
)

func TestBuilderPattrn(t *testing.T) {
	car1 := NewCarBuilder().SetColor("yellow").SetSpeed(200).Build()
	t.Log("car1 look:", car1.Look())
	t.Log("car1 driver:", car1.Drive())
	car2 := NewCarBuilder().SetColor("red").SetSpeed(180).Build()
	t.Log("car2 look:", car2.Look())
	t.Log("car2 driver:", car2.Drive())
}
