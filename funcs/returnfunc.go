package main

import "fmt"

func makeGreet() func() string {
	return func() string {
		return "Hello World!"
	}
}
func main() {
	greet := makeGreet()
	fmt.Println(greet())
	fmt.Printf("%T \n", greet)
}
