package main

import "fmt"

func main() {
	var start int
	var end int
	fmt.Println("Enter the range of numbers: First start and then end")
	fmt.Scan(&start)
	fmt.Scan(&end)
	fmt.Println("Below is the list of even numbers")
	for i := start; i < end; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
