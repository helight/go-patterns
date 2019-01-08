package template

import (
	"fmt"
)

type Game interface {
	initialize() string
	startPlay() string
	endPlay() string
}

type Cricket struct {
}
func (t *Cricket)initialize() string{
	fmt.Println("do initialize")
	return "initialize Cricket"
} 
func (t *Cricket)startPlay() string{
	fmt.Println("do startPlay")
	return "startPlay Cricket"
} 
func (t *Cricket)endPlay() string{
	fmt.Println("do endPlay")
	return "endPlay Cricket"
} 

type Football struct {
}
func (f *Football)initialize() string{
	fmt.Println("do initialize")
	return "initialize Football"
} 
func (f *Football)startPlay() string{
	fmt.Println("do startPlay")
	return "startPlay Football"
} 
func (f *Football)endPlay() string{
	fmt.Println("do endPlay")
	return "endPlay Football"
} 

type Player struct {	
}

func (p *Player)Play(game Game){
	fmt.Println("do Play")
	fmt.Println("do Play： ", game.initialize(), game.startPlay(), game.endPlay())
} 

func main() {
}
