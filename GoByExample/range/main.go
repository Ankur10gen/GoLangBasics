package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	for index, value := range a {
		fmt.Println(index, value)
	}

	b := map[string]string{"name": "Ankur", "class": "3", "lastname": "Raina"}

	for key, value := range b {
		fmt.Println(key, value)
	}

	for index, value := range "Hello World" {
		fmt.Println(index, value, string(value))
	}
}
