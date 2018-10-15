package main

import (
	"fmt"
)

func main() {
	var a [3]int
	a[0] = 15
	// b := [3]int{12, 30, 21}   // short hand declaration to create array [12 30 21]
	// c := [3]int{11}           // [11 0 0]
	d := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(d)
	/*a := [3]int{5, 78, 8}
	  var b [5]int
	  b = a //not possible since [3]int and [5]int are distinct types*/
	// Arrays in Go are value types and not reference types.
	/*a := [...]string{"USA", "China", "India", "Germany", "France"}
	  b := a // a copy of a is assigned to b
	  b[0] = "Singapore"
	  fmt.Println("a is ", a)
		fmt.Println("b is ", b)*/

}
