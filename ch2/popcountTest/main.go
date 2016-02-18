// popcount-test
package main

import "fmt"
import "testing"

import "../popcount"
import "../popcountLoop"
import "../popcountClearLsb"
import "../popcountShift64"

func main() {
	fmt.Println(popcount.PopCount(4))
	fmt.Println(popcountLoop.PopCount(4))
	fmt.Println(popcountClearLsb.PopCount(4))
	fmt.Println(popcountShift64.PopCount(4))
	br := testing.Benchmark(BenchmarkPopcount)
	fmt.Println(br)
	br = testing.Benchmark(BenchmarkPopcountLoop)
	fmt.Println(br)
	br = testing.Benchmark(BenchmarkPopcountShift64)
	fmt.Println(br)
	br = testing.Benchmark(BenchmarkPopcountClearLsb)
	fmt.Println(br)
}

func BenchmarkPopcount(b *testing.B) {
	var x uint64
	x = 0
	for i := 0; i < b.N; i++ {
		_ = popcount.PopCount(x)
	}
}

func BenchmarkPopcountLoop(b *testing.B) {
	var x uint64
	x = 0
	for i := 0; i < b.N; i++ {
		_ = popcountLoop.PopCount(x)
	}
}

func BenchmarkPopcountClearLsb(b *testing.B) {
	var x uint64
	x = 0
	for i := 0; i < b.N; i++ {
		_ = popcountClearLsb.PopCount(x)
	}
}

func BenchmarkPopcountShift64(b *testing.B) {
	var x uint64
	x = 0
	for i := 0; i < b.N; i++ {
		_ = popcountShift64.PopCount(x)
	}
}
