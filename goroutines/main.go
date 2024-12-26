package main

import (
	"fmt"
	"time"
)

func f(t time.Duration) {
	fmt.Printf("f sleep for %ss\n", t)
	time.Sleep(t)
	fmt.Printf("f done sleep for %ss\n", t)
}

func main() {
	go f(2 * time.Second)
	go f(1 * time.Second)

	fmt.Println("wait for 5s")
	time.Sleep(5 * time.Second)
	fmt.Println("done after 5s waiting")
}
