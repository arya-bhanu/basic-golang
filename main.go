package main

import "fmt"

func main() {
	var memA *int
	var numberA = 4
	var numberB = 4

	memA = &numberA
	*memA = 20

	fmt.Println(&numberA)
	fmt.Println(&numberB)
	fmt.Println(*memA)

	numberA = 10
	fmt.Println(*memA)

	// change value with function
	var number = 4
	fmt.Println(number)
	fmt.Println(&number)
	change(&number, 10)
	fmt.Println(number)
	fmt.Println(&number)
}

func change(address *int, number int) {
	*address = number
}
