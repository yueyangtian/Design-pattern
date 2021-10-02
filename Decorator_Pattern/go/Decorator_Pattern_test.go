package Decorator_Pattern

import "testing"

func TestDecorator(t *testing.T) {
	var c Shape
	c = new(Cricle)

	c.draw()

	cd := RedCricle{c}

	cd.draw()
}
