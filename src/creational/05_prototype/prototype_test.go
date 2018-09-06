package prototype

import (
	"fmt"
	"testing"
)

var pm1 *PrototypeManager1
func TestPrototype(t *testing.T){
	c := pm1.Get("t1").Clone()

	t1 := c.(*Circle)

	if t1.name != "circle1" {
		t.Fatal("error")
	}
	t1.Draw()
	t.Log(t1.name)
}

func init() {
	fmt.Println("init")
	pm1 = NewPrototypeManager1()

	t1 := &Circle{
		name: "circle1",
	}
	pm1.Set("t1", t1)

}
