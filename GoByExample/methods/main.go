package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (r *rectangle) perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	rect1 := rectangle{length: 10, width: 12}
	fmt.Println("Area:", rect1.area(), "Perimeter:", rect1.perimeter())
}
