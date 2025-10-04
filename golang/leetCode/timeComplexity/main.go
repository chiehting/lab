package main

import "fmt"

func main() {
	count := linearLogRecur(1.5)
	fmt.Println(count)
}

// 取對數
func linearLogRecur(n float64) int {
	if n <= 1 {
		return 1
	}
	count := linearLogRecur(n/2) + linearLogRecur(n/2)
	for i := 0.0; i < n; i++ {
		count++
	}
	return count
}
