package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	ankur := person{"Ankur", 25}
	fmt.Println(ankur)
	fmt.Println(ankur.name, ankur.age)

	ankit := &person{name: "Ankit", age: 21}
	fmt.Println(ankit)

	ravi := person{age: 26}
	fmt.Println(ravi)

	x := &ravi
	x.name = "ravi"
	fmt.Println(x)
}
