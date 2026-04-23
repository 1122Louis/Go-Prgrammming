package main

import (
	"errors"
	"fmt"
)

// -------------------- FUNCTION f --------------------

// f takes an integer and returns:
// - result (int)
// - error (if something goes wrong)
func f(arg int) (int, error) {

	// Special case: input is 42
	// We return an error instead of crashing
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	// Normal case: return result + no error
	return arg + 3, nil
}

// -------------------- SENTINEL ERRORS --------------------

// Sentinel errors are predefined error values
// They help us identify specific known problems

var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

// -------------------- FUNCTION makeTea --------------------

// makeTea simulates making tea and may return different errors
func makeTea(arg int) error {

	// Case 1: no tea available
	if arg == 2 {
		return ErrOutOfTea
	}

	// Case 2: power issue (wrapped error)
	if arg == 4 {

		// %w wraps the original error inside new context
		// This preserves the original error for later checking
		return fmt.Errorf("making tea: %w", ErrPower)
	}

	// No error
	return nil
}

// -------------------- MAIN FUNCTION --------------------

func main() {

	// -------------------- TEST FUNCTION f --------------------

	// Loop through test inputs
	for _, i := range []int{7, 42} {

		// Call function and check error in one line (idiomatic Go)
		if r, e := f(i); e != nil {

			// Error case
			fmt.Println("f failed:", e)

		} else {

			// Success case
			fmt.Println("f worked:", r)
		}
	}

	// -------------------- TEST FUNCTION makeTea --------------------

	// Loop through values 0 to 4
	for i := range 5 {

		// Call function
		if err := makeTea(i); err != nil {

			// ---------------- ERROR HANDLING ----------------

			// Check if error matches known sentinel errors
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")

			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark!")

			} else {
				// Unknown error type
				fmt.Printf("unknown error: %s\n", err)
			}

			continue
		}

		// No error case
		fmt.Println("Tea is ready!")
	}
}
