package main

import "fmt"

func Sequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {

	nextInt := Sequence()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := Sequence()
	fmt.Println(newInts())

	fmt.Println(Sequence()) // returns function address
	fmt.Println(&newInts)
	fmt.Println(&nextInt)
}
