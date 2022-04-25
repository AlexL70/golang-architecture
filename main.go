package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func main() {
	p1 := Person{FirstName: "Alex", LastName: "Levinson"}
	p2 := Person{FirstName: "Tom", LastName: "Waits"}
	m := map[Person]string{}
	m[p1] = "seems to be an ordinary person"
	m[p2] = "is a celebrity"
	fmt.Println(p1, m[p1])
	fmt.Println(p2, m[p2])
	fmt.Println(p1 == p2)
	p3 := Person{}
	p3.FirstName = p1.FirstName
	p3.LastName = p1.LastName
	fmt.Println(p1 == p3) // do not do it with structures containing pointers; pointers might differ but still refer to the same value
}
