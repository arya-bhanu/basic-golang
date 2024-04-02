package main

import (
	"fmt"
	"strconv"
)

func main() {
	// perulangan di golang hanya for, tidak ada while, foreach, namun bisa membentuk cara kerja yang sama
	// standar perulangan
	for i := 0; i < 10; i++ {
		fmt.Println("Standar looping : ", i)
	}

	// "while" like looping
	angka := 1
	for angka < 5 {
		fmt.Println("While like looping ", angka)
		angka++
	}

	// for without any argument
	for {
		fmt.Println("For loop without argument : ", angka)
		if angka == 0 {
			break
		}
		angka--
	}

	// for looping with range
	const hello = "hello"
	for i, j := range hello {
		fmt.Printf("Indeks : %d, Character : %s\n", i, strconv.QuoteRune(j))
	}

	// for looping with label
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

}
