package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	singlet := GetSingleton()
	singlet.DoSomething()

	singlet2 := GetSingleton2()
	singlet2.DoSomething()
}
