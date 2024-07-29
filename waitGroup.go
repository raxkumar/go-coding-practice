package main

import (
	"fmt"
	"sync"
	"time"
)

func ProcessFile(filepath string, wg *sync.WaitGroup) {

	fmt.Printf("Processing of file %v started, I will take %d sec.\n", filepath, len(filepath))
	time.Sleep(time.Duration(len(filepath)/2) * time.Second)
	fmt.Printf("%v Has been Processed\n", filepath)
	defer wg.Done()

}

func WaitGroup() {

	file := []string{"/usr/local/001.txt", "/usr/local/docs.txt", "/usr/local/readme.md"}

	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go ProcessFile(file[i], wg)
	}

	wg.Wait()

}
