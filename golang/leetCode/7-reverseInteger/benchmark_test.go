/*
go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: bTreeinorderTraversal
cpu: Apple M2 Pro
BenchmarkRecursiveinorderTraversal-12          10735490               113.2 ns/op           248 B/op          5 allocs/op
BenchmarkinorderTraversal-12                   23935214                49.15 ns/op           48 B/op          2 allocs/op
PASS
ok      bTreeinorderTraversal  3.875s

關鍵指標比較：
版本	執行次數 (b.N)	每次操作時間	佔用記憶體		分配次數	相對效能
遞迴	10735490		113.2 ns/op		248 B/op	5 allocs/op	基準
迭代	23935214		49.15 ns/op		48 B/op		2 allocs/op	≈230 % better

相對效能 ← (113.2 ns/op / 49.15 ns/op) ≈ 2.30 倍
*/

package main

import "testing"

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(i)
	}
}

func BenchmarkV1Reverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1Reverse(i)
	}
}
