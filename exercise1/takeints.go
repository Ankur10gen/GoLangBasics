package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Println("Enter one number")
	fmt.Scan(&num1)
	fmt.Println("Enter second number")
	fmt.Scan(&num2)
	fmt.Println("Total is ", num1+num2)
}
