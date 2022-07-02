package main

import (
	"fmt"
)

func LoopExample1() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println("The sum of values from 0-99 is ", sum)
}
