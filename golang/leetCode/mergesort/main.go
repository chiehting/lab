package main

import (
	"fmt"
)

// MergeSort 以遞迴方式實作經典 Merge‑Sort
// input: a slice of ints which will be sorted in place
// MergeSort 原則：分割 → 排序 → 合併
func MergeSort(arr []int) {
	if len(arr) < 2 { // 基本情況：子序列長度不超過 1 時已排好序
		return
	}
	mid := len(arr) / 2
	// 區分為左右兩個子序列
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)
	copy(left, arr[:mid])
	copy(right, arr[mid:])

	// 遞迴給左右子序列排序
	MergeSort(left)
	MergeSort(right)

	// 合併已排序好的左右子序列
	merge(arr, left, right)
}

// merge 兩個已排好序的 slice 合併為一個排好序的結果
// arr: 合併後寫回的 slice
// left, right: 兩個已排好序的子序列
func merge(arr, left, right []int) {
	i, j, k := 0, 0, 0
	// 同時取左 / 右子序列中的小值
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] { // <= 保有穩定性
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	// 一側已全部取完，直接複製剩餘部分
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

// ---------------- 僅作示例：不使用遞迴的迭代 Merge‑Sort ----------------

// IterativeMergeSort 是一個不使用遞迴的版本
func IterativeMergeSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	// size 為子序列長度，初始為 1，逐步翻倍
	for width := 1; width < n; width *= 2 {
		// 每一次外層循環處理整個陣列的兩段子序列
		for i := 0; i < n; i += 2 * width {
			// left 子序列長度
			left := i
			mid := min(i+width, n)
			// right 子序列長度
			right := n

			// 這裡可直接使用一次排序後的臨時 buffer 進行合併
			temp := make([]int, 0, min(n-left, mid-left)+min(n-mid, right-mid))
			a, b := left, mid
			for a < mid && b < right {
				if arr[a] <= arr[b] {
					temp = append(temp, arr[a])
					a++
				} else {
					temp = append(temp, arr[b])
					b++
				}
			}
			for a < mid {
				temp = append(temp, arr[a])
				a++
			}
			for b < right {
				temp = append(temp, arr[b])
				b++
			}
			// 把合併結果寫回原陣列
			copy(arr[left:left+len(temp)], temp)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// --------------- main ---------------

func main() {
	unsorted := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("原始:", unsorted)

	// 遞迴版
	MergeSort(unsorted)
	fmt.Println("MergeSort 遞迴:", unsorted)

	// 迭代版
	// 因遞迴完成後已排好序，為了展示，可以重新製作未排序 slice
	unsorted = []int{38, 27, 43, 3, 9, 82, 10}
	IterativeMergeSort(unsorted)
	fmt.Println("MergeSort 迭代:", unsorted)
}
