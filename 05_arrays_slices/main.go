package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// Slice -- array without fixed length
	fruitSlice := []string{"Apple", "orange", "grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
