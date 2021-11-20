package main

import "fmt"

type hotdog int

func (h hotdog) cook() {
	fmt.Println("I am cooking hotdog")
}

type hotFood interface {
	cook()
}

func main() {
	var food hotFood = hotdog(4)
	fmt.Printf("%T\n", food)
	food.cook()
}
