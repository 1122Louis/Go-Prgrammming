package main

import "fmt"

func main() {

	m := []int{}
	fmt.Println(m)

	var s []string
	fmt.Println(s, len(s), cap(s))

	fmt.Println(s, s == nil, len(s) == 0, cap(s) == 0)

	s = make([]string, 3)
	fmt.Println(s, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s[2], s[len(s)-1])

	s = append(s, "d")
	fmt.Println(s, len(s), cap(s))

	s = append(s, "e", "f", "g")
	fmt.Println(s, len(s), cap(s))

	s = append(s, "h", "i", "j", "k", "l", "m", "n", "o", "p")
	fmt.Println(s, len(s), cap(s))

	//Reverse a slice
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)

	/*// T is a placeholder for any type, and []T is a slice of any type.
	func Reverse[T any](s []T) {
		for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j -1 {
			s[i], s[j] = s[j], s[i]
		}
	}(s)
	fmt.Println(s)

	// T is a placeholder for any type, and []T is a slice of any type.
	func Reverse[T any](s []T) {
		for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j -1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	func main() {
		//Works with a slice of strings
		strs := []string{"a", "b", "c", "d", "e"}
		Reverse(strs)
		fmt.Println(strs)

			//Works with a slice of integers too
			ints := []int{1, 2, 3, 4, 5}
			Reverse(ints)
			fmt.Println(ints)
		}
	}*/

	//Copying a slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy of C:", c)

	// Slicing a slice or to get a slice of an element
	fmt.Println("Slice of s from index 2 to 5:", s[2:5])
	fmt.Println("Slice of s from index 0 to 3:", s[:3])

	//Accessing the last element of a slice
	fmt.Println(s[len(s)-1])

	//Accessing the first and last character of a string
	fmt.Println(c[0], c[len(c)-1])

	fmt.Println(c)

	//Multidimensional slices
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Multidimensional slice (matrix):", matrix)

	//Iterating over a multidimensional slice
	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("Element at [%d][%d] = %d\n", i, j, value)
		}
	}

}
