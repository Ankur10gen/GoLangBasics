package main

import "fmt"

func main() {
	fmt.Println(avgnums(12, 13, 14, 15, 16))
	data := []float64{20, 21, 22, 23}
	fmt.Println(avgnums(data...))
}

func avgnums(nums ...float64) float64 {
	fmt.Println(nums)
	fmt.Printf("%T \n", nums)
	var total float64
	for _, v := range nums {
		total = total + v
	}
	return total / float64(len(nums))
}
