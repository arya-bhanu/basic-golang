package main

import (
	"fmt"
	"math"

)

func main() {
	const phi = math.Phi
	// phi = 12000 --> results error
	fmt.Println(phi)

	// multiple constant declaration
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)
	fmt.Printf("Square : %s\nisToday : %t\nnumeric : %d\nfloatNum : %f\n", square, isToday, numeric, floatNum)

	// multiple constant declaration with inherit
	const (
		a = "1"
		b
		c
	)

	fmt.Println(a, b, c)

	// one line multiple constant declaration
	const satu, dua = 1, 2
	fmt.Println(satu + dua)
}
