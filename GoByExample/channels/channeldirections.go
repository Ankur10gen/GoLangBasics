package main

import "fmt"

func ping(c1 chan<- string, msg string) {
	c1 <- msg
}
func pong(c1 <-chan string, c2 chan<- string) {
	msg := <-c1
	c2 <- msg
}
func main() {
	pings := make(chan string, 1) // Unbuffered channel created deadlock without goroutine here
	pongs := make(chan string, 1)
	ping(pings, "message passed")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
