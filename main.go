package main

import (
	"fmt"
)

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("From a person:", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("This is not my real name:", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	pp1 := &person{
		first: "Jenny",
	}
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	//	here we pass pointer to person
	saySomething(pp1)
	//	here we pass just secretAgend (value)
	saySomething(sa1)
	//	Both work the same way regardless of the fact that receivers are
	//	identical up to type
}
