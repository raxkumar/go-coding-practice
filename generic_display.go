//go:build go1.18
// +build go1.18

package main

import "fmt"

// generics allow us to work with different data types
func Display[T any](stringArr []T) {
	for _, s := range stringArr {
		fmt.Println(s)
	}

}

func GenericDisaply() {
	stringArr := []string{"aa", "bb", "cc", "dd"}
	intArr := []int{1, 2, 3, 4}

	Display(stringArr)
	Display(intArr)

	StackFunc()

}
