package rectangle

import (
	"fmt"
	"math"
)

/*
 * init function added
 */
func init() {
	fmt.Println("rectangle package initialized")
}

// Area returns area of a rectangle
func Area(length, width float64) float64 {
	area := length * width
	return area
}

// Diagonal returns diagonal of a rectangle
func Diagonal(length, width float64) float64 {
	diagonal := math.Sqrt((length * length) + (width * width))
	return diagonal
}
