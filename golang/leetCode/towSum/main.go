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

// 網友寫的
// Benchmark_twoSum-16     11157608                96.1 ns/op 1.418s
func twoSum(nums []int, target int) []int {
	seen := map[int][]int{}
	ans := []int{}

	for i, val := range nums {
		search := target - val
		v, ok := seen[search]

		if ok {
			ans = append(ans, i, v[0])
			break
		}
		_, present := seen[val]
		if present {
			seen[val] = append(seen[val], i)
		} else {
			seen[val] = []int{i}
		}
	}

	return ans
}

// 我寫的
// Benchmark_twoSum-16     39556572                25.3 ns/op 2.526s
// func twoSum(nums []int, target int) []int {
// 	var resoult []int
// 	for i, v := range nums {
// 		for j := (i + 1); j < len(nums); j++ {
// 			if target == (v + nums[j]) {
// 				return append(resoult, i, j)
// 			}
// 		}
// 	}
// 	return nil
// }
