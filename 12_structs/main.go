package main

import (
	"fmt"
	"strconv"
)

// Person ... Define person struct
// Like classes
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// Shorthand: firstName, lastName, city, gender string
	// Shorthand: age in
}

// Greeting method (value reciever)
func (p Person) greet() string {
	// Mismatched data types in return cause error -- enter strconv()
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "F" {
		return
	}

	p.lastName = spouseLastName

}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Brett", lastName: "Hurst", city: "Place", gender: "M", age: 36}
	// No prop name
	person2 := Person{"Bre", "Hur", "Place", "M", 36}

	fmt.Println(person1, person2)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.age)
	fmt.Println(person1.greet())
	person1.getMarried("Hurts")
	fmt.Println(person1.greet())
}
