package Abstract_Factory_Pattern

type Shape interface {
	Draw() string
}

type Circle struct{}
type Square struct{}
type Rectangle struct{}

func (*Circle) Draw() string {
	return "Cricle"
}

func (*Square) Draw() string {
	return "Square"
}

func (*Rectangle) Draw() string {
	return "Rectangle"
}

type Color interface {
	Tint() string
}
type Red struct{}
type Black struct{}

func (*Red) Tint() string {
	return "Red"
}
func (*Black) Tint() string {
	return "Black"
}

type Factory interface {
	CreateShape() Shape
	CreateColor() Color
}

type RedCircleFactory struct{}
type BlackRectangleFactory struct{}

func (*RedCircleFactory) CreateShape() Shape {
	return &Circle{}
}
func (*RedCircleFactory) CreateColor() Color {
	return &Red{}
}

func (*BlackRectangleFactory) CreateShape() Shape {
	return &Rectangle{}
}
func (*BlackRectangleFactory) CreateColor() Color {
	return &Black{}
}
