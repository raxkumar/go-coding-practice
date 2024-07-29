package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func WebScrapper(url string, channel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	timeToprocess, _ := strconv.Atoi(string(url[len(url)-1]))
	fmt.Printf("Reading Page:  %d\n", timeToprocess)
	time.Sleep(time.Duration(timeToprocess) * time.Second)
	channel <- fmt.Sprintf("This is the HTML page from %s", url)
}
func WebScrapperFun() {

	wg := &sync.WaitGroup{}

	webUrl := []string{"https://example.com/page1", "https://example.com/page2", "https://example.com/page3"}
	BufferSize := 3
	channel := make(chan string, BufferSize)

	for _, url := range webUrl {
		wg.Add(1)
		go WebScrapper(url, channel, wg)
	}

	wg.Wait()
	close(channel)

	// This will only work with the waitgroup, and proper closing of channel

	for html := range channel {
		fmt.Println(html)
	}

	// below print would work with out waitgroup

	// for i := 0; i < BufferSize; i++ {
	// 	fmt.Println(<-channel)
	// }

}
