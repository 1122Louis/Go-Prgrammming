package main

import "fmt"

func main() {
	fmt.Println("=== Pointers in Go ===")

	// Declare an integer variable
	var x int = 42
	fmt.Println("Value of x:", x)

	// Declare a pointer to an integer and assign it the address of x
	var p *int = &x
	fmt.Println("Pointer p points to address:", p)
	fmt.Println("Value at pointer p:", *p)

	// Modify the value of x through the pointer
	*p = 100
	fmt.Println("New value of x after modification through pointer:", x)

	// Declare a pointer to a string
	var str string = "Hello, Go!"
	var strPtr *string = &str
	fmt.Println("Pointer strPtr points to address:", strPtr)
	fmt.Println("Value at pointer strPtr:", *strPtr)

	// Modify the string through the pointer
	*strPtr = "Hello, Pointers!"
	fmt.Println("New value of str after modification through pointer:", str)

	// Demonstrating nil pointers
	var nilPtr *int
	fmt.Println("Nil pointer:", nilPtr)
	if nilPtr == nil {
		fmt.Println("nilPtr is a nil pointer")
	}

	// Using pointers with functions
	fmt.Println("\n=== Pointers with Functions ===")
	a := 10
	fmt.Println("Before function call, a:", a)
	double(&a) // Pass the address of a to the function
	fmt.Println("After function call, a:", a)
}

// Function that takes a pointer to an integer and doubles its value
func double(n *int) {
	*n = *n * 2 // Dereference the pointer and modify the value
}
