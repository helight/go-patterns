package factory

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	shapeFactory := ShapeFactory{}

	rectangle := shapeFactory.Produce("rectangle")
	if rectangle != nil {
		rectangle.Draw()
	}

	square := shapeFactory.Produce("square")
	if square != nil {
		square.Draw()
	}

	circle := shapeFactory.Produce("circle")
	if circle != nil {
		circle.Draw()
	}

	errshape := shapeFactory.Produce("cricle")
	if errshape == nil {
		fmt.Println("errshape")
	}
}

func init() {
}
