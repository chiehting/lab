package main

import (
	"fmt"
)

func print(message ...any) {
	fmt.Println(message...)
}

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	l := len(s)
	if l < numRows || numRows == 1 {
		return s
	}

	result := make([]byte, 0, l)
	n := numRows - 1
	mP := n * 2
	nP := 0
	sw := true
	position := 0

	for i := range numRows {
		nP = (n - i) * 2
		if i == 0 {
			nP = 0
		}

		position = i
		sw = true
		for position < l {
			result = append(result, s[position])
			if sw && nP != 0 {
				position = position + nP
			} else {
				position = position + (mP - nP)
			}
			sw = !sw
		}
	}
	return string(result)
}

// @lc code=end
func main() {
	var result string
	result = convert("PAYPALISHIRING", 3)
	// result = convert("PAYPALISHIRING", 4)
	// result = convert("A", 1)
	// result = convert("AB", 1)
	// result = convert("AB", 3)
	// result = convert("ABCDEF", 4)
	print(result)
}

/*
下面可以看到規則
- 迴圈會執行 numRows 次
- 最大跳 mP 個字元, mP = (numRows - 1) * 2
- 接者會跳 nP 個字元, nP = (n - i) * 2
*/

// number rows 3
// P   A   H   N
// A P L S I I G
// Y   I   R

// 4   1 5 9 13
// 2,2 2 4 6 8 10 12 14
// 4   3 7 11

// number rows 4
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

// 6    1 7 13
// 4,2  2 6 8 12 14
// 2,4  3 5 9 11
// 6    4 10

// number rows 5
// P     H
// A   S I
// Y  I  R
// P L   I G
// A     N

// 8   1 9
// 6,2 2 8 10
// 4,4 3 7 11
// 2,6 4 6 12 14
// 8   5 13

// number rows 6
// P       R
// A     I I
// Y    H  N
// P   S   G
// A I
// L

// 10    1 11
// 8,2   2 10 12
// 6,4   3 9  13
// 4,4   4 8  14
// 2,8   5 7
// 10    6

// number rows 7
// P       N
// A     I G
// Y    R
// P   I
// A  H
// L S
// I

// 12     1 13
// 10,2   2 12 14
// 8,4    3 11
// 6,6    4 10
// 4,8    5 9
// 2,10   6 8
// 12     7
