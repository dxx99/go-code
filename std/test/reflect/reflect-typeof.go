package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Say() string
	Eat()
}

type Dog struct {

}

func (d *Dog) Say() string {
	fmt.Println("dog wanging")
	return "I am is a dog"
}

func (d *Dog) Eat() {
	fmt.Println("dog eating")
}

func Run(a Animal) string {

	fmt.Println("animal dynamic type is ", reflect.TypeOf(a))
	s := a.Say()
	a.Eat()
	return s
}

func main() {
	d := new(Dog)
	fmt.Println(Run(d))
	// fmt.Println(Run(Dog{}))	//cannot use Dog{} (type Dog) as type Animal in argument to Run:
							// Dog does not implement Animal (Eat method has pointer receiver)
}
