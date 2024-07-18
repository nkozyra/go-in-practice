package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup // # A
	w := newWords()
	if len(os.Args) < 2 {
		fmt.Println("no files provided")
		os.Exit(1)
	}
	for _, f := range os.Args[1:] { // # B
		wg.Add(1)              // # B
		go func(file string) { // # B
			if err := tallyWords(file, w); err != nil { // # B
				fmt.Println(err.Error()) // # B
			} // # B
			wg.Done() // # B
		}(f) // # B
	}
	wg.Wait()
	fmt.Println("Words that appear more than once:") // # C
	for word, count := range w.found {               // # C
		if count > 1 { // # C
			fmt.Printf("%s: %d\n", word, count) // # C
		} // # C
	} // # C
}

type words struct { // # D
	found map[string]int
}

func newWords() *words { // # E
	return &words{found: map[string]int{}}
}

func (w *words) add(word string) { // # F
	count, ok := w.found[word] // # G
	if !ok {                   // # G
		w.found[word] = 1 // # G
		return            // # G
	} // # G
	w.found[word] = count + 1 // # G
}
func tallyWords(filename string, dict *words) error { // # H
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) // # I
	scanner.Split(bufio.ScanWords)    // # I
	for scanner.Scan() {              // # I
		word := strings.ToLower(scanner.Text()) // # I
		dict.add(word)
	}
	return scanner.Err()
}
