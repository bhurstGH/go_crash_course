package main

import "fmt"

func main() {
	// Define map
	// Can also assign and declare simultaneously
	emails := make(map[string]string)
	simulEmail := map[string]string{"One": "one@one.com", "Two": "two@two.com"}

	// Assign key value
	emails["Brett"] = "brett@brett.com"
	emails["Hurst"] = "hurst@hurst.com"
	emails["Hurst2"] = "hurst2@hurst2.com"

	fmt.Println(emails)
	fmt.Println(simulEmail)
	fmt.Println(emails["Brett"])
	fmt.Println(len(emails))

	// Delete
	delete(emails, "Hurst2")
	fmt.Println(emails)
}
