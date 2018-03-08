package main

import (
	"fmt"
	"github.com/ankurgopher/stringutil"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Printf("%d - %b\n", 42, 42)
	fmt.Printf("%d - %b - %x - %#x - %#X\n", 42, 42, 42, 42, 42)
	for i := 1; i < 100; i++ {
		fmt.Printf("%d %b %x %q\n", i, i, i, i)
	}
}
