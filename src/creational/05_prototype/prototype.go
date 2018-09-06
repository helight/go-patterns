package prototype

import "fmt"

type Cloneable interface {
	Clone() Cloneable
	Hello()
}

type PrototypeManager1 struct {
	prototypes map[string]Cloneable
}

func (p *PrototypeManager1) Get(name string) Cloneable{
	return p.prototypes[name]
}

func (p *PrototypeManager1) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype

}

func NewPrototypeManager1() *PrototypeManager1 {
	return &PrototypeManager1{
		prototypes: make(map[string]Cloneable),
	}
}

type Rectangle struct {
	name string
}

func (r *Rectangle) Draw() {
	fmt.Println("This is rectangle")
}

func (r *Rectangle) Hello() {
	fmt.Println("This is rectangle")
}

func (r *Rectangle) Clone() Cloneable{
	t := *r
	return &t
}

type Square struct {
	name string
}

func (s *Square) Draw() {
	fmt.Println("This is square")
}

func (s *Square) Hello() {
	fmt.Println("This is square")
}

func (s *Square) Clone() Cloneable{
	t := *s
	return &t
}

type Circle struct {
	name string
}

func (c *Circle) Draw() {
	fmt.Println("This is circle")
}

func (c *Circle) Hello() {
	fmt.Println("This is circle")
}

func (c *Circle) Clone() Cloneable{
	t := *c
	return &t
}

func main() {
        pmx := NewPrototypeManager1()

	circle := new(Circle)
	pmx.Set("circle", circle)

// 	rectangle := new(Rectangle)
// 	pmx.Set("rectangle", rectangle)
// 
// 	square := new(Square)
// 	pmx.Set("square", square)
// 
	c1 := pmx.Get("circle")
	c1.Hello()
// 
// 	rectangle1 := Shape(pm.Get("rectangle"))
// 	rectangle1.Draw()
// 
// 	square1 := Shape(pm.Get("square"))
// 	square1.Draw()
}
