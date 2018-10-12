package main

import (
	"fmt"
)

// GENERAL SYNTAX
/*
func functionname(parametername type) returntype {
 //function body
}

func functionname() {
}
*/

func main() {
	/* price, no := 500, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price: ", totalPrice) */

	/* area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %v Perimeter %v", area, perimeter) */

	// Blank Identifier
	area, _ := rectProps(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f ", area)
}

//Sample Function
func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

//Multiple return values
/* func rectProps(length, width float64) (float64, float64) {
	perimeter := (length + width) * 2
	area := length * width
	return perimeter, area
} */

// Named return values
func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}
