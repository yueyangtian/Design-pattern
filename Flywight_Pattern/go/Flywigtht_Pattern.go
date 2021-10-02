package Flywight_Pattern

import "fmt"

type Shape interface {
	draw()
}

type ColorCrile struct {
	color string
}

func (this *ColorCrile) draw() {
	fmt.Printf("%s cricle draw\n", this.color)
}

type ColorCricleFactory struct{}

func (*ColorCricleFactory) createColorCricle(c string) *ColorCrile {
	return &ColorCrile{c}
}

type Flywight struct {
	Map     map[string]*ColorCrile
	Factory ColorCricleFactory
}

func (this *Flywight) getCricle(color string) *ColorCrile {
	if c, ok := this.Map[color]; ok {
		fmt.Printf("flyweight color %s\n", color)
		return c
	}
	this.Map[color] = this.Factory.createColorCricle(color)
	fmt.Printf("new color %s\n", color)
	return this.Map[color]
}
