package main

import "fmt"

type Sayer interface {
	Say(name string) string
}

type Dog struct {
}

func (d Dog) Say(name string) string {
	return fmt.Sprintf("dog say %s\n", name)
}

func main() {
	d := new(Dog)
	var _ Sayer = Dog{}				// verify that Dog implements I.
	var _ Sayer = d					// verify that *Dog implements I.
	var _ Sayer = (*Dog)(nil)


}
