package composite

import (
	"testing"
)

func TestMain(t *testing.T) {

	ceo := Person{"Robert0", "Male", "Single", make([]Component, 0)}
	head1 := Person{"Robert1", "Male", "Married", make([]Component, 0)}
	head2 := Person{"Robert2", "Female", "Married", make([]Component, 0)}
	sales1 := Person{"Robert3", "Female", "Single", make([]Component, 0)}
	sales2 := Person{"Robert4", "Female", "Single", make([]Component, 0)}

	ceo.Add(&head1)
	ceo.Add(&head2)
	head1.Add(&sales2)
	head2.Add(&sales1)

	ceo.Print()
}
