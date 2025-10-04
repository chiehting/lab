package main

import (
	"fmt"
)

// ------------- 1. 正確的遞迴 QuickSort -------------
func quickSort(a []int) {
	if len(a) < 2 { // 完全有序或空陣列直接回傳
		return
	}

	// ① 把 pivot 的值儲存起來，pivot 本身仍保留在 a[0]
	pivot := a[0]
	// ② 迴圈中，i 只指向「下一個 < pivot 的位置」。
	//     j 則掃描整個陣列。
	i := 1
	for j := 1; j < len(a); j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i] // 把 < pivot 的元素往左移
			i++                     // 下個 < pivot 的位置
		}
	}
	// ③ 把 pivot 放到正確的位置（i-1）
	a[0], a[i-1] = a[i-1], a[0]

	// ④ 分別遞迴排序左右兩段
	quickSort(a[:i-1]) // 左邊
	quickSort(a[i:])   // 右邊
}

// ------------- 2. 迭代版 QuickSort (避免遞迴深度過大) -------------
func quickSortIter(a []int) {
	// 手動堆疊：每個項目代表一個子區間 [lo, hi]
	stack := []struct{ lo, hi int }{{0, len(a) - 1}}
	for len(stack) > 0 {
		fmt.Println(a, stack)
		// 取得最上層區間
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		lo, hi := p.lo, p.hi
		if lo >= hi {
			continue
		}

		// Partition (兩指針版本，pivot 取左端)
		i, j := lo+1, hi
		pivot := a[lo]
		for i <= j {
			for i <= j && a[i] <= pivot {
				i++
			}
			for i <= j && a[j] >= pivot {
				j--
			}
			if i < j {
				a[i], a[j] = a[j], a[i]
				i++
				j--
			}
		}
		a[lo], a[j] = a[j], a[lo] // 把 pivot 放到正確位置

		// 將左右子區間壓入堆疊
		stack = append(stack, struct{ lo, hi int }{lo, j - 1})
		stack = append(stack, struct{ lo, hi int }{j + 1, hi})
	}
}

// ------------- 3. 主函式 -------------------
func main() {
	arr := []int{7, 2, 9, 4, 5, 1, 3, 6}
	fmt.Println("原始:", arr)

	// 遞迴版
	quickSort(arr)
	fmt.Println("遞迴 QuickSort:", arr)

	// 迭代版（重置順序）
	arr = []int{7, 2, 9, 4, 5, 1, 3, 6}
	quickSortIter(arr)
	fmt.Println("迭代 QuickSort:", arr)
}
