package firsttry

import "fmt"

func main() {
	x := 10
	y := 20
	fmt.Println(x + y)
	fmt.Printf("%b %b \n", x, y)
	fmt.Printf("%b %b \n", x<<2, y>>2)
	fmt.Println(x<<5 + y>>5)
}
