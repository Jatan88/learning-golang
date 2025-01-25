package main

import (
	"fmt"
	"math/cmplx"
)

// Datatypes:
// 1. Basic -> number, string, bool
// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte -> alias for uint8

// rune -> alias for int32
// 	    -> represents a Unicode code point

// float32 float64

// complex64 complex128

// 2. Aggregate -> array, struct
// 3. Reference -> pointers, maps, channels, slices, functions

var a, b, c int

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	s      string
)

const cc = 4 // constant should be initialized

func main() {
	a = 4
	b = 5
	c = 6
	var d int = 7
	e := 8 // Can't be declared outside of a function
	const dd = 8
	//dd = 10 // cannot assign as the value is constant
	fmt.Println(a, b, c, d, e)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %q\n", s, s)
	fmt.Printf("Type: %T Value: %v\n", cc, cc)
	fmt.Printf("Type: %T Value: %v\n", dd, dd)
}
