package main

import (
	"testing"
)

func Test_FizzBuzz(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{37, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19, Buzz, Fizz, 22, 23, Fizz, Buzz, 26, Fizz, 28, 29, Fizz Buzz, 31, 32, Fizz, 34, Buzz, Fizz"},
	}
	for i := range tests {
		test := tests[i]
		res := fizzbuzz(test.input)
		if res != test.expected {
			t.Fatalf("\ngot \n%s \nexpected \n%s", res, test.expected)
		}
	}
}
