package observer

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	fmt.Println("do test")  
	subject := Subject{}
	x1 := X1Observer{}
	x2 := X2Observer{}

	subject.Attach(&x1)
	subject.Attach(&x2)

	subject.NotifyAllObservers("hello news")
	subject.NotifyAllObservers("hello news 2")
}

func init() {
}
