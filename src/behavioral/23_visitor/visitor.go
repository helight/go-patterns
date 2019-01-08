package visitor

import (
	"fmt"
)

type ComputerPart  interface {
	Accept(computerPartVisitor ComputerPartVisitor)
}

type ComputerPartVisitor interface {
	Visitor(cpart ComputerPart)
}

type Keyboard struct {
}
func (k *Keyboard)Accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.Visitor(k);
} 

type Monitor struct {
}
func (m *Monitor)Accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.Visitor(m);
} 

type Mouse struct {
}
func (mo *Mouse)Accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.Visitor(mo);
} 

type Computer struct {
	parts []ComputerPart
}
func (c *Computer)Add(part ComputerPart) {
	c.parts = append(c.parts, part)
	fmt.Println("add ComputerPart ", part)  
}

func (c *Computer)Accept(computerPartVisitor ComputerPartVisitor) {
	for _, part := range c.parts {
		part.Accept(computerPartVisitor)
	}
	computerPartVisitor.Visitor(c);
} 

type ComputerPartDisplayVisitor struct {
}

func (com *ComputerPartDisplayVisitor)Visitor(cpart ComputerPart) {
	switch c := cpart.(type) {
	case *Keyboard:
		fmt.Println("Displaying Keyboard.", c)
	case *Monitor:
		fmt.Println("Displaying Monitor.", c)
	case *Mouse:
		fmt.Println("Displaying Mouse.", c)
	case *Computer:
		fmt.Println("Displaying Computer.", c)
	}	
}

func main() {
}
