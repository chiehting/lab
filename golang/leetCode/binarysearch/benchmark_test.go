/*
go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: binarysearch
cpu: Apple M2 Pro
BenchmarkBinarySearchOriginal-12        520466720                2.062 ns/op           0 B/op          0 allocs/op
BenchmarkBinarySearchOptimized-12       682255736                1.759 ns/op           0 B/op          0 allocs/op
PASS
ok      binarysearch    2.985s

關鍵指標比較：
版本		執行次數	每次操作時間	相對效能
原始版本	5.2億次		2.062 ns/op		基準
優化版本	6.8億次		1.759 ns/op		+17% 更快

*/

package main

import "testing"

func BenchmarkBinarySearchOriginal(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	for i := 0; i < b.N; i++ {
		binarySearchOriginal(arr, 12)
	}
}

func BenchmarkBinarySearchOptimized(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	for i := 0; i < b.N; i++ {
		binarySearchOptimized(arr, 12)
	}
}

func binarySearchOriginal(list []int, target int) int {
	left, right, mid := 0, len(list)-1, 0
	result := -1

	for left <= right {
		mid = left + (right-left)/2

		if target < list[mid] {
			right = mid - 1
		} else if target > list[mid] {
			left = mid + 1
		} else if target == list[mid] {
			result = mid
			break
		}
	}

	return result
}

func binarySearchOptimized(list []int, target int) int {
	if len(list) == 0 {
		return -1
	}

	left, right := 0, len(list)-1

	for left <= right {
		mid := left + (right-left)/2
		if list[mid] == target {
			return mid
		}
		if target < list[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
