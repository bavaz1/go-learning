package main

import "fmt"

func main() {
	/* const a = 55 //allowed
	a = 89       //reassignment not allowed */
	/* var a = math.Sqrt(4)   //allowed
	const b = math.Sqrt(4) //not allowed      b is a constant and the value of b needs to be know at compile time. The function math.Sqrt(4) will be evaluated only during run time */

	/* var name = "Sam"
	fmt.Printf("type %T value %v", name, name)
	*/

	/* var defaultName = "Sam" //allowed
	type myString string
	var customName myString = "Sam" //allowed
	customName = defaultName        //not allowed */

	/* const trueConst = true
	type myBool bool
	var defaultBool = trueConst       //allowed
	var customBool myBool = trueConst //allowed
	defaultBool = customBool          //not allowed */

	/* const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var) */

	/* var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("i's type %T, f's type %T, c's type %T", i, f, c) */

	var a = 5.9 / 8
	fmt.Printf("a's type %T value %v", a, a)
}
