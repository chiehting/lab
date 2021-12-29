package main

import (
	"fmt"
	"time"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var left, max int = 0, 0
	var seen = make(map[rune]int)

	for k, v := range s {
		if idx, ok := seen[v]; ok && idx >= left {
			left = idx + 1
		}
		seen[v] = k

		tmp := k - left + 1
		if tmp > max {
			max = tmp
		}
	}
	fmt.Println(max)
	return max
}

func main() {
	t1 := time.Now()
	lengthOfLongestSubstring("abcabcbb") // 3
	lengthOfLongestSubstring("bbbbb")    // 1
	lengthOfLongestSubstring(" ")        // 1
	lengthOfLongestSubstring("au")       // 2
	lengthOfLongestSubstring("dvdf")     // 3
	fmt.Println(t1.Sub(time.Now()))
}

