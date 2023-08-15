package main

import (
	"fmt"
	"net/http"
)

func main() {
	count := 10

	ch := make(chan string)
	for i := 0; i < count; i++ {
		go myfun(ch)

		if i >= count {
			close(ch)
		}

	}

	fmt.Println("Channel In ")

	for {
		res, ok := <-ch
		if ok == false {
			fmt.Println("break")
			break
		}
		fmt.Println(res, ok)

	}
}

func myfun(ch chan string) {
	resp, err := http.Get("https://google.com/")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 4033 {
		ch <- resp.Status
	}
}
