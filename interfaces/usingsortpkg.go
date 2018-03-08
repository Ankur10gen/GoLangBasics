package main

import (
	"fmt"
	"sort"
)

func main() {
	//type people []string
	//studygroup1 := []people{"Ankur", "Abhishek", "Vivek", "Ravi"}
	//	func(p people) Len() int {return len(p)}
	//	func(p people) Swap(i, j int) bool {p[i], p[j] = p[j], p[i]}
	//	func(p people) Less(i, j int) bool {return a[i] < a[j]}

	studygroup := []string{"Ankur", "Abhishek", "Vivek", "Ravi"}
	sort.Strings(studygroup)
	fmt.Println(studygroup[0])
	//sort.Strings(studygroup1)
	//fmt.Println(studygroup1[0])
	sort.Sort(sort.Reverse(sort.StringSlice(studygroup)))

	fmt.Println(studygroup)

}
