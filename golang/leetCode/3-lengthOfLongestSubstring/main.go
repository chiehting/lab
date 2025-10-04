/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

package main

import (
	"fmt"
)

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var leftIdx, max int = 0, 0
	var seen = make(map[rune]int)

	for rightIdx, v := range s {
		if findIdx, ok := seen[v]; ok && findIdx > leftIdx {
			leftIdx = findIdx
		}
		tmp := rightIdx - leftIdx + 1
		if tmp > max {
			max = tmp
		}

		seen[v] = rightIdx + 1
	}
	return max
}

// @lc code=end
func main() {
	testCase := []string{
		"abcabcbb", //3
		"abba",     //2
		"bbbbb",    //1
		" ",        //1
		"au",       //2
		"dvdf",     //3
	}

	for _, v := range testCase {
		// t1 := time.Now()
		fmt.Println(lengthOfLongestSubstring(v))
		fmt.Println("---")
		// fmt.Println(t1.Sub(time.Now()))
	}

}
