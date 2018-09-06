package flyweight

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

type ShapeMaker struct {
	circles map[string]Circle
}

func (d *ShapeMaker) GetCircle(color string) Circle{
	circle,ok := d.circles[color]
	if ok == false{
		circle := Circle{color, 0, 0, 0}
		circle.SetColor(color)
		d.circles[color] = circle
		fmt.Printf("Creating circle of color: %s \r\n", color)
	}

	return circle
}

func main() {
}
