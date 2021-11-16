package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("Hello,", p.first)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent. My name is", sa.first)
}

type human interface {
	speak()
}

func foo(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Miss Moneypenny",
	}

	sa1 := secretAgent{
		person: person{first: "James"},
		ltk:    true,
	}

	//	p1 is of type "person", but it is also of type "human" (concrete)
	//	because whatever has "speak()" method is of type "human" (abstract)
	var x, y human
	x = p1
	y = sa1
	x.speak()
	fmt.Printf("%T\n", x)
	y.speak()
	fmt.Printf("%T\n", y)
	fmt.Println("----------")
	foo(x)
	foo(y)
	foo(p1)
	foo(sa1)
}
