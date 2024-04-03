package main

import "fmt"

func main() {
	// urutan deklarasi array -> [nama][[jumlah]][tipe_data]
	var fruits [3]string
	fruits[0] = "Apel"
	fruits[1] = "Jeruk"
	fruits[2] = "Mangga"

	fmt.Println(fruits)

	// deklarasi dan inisialisasi
	cars := [4]string{"Hyundai", "Toyota", "Nissan"}
	fmt.Println(cars)

	// dynamic array tanpa deklarasi panjang array
	hyundai := [...]string{"Palisade", "Creta"}
	fmt.Println(hyundai)

	// multidimensional array
	// -- cara 1
	numbers1 := [3][2]int64{{2, 3}, {2, 3}, {4, 5}}
	fmt.Println("Number 1 ", numbers1)
	// -- cara 2
	// warning redundant declaration
	// numbers2 := [2][3]int{[3]int{1, 2, 4}, [3]int{}}
	// fmt.Println(numbers2)

	// loop array with regular for
	for i := 0; i < len(numbers1); i++ {
		for j := 0; j < len(numbers1[i]); j++ {
			fmt.Print(numbers1[i][j])
		}
		fmt.Println()
	}

	// loop with range
loopStop:
	for _, j := range numbers1 {
		fmt.Println(j)
		for _, i := range j {
			if i == 4 {
				break loopStop
			}
		}
	}
	// allocation array element with "make" keyword
	fruits2 := make([]string, 0)
	fruits3 := [...]string{}
	fmt.Println(fruits2)
	fmt.Println(fruits3)
}
