package main

import "fmt"

type person struct {
	name string
	age  int8
}

func main() {
	var secret interface{} = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	// any keyword
	var anything any
	anything = "Hello World"
	anything = 23
	fmt.Println(anything)

	// arrayOfAny := []interface{}{"Hello", 23, 100}
	// OR
	arrayOfAny := []any{"World", "Hello", 100, 30}
	fmt.Println(arrayOfAny)

	// map of any
	mapOfArray := map[any]interface{}{23: "Hello World", "Hai": 100, "apple": 2500, 10_000: 100_000}
	for key, val := range mapOfArray {
		fmt.Printf("Key: %v, Value: %v\n", key, val)
	}

	// assertion type interface any variable
	var anyType any = 10
	var calculated = anyType.(int) * 100
	fmt.Println(calculated)

	// assertion type interface into pointer object
	var secretAssert any = &person{name: "Putu Arya", age: 12}
	var assertedType = secretAssert.(*person)
	fmt.Println(assertedType)

}
