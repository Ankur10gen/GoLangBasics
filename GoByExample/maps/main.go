package main

import "fmt"

func main() {
	a := make(map[string]int)
	a["age"] = 20
	a["rollno"] = 10

	fmt.Println(a)
	fmt.Println(a["age"])

	_, exists := a["marks"]
	fmt.Println(exists)

	delete(a, "age")
	fmt.Println(a)

	b := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(b)
}
