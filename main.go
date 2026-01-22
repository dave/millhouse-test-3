package main

import "fmt"

func main() {
	// Print greeting
	fmt.Println(Greet("World"))

	// Print math results
	fmt.Printf("%d + %d = %d\n", 2, 3, Add(2, 3))
	fmt.Printf("%d * %d = %d\n", 4, 5, Multiply(4, 5))
}