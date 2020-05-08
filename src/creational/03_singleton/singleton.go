package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var (
	singleton *Singleton
	singlock  sync.Mutex
)

// DoSomething just for test
func (s *Singleton) DoSomething() {
	fmt.Println("this is a Singleton")
}

// Singleton return Singleton
func GetSingleton() *Singleton {
	singlock.Lock()
	if singleton == nil {
		singleton = &Singleton{}
	}
	singlock.Unlock()
	return singleton
}

var (
	singleton2 *Singleton
	done       sync.Once
)

// GetSingleton2 return Singleton
func GetSingleton2() *Singleton {
	done.Do(func() {
		singleton2 = &Singleton{}
	})
	return singleton2
}

func main() {
	singlet := GetSingleton()
	singlet.DoSomething()

	singlet2 := GetSingleton2()
	singlet2.DoSomething()
}
