package main

import (
	"fmt"
	"time"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var take []int
	var i = 1
	for {
		take = append(take, (x / i % 10))
		i = i * 10
		if i > x {
			break
		}
	}

	for i := 0; i < len(take)/2; i++ {
		if take[i] != take[len(take)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	t1 := time.Now()
	isPalindrome(-121)
	isPalindrome(121)
	isPalindrome(0)
	fmt.Println(t1.Sub(time.Now()))
}

