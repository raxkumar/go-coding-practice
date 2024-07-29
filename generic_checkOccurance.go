package main

import (
	"fmt"
)

func checkOccurance[T any](arr []T) {
	arrMap := make(map[interface{}]int)

	for _, value := range arr {
		arrMap[value] = arrMap[value] + 1
	}

	fmt.Println(len(arrMap))

	for key, value := range arrMap {
		fmt.Println(key, ": ", value)
	}

}

func GenericCheckOccurance() {

	arrInt := []int{1, 2, 3, 4, 5, 1, 2, 6, 7}

	arrString := []string{"a", "b", "c", "d", "e", "a"}

	checkOccurance(arrInt)
	checkOccurance(arrString)
}
