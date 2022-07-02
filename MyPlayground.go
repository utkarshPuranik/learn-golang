package main

import (
	"fmt"
	"math"
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
