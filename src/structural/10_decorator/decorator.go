package decorator

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

type RedShapeDecorator struct {
	decoratedshape Shape
}

func (d *RedShapeDecorator) Draw() {
	d.decoratedshape.Draw()
	d.SetRedBorder(d.decoratedshape)
}

func (d *RedShapeDecorator) SetRedBorder(shape Shape){
	fmt.Println("Border Color: Red")
}

func main() {
}
