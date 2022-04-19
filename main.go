package main

import (
	"errors"
	"fmt"
	"time"
)

type ErrFileNotFound struct {
	FileName string
	When     time.Time
}

func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("File %s not found at %v.", e.FileName, e.When)
}

//  In this implementation "Is" compares error type instead of string error returns
func (e ErrFileNotFound) Is(other error) bool {
	_, ok := other.(ErrFileNotFound)
	return ok
}

func main() {
	err := ErrFileNotFound{
		FileName: "mytext.txt",
		When:     time.Now(),
	}

	fmt.Println(err)

	isTheSameError := errors.Is(err, ErrFileNotFound{})
	fmt.Println(isTheSameError)
}
