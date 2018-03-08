package main

import "fmt"

func main() {
	var x [10]string
	var y [10]byte
	fmt.Println("x =", x, ", len x =", len(x), ", x[5] =", x[5])
	for i := 65; i < 75; i++ {
		x[i-65] = string(i)
		y[i-65] = byte(i)
	}
	fmt.Println("x =", x, ", len x =", len(x), ", x[5] =", x[5])
	fmt.Printf("y = %b", y)
}
