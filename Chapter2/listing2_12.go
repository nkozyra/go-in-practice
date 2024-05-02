package main

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

func main() {
	config, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Println("conf.ini not found")
		os.Exit(1)
	}

	fmt.Println(config.Section("Section").Key("path").String())

	enabled, err := config.Section("Section").Key("enabled").Bool()
	if err != nil {
		fmt.Println("enabled is not set!")
		os.Exit(1)
	}
	fmt.Println(enabled)
}
