package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	mealbuilder := new(MealBuilder)

	vegmeal := mealbuilder.PrepareVegMeal()
	vegmeal.ShowItems()
	fmt.Printf("Price: %f \r\n", vegmeal.GetPrice())

	nonvegmeal := mealbuilder.PrepareNonVegMeal()
	nonvegmeal.ShowItems()
	fmt.Printf("Price: %f \r\n", nonvegmeal.GetPrice())
}
