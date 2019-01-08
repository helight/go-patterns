package strategy

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	fmt.Println("do test") 
	player := Player{} 

	cricket := Cricket{}
	player.Play(&cricket)

	football := Football{}
	player.Play(&football)
}

func init() {
}
