package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	t.Run("testing foo", func(t *testing.T) {
		foo()
	})
}

func TestBar(t *testing.T) {
	t.Run("testing foo", func(t *testing.T) {
		bar()
	})
}
