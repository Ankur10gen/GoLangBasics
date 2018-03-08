package main

import "fmt"

func sumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	fmt.Println(sumAndProduct(3, 4))
}
