package main

import "fmt"

func binarySearch(list []int, target int) int {
	left, right, mid := 0, len(list)-1, 0
	result := -1

	for left <= right {
		mid = left + (right-left)/2 // 避免整數溢出

		/*
			Q: 因為找到目標的可能性較低, 所以把找到目標放在最後一個 else if 的判斷中, 會不會增加效率?
			A: 想法是正確的！將最可能發生的情況放在前面確實可以提高效率。但在 binary search 的情況下，找到目標的機率其實相對較高，特別是當目標存在於數組中時。

			效率分析:
			在 binary search 中，三種情況的發生機率：

			target < list[mid] - 約 33% 機率
			target > list[mid] - 約 33% 機率
			target == list[mid] - 約 33% 機率
		*/

		// 先檢查是否找到目標（最可能提前結束的情況）
		if list[mid] == target {
			return mid
		}

		if target < list[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func print(message ...any) {
	fmt.Println(message...)
}

func main() {
	// 測試案例
	testCases := [][]int{
		{1, 3, 5, 7, 9, 11, 13, 15},  // 目標不存在
		{2, 4, 6, 8, 10, 12, 14, 16}, // 目標存在，在中間
		{12, 13, 14, 15, 16, 17, 18}, // 目標存在，在開頭
		{1, 2, 3, 4, 5, 6, 12},       // 目標存在，在末尾
		{12},                         // 只有一個元素且是目標
		{13},                         // 只有一個元素但不是目標
		{},                           // 空數組
		{10, 11, 12, 12, 12, 13, 14}, // 重覆元素，目標存在
	}

	for _, v := range testCases {
		print("測試:", v)
		print("結果:", binarySearch(v, 12))
		print("---")
	}
}
