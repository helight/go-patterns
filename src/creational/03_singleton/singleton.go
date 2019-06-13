package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

func (s *Singleton) DoSomething() {
	fmt.Println("Tish is a Singleton")
}

var singleton *Singleton
var singlock sync.Mutex

//
func GetSingleton() *Singleton {
	singlock.Lock()
	if singleton == nil {
		singleton = &Singleton{}
	}
	singlock.Unlock()
	return singleton

}

func main() {
	singlet := GetSingleton()
	singlet.DoSomething()
}
