package facade

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

type ShapeMaker struct {
	circle Shape
	rectangle Shape
	square Shape
}

func (d *ShapeMaker) DrawCircle() {
	d.circle.Draw()
}

func (d *ShapeMaker) DrawRectangle(){
	d.rectangle.Draw()
}

func (d *ShapeMaker) DrawSquare(){
	d.square.Draw()
}

func main() {
}
