package observer

import (
	"fmt"
)

type Subject struct {
	observers []Observer
	state string
}

func (s *Subject)GetState() string{
	return s.state
}

func (s *Subject)SetState(state string) {
	s.state = state
}

func (s *Subject)Attach(o Observer){
	s.observers = append(s.observers, o)
	fmt.Println("add Observer ", o)  
}

func (s *Subject)NotifyAllObservers(event string) {
	for _, observer := range s.observers {
		observer.Update(event)
	}
}  

type Observer interface {
	Update(event string)
}

type X1Observer struct {
}

func (x *X1Observer)Update(event string) {
	fmt.Println("do Observer1 update ", event) 
} 


type X2Observer struct {
}

func (x *X2Observer)Update(event string) {
	fmt.Println("do Observer2 update ", event) 
} 

func main() {
}
