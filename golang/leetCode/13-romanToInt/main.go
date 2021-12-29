package main

import (
	"fmt"
	"time"
)

var symbolToInt = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum := 0
	symbols := []rune(s)
	start, end := len(symbols)-1, 0
	for i := start; i >= end; i-- {
		currentValue := symbolToInt[symbols[i]]
		if i < start && symbolToInt[symbols[i+1]] > currentValue {
			sum -= currentValue
		} else {
			sum += currentValue
		}
	}

	return sum
}

func main() {
	t1 := time.Now()
	romanToInt("MCMXCIV")
	fmt.Println(t1.Sub(time.Now()))
}

