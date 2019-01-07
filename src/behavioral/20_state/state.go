package state

import (
	"fmt"
)

type State interface {
	GetState() string
}

type StartState struct {
}

func (t *StartState)GetState() string {
	fmt.Println("do GetState StartState ")
	return "Start State"
} 

type StopState  struct {
}

func (p *StopState)GetState() string {
	fmt.Println("do GetState StopState ") 
	return "Stop State"
} 

type Context struct {
	state State
}

func (s *Context)GetState() State{
	return s.state
}

func (s *Context)setState(state State) {
	s.state = state
}

func (s *Context)DoThings() {
	fmt.Println("do DoThings") 
	stopState := StopState{}
	s.setState(&stopState)
	fmt.Println("Finish DoThings and to StopState ")
}

func main() {
}
