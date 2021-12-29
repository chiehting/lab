package main

import (
	"fmt"
	"strings"
	"time"
)

func longestCommonPrefix(strs []string) string {
	var sb strings.Builder
	for i := 0; ; i++ {
		for _, v := range strs {
			if i == len(v) || strs[0][i] != v[i] {
				return sb.String()
			}
		}
		sb.WriteByte(strs[0][i])
	}
}

func main() {
	t1 := time.Now()
	longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(t1.Sub(time.Now()))
}

