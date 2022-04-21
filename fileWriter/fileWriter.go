package fileWriter

import (
	"fmt"
	"io"
	"os"
)

type WriteFileError struct {
	Op       string
	FileName string
	err      error
}

func (e WriteFileError) Error() string {
	return fmt.Sprintf("Error in %s method working with file \"%s\": %v", e.Op, e.FileName, e.err)
}

func (e WriteFileError) Unwrap() error {
	return e.err
}

type WriteFile struct {
	f   *os.File
	err *WriteFileError
}

func NewWriteFile(fileName string) *WriteFile {
	f, err := os.Create(fileName)
	var wrapped *WriteFileError = nil
	if err != nil {
		wrapped = &(WriteFileError{Op: "newWriteFile", FileName: fileName, err: err})
		return &WriteFile{f: nil, err: wrapped}
	}
	return &WriteFile{f: f, err: nil}
}

func (wf *WriteFile) WriteString(s string) {
	if wf.err != nil {
		return
	}
	_, err := io.WriteString(wf.f, s)
	if err != nil {
		wf.err = &(WriteFileError{Op: "WriteString", FileName: wf.f.Name(), err: err})
	}
}

func (wf *WriteFile) Close() {
	if wf.f == nil {
		return
	}
	err := wf.f.Close()
	if err != nil {
		wf.err = &(WriteFileError{Op: "Close", FileName: wf.f.Name(), err: err})
	}
}

func (wf *WriteFile) Err() error {
	if wf.err != nil {
		return *wf.err
	}
	return nil
}
