package main

import "fmt"

func main() {
	myMap := map[int]string{
		0: "Zero",
		1: "One",
		2: "Two",
		3: "Three",
	}
	for key, val := range myMap {
		fmt.Println("key = ", key, "val = ", val)
	}
}
