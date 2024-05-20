package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
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
func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	// Again as above.
	*p = slice
	return len(data), nil
}

// Interfaces and other types
// Interfaces

type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
	s = s.Copy() // Make a copy; don't overwrite argument.
	sort.Sort(s)
	str := "["
	for i, elem := range s { // Loop is O(N²); will fix that in next example.
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

// Conversions
/*
func (s Sequence) String() string {
	s = s.Copy()
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}
*/
// Method for printing - sorts the elements before printing.
/*
func (s Sequence) String() string {
	s = s.Copy()
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}
*/

// Interface conversions and type assertions
func test01() string {
	type Stringer interface {
		String() string
	}

	var value interface{} // Value provided by caller.
	switch str := value.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}

	str, ok := value.(string)
	if ok {
		fmt.Printf("string value is: %q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}

	if str, ok := value.(string); ok {
		return str
	} else if str, ok := value.(Stringer); ok {
		return str.String()
	}
	return ""
}

// The crypto/cipher interfaces look like this:
type Block interface {
	BlockSize() int
	Encrypt(dst, src []byte)
	Decrypt(dst, src []byte)
}

type Stream interface {
	XORKeyStream(dst, src []byte)
}

// Interface and methods

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/

/* Simple counter server.
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)
*/

// Simple counter server.
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

// A channel that sends a notification on each visit.
// (Probaly want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprint(w, "notification sent")
}

func ArgServer() {
	fmt.Println(os.Args)
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.

// type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, req).

/*
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
	f(w, req)
}
*/

// Argument server.

/*
func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fpriintln(w, os.Args)
}
*/

// AgrServer can be converted to a HandlerFunc.

// http.Handle("/args", http.HandlerFunc(ArgServer))

// The blank identififer

// The blank identifier in multiple assignment
/*
if _, err := os.Stat(path); os.IsNotExist(err) {
	fmt.Printf("%s does not exist\n", path)
}
*/

// Bad! This code will crash if path does not exist.
/*
fi, _ := os.Stat(path)
if fi.IsDir() {
	fmt.Printf("%s is a directory\n", path)
}
*/

// Unused imports and variables

/*
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fd, err := os.Open("test.go")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: use fd.
}
*/

// To silence complaints about the unused imports, use a blank identifier to refer to a symbol from the imported package.

/*
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader // For debugging: delete when done.

func main() {
	fd, err := os.Open("test.go")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: use fd.
	_ = fd
}
*/

// Import for side effect

// To import the package only for its side effects, rename the package to the blank identifier

// import _ "net/http/pprof"

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

	// Pointers vs. Values
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
}
