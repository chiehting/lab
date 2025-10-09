package main

import (
	"fmt"
	"strconv"
)

func print(message ...any) {
	fmt.Println(message...)
}

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	var result int
	for x > 0 {
		result = (result * 10) + x%10
		x = x / 10

	}

	if result > 2147483647 {
		result = 0
	}

	return sign * result
}

// @lc code=end

func v1Reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	s := strconv.Itoa(x)
	n := len(s)
	list := make([]int, n)
	pos := 0
	var result int

	for i := n - 1; i >= 0; i-- {
		pow := 1
		for j := range i {
			_ = j
			pow *= 10
		}
		list[pos] = x / pow
		pos++
		x %= pow
	}

	for i := n - 1; i >= 0; i-- {
		pow := 1
		for j := range i {
			_ = j
			pow *= 10
		}
		result += list[i] * pow
		if result >= 2147483647 {
			result = 0
			break
		}
	}

	return sign * result
}

func main() {
	var result int
	input := 1534236469
	result = reverse(input)
	print(result)

	result = v1Reverse(input)
	print(result)
}
