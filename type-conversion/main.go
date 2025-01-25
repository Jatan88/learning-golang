package main

import "fmt"

func main() {
	f := 3.3
	i := int(f)

	var ff float32 = 3.3333
	var ii int32 = int32(ff)
	fmt.Println(i)
	fmt.Println(ii)
}
