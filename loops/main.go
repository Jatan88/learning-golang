package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// There is no loop other than for loop in Go
	// But if you want to implement while loop, check below

	sum1 := 0
	for sum1 < 10 {
		sum1++
	}

	fmt.Println(sum1)
}
