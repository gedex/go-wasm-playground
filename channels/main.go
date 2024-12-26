package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)

	go func() {
		time.Sleep(1 * time.Second)
		messages <- "message #1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		messages <- "message #2"
	}()

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
