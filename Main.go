package main

import (
	"fmt"
	"math"
)

func main() {
	// Print Hello
	fmt.Println("Hello, World!")

	//Swap - multiple variable return
	a, b := swap("Utkarsh", "Puranik")
	fmt.Println(a, b)

	//Data Types
	fmt.Println("============ Testing Data types ============")
	types()

	//Loops
	fmt.Println("============ Testing Loops ============")
	LoopExample1()

	//Switch
	fmt.Println("============ Testing Switch ============")
	switchTest()

	//Defer
	fmt.Println("============ Testing Defer ============")
	TestDefers()

	fmt.Println("============ Square Root Excercise ============")
	mySqrt, iterations := Sqrt(2)
	sysSqrt := math.Sqrt(2)
	fmt.Printf("Value of my sqrt is %g with %v iterations and main sqrt is %g", mySqrt, iterations, sysSqrt)
}
