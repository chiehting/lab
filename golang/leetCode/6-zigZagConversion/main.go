package main

import (
	"fmt"
	"time"
)

func convert(s string, numRows int) string {

	return ""
}

func main() {
	t1 := time.Now()
	a := convert("PAYPALISHIRING", 3)
	fmt.Println(a)
	fmt.Println(time.Now().Sub(t1))
}
