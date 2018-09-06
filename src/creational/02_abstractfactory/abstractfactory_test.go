package abstractfactory

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	factory := FactoryProduce{}

	shapeFactory := factory.ProduceFactory("shape")
	if shapeFactory != nil {
		rectangle := shapeFactory.ProduceShape("rectangle")
		if rectangle != nil {
			rectangle.Draw()
		}

		square := shapeFactory.ProduceShape("square")
		if square != nil {
			square.Draw()
		}

		circle := shapeFactory.ProduceShape("circle")
		if circle != nil {
			circle.Draw()
		}

		errshape := shapeFactory.ProduceShape("cricle")
		if errshape == nil {
			fmt.Println("errshape")
		}
	}

	colorFactory := factory.ProduceFactory("color")
	if colorFactory != nil {
		red := colorFactory.ProduceColor("red")
		if red != nil {
			red.Fill()
		}
		green := colorFactory.ProduceColor("green")
		if green != nil {
			green.Fill()
		}
		blue := colorFactory.ProduceColor("blue")
		if blue != nil {
			blue.Fill()
		}
	}

}
