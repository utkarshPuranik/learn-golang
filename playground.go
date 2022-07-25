package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func function1() {

	fmt.Println("This is test of packages, number is ", rand.Intn(10), " and also value of pi is ", math.Pi)
}

func addition(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func LoopExample1() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println("The sum of values from 0-99 is ", sum)
}

var (
	ToBe     bool   = false
	MaxInt8  byte   = 1<<8 - 1
	MaxInt16 uint16 = 1<<16 - 1
	MaxInt32 uint32 = 1<<32 - 1
	MaxInt64 uint64 = 1<<64 - 1

	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func types() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt8, MaxInt8)
	fmt.Printf("Type: %T Value: %v\n", MaxInt16, MaxInt16)
	fmt.Printf("Type: %T Value: %v\n", MaxInt32, MaxInt32)
	fmt.Printf("Type: %T Value: %v\n", MaxInt64, MaxInt64)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Printf("The value is %v\n", float64(MaxInt16))
}
