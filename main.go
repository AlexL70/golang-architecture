package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type writeFileError struct {
	Op       string
	FileName string
	Err      error
}

func (e writeFileError) Error() string {
	return fmt.Sprintf("Error in %s method working with file \"%s\": %v", e.Op, e.FileName, e.Err)
}

func (e writeFileError) Unwrap() error {
	return e.Err
}

type writeFile struct {
	f   *os.File
	err *writeFileError
}

func newWriteFile(fileName string) *writeFile {
	f, err := os.Create(fileName)
	var wrapped *writeFileError = nil
	if err != nil {
		wrapped = &(writeFileError{Op: "newWriteFile", FileName: fileName, Err: err})
		return &writeFile{f: nil, err: wrapped}
	}
	return &writeFile{f: f, err: nil}
}

func (wf *writeFile) WriteString(s string) {
	if wf.err != nil {
		return
	}
	_, err := io.WriteString(wf.f, s)
	if err != nil {
		wf.err = &(writeFileError{Op: "WriteString", FileName: wf.f.Name(), Err: err})
	}
}

func (wf *writeFile) Close() {
	if wf.f == nil {
		return
	}
	err := wf.f.Close()
	if err != nil {
		wf.err = &(writeFileError{Op: "Close", FileName: wf.f.Name(), Err: err})
	}
}

func (wf *writeFile) Err() error {
	if wf.err != nil {
		return *wf.err
	}
	return nil
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
