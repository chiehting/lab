package main

import (
	"testing"
)

// go test -v -bench=. -run=none .

func Benchmark_twoSum(b *testing.B) {
	nums := []int{3, 2, 4}
	target := 6
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
