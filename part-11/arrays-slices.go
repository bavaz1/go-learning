package main

import (
	"fmt"
)

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	// var a [3]int
	// a[0] = 15
	// b := [3]int{12, 30, 21}   // short hand declaration to create array [12 30 21]
	// c := [3]int{11}           // [11 0 0]
	// d := [...]int{12, 78, 50} // ... makes the compiler determine the length
	// fmt.Println(d)
	/*a := [3]int{5, 78, 8}
	  var b [5]int
	  b = a //not possible since [3]int and [5]int are distinct types*/

	// Arrays in Go are value types and not reference types.
	/*a := [...]string{"USA", "China", "India", "Germany", "France"}
	  b := a // a copy of a is assigned to b
	  b[0] = "Singapore"
	  fmt.Println("a is ", a)
		fmt.Println("b is ", b)*/

	// Local change
	/* num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num) */

	// Length of an array
	/* a := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(a)) */

	// Iterating arrays using length
	/* a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	} */

	// Iterating arrays using range
	/* a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum) */

	// Ignore index
	/* a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	fmt.Printf("The values: ")
	for _, v := range a { //ignores index
		fmt.Printf(" %.2f", v)
		sum += v
	}
	fmt.Println()
	fmt.Println("The sum of values: ", sum) */

	// Multidimensional arrays
	/* a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b) */

	// Creating a slice
	/* a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b) */

	/* c := []int{6, 7, 8} //creates an array and returns a slice reference
	fmt.Println(c) */

	// Slice is just a representation of the underlying array.
	/* darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr) */

	/* numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa) */
}
