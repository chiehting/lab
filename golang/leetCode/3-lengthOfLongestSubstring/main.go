/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

package main

import (
	"fmt"
	"time"
)

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var left, max int = 0, 0
	var seen = make(map[rune]int)

	for k, v := range s {
		if idx, ok := seen[v]; ok && idx > left {
			left = idx
		}
		tmp := k - left + 1
		if tmp > max {
			max = tmp
		}

		seen[v] = k + 1
	}
	return max
}

// @lc code=end
func main() {
	t1 := time.Now()
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("abba"))     // 2
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring(" "))        // 1
	fmt.Println(lengthOfLongestSubstring("au"))       // 2
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // 3
	fmt.Println(t1.Sub(time.Now()))
}
