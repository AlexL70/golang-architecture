package main

import (
	"errors"
	"fmt"
)

type ErrFile struct {
	FileName string
	Base     error
}

func (e ErrFile) Error() string {
	return fmt.Sprintf("File %s: %v.", e.FileName, e.Base)
}

func (e ErrFile) Unwrap() error {
	return e.Base
}

var ErrNotExist = errors.New("File does not exist")

func OpenFile(fileName string) (string, error) {
	return "", ErrNotExist
}

func OpenFile2(fileName string) (string, error) {
	return "", ErrFile{
		FileName: fileName,
		Base:     ErrNotExist,
	}
}

func main() {
	//	Easy way. Wrapping error using Errorf function with %w format option
	_, err := OpenFile("mytextfile.txt")
	if err != nil {
		wrappedError := fmt.Errorf("Unable to open %v: %w", "mytextfile.txt", err)
		//	fmt.Errors keeps old error so comparison shows it is the same error as it was before wrapping
		if errors.Is(wrappedError, ErrNotExist) {
			fmt.Printf("This is still \"%e\" after wrapping.\n", ErrNotExist)
		}
		fmt.Println(wrappedError)
	}

	//	More complicated way. Wrapping error into custom error type that implements Unwrap function
	//	so that Is function could still compare with original error
	_, err = OpenFile2("MySecondTextFile.txt")
	if err != nil {
		if errors.Is(err, ErrNotExist) {
			fmt.Printf("This is still \"%e\" after wrapping.\n", ErrNotExist)
		}
		fmt.Println(err)
	}
}
