package Factory_Pattern

import (
	"errors"
	"fmt"
)

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

func Create(shape string) (Shape, error) {
	switch shape {
	case "Cricle":
		return new(Circle), nil
	case "Square":
		return new(Square), nil
	case "Rectangle":
		return new(Square), nil
	default:
		return nil, errors.New(fmt.Sprintf("%s not in Shape", shape))

	}
}
