package main

import (
	"fmt"
	"math"
)

func main() {
	const a int = 100
	const b = 200

	var c, d int = 100, 200

	fmt.Println(a, b)
	fmt.Printf("%T %T \n", a, b)
	fmt.Println(a * b)
	fmt.Println(c * d)

	// math.Sin expects float64 value. c cannot satisfy that as it is declared int
	// while b does satisfy this although it earlier showed type as int

	fmt.Println(math.Sin(float64(c)))
	fmt.Println(math.Sin(b))
}
