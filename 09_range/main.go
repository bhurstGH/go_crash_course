package main

import "fmt"

func main() {
	ids := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}

	// Loop through IDs
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// No index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum:", sum)

	// Range with map
	emails := map[string]string{"one": "one@one.com", "two": "two@two.com"}

	// Sometimes they print in different order....?
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
