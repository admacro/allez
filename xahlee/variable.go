// http://xahlee.info/golang/golang_variable.html
package main

import "fmt"

// package level variables are local to the package
// this is also a constant
//   constant must have a value
//   constant values can be string, rune, bool, or numeric values
const package_version = "1.2.3"

// variables inside function are local to the function
func local_func() {
	// inside a function, := is same as var
	local_variable := "hello local"
	fmt.Printf("%v\n", local_variable)
}

func main() {
	// fmt.Printf("%v\n", local_variable) // undefined: local_variable
	fmt.Printf("Package version: %v\n", package_version)

	local_func()

	var name string
	var age int

	// grouping variable declaration
	var (
		height float32; balance float64
		married bool
		titles []string							// ; can be omitted when it's at the end of line
	)

	var address, city string // declare multiple variables of one type

	// in golang, default values have another name: zero value
	// the zero values are:
	//   0 => for numeric types
	//   "" => for string type
	//   false => for bool type
	//   nil => for pointers, functions, interfaces, slices, channels, and maps
	fmt.Printf("%#v\n", name)			// ""
	fmt.Printf("%#v\n", age)			// 0
	fmt.Printf("%#v\n", height)		// 0
	fmt.Printf("%#v\n", balance)	// 0
	fmt.Printf("%#v\n", married)	// false
	fmt.Printf("%#v\n", titles)		// []string(nil)

	fmt.Printf("%#v, %#v\n", address, city) // "", "", ""

	// possible values of various golang types
	// name = nil										// cannot use nil as type string in assignment
	// name = 'a'										// cannot use 'a' (type rune) as type string in assignment
	name = "James"

	// age = 35.1										// constant 35.1 truncated to integer
	age = 35											// valid
	age = 35.0											// valid
	fmt.Printf("%#v\n", age)			// 35

	// floating point numbers
	// check IEEE Standard 754 for more info on floating point numbers
	// https://steve.hollasch.net/cgindex/coding/ieeefloat.html
	// https://en.wikibooks.org/wiki/Floating_Point
	// http://en.wikipedia.org/wiki/Floating-point_arithmetic
	// https://floating-point-gui.de/formats/fp/
	// https://www.cs.cornell.edu/~tomf/notes/cps104/floating.html

	height = 123456.123456
	fmt.Printf("%#v\n", height)		// 123456.125
	fmt.Printf("%e\n", height)		// 1.234561e+05 (%e => scientific notation)
	fmt.Printf("%f\n", height)		// 123456.125000 (decimal point but no exponent)
	fmt.Printf("%g\n", height)		// 123456.125 (%e for large exponents, %f otherwise)

	balance = 1234567890.123456789
	fmt.Printf("%#v\n", balance)		// 1.2345678901234567e+09
	fmt.Printf("%e\n", balance)		// 1.234568e+09
	fmt.Printf("%f\n", balance)		// 1234567890.123457
	fmt.Printf("%g\n", balance)		// 1.2345678901234567e+09

	var arbitrary_number = 3.14
	fmt.Printf("%T\n", arbitrary_number) // float64 (golang infers any float number as float64)

	// parallel assignment
	// but what's more common is to declare multiple variables to
	// take multiple return values of a function
	//   var x, y = f(a)
	address, city = "2000 shoreline CT", "Brisbane"
	fmt.Printf("%#v\n", address)	// "2000 shoreline CT"
	fmt.Printf("%#v\n", city)			// "Brisbane"
}

