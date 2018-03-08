package main

import "fmt"

type Square struct {
	side float32
}

func (s Square) area() float32 {
	return s.side * s.side
}

func (s Square) additionalMethod() {
	fmt.Println("I'm useless")
}

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

type Shape interface {
	area() float32
	//additionalMethod() float32 - Subset reqd
}

func areaOfShape(s Shape) {
	//fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println("Shape I got was", s)
}
func main() {
	//x := int(fmt.Scan())
	s := Square{5}
	areaOfShape(s)
	c := Circle{5}
	areaOfShape(c)
	s.additionalMethod()
}
