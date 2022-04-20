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
	var perr *os.PathError
	if errors.As(err, &perr) {
		switch {
		case errors.Is(perr, os.ErrNotExist):
			err = fmt.Errorf("File \"%s\" does not exists: %w", fName, err)
		case errors.Is(perr, os.ErrPermission):
			err = fmt.Errorf("You do not have a permission to open \"%s\" file: %w", fName, err)
		default:
			err = fmt.Errorf("File \"%s\" cannot be opened: %w", fName, err)
		}
		log.Printf("%s\nOperation: %s\nPath: %s\nIs a timeout: %t\n", err, perr.Op, perr.Path, perr.Timeout())
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
