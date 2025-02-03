package main

import (
	// Same as before…
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	w := newWords()
	if len(os.Args) < 2 {
		fmt.Println("no files provided")
		os.Exit(1)
	}
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Println("Words that appear more than once:")
	w.Lock()         // # A
	defer w.Unlock() // # A
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}

type words struct {
	sync.Mutex // # B
	found      map[string]int
}

func newWords() *words {
	return &words{found: map[string]int{}}
}
func (w *words) add(word string) {
	w.Lock() // # C
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = 1
		return
	}
	w.found[word] = count + 1
}
func tallyWords(filename string, dict *words) error {
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
