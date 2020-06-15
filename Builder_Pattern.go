package DesginPattern

import "fmt"

type Car interface {
	Drive() string
	Look() string
}

type CarBuilder interface {
	Speed(int) CarBuilder
	Look(string) CarBuilder
	Build() Car
}

type cb struct {
	speed int
	color string
}

func (c *cb) Speed(sp int) CarBuilder {
	c.speed = sp
	return c
}

func (c *cb) Look(color string) CarBuilder {
	c.color = color
	return c
}

func (c *cb) Build() Car {
	return &car{
		speed: c.speed,
		color: c.color}
}

type car struct {
	speed int
	color string
}

func (c *car) Drive() string {
	return fmt.Sprintf("Speed is %d m/s", c.speed)
}
func (c *car) Look() string {
	return fmt.Sprintf("Color is %s", c.color)
}

func NewCarBuilder() CarBuilder {
	return &cb{}
}
