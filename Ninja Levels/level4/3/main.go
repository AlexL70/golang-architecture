package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type CopyFile struct {
	err     error
	srcFile *os.File
	dstFile *os.File
}

func (cf *CopyFile) OpenSource(src string) {
	if cf.err != nil {
		return
	}
	srcF, err := os.Open(src)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			cf.err = fmt.Errorf("Source file \"%s\" does not exist! Error: %w", src, err)
			return
		}
		cf.err = fmt.Errorf("Error opening source file \"%s\": %w", src, err)
		return
	}
	cf.srcFile = srcF
}

func (cf *CopyFile) CreateDest(dst string) {
	if cf.err != nil {
		return
	}
	dstF, err := os.Create(dst)
	if err != nil {
		cf.err = fmt.Errorf("Error creating destination file \"%s\": %w", dst, err)
		return
	}
	cf.dstFile = dstF
}

func (cf *CopyFile) DoCopy() int64 {
	if cf.srcFile != nil {
		defer cf.srcFile.Close()
	}
	if cf.dstFile != nil {
		defer cf.dstFile.Close()
	}
	if cf.err != nil {
		return 0
	}
	nBytes, err := io.Copy(cf.dstFile, cf.srcFile)
	if err != nil {
		cf.err = err
		return 0
	}
	return nBytes
}

func (cf CopyFile) Err() error {
	return cf.err
}

func main() {
	cf := CopyFile{}
	cf.OpenSource("rumi.txt")
	cf.CreateDest("rumiCopy.txt")
	copied := cf.DoCopy()
	if cf.Err() != nil {
		fmt.Println(cf.Err())
		return
	}
	fmt.Printf("%d bytes copied\n", copied)
}
