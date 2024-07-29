package main

import (
	"fmt"
	"time"
)

func DownloadFile(file string, fileChannel chan<- string) {

	for i := 0; i < 5; i++ {
		if file == "file1" {
			time.Sleep(1 * time.Second)
			downloadProgress := fmt.Sprintf("File: %s downloading... %d%%", file, (20 * (i + 1)))
			fileChannel <- downloadProgress
		}
		if file == "file2" {
			time.Sleep(2 * time.Second)
			downloadProgress := fmt.Sprintf("File: %s downloading... %d%%", file, (20 * (i + 1)))
			fileChannel <- downloadProgress
		}
		if file == "file3" {
			time.Sleep(3 * time.Second)
			downloadProgress := fmt.Sprintf("File: %s downloading... %d%%", file, (20 * (i + 1)))
			fileChannel <- downloadProgress
		}
	}

}

func FileDownloader() {

	arr := []string{"file1", "file2", "file3"}
	BufferSize := 100
	channel := make(chan string, BufferSize)

	for _, v := range arr {
		go DownloadFile(v, channel)
	}

	for i := 0; i < 3*5; i++ {
		outputProgress := <-channel
		fmt.Println(outputProgress)
	}

	// this will cause deadloack
	for outputProgresss := range channel {
		fmt.Print(outputProgresss)
	}
}
