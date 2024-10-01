package main

import "testing"

func TestMain(t *testing.T) {
	if getName() != "World!" {
		t.Error("Test failed")
	}
}
