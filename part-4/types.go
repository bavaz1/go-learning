package main

import (
	"fmt"
)

func main() {
	// BOOL

	/* a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d) */

	// INTEGER
	/* var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b */

	// FLOAT
	/* a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2) */

	// COMPLEX
	/* c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul) */

	// STRINGS
	/* first := "Gergo"
	last := "Madar"
	name := first + " " + last
	fmt.Println("My name is", name) */

	// TYPE CONVERSION
	/* i := 55      //int
	j := 67.8    //float64
	sum := i + j //int + float64 not allowed
	fmt.Println(sum) */

	/* i := 55           //int
	j := 67.8         //float64
	sum := i + int(j) //j is converted to int
	fmt.Println(sum) */

	i := 10
	var j float64 = float64(i) //this statement will not work without explicit conversion
	fmt.Println("j", j)
}
