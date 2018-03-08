package main

import "fmt"

func main() {
	b := 44
	fmt.Println("a -", b)
	fmt.Println("a's memory address is ", &b)
	fmt.Printf("%d \n", &b)
}
