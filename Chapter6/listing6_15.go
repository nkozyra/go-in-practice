package main

import (
	"fmt"
	"strings"
)

func summedRuneCodes(input string) int16 {
	value := 0
	inRunes := strings.Map(func(r rune) rune { // # A
		return r
	}, input)

	for r := range inRunes {
		value += int(inRunes[r]) // # B
	}

	return int16(value)
}

func main() {
	var testString = "i am a test string"
	output := summedRuneCodes(testString)
	fmt.Println(output)
}
