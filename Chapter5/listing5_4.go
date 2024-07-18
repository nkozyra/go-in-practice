package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup // # A
	for i, file := range os.Args[1:] {
		wg.Add(1)                  // # B
		go func(filename string) { // # C
			compress(filename) // # C
			wg.Done()          // # C
		}(file) // # D
	}
	wg.Wait() // # E
	fmt.Printf("Compressed %d files\n", len(os.Args[1:]))
}
func compress(filename string) error {
	// Unchanged from above
}
