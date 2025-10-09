/*
go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: binaryTreePreorderTraversal
cpu: Apple M2 Pro
BenchmarkRecursivePostorderTraversal-12         19099137                62.82 ns/op           80 B/op          1 allocs/op
BenchmarkPostorderTraversal-12                  12126958                86.80 ns/op          160 B/op          2 allocs/op
BenchmarkBreakPostorderTraversal-12             51215841                22.74 ns/op           16 B/op          2 allocs/op
PASS
ok      binaryTreePreorderTraversal     4.699s

關鍵指標比較：
版本	執行次數 (b.N)	每次操作時間	佔用記憶體		分配次數	相對效能
遞迴	19099137		62.82 ns/op		80 B/op		1 allocs/op 基準 better
迭代	12126958		86.80 ns/op		160 B/op	2 allocs/op	≈0.72%
破壞	51215841		22.74 ns/op		16 B/op		2 allocs/op

相對效能 ← (113.2 ns/op / 49.15 ns/op) ≈ 2.30 倍
*/

package main

import "testing"

func BenchmarkRecursivePostorderTraversal(b *testing.B) {
	list := []any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, nil, nil, 9}
	root := setBinaryTree(list)
	for i := 0; i < b.N; i++ {
		recursivePostorderTraversal(root)
	}
}

func BenchmarkPostorderTraversal(b *testing.B) {
	list := []any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, nil, nil, 9}
	root := setBinaryTree(list)
	for i := 0; i < b.N; i++ {
		postorderTraversal(root)
	}
}

func BenchmarkBreakPostorderTraversal(b *testing.B) {
	list := []any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, nil, nil, 9}
	root := setBinaryTree(list)
	for i := 0; i < b.N; i++ {
		breakPostorderTraversal(root)
	}
}
