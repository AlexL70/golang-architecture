package main

import "log"

type MyError struct {
	err string
}

func NewMyError(s string) error {
	return MyError{err: s}
}

func (e MyError) Error() string {
	return e.err
}

func main() {
	newError := NewMyError("This is new error created for demo purposes only.")
	log.Printf("New Error: %v\n", newError)
}
