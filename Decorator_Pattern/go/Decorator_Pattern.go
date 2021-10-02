package Decorator_Pattern

import (
	"fmt"
)

type Shape interface {
	draw()
}

type Cricle struct{}

func (*Cricle) draw() {
	fmt.Println("cricle draw")
}

type CricleDecorator interface {
	draw()
}

type RedCricle struct {
	shape Shape
}

func (s *RedCricle) draw() {
	fmt.Print("draw in decorator ")
	s.shape.draw()
}
