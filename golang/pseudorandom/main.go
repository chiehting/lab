package main

import (
	"fmt"
	"time"
)

// 種子
var next uint64 = 1

// 生成偽隨機數
func rand() uint64 {
	next = (next*1103515245 + 12345)
	return (next / 65536) % 32768
}

// 修改種
func srand(seed uint64) {
	next = seed
}

func main() {
	fmt.Printf("Type: %T Value: %v\n", next, next)
	i := 1
	for i < 10 {
		reseed := (int(time.Now().Unix()) * i) % 100
		srand(uint64(reseed))
		fmt.Printf("next %v\n", next)
		fmt.Printf("rand Type: %T Value: %v\n", rand(), rand())
		i++
	}
}
