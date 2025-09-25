package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Address struct {
	Street string
	City   string
}

type Contact struct {
	Name    string
	Address Address
	Phone   string
}

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

	// For loop

	for i := 0; i < 5; i++ {
		fmt.Println("This is i", i)
	}

	// while loop does not exist in go

	/**
	* This can be a infinite loop
	* For loop can be conditioned as while loop
	 */

	// counter := 0
	// for counter < 3 {
	// 	fmt.Println("This is the counter", counter)
	// 	counter++
	// }

	iterations := 0
	for {
		if iterations > 3 {
			fmt.Println("Iteration", iterations)
			break
			// continue can also be used based on a condition
		}
		iterations++
	}

	// Arrays and slices

	// arrays

	numbers := [5]int{10, 20, 30, 40, 50} // This tells the compilers the length the array can hold

	// arrays can hold only one data type
	// arrays of the value can be changed by the capacity of the array can't be increased

	numbers[0] = 70

	fmt.Println("This is an array", numbers)
	fmt.Println("This is the length of the array", len(numbers))
	fmt.Println("This is the last value of the array", numbers[len(numbers)-1])

	// this is an array without fixed number of length

	numbersAtInt := [...]int{10, 12, 34}

	fmt.Println("This is an array", numbersAtInt)

	//matrix

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Printf("This is the matrix: %v\n", matrix)

	// arrays are memory efficent

	// slices --- This could be an portion of an array + Dynamic

	// slices are dynamic arrays
	// slices creates a new array to it's capacity when an new element is added to an slice, This is like doubling [1] , [1,2] --> creates new one in the memory

	allNumbers := numbers[:]

	firstThree := numbers[0:3]

	fmt.Println("This is all numbers", allNumbers)
	fmt.Println("This is first three numbers", firstThree)

	// declaring an slice

	fruits := []string{"apple", "bannana", "straberry"}

	fmt.Printf("these are my fruits %v\n", fruits)

	fruits = append(fruits, "kiwi")
	fmt.Printf("these are my fruits with kiwi %v\n", fruits)

	fruits = append(fruits, "mango", "pineapple")

	fmt.Printf("these are my fruits with more fruits %v\n", fruits)

	moreFruits := []string{"blueberries", "tomato"}
	// This appends and slice with another slice
	fruits = append(fruits, moreFruits...)

	fmt.Printf("Appending fruits slice with more fruits slice %v\n", fruits)

	// short form to iterate to all the values in an array or slice
	// for range
	for index, value := range numbers {
		fmt.Printf("index %d  and the value %d\n", index, value)
	}

	//maps

	capitalCities := map[string]string{
		"USA":   "Washington D.C.",
		"India": "New Delhi",
		"UK":    "London",
	}

	// To find if an key exists in a map

	capital, exists := capitalCities["Germany"]

	if exists {
		fmt.Println("This is the capital", capital)
	} else {
		fmt.Println("Does not exist")
	}

	fmt.Println(capitalCities["USA"])

	delete(capitalCities, "UK")

	fmt.Println("This is new deleted map: %v\n", capitalCities)

	// struct -- custom data structures ( solves OOP - Object Oriented Programming)

	person := Person{Name: "John", Age: 25} // added in the top of the code

	fmt.Printf("This is out person %v\n", person)

	fmt.Printf("This is out person %+v\n", person) // adding a + before v can make the output more explict

	/*
	* anonymous struct
	* The struct type is not defined in the top
	 */

	employee := struct {
		name string
		id   int
	}{
		name: "alice",
		id:   123,
	}

	fmt.Printf("This is an anonymous struct % v\n", employee)

	// can avoid having an element inside the struct
	contact := Contact{
		Name: "Marc",
		Address: Address{
			Street: "123 Main Street",
			City:   "Anytown",
		},
		Phone: "123",
	}

	fmt.Printf("Contact %+v\n", contact)

	fmt.Println("Name before", person.Name)
	modifyPersonName(person)
	fmt.Println("name after :", person.Name)

	// adding an & here and * infront of the struct in the function will turn the memory pointer to this to have the value globally

	modifyPersonNameWithPoiner(&person)

	person.modifyPersonNameWithPoinerReceiver("Josh")

	fmt.Println("name after pointer :", person.Name)

	x := 20

	ptr := &x // ability to set address for an variable

	fmt.Printf("value of x: %d and address of x %p\n", x, ptr)

	*ptr = 30 // ability to de-extract the value associated with the memory address and modify the value

	fmt.Printf("value of new x: %d and address of x %p\n", x, ptr)

}

func modifyPersonName(person Person) {
	person.Name = "Mlekey"
	fmt.Println("Inside scope: new name", person.Name)
}

func modifyPersonNameWithPoiner(person *Person) {
	person.Name = "Mlekey"
	fmt.Println("Inside scope: new name", person.Name)
}

// Doing this enables struct to pass methods (receivers) on them

func (person *Person) modifyPersonNameWithPoinerReceiver(name string) {
	person.Name = name
	fmt.Println("Inside scope: new name", person.Name)
}

// note := go dosen't use OOP it uses interface and struct
