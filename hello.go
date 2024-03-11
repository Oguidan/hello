package main

import (
	"flag"
	"fmt"
	"log"
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

// MAaps

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
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

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}

// Constants
type ByteSize float64

const (
	_ = iota // ignore first value by assigning to blank
	identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2FGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

// Variables
var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func init() {
	if user == "" {
		log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
	// gopath may be overridden by --gopath flag on command line.
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

// Methods
// Pointers vs. Values
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
	// Body exactly the same as the Append function defined above.
	return append(slice, data...)
}

// This still requires the method to return the updated slice. We can eliminate that clumsiness by redefining the method to take a pointer to a ByteSlice as its receiver, so the method can overwrite the caller's slice.
// Another version
/*
func (p *ByteSlice) Append(data []byte) {
	slice := *p
	// Body as above, without the return.
	*p = slice
}
*/

// If we modify our function so it looks like a Write method, like this,
func (p *ByteSlice) Write(data []byte) (n int, err error)

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

	var xx uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", xx, xx, int64(xx), int64(xx))

	fmt.Printf("%v\n", timeZone)

	// Append
	// xxx := []int{1, 2, 3}
	// xxx = append(xxx, 4, 5, 6)
	// fmt.Println(xxx)

	xxx := []int{1, 2, 3}
	y := []int{4, 5, 6}
	xxx = append(xxx, y...)
	fmt.Println(xxx)
}
