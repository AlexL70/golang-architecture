package main

import (
	"fmt"
	"log"

	"github.com/AlexL70/golang-architecture/fileWriter"
)

func main() {
	wf := fileWriter.NewWriteFile("File.txt")
	wf.WriteString("Hello World!\n")
	wf.WriteString("Second line.\n")
	wf.WriteString("Third line.\n")
	wf.Close()
	if wf.Err() != nil {
		log.Println(fmt.Errorf("Error: %w\n", wf.Err()))
	}

}
