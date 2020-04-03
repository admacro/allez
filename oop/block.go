// https://golang.org/ref/spec#Blocks
// A block is a possibly empty sequence of declarations
// and statements within matching brace brackets

// No identifier may be declared twice in the same block, and
// no identifier may be declared in both the file and package block

// Explicit blocks: {}, function body
// Implicit blocks in Go: universe > package > file > if/for/switch > switch/select clause
package main

import "fmt"

// The scope of an identifier denoting a constant, type, variable, or function
// (but not method) declared at top level (outside any function) is the package
// block.

// identifiers in package scope: Data, ConstantData, data, dataFunc, choose, main
type Data struct {
	x, y int
}

// method is not in package scope
func (d *Data) DataMethod() { fmt.Println("DataMethod") }

const ConstantData = "Constant Data"

var data Data

var dataFunc = func() {
	fmt.Printf("%#v\n", data)
}

func choose(c int) {
	switch c { // switch block
	case 1:
		{ // case clause block
			func() { // function block
				fmt.Printf("temp func in case-clause block, case: %d\n", c)
				// more blocks can be added here
				// for instance, another explicit block or if/for/switch block
				{
					fmt.Printf("explicit block inside temp func in case-clause block, case: %d\n", c)
				}
			}()
			fmt.Println("1")
			fmt.Println("11")
		}
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}
}

func main() {
	s := "function block of main()"
	fmt.Println(s)

	{ // explicit block
		s = "explicit block inside main()"
		fmt.Println(s)

		f := func() { fmt.Println("temp func inside explicit block") }
		f()
	}

	choose(1)
	choose(3)
}