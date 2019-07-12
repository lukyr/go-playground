/*
  hands-on exercise #4
  1. Create your own type. Have underlying type be an int.
  2. Create a VARIABLE of your new TYPE with IDENTIFIER "x" using he "VAR" keyword
  3. in func main
    a. print out the value of the variable "x"
    b. print out the type of the variable "x"
    c. assign 42 to the VARIABLE "x" using the "=" OPERATOR
    d. print out the value of the variable "X"
*/

package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
