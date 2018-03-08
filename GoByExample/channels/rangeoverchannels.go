package main

import "fmt"
import "testing"

func main() {
	c1 := make(chan string, 3)

	c1 <- "Hello"
	c1 <- "how"
	c1 <- "are"

	close(c1)

	for elem := range c1 {
		fmt.Println(elem)
	}
}
