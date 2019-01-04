package interpreter

import (
	"fmt"
	"testing"
)

func getMaleExpression() Expression {
	robert := NewTExpression("Robert")
	john := NewTExpression("John")
	return NewOrExpression(robert, john)
}

func getMarriedWomanExpression() Expression {
	julie := NewTExpression("Julie")
	married := NewTExpression("Married")
	return NewAndExpression(julie, married)
}

func TestChain(t *testing.T) {
	isMale := getMaleExpression()
	isMarriedWoman := getMarriedWomanExpression()

	fmt.Println("John is mail ? ", isMale.Interpret("John"))
	fmt.Println("Julie is a married women", isMarriedWoman.Interpret("Married Julie"))
}

func init() {
}
