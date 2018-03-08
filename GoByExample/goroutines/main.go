package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printseq(start int, end int, caller string) {
	for i := start; i < end; i++ {
		fmt.Printf("Val %d of caller %s \n", i, caller)
		time.Sleep(time.Duration(rand.Int31n(3)))
	}
}
func main() {
	go printseq(1, 10, "a")
	go printseq(20, 30, "b")
	fmt.Scanln()
	fmt.Println("done")
}
