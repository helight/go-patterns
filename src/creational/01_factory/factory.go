package factory

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct {
}
func (r *Rectangle) Draw() {
	fmt.Println("This is rectangle")
}

type Square struct {
}
func (s *Square) Draw() {
	fmt.Println("This is square")
}

type Circle struct {
}
func (c *Circle) Draw() {
	fmt.Println("This is circle")
}


// factory
type ShapeFactory struct {
}

func (s *ShapeFactory) Produce(itype string) Shape {
	switch itype {
		case "rectangle":
			return new(Rectangle)
		case "square":
			return new(Square)
		case "circle":
			return new(Circle)
		default:
			return nil
	}
}

func main() {
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
