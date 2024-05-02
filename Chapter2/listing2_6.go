package main

import (
	"fmt"
	"log"
	"reflect"

	flags "github.com/spf13/pflag"
)

var opts struct {
	Name    string
	Spanish bool
}

func init() {
	flags.StringVarP(&opts.Name, "name", "n", "", "A name to say hello to.")
	flags.BoolVarP(&opts.Spanish, "spanish", "s", false, "Use Spanish language.")
	flags.Parse()
}

type Animal struct {
	name string
}

func (a Animal) speak() string {
	switch a.name {
	case "cat":
		return "meow"
	case "dog":
		return "woof"
	default:
		if member, ok := reflect.TypeOf(a).FieldByName("name"); ok {
			return fmt.Sprintf("Invalid animal name: %s", member.Tag.Get("help"))
		}
		return "nondescript animal noise?"
	}
}

func main() {

	a := Animal{"cat"}
	log.Println(a.speak())

	flags.Parse()
	if opts.Spanish {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
