package chain

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct {
	color string
	x int
	y int
	radius int
}

func (c *Circle) SetColor(color string) {
	c.color = color
}

func (c *Circle) Setx(x int) {
	c.x = x
}

func (c *Circle) Sety(y int) {
	c.y = y
}

func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

func (c *Circle) Draw() {
	fmt.Println("color: ", c.color)
	fmt.Printf("This is circle color: %s x: %d y: %d radius: %d \r\n", c.color, c.x, c.y, c.radius)
}

type ChainShape struct {
	circle *Circle
}

func (p *ChainShape) Draw() {
	if nil == p.circle {
		p.circle = &Circle{"red", 3, 9, 8}
	}
	p.circle.Draw()
}

func main() {
}
