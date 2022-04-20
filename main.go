package main

import (
	"errors"
	"fmt"
)

func cat() error {
	return errors.New("Cat is an error!")
}

func moo() error {
	return fmt.Errorf("Moo is an error: %w", cat())
}

func bar() error {
	return fmt.Errorf("Bar is n error: %w", moo())
}

func foo() error {
	return fmt.Errorf("Foo is an eror: %w", bar())
}

func main() {
	err := foo()
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
}
