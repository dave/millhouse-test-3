package main

import (
	"fmt"

	"millhouse-test-3/greeting"
	"millhouse-test-3/math"
)

func main() {
	// Print greeting
	fmt.Println(greeting.Greet("World"))

	// Print math results
	fmt.Printf("%d + %d = %d\n", 2, 3, math.Add(2, 3))
	fmt.Printf("%d * %d = %d\n", 4, 5, math.Multiply(4, 5))
}