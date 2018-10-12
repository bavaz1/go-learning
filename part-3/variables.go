package main

import (
	"fmt"
	"math"
)

func main() {
	// var age int = 29 // variable declaration with initial value

	// fmt.Println("my age is", age)

	// var width, height int = 100, 50 //declaring multiple variables

	// fmt.Println("width is", width, "height is", height)

	/* var (
		name   = "naveen"
		age    = 29
		height int
	)
	fmt.Println("my name is", name, ", age is", age, "and height is", height) */

	/* name, age := "naveen", 29 //short hand declaration

	fmt.Println("my name is", name, "age is", age) */

	/* a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c) */

	/* a, b := 20, 30 //a and b declared
	fmt.Println("a is", a, "b is", b)
	a, b := 40, 50 //error, no new variables */

	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("minimum value is ", c)

	/* age := 29      // age is int
	age = "naveen" // error since we are trying to assign a string to a variable of type int */
}
