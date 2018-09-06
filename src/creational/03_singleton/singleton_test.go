package singleton

import (
	"testing"
)


func TestSingleton(t *testing.T) {
	singlet := GetSingleton()
	singlet.DoSomething()
}
