package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	case 3:
		fmt.Println("Value is 3")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Looks like a weekend")
	default:
		fmt.Println("Get back to work")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's still morning")
	default:
		fmt.Println("Day is about to end.")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Type is int")
		case bool:
			fmt.Println("Type is bool")
		default:
			fmt.Printf("Unknown type %T \n", i)
		}
	}

	whatAmI(1)
	whatAmI(true)
	whatAmI("Ankur")
}
