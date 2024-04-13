package main

import (
	"fmt"
	"reflect"

)

func main() {
	var number = 23
	var reflectValueNumber = reflect.ValueOf(number)
	if reflectValueNumber.Kind() == reflect.Int {
		fmt.Println("tipe variable: ", reflectValueNumber.Type())
		fmt.Println("nilai variable: ", reflectValueNumber.Int())
	}

	
}
