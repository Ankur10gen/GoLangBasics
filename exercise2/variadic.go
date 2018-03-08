package main

import "fmt"

func main() {
	greatest := findGreatest(7, 1, 2, 3, 4, 6, 6)
	fmt.Println(greatest)
}
func findGreatest(x ...int) int {
	var max int
	for _, item := range x {
		if item > max {
			max = item
		}
		//fmt.Println("max ", max)
	}
	//fmt.Println(max)
	return max
}
