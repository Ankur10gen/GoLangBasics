package main

import "fmt"

func takeAPointer(ptr *int) { // accepts pointer to the variable
	fmt.Println(ptr)
	*ptr = 0 // changes the value at pointer to 0
}

func takeAValue(val int) {
	val = 0
}
func main() {
	a, b := 1, 1

	takeAPointer(&a) // sends reference of variable a
	fmt.Println(a)
	fmt.Println(&a)

	takeAValue(b)
	fmt.Println(b)
}
