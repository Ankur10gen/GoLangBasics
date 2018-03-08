package main

import "fmt"

func main() {

	i := 10
	if i/5 == 2 {
		fmt.Println("10/5 == 2")
	}

	if i%2 == 0 {
		fmt.Println("10%2 == 0")
	} else {
		fmt.Println("10%2 != 0")
	}
	if cnt := 10; cnt == 10 {
		fmt.Println("cnt == 10")
	}
}
