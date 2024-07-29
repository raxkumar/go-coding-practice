package main

import (
	"fmt"
	"time"
)

func Square(value int, squareValue chan<- int) {
	time.Sleep(3 * time.Second)
	squareValue <- value * value
}

func Factorial(v int) int {
	if v < 2 {
		return 1
	}
	return v * Factorial(v-1)
}

func CalculateFactorial(value int, channel chan<- int) {
	time.Sleep(3 * time.Second)
	channel <- Factorial(value)
}

func ChannelGoroutinesExample() {

	arr := []int{1, 2, 3, 4, 5}

	FactorialChannel := make(chan int, len(arr))
	SquareChannel := make(chan int, len(arr))

	for _, value := range arr {
		go CalculateFactorial(value, FactorialChannel)
		go Square(value, SquareChannel)
	}

	for range arr {
		fact := <-FactorialChannel
		fmt.Println(fact)
	}

	for range arr {
		squared := <-SquareChannel
		fmt.Println(squared)
	}
	close(FactorialChannel)
	close(SquareChannel)

}
