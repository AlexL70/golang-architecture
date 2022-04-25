package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	err := copyFile("kabirCopy.txt", "kabir.txt")
	var pe *os.PathError
	if errors.Is(err, os.ErrNotExist) && errors.As(err, &pe) {
		fmt.Printf("Error occured: %v\nOperation: %s\nPath: %s\n", err, pe.Op, pe.Path)
		return
	} else if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File has been copied successfully")
}

func copyFile(dst, src string) error {
	srcF, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Could not open file \"%s\": %w", src, err)
	}
	defer srcF.Close()

	dstF, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Could not create file \"%s\": %w", dst, err)
	}
	defer dstF.Close()

	nBytes, err := io.Copy(dstF, srcF)
	if err != nil {
		return fmt.Errorf("Could not copy \"%s\" to \"%s\": %w", src, dst, err)
	}
	fmt.Printf("Just in development nice to see. Bytes written: %d\n", nBytes)

	return nil
}
