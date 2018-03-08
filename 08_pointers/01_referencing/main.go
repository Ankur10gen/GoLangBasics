package main

import "fmt"

func main() {
	a := "hello"
	fmt.Println(a)
	fmt.Println(&a)
	var b = &a
	fmt.Println(b)
	fmt.Println(*b)
	*b = "yellow"
	fmt.Println(b)
	fmt.Println(a)
}
