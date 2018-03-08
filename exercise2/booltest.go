package main

import "fmt"

func main() {
	if (true && false) || (false && true) || !(false && false) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
