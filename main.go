package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Declaration With Variable Type [manifest typing]
	var firstName string = "Putu"
	// --this will result warning, go compiler is smart enough to tell declaration and initialization value must on the same line
	var lastName string
	lastName = "Bhanuartha"
	fmt.Println(firstName + lastName)

	// Declaration Without Variable Type [type inference]
	// --declaration only on function scope, not in global scope
	firstName1 := "Putu"
	lastName1 := "Bhanuartha"
	//  print using Printf
	fmt.Printf("Hallo %s %s !, selamat datang\n", firstName1, lastName1)

	// Multiple Declaration Variable
	// --using [manifest typing]
	var first, second, third = 12, 23, 12
	fmt.Println(first + second + third)
	// --using [type inference]
	seventh, eight, nine := 12, "8", 9.0
	fmt.Printf("Seventh : %s, eight : %s , nine : %s ", strconv.Itoa(seventh), eight, fmt.Sprintf("%f", nine))

	// Reserved Variable
	// -- just like in javascript, it will ignore unused variable, but in golang, it will throw it away
	var first1, _ = first, second
	fmt.Println(first1)

	// Declare Variable Using "new" Keyword
	stringMem := new(string)
	// --return it's memory address
	fmt.Println("Memory address : " + fmt.Sprintf("%p", stringMem))
	// --dereference
	fmt.Println("Dereference memory : " + *stringMem)

	// Declare Variable Using Keyword "make"
}
