package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml" // #A
)

func main() {
	config, err := yaml.ReadFile("conf.yaml") // #B
	if err != nil {
		fmt.Println(err)
		return
	}
	var path string
	var enabled bool

	path, err = config.Get("path")
	if err != nil {
		fmt.Println("`path` flag not set in conf.yaml", err)
		return
	}

	enabled, err = config.GetBool("enabled")
	if err != nil {
		fmt.Println("`enabled` flag not set in conf.yaml", err)
		return
	}

	fmt.Println("path", path)       // #C
	fmt.Println("enabled", enabled) // #C

}
