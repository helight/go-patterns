package strategy

import (
	"fmt"
)

type Strategy interface {
	DoOperation(num1, num2 int) int 
}

type OperationAdd  struct {
}

func (t *OperationAdd)DoOperation(num1, num2 int) int {
	fmt.Println("do OperationAdd")
	return num1 + num2
} 

type OperationSubstract struct {
}

func (t *OperationSubstract)DoOperation(num1, num2 int) int {
	fmt.Println("do OperationSubstract")
	return num1 - num2
} 

type OperationMultiply struct {
}

func (t *OperationMultiply)DoOperation(num1, num2 int) int {
	fmt.Println("do OperationMultiply")
	return num1 * num2
} 
type Context struct {
	strategy Strategy 
}

func (s *Context)ExecuteStrategy(num1, num2 int) int{
	return s.strategy.DoOperation(num1, num2)
}

func (s *Context)SetStrategy(strategy Strategy) {
	s.strategy = strategy
}

func main() {
}
