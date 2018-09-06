package singleton

import (
	"testing"
)


func TestMain(t *testing.T) {
	singlet := GetSingleton()
	singlet.DoSomething()
}
