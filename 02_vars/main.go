package main

import "fmt"

// Var
var name string = "Go"
var age = 36

// Const
const isConst = true

func main() {

	// Shorthand
	shorthand := "Only in a function"
	multi1, multi2 := "One", "Two"
	size := 1.3

	fmt.Println(name, age, shorthand, multi1, multi2)
	fmt.Printf("%T\n%T", name, size)
}
