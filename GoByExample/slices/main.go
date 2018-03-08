package main

import "fmt"

func main() {
	a := make([]int, 3)
	fmt.Println(a, len(a), cap(a))

	for i := 0; i < 20; i++ {
		a = append(a, i)
		fmt.Println(a, len(a), cap(a))
	}

	a[22] = 50
	fmt.Println(a, len(a), cap(a))

	// cut a slice
	fmt.Println("Slice from index 2 to 10", a[2:10])

	b := make([]int, 5)

	// copy(target,source)
	copy(b, a[10:15])

	fmt.Println("Copied elements from a in b", b)

	b[1] = 100 // Value in a doesn't change as we created a new copy

	fmt.Println(a, b)

	c := a[15:18]
	c[2] = 150 // Value in a also changes as it is referring to the same elements

	fmt.Println(a, c)

	twoD := make([][]string, 3)
	fmt.Println(twoD) // [[] [] []] inner slice empty
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoD[i] = make([]string, innerLength) // Different lengths of inner slices
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = string(65 + i + j)
		}
	}
	fmt.Println(twoD)

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:] // a slice referencing the storage of x
	fmt.Println(s)
}
