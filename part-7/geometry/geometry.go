package main

import (
	"fmt"
	"golangbot/part-7/geometry/rectangle" //importing custom package
	"log"
)

/*
 * 1. package variables
 */
var rectLength, rectWidth float64 = 6, 7

/*
*2. init function to check if length and width are greater than zero
 */
func init() {
	println("main package initialized")
	if rectLength < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLength, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLength, rectWidth))
}
