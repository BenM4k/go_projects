package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("index.html")
	if err != nil {
		panic(err)
	}

	oldTime := fileInfo.ModTime()

	for {
		fileInfo, err := os.Stat("index.html")
		if err != nil {
			panic(err)
		}
		newTime := fileInfo.ModTime()

		if newTime.After(oldTime) {
			fmt.Printf("File %v has been modified", fileInfo.Name())
			oldTime = newTime
		}
	}
}
