package main

import "fmt"

var w = 23

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := "Do you like my hat?"
	g := 'M'
	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
	fmt.Printf("%v %T \n", c, c)
	fmt.Printf("%v %T \n", d, d)
	fmt.Printf("%v %T \n", e, e)
	fmt.Printf("%v %T \n", f, f)
	fmt.Printf("%v %T \n", g, g)

	var i int
	var j string
	var k float64
	var l bool
	fmt.Printf("%v \n", i)
	fmt.Printf("%v \n", j)
	fmt.Printf("%v \n", k)
	fmt.Printf("%v \n", l)

	j = "Hello, I've been assigned now"
	fmt.Println(j)

	var m, n, o int
	m = 1
	fmt.Println(m, n, o)

	var p, q, r int = 1, 2, 3
	fmt.Println(p, q, r)

	var s, t, u = 100, false, "Awesome"
	fmt.Println(s, t, u)
	t = true

	var v string = "Checking..."
	fmt.Println(v)

	foo()
}

func foo() {
	fmt.Println("You just entered foo")
	fmt.Println("The value of w is ", w)
}
