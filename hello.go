package main

import (
	"fmt"
	"os"
)

// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

// Defer function
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

// Switch
func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func main() {
	// Example usage
	a := []byte{1, 2, 3}
	c := []byte{1, 2, 4}
	result := Compare(a, c)
	fmt.Println(result)

	// Defers functions
	b()

	// Arrays
	array := [...]float64{7.0, 8.5, 9.1}
	x := Sum(&array)
	fmt.Println(x)

	// Example for func unhex
	fmt.Println(unhex('A'))

	// Example for func shouldEscape
	fmt.Println(shouldEscape(' '))

	// Printing
	fmt.Printf("Hello %d\n", 23)              // Works as C's printf function
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n") // print of Fprint only add blank between operand when neither side is a string
	fmt.Println("Hello", 23)                  // Println add blank between operand and newline at the end
	fmt.Println(fmt.Sprint("Hello ", 23))     // The print of Sprint only add blank between operand when neither side is a string and the first Ptintln will add a new line
}
