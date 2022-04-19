package main

import (
	"errors"
	"fmt"
	"net"
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

func openFile(fileName string) (string, error) {
	return "", ErrFile{
		FileName: fileName,
		Base:     ErrNotExist,
	}
}

func processFile(fileName string) error {
	_, err := openFile(fileName)
	if err != nil {
		return fmt.Errorf("Error while opening file: %w", err)
	}
	//  Do work on stuff
	return nil
}

func main() {
	err := processFile("MySecondTextFile.txt")
	if err != nil {
		var fErr ErrFile
		if errors.As(err, &fErr) {
			fmt.Printf("Was unable to do something with file \"%s\"\n", fErr.FileName)
		}
		fmt.Println(err)
		var netErr net.Error
		if errors.As(err, &netErr) {
			if netErr.Temporary() {
				//  Retry
			}
		}
	}
}
