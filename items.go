package main

var items = []string{"Kapak", "Palu"}

func CountItems() int {
	return len(items)
}

func countKeliling(panjang int, lebar int) int {
	return 2*panjang + 2*lebar
}
