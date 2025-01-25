package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}

func main() {
	fmt.Println("Addition:", add(5, 5))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(30))
}
