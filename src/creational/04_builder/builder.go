package builder

import "fmt"

type Item interface {
	Name() string
	Price() float64
}

type Packing interface {
	Pack() string
}

type VegBurger struct {
}

func (r *VegBurger) Name() string{
	return "vegburger"
}

func (r *VegBurger) Price() float64{
	return 1.5
}

type ChickenBurger struct {
}

func (r *ChickenBurger) Name() string{
	return "chickenburger"
}

func (r *ChickenBurger) Price() float64{
	return 5.5
}

type Coke struct {
}

func (r *Coke) Name() string{
	return "coke"
}

func (r *Coke) Price() float64{
	return 2.5
}

type Pepsi struct {
}

func (r *Pepsi) Name() string{
	return "pepsi"
}

func (r *Pepsi) Price() float64{
	return 3.5
}

type Imeal interface {
	AddItem(item Item)
	GetPrice() float64
	ShowItems()
}
type Meal struct {
	meal map[int]Item
	count int
}

func (m *Meal)Init() {
	m.count = 0
	m.meal = make(map[int]Item)
}

func (m *Meal)AddItem(item Item) {
	m.count++
	m.meal[m.count] = item
}

func (m *Meal)GetPrice() float64 {
	price := 0.0
	for _, item := range m.meal {
		price += item.Price()
	}
	return price
}

func (m *Meal)ShowItems() {
	for _, item := range m.meal {
		fmt.Printf("name: %s Pricce: %f \r\n", item.Name(), item.Price())
	}
}

type MealBuilder struct {
}

func (m *MealBuilder) PrepareVegMeal() Imeal {
	meal := new(Meal)
	meal.Init()
	meal.AddItem(new(VegBurger))
	meal.AddItem(new(Coke))
	return meal
}

func (m *MealBuilder) PrepareNonVegMeal() Imeal {
	meal := new(Meal)
	meal.Init()
	meal.AddItem(new(ChickenBurger))
	meal.AddItem(new(Pepsi))
	return meal
}

func main() {
	mealbuilder := new(MealBuilder)

	vegmeal := Imeal(mealbuilder.PrepareVegMeal())
	vegmeal.ShowItems()

	nonvegmeal := mealbuilder.PrepareNonVegMeal()
	nonvegmeal.ShowItems()
}
