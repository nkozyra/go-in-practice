package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("checking for environment variable CLEANUP") // # A

	if envvar := os.GetEnv("CLEANUP"); envvar != "" {
		fmt.Println(" did not find it, value is:", envvar) // # B
	}

}
