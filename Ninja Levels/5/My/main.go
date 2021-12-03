package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	source, err := os.Open("Example.txt")
	if err != nil {
		log.Fatal(err)
	}
	target, crErr := os.Create("Copy.txt")
	if crErr != nil {
		log.Fatal(err)
	}
	copied, cpErr := io.Copy(target, source)
	if cpErr != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Copied %d bytes.", copied)
	}
	closeErr := target.Close()
	if closeErr != nil {
		log.Fatal(closeErr)
	}
}
