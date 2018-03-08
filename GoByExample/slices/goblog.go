package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	//fmt.Println(string(b))
	b = digitRegexp.Find(b)
	fmt.Println(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// Exercise to compact the above function using append - unable to think of a way
func CopyDigitsModified(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func main() {
	fmt.Println([]byte(CopyDigits("./dummy.txt")))

}
