package visitor

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	fmt.Println("do test") 
	key := Keyboard{}
	monitor := Monitor{}
	mouse := Mouse{}
	computer := Computer{}
	computer.Add(&key)
	computer.Add(&monitor)
	computer.Add(&mouse)

	comvistor := ComputerPartDisplayVisitor{}
	computer.Accept(&comvistor);
}

func init() {
}
