package main

import (
	"fmt"
	"time"
)

func worker(c1 chan string) {
	time.Sleep(time.Second)
	fmt.Println("done")
	c1 <- "Hello World"

}

func main() {
	channel1 := make(chan string)
	go worker(channel1)
	fmt.Println(<-channel1)
}
