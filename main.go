package main

import (
	"fmt"
	"runtime"
)

func print(amount int, input string) {
	for val := range amount {
		fmt.Println((val + 1), input)
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	go print(5, "Halo")
	print(5, "Apa kabar?")

	var word string
	var count int
	fmt.Scanln(&word)
	fmt.Scanln(&count)
	fmt.Println(word)
	fmt.Println(count)
	for val := range count {
		_ = val
		fmt.Println(word)
	}
}
