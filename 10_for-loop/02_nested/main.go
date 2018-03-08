package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println("i * j = ", i*j)
		}
	}
	x := 1
	for x < 5 {
		fmt.Println("Hello ", x)
		x++
	}
	y := 1
	for {
		y++
		if y%2 == 1 {
			continue
		}
		fmt.Println("value of y is ", y)
		if y > 10 {
			break
		}
	}
	for j := 1; j < 340; j++ {
		fmt.Println(j, "-", string(j), "-", []byte(string(j)))
		fmt.Println(j, "-", string(j), "-", []int32(string(j)))
	}
}
