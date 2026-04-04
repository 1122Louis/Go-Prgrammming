package main

import (
	"fmt"
)

/*func SquareAdd(p *int) {
	*p += *p * *p
	fmt.Println(p, *p)
}
func squareVal(v int) {
	v *= v
	fmt.Println(&v, v)
}
func main() {
	/*i, j := 42, 2701
	fmt.Println(&i, &j)

	p := &i
	q := &j

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(q)
	fmt.Println(*q)

	*p = 21
	fmt.Println(*p)

	*q = *q / 37
	fmt.Println(*q)*/

/*i := 3
	p := &i
	q := &i

	fmt.Println(&p, &q)

	squareVal(*q)
	fmt.Println(*q)
	SquareAdd(p)
	fmt.Println(*p)
}*/

func main() {

	var myNumber = 40

	var ptr = &myNumber

	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of myNumber is ", *ptr)

	*ptr = 50

	fmt.Println("Value of myNumber after change is ", myNumber)

	*ptr = *ptr + 10

	fmt.Println("Value of myNumber after addition is ", myNumber)

}
