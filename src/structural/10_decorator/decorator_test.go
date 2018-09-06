package decorator

import (
//	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {

	circle := Circle{}
	redcircle := RedShapeDecorator{&circle}
	redcircle.Draw()

	rectangle := Rectangle{}
	redrectangle := RedShapeDecorator{&rectangle}
	redrectangle.Draw()
}

func init() {
}
