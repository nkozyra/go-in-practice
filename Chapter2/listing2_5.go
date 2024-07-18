package main

import (
	"flag"
	"log"
)

var name = flag.String("name", "World", "A name to say hello to.")

type language = string // #A

var userLanguage language // #B

const ( // #C
	English language = "en"
	Spanish          = "sp"
	French           = "fr"
	German           = "de"
)

func init() {
	flag.StringVar(&userLanguage, "lang", "en", "language to use (en, sp, fr, de).") // #D
	flag.Parse()
}

func main() {
	switch userLanguage { // #E
	case English: // #E
		log.Printf("Hello %s!\n", *name)
	case Spanish:
		log.Printf("Hola %s!\n", *name)
	case French:
		log.Printf("Bonjour %s!\n", *name)
	case German:
		log.Printf("Hallo %s!\n", *name)
	}
}
