package main

import (
	"fmt"
	"strings"
)

func print(message ...any) {
	fmt.Println(message...)
}

/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	var negative bool

	switch s[0] {
	case '-':
		negative = true
		s = s[1:]
	case '+':
		s = s[1:]
	}

	var result int
	for i, c := range s {
		if c < 48 || c > 57 { // 48 is ascii code 0, 57 is ascii code 9
			break
		}

		result = result*10 + int(s[i]-48)

		if result > 2147483647 && !negative {
			return 2147483647
		}

		if -result < -2147483648 && negative {
			return -2147483648
		}
	}

	if negative {
		result = -result
	}

	return result
}

// @lc code=end

func v1MyAtoi(s string) int {
	if len(s) > 200 {
		return 0
	}
	var result uint

	zero := true
	space := true
	sign := true
	negative := 1

	for i := 0; i < len(s); i++ {
		if space && s[i] == 32 {
			continue
		}

		if sign && s[i] == 43 {
			space = false
			sign = false
			continue
		}

		if sign && s[i] == 45 {
			space = false
			sign = false
			negative = -1
			continue
		}

		if zero && s[i] == 48 {
			space = false
			sign = false
			zero = false
			continue
		}

		if s[i] > 47 && s[i] < 58 && result < 2147483647 {
			space = false
			sign = false
			zero = false
			result = result*10 + uint(s[i]) - 48
			continue
		}
		break
	}
	if result > 2147483646 && negative > 0 {
		result = 2147483647
	}
	if result > 2147483647 && negative < 0 {
		result = 2147483648
	}

	return negative * int(result)
}

func main() {
	var result int
	var input string
	// input = "42"
	input = "   -042"
	// input = "1337c0d3"
	// input = "0-1"
	// input = "-91283472332"
	// input = "+1"
	// input = "9223372036854775808"
	// input = "18446744073709551617"
	// input = "-2147483647"
	result = v1MyAtoi(input)
	print(result)
	result = myAtoi(input)
	print(result)

}
