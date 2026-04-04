package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go" + " is " + "awesome!")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7.0 / 3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a string = "Hello"
	fmt.Println(a)

	var b, c int = 2, 4
	fmt.Println(b, c)

	var d int
	d = 10
	fmt.Println(d)

	var e = true
	fmt.Println(e)

	const s string = "Constant value"
	fmt.Println(s)

	i := 1
	for i <= 10 {
		fmt.Print(i)
		i++
	}

	j := 0

	for j = 0; j <= 8; j++ {
		fmt.Println(j)

	}
	fmt.Println("----------------------------")

	for i := range 5 {
		fmt.Println(i)
	}

	for {
		fmt.Println("Infinite loop")
		fmt.Print(len("This is just the beginning"))
		break
	}

	for n := range 10 {
		if n%2 == 0 {
			fmt.Println(n, "is even")
			continue
		} else {
			fmt.Println(n, "is odd")
		}
	}

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Print("8 is divisible by 4/n")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "Num is Negative")
	} else if num < 10 {
		fmt.Println(num, "The is positive single digit number")
	} else {
		fmt.Println("The num is a multiple digit number")
	}

	x := 2
	fmt.Println("Write", x, "as")
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	switch time.Now().Weekday() {
	case time.Monday, time.Thursday:
		fmt.Println("It is a weekend")
	default:
		fmt.Println("It is a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	default:
		fmt.Println("Good afternoon")
	}

}
