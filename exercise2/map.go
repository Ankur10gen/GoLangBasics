package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 100
	fmt.Println("map: ", m)
	a, b := m["k1"]
	fmt.Println(a, b)
	c, d := m["k3"]
	fmt.Println(c, d)
	if val, exists := m["k2"]; exists {
		delete(m, "k2")
		fmt.Println("Deleted value ", val)
	}
	fmt.Println(m)
}
