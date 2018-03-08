package main

import "fmt"

func sumOfNums(nums ...int) int {
	if len(nums) == 0 {
		panic("No numbers passed")
	}
	var sum int
	for _, num := range nums { // Note it returns index (here in _ blank variable) and then value
		sum = sum + num
	}
	return sum
}

func main() {
	fmt.Println(sumOfNums(1, 2, 3, 4, 5))
	//fmt.Println(sumOfNums())

	listOfNums := []int{1, 10, 100, 1000, 10000}
	fmt.Println(sumOfNums(listOfNums...))
}
