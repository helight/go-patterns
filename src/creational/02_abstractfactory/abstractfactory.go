package abstractfactory

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

type Color interface {
	Fill()
}

type Red struct {
}

func (r *Red) Fill() {
	fmt.Println("This is red")
}

type Green struct {
}

func (g *Green) Fill() {
	fmt.Println("This is green")
}

type Blue struct {
}

func (b *Blue) Fill() {
	fmt.Println("This is blue")
}

// Abstractfactory
type AbstractFactory interface {
	ProduceShape(itype string) Shape
	ProduceColor(itype string) Color
}

// factory
type ShapeFactory struct {
}

func (s *ShapeFactory) ProduceShape(itype string) Shape {
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
func (s *ShapeFactory) ProduceColor(itype string) Color {
	return nil
}

// factory
type ColorFactory struct {
}

func (s *ColorFactory) ProduceColor(itype string) Color {
	switch itype {
	case "red":
		return new(Red)
	case "green":
		return new(Green)
	case "blue":
		return new(Blue)
	default:
		return nil
	}
}

func (s *ColorFactory) ProduceShape(itype string) Shape {
	return nil
}

type FactoryProduce struct {
}

func (f *FactoryProduce) ProduceFactory(itype string) AbstractFactory{
	switch itype {
	case "shape":
		return new(ShapeFactory)
	case "color":
		return new(ColorFactory)
	default:
		return nil
	}

}
func main() {
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
