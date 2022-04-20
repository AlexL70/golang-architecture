package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type writeFile struct {
	f   *os.File
	err error
}

func newWriteFile(fileName string) *writeFile {
	f, err := os.Create(fileName)
	return &writeFile{
		f:   f,
		err: err,
	}
}

func (wf *writeFile) WriteString(s string) {
	if wf.err != nil {
		return
	}
	_, err := io.WriteString(wf.f, s)
	if err != nil {
		wf.err = err
	}
}

func (wf *writeFile) Close() {
	if wf.err != nil {
		return
	}
	err := wf.f.Close()
	if err != nil {
		wf.err = err
	}
}

func (wf *writeFile) Err() error {
	return wf.err
}

func main() {
	wf := newWriteFile("File.txt")
	wf.WriteString("Hello World!\n")
	wf.WriteString("Second line.\n")
	wf.WriteString("Third line.\n")
	wf.Close()
	if wf.Err() != nil {
		log.Println(fmt.Errorf("Error: %w\n", wf.Err()))
	}

}
