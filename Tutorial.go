package main

// I am aware that Map is unordered just too lazy to fix it currently
import (
	"fmt"
	"strconv"
)

// To create a function just use func
func main() {
	fmt.Println("Welcome to the quiz game")
	/*


		var name = "Tim" // You can use name := Tim
		name = "Rayan"
		fmt.Println("Hello "+name+", How are you", name) // To add it in a print statement
	*/
	var name string
	var answer string
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("")
	fmt.Println("Hello " + name + ", welcome to the game!")

	fmt.Println("Enter your age")

	var age uint
	rightAnswer := 0
	arr := [5]string{"Algeria", "Italy", " India", " Russia", " South Africa"}

	fmt.Scan(&age)
	fmt.Println("")
	// To make an if loop
	if age >= 10 {
		fmt.Println("You are the correct age to play")
	} else {
		fmt.Println("You cannot play")
	}
	fmt.Println("Quiz is on Geography\n")

	option := map[string]string{
		"Algeria":      "Morocco",
		"Italy":        "England",
		"Ireland":      "India",
		"Egypt":        "Russia",
		"South Africa": "Sudan",
	}

	fmt.Println("Quiz is starting\n")
	num := 0
	for option1, option2 := range option {

		fmt.Println(" Which one is bigger "+option1+" or ", option2+"\n")
		fmt.Scan(&answer)
		if answer == arr[num] {
			fmt.Println("You got the correct answer" + "\n")
			rightAnswer++
		} else {
			fmt.Println(" Sorry, you got the wrong answer" + "\n")
		}

		num++

	}

	Finalscore := rightAnswer / 5
	fmt.Println("The total score you got is " + strconv.Itoa(Finalscore) + "\n")

}
