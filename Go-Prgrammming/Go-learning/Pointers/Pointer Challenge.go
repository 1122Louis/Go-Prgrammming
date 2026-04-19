package main

import "fmt"

// TODO: Define the Employee struct here
// Fields: name (string), id (int), salary (float64)

type Employee struct {
	name   string
	id     int
	salary float64
}

// This function gives a raise to an employee by a certain percentage and takes a pointer to an Employee struct as an argument and percentage as a float64 parameter
func giveRaise(emp *Employee, percentage float64) {
	emp.salary += emp.salary * percentage / 100
}

// This function prints the details of an employee and takes a pointer to an Employee struct as an argument
func printEmployee(emp *Employee) {
	fmt.Printf("Employee Name: %s, ID: %d, Salary: %.2f\n", emp.name, emp.id, emp.salary)
}

// This function swaps the details of two employees and takes pointers to two Employee structs as arguments
func swapEmployees(emp1, emp2 *Employee) {
	emp1.name, emp2.name = emp2.name, emp1.name
	emp1.id, emp2.id = emp2.id, emp1.id
	emp1.salary, emp2.salary = emp2.salary, emp1.salary
}

// This function takes a slice of pointers to Employee structs and returns a pointer to the Employee with the highest salary
func highestPaidEmployee(employees []*Employee) *Employee {
	highest := employees[0]
	for _, emp := range employees {
		if emp.salary > highest.salary {
			highest = emp
		}
	}
	return highest
}

func main() {
	fmt.Println("=== Employee Salary Manager ===\n")

	// TODO: Create 2-3 employee instances

	//This creates three employee instances with different names, IDs, and salaries. You can modify the details as needed.
	// Example: emp1 := Employee{name: "Alice", id: 1, salary: 50000}
	emp1 := Employee{name: "Alice", id: 1, salary: 50000}
	emp2 := Employee{name: "Bob", id: 2, salary: 60000}
	emp3 := Employee{name: "Charles", id: 3, salary: 55000}

	// TODO: Test giveRaise function - give raises to employees
	//This tests the giveRaise function by giving a 10% raise to each employee. You can adjust the percentage as needed.
	// Example: giveRaise(&emp1, 10) // 10% raise
	giveRaise(&emp1, 10)
	giveRaise(&emp2, 10)
	giveRaise(&emp3, 10)

	// TODO: Implement printEmployee function and test it
	//this prints the details of each employee after giving them a raise
	printEmployee(&emp1)
	printEmployee(&emp2)
	printEmployee(&emp3)

	// TODO: Test swapEmployees function
	//This swaps the details of emp1 and emp2, and then prints their details to verify the swap
	swapEmployees(&emp1, &emp2)
	printEmployee(&emp1)
	printEmployee(&emp2)
	printEmployee(&emp3)
	// TODO: Create a slice of employee pointers
	// This slice will hold pointers to the employee instances created earlier. You can create it like this:
	//This prints the details of each employee after swapping their details
	// Example: employees := []*Employee{&emp1, &emp2}
	employees := []*Employee{&emp1, &emp2, &emp3}

	// TODO: Iterate through slice and give everyone 10% bonus
	// This loop iterates through the slice of employee pointers and gives each employee a 10% raise using the giveRaise function. You can adjust the percentage as needed.
	for _, emp := range employees {
		giveRaise(emp, 10)
	}

	// BONUS: Test highestPaidEmployee function
	// This finds the employee with the highest salary using the highestPaidEmployee function and prints their details. You can modify the print statement as needed.
	highestPaid := highestPaidEmployee(employees)
	fmt.Printf("\nHighest Paid Employee:\n")
	printEmployee(highestPaid)

	// BONUS: Try using pointer receiver methods
	// If you implemented the GiveRaise and Print methods as pointer receiver methods, you can test them like this:
	// Example: emp1.GiveRaise(10) // 10% raise using method
	//          emp1.Print() // Print details using method
	emp1.GiveRaise(10)
	emp1.Print()
}

// BONUS: Pointer receiver methods
// GiveRaise method gives an employee a raise
func (e *Employee) GiveRaise(percentage float64) {
	e.salary += e.salary * percentage / 100
}

// Print method prints employee details
func (e *Employee) Print() {
	fmt.Printf("Employee Name: %s, ID: %d, Salary: %.2f\n", e.name, e.id, e.salary)
}

// BONUS: Implement these as pointer receiver methods:
// - (e *Employee) GiveRaise(percentage float64)
// - (e *Employee) Print()
