package main

import "fmt"

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	i, j, k := 0, 0, 0
	arr := make([]int, l1+l2)
	for i < l1 || j < l2 {

		if i < l1 && j < l2 {
			if nums1[i] < nums2[j] {
				arr[k] = nums1[i]
				i++
			} else {
				arr[k] = nums2[j]
				j++
			}
		} else if i < l1 {
			arr = append(arr[:k], nums1[i:]...)
			i = l1
		} else if j < l2 {
			arr = append(arr[:k], nums2[j:]...)
			j = l2
		}
		k++
	}

	l := len(arr)
	var result float64
	if l%2 == 0 {
		result = (float64(arr[l/2]+arr[l/2-1]) / 2)
	} else {
		result = float64(arr[l/2])
	}
	return result
}

// @lc code=end
func print(m ...any) {
	fmt.Println(m...)
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 3, 6}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
