package facade

import (
//	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {

	circle := Circle{}
	rectangle := Rectangle{}
	square := Square{}

	shapemaker := ShapeMaker{&circle, &rectangle, &square}
	shapemaker.DrawCircle()
	shapemaker.DrawRectangle()
	shapemaker.DrawSquare()
}

func init() {
}
