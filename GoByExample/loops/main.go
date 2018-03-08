package main

import "fmt"

func main() {
	i := 1
	// Type 1 : range only
	for i < 5 {
		fmt.Println("Value of i is", i)
		i++
	}

	// Type 2 : Initialisation, Range, Increment
	for j := 5; j < 10; j++ {
		fmt.Println("Value of j is", j)
	}

	cnt := 1

	// Type 3 : Infinite loop with break statement

	for {
		cnt++
		if cnt > 10 {
			break
		}
		fmt.Println("Counter is now:", cnt)
	}

	// Type 4 : continue statement

	for k := 10; k < 20; k++ {
		if k == 15 {
			fmt.Println("Skipping 15")
			continue
		}
		fmt.Println("Value of k is", k)
	}

}
