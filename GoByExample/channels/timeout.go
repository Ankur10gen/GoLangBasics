package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "sent message to channel"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received the", msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("Timing out before receiving anything from channel")
	}
}
