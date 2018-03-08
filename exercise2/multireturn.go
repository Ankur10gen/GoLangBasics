package main

import "fmt"

func main() {
	fmt.Println(multipleReturns(3))
	fmt.Println(multipleReturns(4))

	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	fmt.Println(half(10))
}
func multipleReturns(x int) (int, bool) {
	return x / 2, x%2 == 0
}
