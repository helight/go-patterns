package filter

import "fmt"

type Person struct {
	name string
	gender string
	maritalStatus string
}
func (p *Person) GetName() string {
	return p.name
}
func (p *Person) GetGender() string {
	return p.gender
}
func (p *Person) GetMaritalStatus() string {
	return p.maritalStatus
}

type filter interface {
	DoFilter(persons map[int]Person) map[int]Person
}

type FilterMale struct {
	malePersons map[int]Person
}

func (f *FilterMale) DoFilter(persons map[int]Person) {
	f.malePersons = make(map[int]Person)
	i := 0
	for _, person := range persons {
		if (person.GetGender() == "Male") {
			f.malePersons[i] = person
			i++
		}
	}
	fmt.Println("Male person num: ", i)
}

type FilterFemale struct {
	malePersons map[int]Person
}

func (f *FilterFemale) DoFilter(persons map[int]Person) {
	f.malePersons = make(map[int]Person)
	i := 0
	for _, person := range persons {
		if (person.GetGender() == "Female") {
			f.malePersons[i] = person
			i++
		}
	}
	fmt.Println("Female person num: ", i)
}

type FilterSingle struct {
	malePersons map[int]Person
}

func (f *FilterSingle) DoFilter(persons map[int]Person) {
	f.malePersons = make(map[int]Person)
	i := 0
	for _, person := range persons {
		if (person.GetMaritalStatus() == "Single") {
			f.malePersons[i] = person
			i++
		}
	}
	fmt.Println("Single person num: ", i)
}

func main() {

}
