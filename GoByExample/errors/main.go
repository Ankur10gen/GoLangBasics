package main

import (
	"errors"
	"fmt"
)

func division(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return -1, errors.New("Division by 0 is not possible")
	}
	return dividend / divisor, nil
}

type argError struct {
	arg  int
	desc string
}

func (a *argError) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.desc)
}

func f1(arg int) (int, error) {
	if arg == 10 {
		return -1, &argError{arg, "You know 10 is not an acceptable value for me."}
	}
	return arg + 10, nil
}

func main() {
	fmt.Println(division(10, 0))
	fmt.Println(division(10, 5))

	a, b := f1(10)
	fmt.Println(a, b)

	fmt.Println(b.(*argError))
}
