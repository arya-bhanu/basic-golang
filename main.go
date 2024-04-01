package main

import "fmt"

func main() {
	// non floating point number (int)
	var intA int32 = -1023901931
	// var intB uint32 = 1902301923
	fmt.Printf("Number integer : %d\n", intA)

	// floating point number (float)
	var floatA float32 = 12310231920389123.1203123
	// var floatB float64 = -12031231231231200000123123.1230123123123123
	fmt.Printf("Floating point float 32 : %f\n", floatA)

	// boolean data type
	var exist = true
	fmt.Println(exist)

	// string data type
	var message string = "Hello World"
	fmt.Printf("Your message is : %s\n", message)
}
