package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(x ...int) {
	if len(x) > 0 {
		fmt.Println("Here is the list of items: ")
	} else {
		fmt.Println("Nothing found in the list.")
		return
	}
	fmt.Println()
	for _, item := range x {
		fmt.Println(item)
	}
}
