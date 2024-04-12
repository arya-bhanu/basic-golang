package main

import (
	"fmt"
	"strings"
)

var funcA = func(a string, b ...int) string {
	counter := 0
	for _, val := range b {
		counter += val
	}
	return fmt.Sprintf("Word : %s\nNumber sum : %d\n", a, counter)
}

var funcMap = func(data []int, callback func(val int, index int) int) []int {
	newVal := data[:]
	for i, val := range newVal {
		newVal[i] = callback(val, i)
	}
	return newVal
}

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
	// array must be converted into slice first
	arr := [5]int64{2, 3, 4, 5, 12}
	fmt.Println(createArray(arr[:]...))

	// closure function
	// assign function into variable so that the function doesn't have it's own name (anonymous)
	// closure only able to be called if it located above the invoker
	fmt.Println(funcA("Hello World", 12, 23, 20))

	// usage of dynamic template string
	var hurufA = "Huruf"
	var angkaA = 23
	var arrA = []int{2, 3, 4, 5}
	fmt.Printf("Huruf : %v\nAngka : %v\nArray: %v\n", hurufA, angkaA, arrA)

	// IIFE
	func() {
		fmt.Println("IIFE Invoked")
	}()

	// mapping function (Function as Arguments)
	var toBeMap = []int{1, 2, 3, 4, 5, 6}
	var resultMapped = funcMap(toBeMap, func(val, index int) int {
		if val > 3 {
			return val * 2
		}
		return val
	})
	fmt.Println(resultMapped)

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
