package main

import "fmt"

func main() {
	// go run towSum.go
	// go test -v -bench=. -run=none .

	nums := []int{3, 2, 1, 12, 8, 4}
	target := 7
	resoult := twoSum(nums, target)
	fmt.Println(resoult)
}

// Benchmark_twoSum-16     11157608                96.1 ns/op 1.418s
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // A mapping to store numbers and their indices
	for i, num := range nums {
		complement := target - num // Find the required number to reach the target
		if index, found := numMap[complement]; found {
			return []int{index, i} // Return indices of the complement and current number
		}
		numMap[num] = i // Store the number with its index
	}
	return nil // This line is never reached due to the problem guarantee
}

// // Benchmark_twoSum-16     39556572                25.3 ns/op 2.526s
// func twoSum(nums []int, target int) []int {
// 	n := len(nums)
// 	var resoult []int
// 	for i, v := range nums {
// 		for j := (i + 1); j < n; j++ {
// 			if target == (v + nums[j]) {
// 				return append(resoult, i, j)
// 			}
// 		}
// 	}
// 	return nil
// }
