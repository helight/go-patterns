package flyweight

import (
	"math/rand"
	"testing"
)

func TestFlyweight(t *testing.T) {
	colors := []string{"red", "green", "blue", "white", "black"}

	shapemaker := ShapeMaker{make(map[string]Circle)}
	circle := shapemaker.GetCircle(colors[2])
	circle.Setx(rand.Intn(100))
	circle.Sety(rand.Intn(100))
	circle.SetRadius(rand.Intn(100))
	circle.Draw()
}

func init() {
}
