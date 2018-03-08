package main

import "fmt"

func main() {
	students := make([][]string, 0)
	student1 := make([]string, 4)
	student1[0] = "Ankur"
	student1[1] = "Raina"
	student1[2] = "94"
	student1[3] = "95"
	student2 := make([]string, 4)
	student2[0] = "Ravi"
	student2[1] = "Kumar"
	student2[2] = "100"
	student2[3] = "99"
	students = append(students, student1)
	students = append(students, student2)
	fmt.Println(students)
}
