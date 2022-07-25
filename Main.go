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
	types()

	//Loops
	LoopExample1()

	mySqrt, iterations := Sqrt(2)
	sysSqrt := math.Sqrt(2)
	fmt.Printf("Value of my sqrt is %g with %v iterations and main sqrt is %g", mySqrt, iterations, sysSqrt)
}
