package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func1(ch)

	go func2(ch)
	time.Sleep(2 * 1e9)
}

func func1(ch chan int) {
	for i := 0; i < 100; i+=10 {
		ch <- i
	}
}

func func2(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
