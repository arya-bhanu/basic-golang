package main

import "fmt"

func main() {
	// default is nil, cannot assigned with any key-value
	// it will also return warning, should merge into one line
	// var fruits map[string]int
	// init with emtpy
	// fruits = map[string]int{}

	var fruits = map[string]int{}

	// assignment one by one
	fruits["apple"] = 12000
	fruits["orange"] = 25000

	fmt.Println(fruits)

	// declare and fill with value
	cars := map[string]int{"Hyundai": 12, "Toyota": 30}
	fmt.Println(cars)

	// different type of declaration
	chicken1 := make(map[string]int)
	chicken2 := map[int]int{}
	chicken3 := *new(map[string]string)

	// must be initialized with empty val
	chicken3 = map[string]string{}

	chicken1["banner"] = 3
	chicken2[3] = 10
	chicken3["alpha"] = "alpha"

	chicken := map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, value := range chicken {
		fmt.Printf("Key : %s\nValue : %d\n", key, value)
	}
	fmt.Println(len(chicken))
	delete(chicken, "januari")
	fmt.Println(chicken)

	// detect exitence of item
	var val, isExist = chicken["februari"]
	fmt.Printf("Value : %d\nExist ? : %t\n", val, isExist)

	// slice and map
	var arrMap = []map[string]int{
		{
			"Hero Alpha": 20,
			"Hero Beta":  10,
		},
		{
			"Enemy Alpha": 10,
			"Enemy Beta":  50,
		},
		{
			"King Alpha": 25,
			"King Beta":  30,
		},
	}

	for _, mapItem := range arrMap {
		for key, val := range mapItem {
			fmt.Println("Key : ", key)
			fmt.Println("Val : ", val)
		}
	}
}
