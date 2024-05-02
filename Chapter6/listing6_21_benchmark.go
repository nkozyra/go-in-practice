package main

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func runBenchmark(arr []int, runs int) {
	for i := 0; i < runs; i++ {
		bubbleSort(arr)
	}
}

func generateRandoms(num int) []int {
	out := make([]int, num)
	for k := range out {
		out[k] = rand.Intn(num-1) + 1
	}
	return out
}

func BenchmarkBubbleSort10(b *testing.B) {
	runBenchmark(generateRandoms(10), b.N)
}

func BenchmarkBubbleSort100(b *testing.B) {
	runBenchmark(generateRandoms(100), b.N)
}

func BenchmarkBubbleSort1000(b *testing.B) {
	runBenchmark(generateRandoms(1000), b.N)
}

func BenchmarkBubbleSort10000(b *testing.B) {
	runBenchmark(generateRandoms(10000), b.N)
}

func BenchmarkBubbleSort100000(b *testing.B) {
	runBenchmark(generateRandoms(100000), b.N)
}
