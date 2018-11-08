package composite

import "fmt"

type Component interface {
	Add(Component)
	Print()
	GetName() string
	GetGender() string
}

type Person struct {
	name string
	gender string
	maritalStatus string
	persons []Component
}

func (p *Person) Add(component Component) {
	p.persons = append(p.persons, component)
}

func (p *Person) Print() {
	for _, person := range p.persons {
		person.Print()
	}
	fmt.Printf("Name: %s", p.GetName())
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetGender() string {
	return p.gender
}

func main() {

}
