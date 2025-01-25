package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, 世界") // Written in Chinese to showcase that Go can handle UTF-8
	fmt.Println("Random number : ", rand.Intn(10))
}
