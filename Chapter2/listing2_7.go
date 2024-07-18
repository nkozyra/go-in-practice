package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCommand *cobra.Command // #A

func init() {
	helloCommand = &cobra.Command{ // #B
		Use:   "hello",
		Short: "Print hello world",
		Run:   sayHello,
	}
	helloCommand.Flags().StringP("name", "n", "World", "Who to say hello to.")             // #C
	helloCommand.MarkFlagRequired("name")                                                  // #D
	helloCommand.Flags().StringP("language", "l", "en", "Which language to say hello in.") // #E
}

func sayHello(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name") // #F
	greeting := "Hello"
	language, _ := cmd.Flags().GetString("language") // #F
	switch language {                                // #G
	case "en":
		greeting = "Hello"
	case "es":
		greeting = "Hola"
	case "fr":
		greeting = "Bonjour"
	case "de":
		greeting = "Hallo"
	}
	fmt.Printf("%s %s!\n", greeting, name)
}

func main() {
	helloCommand.Execute() // #H
}
