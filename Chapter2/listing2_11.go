package main

import (
	"fmt"
	"os"

	"github.com/go-ini/ini" // #A
)

func main() {
	config, err := ini.Load("conf.ini") // #A
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(config.Section("Section").Key("path").String()) // #B

	enabled, err := config.Section("Section").Key("enabled").Bool() // #C
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(enabled) // #D
}
