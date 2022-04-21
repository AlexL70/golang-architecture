package main

import (
	"fmt"
	"log"
)

type errorString string

func (es errorString) Error() string {
	return fmt.Sprintf("Error string error: %s", string(es))
}

func sum(i, j int) (int, error) {
	n := i + j
	if n != i+j {
		var sErr errorString = "the numbers did not add up"
		return 0, sErr
	}
	return n, nil
}

func main() {
	s, err := sum(12, 25)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("12 + 25 = %d\n", s)
}
