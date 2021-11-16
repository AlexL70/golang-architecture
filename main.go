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

func main() {
	p1 := person{
		first: "Miss Moneypenny",
	}

	sa1 := secretAgent{
		person: person{first: "James"},
		ltk:    true,
	}

	fmt.Printf("%T\n", p1)

	//	p1 is of type "person", but it is also of type "human" (concrete)
	//	because whatever has "speak()" method is of type "human" (abstract)
	var x human = p1
	x.speak()
	fmt.Printf("%T\n", x)
	x = sa1
	x.speak()
	fmt.Printf("%T\n", x)

}
