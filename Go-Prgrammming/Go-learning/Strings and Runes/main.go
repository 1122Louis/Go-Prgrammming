package main

import (
	"fmt"
	"unicode/utf8"
)

/*
A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of "characters". In Go, the concept of a character is called a rune - it's an integer that represents a Unicode code point. This Go blog post is a good introduction to the topic.
*/

func main() {
	fmt.Println("=== Original Examples ===")
	s := "Hello, 世界"
	fmt.Println("String:", s)

	// Length of the string in bytes
	fmt.Println("Length in bytes:", len(s))

	fmt.Println("Printing character at index 7:", string(s[7]))

	for i, r := range s {
		fmt.Printf("Index: %d, Rune: %c\n", i, r)
	}

	// Iterate over the string using range (handles Unicode correctly)
	fmt.Println("Iterating over string with range:")
	for i, r := range s {
		fmt.Printf("Index: %d, Rune: %c\n", i, r)
	}

	// Convert string to a slice of runes to access individual characters
	runes := []rune(s)
	fmt.Println("Runes:", runes)

	// Accessing specific runes
	fmt.Printf("First rune: %c\n", runes[0])           // 'H'
	fmt.Printf("Last rune: %c\n", runes[len(runes)-1]) // '界'

	c := "Hello"
	fmt.Println("First byte of string c:", c[0]) // Output: 'H' (the first byte of the string)

	b := c[0]                                             // This is a byte, not a rune
	fmt.Printf("First byte of string c as byte: %d\n", b) // Output: 72 (ASCII value of 'H')
	fmt.Println(string(b))
	fmt.Printf("First byte of string c as character: %c\n", b) // Output: 'H'

	fmt.Println("\n=== Thai Language UTF-8 Examples ===")
	// s is a string assigned a literal value representing the word "hello" in the Thai language. Go string literals are UTF-8 encoded text.
	const thaiString = "สวัสดี"

	// Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
	fmt.Println("Len:", len(thaiString))

	// Indexing into a string produces the raw byte values at each index. This loop generates the hex values of all the bytes that constitute the code points in thaiString.
	for i := 0; i < len(thaiString); i++ {
		fmt.Printf("%x ", thaiString[i])
	}
	fmt.Println()

	// To count how many runes are in a string, we can use the utf8 package. Note that the run-time of RuneCountInString depends on the size of the string, because it has to decode each UTF-8 rune sequentially. Some Thai characters are represented by UTF-8 code points that can span multiple bytes, so the result of this count may be surprising.
	fmt.Println("Rune count:", utf8.RuneCountInString(thaiString))

	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	for idx, runeValue := range thaiString {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// We can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(thaiString); i += w {
		runeValue, width := utf8.DecodeRuneInString(thaiString[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// This demonstrates passing a rune value to a function.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// Values enclosed in single quotes are rune literals. We can compare a rune value to a rune literal directly.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}

	//CODING CHALLENGE
	//Write a program that:
	// Stores a string containing at one Unicode character that is represented by multiple bytes in UTF-8 (e.g., a Chinese character, an emoji, etc.).
	//Prints the length of the string in bytes using len().
	//Iterates over the string using a for loop with an index to print each byte in hexadecimal format.
	//Iterates over the string using a range loop to print each rune and its corresponding index.
	//Print each rune as a character.

	h := "Louis, 你好" // A string containing a Chinese character, which is represented by multiple bytes in UTF-8.
	fmt.Println("Length of string in bytes........:", len(h))
	fmt.Println("Iterating over string with index:")
	for i := 0; i < len(h); i++ {
		fmt.Printf("%x ", h[i])
	}
	fmt.Println()
	fmt.Println("Iterating over string with range:")
	for i, r := range h {
		fmt.Printf("Index: %d, Rune: %c\n", i, r)
	}

	// Print each rune as a character
	fmt.Println("Printing each rune as a character:")
	for _, r := range h {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
}
