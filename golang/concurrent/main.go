package main

import (
	"fmt"
	"time"
)

func printnum(a int){
	fmt.Println(a)
}

func main(){
	for i:=0;i<10;i++ {
		go printnum(i)
	}
	time.Sleep(time.Second)
}
