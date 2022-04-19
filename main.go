package main

import (
	"errors"
	"fmt"
)

var ErrNotExist = fmt.Errorf("File does not exist: %s", "myfile.txt")
var ErrUserNotExist = errors.New("User does not exist")

func main() {
	err := ErrUserNotExist

	if err == ErrUserNotExist {
		fmt.Println("You need to register first!")
	} else {
		fmt.Println("Unknown error")
	}
}
