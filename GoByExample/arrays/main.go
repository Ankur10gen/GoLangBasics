package main

import "fmt"

func main() {
	var a [5]int

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var c [3][2]int
	fmt.Println(c)

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = i * j
		}
	}

	fmt.Println(c)
}
