package main

import "fmt"

func main() {
	// declaration with variable type [manifest typing]
	var firstName string = "Putu"
	// this will result warning, go compiler is smart enough to tell declaration and initialization value must on the same line
	var lastName string
	lastName = "Bhanuartha"
	fmt.Println(firstName + lastName)

	// declaration without variable type [type inference]
	// --deklarasi ini hanya bisa dilakuka dalam blok fungsi, tidak dalam scope global
	firstName1 := "Putu"
	lastName1 := "Bhanuartha"
	//  print using Printf
	fmt.Printf("Hallo %s %s !, selamat datang\n", firstName1, lastName1)

	// multiple declaration variable
}
