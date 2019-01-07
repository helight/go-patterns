package memento

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	originator := Originator{}
	careTaker := CareTaker{}

	originator.SetName("State #1");
	originator.SetName("State #2");
	careTaker.Add(originator.SaveNameToMemento());
	originator.SetName("State #3");
	careTaker.Add(originator.SaveNameToMemento());
	originator.SetName("State #4");

	fmt.Println("Current State: ", originator.GetName());    
	originator.GetNameToMemento(careTaker.Get(0));
	fmt.Println("First saved State: ", originator.GetName());
	originator.GetNameToMemento(careTaker.Get(1));
	fmt.Println("Second saved State: ", originator.GetName());
}

func init() {
}
