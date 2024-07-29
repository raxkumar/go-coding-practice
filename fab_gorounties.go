package main

import "fmt"

func Fibonacci(n int, channel chan<- int) {
	x, y := 0, 1

	for i := 1; i <= n; i++ {
		channel <- x
		x, y = y, x+y
	}
	close(channel)

}
func GoRoutineFab() {
	ch := make(chan int, 10)
	go Fibonacci(10, ch)

	for i := range ch {
		fmt.Println(i)
	}

}
