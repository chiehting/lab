package main

import "fmt"

// 遞迴
func recursive_heapify(list []int, listLength int, idx int) {
	leftIdx, rightIdx, targetIdx := idx*2+1, idx*2+2, idx

	if rightIdx < listLength && list[rightIdx] > list[targetIdx] {
		targetIdx = rightIdx
	}

	if leftIdx < listLength && list[leftIdx] > list[targetIdx] {
		targetIdx = leftIdx
	}

	if targetIdx != idx {
		swap(list, idx, targetIdx)
		heapify(list, listLength, targetIdx)
	}
}

// 迭代
func heapify(list []int, listLength int, idx int) {
	stack := []struct{ n, idx int }{{listLength, idx}}
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		leftIdx, rightIdx, targetIdx := p.idx*2+1, p.idx*2+2, p.idx

		if rightIdx < listLength && list[rightIdx] > list[targetIdx] {
			targetIdx = rightIdx
		}

		if leftIdx < listLength && list[leftIdx] > list[targetIdx] {
			targetIdx = leftIdx
		}

		if targetIdx != p.idx {
			swap(list, p.idx, targetIdx)
			// heapify(list, listLength, targetIdx)
			stack = append(stack, struct{ n, idx int }{listLength, targetIdx})
		}
	}
}

func heapSort(list []int) (result []int) {
	listLength := len(list)
	subRoot := listLength/2 - 1

	// 建立最大堆（從最後一個非葉子節點開始）
	for i := subRoot; i >= 0; i-- {
		heapify(list, listLength, i)
	}

	elementCount := listLength

	for elementCount > 0 {
		elementCount = elementCount - 1
		swap(list, 0, elementCount)
		heapify(list, elementCount, 0)
	}

	return list
}

func swap(list []int, sorceIdx int, targetIdx int) {
	tmp := list[sorceIdx]
	list[sorceIdx] = list[targetIdx]
	list[targetIdx] = tmp
}

func print(message ...any) {
	fmt.Println(message...)
}

func main() {
	// 測試案例
	testCases := [][]int{
		{1},
		{2, 1},
		{5, 2, 8, 1, 9},
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{3, 1, 5, 0, 2, 3, 9, 10, 12, 15, 29, 41, 38},
		{1, 2, 3, 4, 5, 6, 7},
		{10, 20, 30, 40, 50, 60},
	}

	for _, v := range testCases {
		print("測試:", v)
		print("結果:", heapSort(v))
		print("---")
	}
}
