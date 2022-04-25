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
	fmt.Println(Person{FirstName: "Alex", LastName: "Levinson"})
}
