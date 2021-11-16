package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("Hello, ", p.first)
}

type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	fmt.Printf("%T\n", p1)

	//	p1 is of type "person", but it is also of type "human" (concrete)
	//	because whatever has "speak()" method is of type "human" (abstract)
	var x human = p1
	fmt.Printf("%T\n", x)
}
