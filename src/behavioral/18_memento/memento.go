package memento

import (
	"fmt"
)

type Memento struct {
	name string
}

func (m *Memento)GetName() string{
	return m.name
}

type Originator struct{
	name string
}

func (o *Originator)SetName(name string) {
	   o.name = name
}

func (o *Originator)GetName() string {
	   return o.name
}

func (o *Originator)SaveNameToMemento() Memento{
	return Memento{o.name}
}

func (o *Originator)GetNameToMemento(m Memento) {
	o.name = m.GetName()
}

type CareTaker struct {
	mementoList []Memento
}
func (c *CareTaker)Add(m Memento) {
	c.mementoList = append(c.mementoList, m)
	fmt.Println("add Memento ", m.GetName())
}

func (c *CareTaker)Get(index int) Memento {
	return c.mementoList[index]
}

func main() {
}
