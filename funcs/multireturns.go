package main

import "fmt"

func main() {
	fmt.Println(multiplereturns("Ankur", "Raina"))
}
func multiplereturns(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
