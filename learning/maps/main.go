package main

import "fmt"

func main() {
	// Create an empty map

	m := make(map[string]int)
	fmt.Println(m)

	// Add some key-value pairs to the map
	m["Alice"] = 25
	m["Bob"] = 30
	m["Charlie"] = 35
	fmt.Println(m)

	// Access values in the map using keys
	fmt.Println("Alice's age:", m["Alice"])
	fmt.Println("Bob's age:", m["Bob"])
	fmt.Println(m["Charlie"])

	// Check if a key exists in the map
	age, exists := m["Diana"]
	if exists {
		fmt.Println("Diana's age:", age)
	} else {
		fmt.Println("Diana not found in the map")
	}

	age, exists = m["Alice"]
	if exists {
		fmt.Println("Alice's age is:", age)
	} else {
		fmt.Println("Alice not found in the map")
	}

	// Delete a key-value pair from the map
	delete(m, "Bob")
	fmt.Println(m)

	// Iterate over the map
	for name, age := range m {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Get the number of key-value pairs in the map
	fmt.Println("Number of people in the map:", len(m))

	// Clear the map by reinitializing it
	clear(m)
	fmt.Println("Map after clearing:", m)

	// Print only the values in the map
	for _, age := range m {
		fmt.Println("Age:", age)
	}

	// Print only the keys in the map
	for name := range m {
		fmt.Println("Name:", name)
	}

	_, prs := m["Alice"]
	fmt.Println("Is Alice in the map?", prs)
}
