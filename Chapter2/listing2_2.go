package main

import (
	"flag" // #A
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.") // #B
var inSpanish bool                                                 // #C
func init() {
	flag.BoolVar(&inSpanish, "spanish", false, "Use Spanish language.") // #C
	flag.BoolVar(&inSpanish, "s", false, "Use Spanish language.")       // #D
	flag.Parse()                                                        //  #E
}

func main() {
	if inSpanish {
		fmt.Printf("Hola %s!\n", *name) // #F
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
