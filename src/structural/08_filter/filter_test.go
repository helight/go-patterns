package filter

import (
	"testing"
)

func TestMain(t *testing.T) {
	persons := make(map[int]Person)
	persons[0] = Person{"Robert0", "Male", "Single"}
	persons[1] = Person{"Robert1", "Male", "Married"}
	persons[2] = Person{"Robert2", "Female", "Married"}
	persons[3] = Person{"Robert3", "Female", "Single"}
	persons[4] = Person{"Robert4", "Female", "Single"}

	malefilter := new(FilterMale)
	femalefilter := new(FilterFemale)
	singlefilter := new(FilterSingle)
	malefilter.DoFilter(persons)
	femalefilter.DoFilter(persons)
	singlefilter.DoFilter(persons)
}
