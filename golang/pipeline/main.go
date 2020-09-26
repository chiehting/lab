package main

import "fmt"

func main(){
	pipe := make(chan int,3)
	pipe <- 10
	pipe <- 20
	pipe <- 30

	fmt.Println(len(pipe))

	var t int
	t =<- pipe
	fmt.Println(t)

	fmt.Println(len(pipe))

	t =<- pipe
	fmt.Println(t)
	fmt.Println(len(pipe))
}

