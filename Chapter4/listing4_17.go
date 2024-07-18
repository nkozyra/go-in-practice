package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var file io.ReadCloser
	file, err := OpenCSV("data.csv") // # A
	if err != nil {                  // # A
		fmt.Printf("Error: %s", err) // # A
		return                       // # A
	} // # A
	defer file.Close() // # B
	// Do something with file.  // # C
}
func OpenCSV(filename string) (file *os.File, err error) { // # D
	defer func() { // # E
		if r := recover(); r != nil { // # E
			file.Close()    // # E
			err = r.(error) // # E
		} // # E
	}() // # E
	file, err = os.Open(filename) // # F
	if err != nil {               // # F
		fmt.Printf("Failed to open file\n") // # F
		return file, err                    // # F
	} // # F
	RemoveEmptyLines(file) // # G
	return file, err
}
func RemoveEmptyLines(f *os.File) {
	panic(errors.New("failed parse")) // # H
}
