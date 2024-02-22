package main

import "fmt"

// To create a function just use func
func main() {
	fmt.Println("Welcome to the quiz game")
	/*


		var name = "Tim" // You can use name := Tim
		name = "Rayan"
		fmt.Println("Hello "+name+", How are you", name) // To add it in a print statement
	*/
	var name string
	fmt.Println("Enter your name: R")
	fmt.Scan(&name)
	fmt.Println("")
	fmt.Println("Hello " + name + ", welcome to the game!")

	fmt.Println("Enter your age")

	var age uint
	fmt.Scan(&age)

	// To make an if loop
	if age >= 10 {
		fmt.Println("You are the correct age to play")
	} else {
		fmt.Println("You cannot play")
	}
	fmt.Println("Quiz is on Geography")

	option := map[string]string{
		"Algeria":      "Morocco", // Also, there was a typo in "Morocco".
		"Italy":        "England",
		"Ireland":      "India",
		"Egypt":        "Russia",
		"South Africa": "Sudan",
	}

	fmt.Println("")
	for option1, option2 := range option {

		fmt.Println(" Which one is bigger")
	}
}
