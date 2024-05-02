package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var path string
	var enabled bool

	path, err = config.Get("path")
	if err != nil {
		fmt.Println("`path` flag not set in conf.yaml")
		return
	}

	enabled, err = config.GetBool("enabled")
	if err != nil {
		fmt.Println("`enabled` flag not set in conf.yaml")
		return
	}

	fmt.Println("path", path)
	fmt.Println("enabled", enabled)

}
