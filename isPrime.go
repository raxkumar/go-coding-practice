package main

import "fmt"

func isPrime(value int) bool {
	var isPrime bool = true
	if value == 0 || value == 1 {
		isPrime = false
	}
	if value == 2 {
		isPrime = true
	}

	for i := 2; i <= value/2; i++ {
		if (value % i) == 0 {
			isPrime = false
		}
	}
	return isPrime
}

func Prime() {

	fmt.Println(isPrime(5))
	fmt.Println(isPrime(10))

}
