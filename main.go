package main

import (
	"fmt"
	"strings"
)

func main() {
	// basic void function
	printMessage("Hai", []string{"John", "Wick"})

	// function with return, with multiple same datatype declaration
	fmt.Println(calculateA(2, 3, 5))

	// multiple return function
	var n, m = calculateB(2, 3, 4, 10)
	fmt.Println(n + 32)
	fmt.Println("Nilainya adalah ", m)

	// variadic function -> spread operators
	fmt.Println(createArray(100, 20, 30, 23, 12, 11))

	// variadic function with slice
	sliceArr := []int64{2, 13, 233, 1222, 323}
	fmt.Println(createArray(sliceArr...))

	// variadic function with array
	arr := [5]int64{2, 3, 4, 5, 12}
	fmt.Println(createArray(arr[:]...))
}
func printMessage(message string, arr []string) {
	nameString := fmt.Sprintf("Your message %s with %s", message, strings.Join(arr, " "))
	fmt.Println(nameString)
}

func calculateA(a, b, c int) float32 {
	return float32(a + b + c)
}

func calculateB(a, b, c, d int) (int32, string) {
	var calculated = a + b + c + d
	var stringVal = fmt.Sprintf("%d", calculated)
	return int32(calculated), stringVal
}

func createArray(input ...int64) int {
	var total int64
	for _, val := range input {
		total += val
	}
	return int(total)
}
