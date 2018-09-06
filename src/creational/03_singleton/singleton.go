package singleton

import "fmt"


type Singleton struct {
}

func (s *Singleton) DoSomething() {
	fmt.Println("Tish is a Singleton")
}

var singleton *Singleton

//
func GetSingleton() *Singleton {
	if singleton == nil {
		singleton = &Singleton{}
	}
	return singleton

}

func main() {
	singlet := GetSingleton()
	singlet.DoSomething()
}
