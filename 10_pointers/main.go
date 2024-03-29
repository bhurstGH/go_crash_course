package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n%T\n", a, b)

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)

	// Passing address instead of data can be more performant

}
