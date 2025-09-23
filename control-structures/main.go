package main

import "fmt"

func main() {

	age := 32

	if age >= 18 {
		fmt.Println("you are an adult")
	} else if age >= 13 {
		fmt.Println("You are an teenager")
	} else {
		fmt.Println("You are a child")
	}

	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Tuesday", "Wednesday", "Thursday": // multiple cases with same logic
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("TGIF")
	default:
		fmt.Println("it's the weeked")

	}
}

func add(a int, b int) int {
	return a + b
}

func calculateSunAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
