package strategy

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	fmt.Println("do test") 
	add := OperationAdd{} 
	sub := OperationSubstract{}
	mul := OperationMultiply{}

	context := Context{&add}
	fmt.Println("10 + 5 = " , context.ExecuteStrategy(10, 5))
	context.SetStrategy(&sub)
	fmt.Println("10 - 5 = " , context.ExecuteStrategy(10, 5))
	context.SetStrategy(&mul)
	fmt.Println("10 * 5 = " , context.ExecuteStrategy(10, 5))
}

func init() {
}
