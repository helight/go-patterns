package decorator

import (
//	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {

	circle := Circle{}
	redcircle := RedShapeDecorator{&circle}
	redcircle.Draw()

	rectangle := Rectangle{}
	redrectangle := RedShapeDecorator{&rectangle}
	redrectangle.Draw()
}

func init() {
}
