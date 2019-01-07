package state

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	fmt.Println("do test") 
	startState := StartState{} 

	context := Context{&startState}
	fmt.Println(context.GetState().GetState())
	
	context.DoThings();
	fmt.Println(context.GetState().GetState())
}

func init() {
}
