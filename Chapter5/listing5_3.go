package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	for _, file := range os.Args[1:] { // # A
		compress(file) // # A
	} // # A
}
func compress(filename string) error {
	in, err := os.Open(filename) // # B
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(filename + ".gz") // # C
	if err != nil {
		return err
	}
	defer out.Close()
	gzout := gzip.NewWriter(out) // # D
	_, err = io.Copy(gzout, in)  // # E
	gzout.Close()
	return err
}
