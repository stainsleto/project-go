package main // compiler searches for this package to execute the program

// Function in go

import (
	"fmt"
)

func main() {
	fmt.Println("First line in Go")

	// variables

	var name string = "Stains"
	fmt.Printf("This is my name %s\n", name) // This printf prints with a format specifier

	age := 13. // declaring a variable without specifing the type
	fmt.Printf("This is my age %d\n", age)

	var city string // Zero Value
	city = "Seatle"

	fmt.Printf("This is my city %s\n", city)

	var country, continent string = "USA", "North America"
	fmt.Printf("This is my country %s This is my continent %s\n", country, continent)

	// way to declare different variables in a single line

	var (
		isEmployed bool   = true
		salary     int    = 50000
		position   string = "developer"
	)

	fmt.Printf("is employed : %t this is my salary %d and this is my position: %s\n", isEmployed, salary, position)

	//Zero Values
	var defaultint int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("This is the default int %d this is the default float %f this is the defaultString %s This is the default boolean %t\n", defaultint, defaultFloat, defaultString, defaultBool)

	const pi = 3.14 // constant

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	fmt.Printf("Monday %d - Tuesday %d - Wednesday %d\n", Monday, Tuesday, Wednesday)

	const typedAge int = 10
	const untypedAge = 10

	fmt.Printf("This is typed age %d - This is un-typed age %d\n", typedAge, untypedAge)

	// Surface level ENUMS

	const (
		Jan   int = iota + 1. //1.  --> iota helps in incremental values
		Feb                   //2
		March                 //3
		Apr                   //4
	)

	fmt.Printf("jan - %d feb - %d mar - %d apr - %d\n", Jan, Feb, March, Apr)
	result := add(3, 4)

	fmt.Println("This is the result", result)

	sum, product, name, value := calculateSumAndProduct(10, 10, 10.4, "test")

	fmt.Printf("This is sum %d - This is product %d - This is name %s - this is value %f\n", sum, product, name, value)

}

// to run the code --> go run main.go

// Note ---> Package main can import from other package but can't export (main package can't be imported by other packages)

// go build -o main main.go ---> This helps to create a binary file (build version)
// just run ./main

// functions -- first letter of function is neeed to be capital if we need to export it

func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int, float float64, name string) (int, int, string, float64) { // if a function is going to return more than one
	return a + b, a * b, name, float
}
