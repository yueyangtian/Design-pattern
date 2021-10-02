package Flywight_Pattern

import "testing"

func TestFlywight(t *testing.T) {
	m := make(map[string]*ColorCrile)
	flyweight := Flywight{Map: m}

	flyweight.getCricle("red").draw()
	flyweight.getCricle("yellow").draw()

	flyweight.getCricle("red").draw()
}
