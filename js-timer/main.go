package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start 2s timer")
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Println("2s has passed")
}
