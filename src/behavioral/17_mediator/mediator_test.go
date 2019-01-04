package interpreter

import (
	"testing"
)

func TestChain(t *testing.T) {
	robert := User{"Robert"};
	john := User{"John"};

	robert.SendMessage("Hi! John!");
	john.SendMessage("Hello! Robert!");
}

func init() {
}
