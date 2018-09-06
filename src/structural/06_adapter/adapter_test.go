package adapter

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)

	ret := target.Request()

	fmt.Println("ret: ", ret)

	target2 := NewAdapter(new(adapteeImpl2))

	ret2 := target2.Request()
	fmt.Println("ret: ", ret2)
}
