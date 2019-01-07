package interpreter

import (
	"fmt"
	"strings"
)

type Expression interface {
	Interpret(context string) bool
}

type TExpression struct {
	data string
}

func (t *TExpression) Interpret(context string) bool {
	fmt.Println("do TExpression Interpret")
	return strings.Contains(t.data, context)
}

func NewTExpression(context string) Expression {
	return &TExpression {
		data: context,
	}
}

type AndExpression struct {
	express1, express2 Expression
}

func (a *AndExpression) Interpret(context string) bool {
	fmt.Println("do AndExpression Interpret")
	return a.express1.Interpret(context) && a.express2.Interpret(context)
}

func NewAndExpression(exp1, exp2 Expression) Expression {
	return &AndExpression {
		express1: exp1,
		express2: exp2,
	}
}

type OrExpression struct {
	express1, express2 Expression
}

func (o *OrExpression) Interpret(context string) bool {
	fmt.Println("do OrExpression Interpret")
	return o.express1.Interpret(context) || o.express2.Interpret(context)
}

func NewOrExpression(exp1, exp2 Expression)  Expression {
	return &OrExpression {
		express1: exp1,
		express2: exp2,
	}
}

func main() {
}
