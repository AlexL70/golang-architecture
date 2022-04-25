package main

import (
	"errors"
	"fmt"
)

func cat() error {
	return errors.New("Error: cat calls nothing!")
}

func moo() error {
	err := cat()
	return fmt.Errorf("Error: moo calls cat - %w", err)
}

func bar() error {
	err := moo()
	return fmt.Errorf("Error: bar calls moo - %w", err)
}

func foo() error {
	err := bar()
	return fmt.Errorf("Error: foo calls bar - %w", err)
}

func main() {
	err := foo()
	if err != nil {
		err = fmt.Errorf("Error: main calls foo - %w", err)
	}
	fmt.Println(err)
	for err != nil {
		err = errors.Unwrap(err)
		fmt.Println(err)
	}
}
