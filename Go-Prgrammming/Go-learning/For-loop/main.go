package main

import (
	"fmt"
)

func main() {

	/*for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------------------")

	m := []int{1, 2, 3, 4, 5}

	for i, v := range m {
		fmt.Println(i, v)
	}

	fmt.Println("------------------------")

	for i, v := range "Hello" {
		fmt.Println(i, v)
	}

	fmt.Println("------------------------")

	for i := 0; i < len(m); i++ {
		fmt.Println(i, m[i])
	}

	for _, v := range m {
		fmt.Println(v)
	}*/

	// Print multiplication table for 7
	table := 7
	for i := 1; i <= 10; i++ {
		m := i * table
		fmt.Printf("%d x %d = %d\n", i, table, m)
	}

	// Print even numbers from 0 to 20
	for n := 0; n <= 20; n++ {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		}
	}

	// Print odd numbers from 0 to 20
	for n := 0; n <= 20; n++ {
		if n%2 != 0 {
			fmt.Println(n, "is odd")
		}
	}

	fmt.Println("----------------------------")

	names := []string{"Alice", "Bob", "Charlie", "Diana"}

	for i, name := range names {
		if name == "Charlie" {
			fmt.Println(name, "Found Charlie at index", i)
			break
		}
	}

	for i := 0; i <= 20; i++ {
		if i%3 == 0 {
			fmt.Println("Fix")

		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FixBuzz")
		} else {
			fmt.Println(i)
		}
	}
}
