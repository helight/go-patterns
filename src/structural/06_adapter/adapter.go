package adapter

import "fmt"

type Target interface {
	Request() string
}


type Adaptee interface {
	SpecificRequest() string
}

type adapteeImpl1 struct {
}

func (r *adapteeImpl1) SpecificRequest() string {
	fmt.Println("This is adaptee method1")
	return "adaptee method1"
}

type adapteeImpl2 struct {
}

func (r *adapteeImpl2) SpecificRequest() string {
	fmt.Println("This is adaptee method2")
	return "adaptee method2"
}

func NewAdaptee() Adaptee {
	return &adapteeImpl1{}
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}

func main() {

}
