package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(mySlice[1:3])
	fmt.Println(mySlice[2])
	// Not working: fmt.Println(mySlice[-1:3])
}
