package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	const fName = "rumi.txt"
	f, err := os.Open(fName)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			err = fmt.Errorf("File \"%s\" does not exists: %w", fName, err)
		case errors.Is(err, os.ErrPermission):
			err = fmt.Errorf("You do not have a permission to open \"%s\" file: %w", fName, err)
		default:
			err = fmt.Errorf("File \"%s\" cannot be opened: %w", fName, err)
		}
		log.Println(err)
		return
	}
	defer f.Close()
	data := make([]byte, 1000)
	count, err := f.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
