package main

import (
	"fmt"
	"github.com/ankurgopher/04_scope/01_package-scope/02_visibility/vis"
)

var i = 0

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
	x := 42
	{
		y := 44
		fmt.Println(x, y)
	}
	//fmt.Println(y)
	for j := 0; j < 10; j++ {
		fmt.Println(increment() + j)
		fmt.Println("i = ", i)
	}

	z := 100

	decrement := func() int {
		z--
		return z
	}
	fmt.Println("Decremented z ", decrement())
	inc := wrapper()
	fmt.Println("Calling Wrapper()", inc())

	fmt.Println("Calling aftereverything part ", aftereverything)

}

func increment() int {
	i++
	return i
}

var aftereverything = "Not sure you can reach here"
