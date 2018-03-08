package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func readAFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print
	file.Close()
}

func main() {
	x := 1
	y := 2
	if x > y {
		fmt.Println("Successful")
	}
	// if x left shifted by 2 is greater than y
	if z := x << 2; z > y {
		fmt.Println("z turned out to be greater. Value is: ", z)
	}

	readAFile("./dummy.txt")
}
