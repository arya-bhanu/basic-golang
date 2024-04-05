package main

import "fmt"

func main() {
	// slice adalah bagian dari array
	// tiap nilai dari array memiliki alamat (reference) dan itu dapat dikatakan sebagai slice

	fmt.Println("Hello World")
	// it's a slice
	sliceA := []string{"Hello", "World"}
	fmt.Println(sliceA)

	fruits := [5]string{"Apple", "Mango", "Grape", "Banana", "Melon"}
	fruitsSlice1 := fruits[1:3]
	fruitsSlice2 := fruits[1:]
	fruitsSlice3 := fruitsSlice2[0:2]

	// change value reference slice
	fruitsSlice3[1] = "Orange"

	fmt.Println(fruitsSlice1)
	fmt.Println(fruitsSlice2)
	fmt.Println(fruitsSlice3)

	// len () -> menghitung panjang element
	fmt.Println(len(fruits))

	// cap () -> menghitung kapasitas maksimum slice
	fmt.Println("Length of fruits : ", len(fruits))
	fmt.Println("Cap of fruits : ", cap(fruits))

	fruitsCap1 := fruits[0:3]
	fmt.Println("Length of fruitsCap1 : ", len(fruitsCap1))
	// cap ketika diambil dari index = 0, cap masih sama dengan awal
	fmt.Println("Cap of fruitsCap1 : ", cap(fruitsCap1))

	fruitsCap2 := fruits[1:4]
	fmt.Println("Length of fruitsCap2 : ", len(fruitsCap2))
	// cap ketika diambil dari index > 0, cap akan berkurang dari panjang awal - index start
	fmt.Println("Cap of fruitsCap2 : ", cap(fruitsCap2))

	// append ()
	// Ketika jumlah elemen dan lebar kapasitas adalah sama (len(fruits) == cap(fruits)), maka elemen baru hasil append() merupakan referensi baru
	cars := []string{"Honda", "Hyundai"}

	// cap carsA === len carsA
	carsA := cars[0:]
	carsC := append(carsA, "Jeep")

	fmt.Println(cap(cars))
	fmt.Println(cap(carsA))
	// kapasitasnya double, menambahkan reference baru
	fmt.Println(cap(carsC))

	fmt.Println(cars)
	fmt.Println(carsA)
	fmt.Println(carsC)

	// Ketika jumlah elemen lebih kecil dibanding kapasitas (len(fruits) < cap(fruits)), elemen baru tersebut ditempatkan ke dalam cakupan kapasitas, menjadikan semua elemen slice lain yang referensi-nya sama akan berubah nilainya.
	animals := []string{"tiger", "cat", "elephant"}
	animalsA := animals[0:2]

	// len < cap
	fmt.Println("len animalsA : ", len(animalsA))
	fmt.Println("cap animalsA : ", cap(animalsA))

	animalsC := append(animalsA, "lion")
	animalsD := append(animalsC, "Mongoose")

	fmt.Println(animals)
	fmt.Println(animalsA)
	fmt.Println(animalsC)

	// animals D add their new chunk of mem
	fmt.Println(animalsD)
	fmt.Println(cap(animalsC))

	// copy ()
	// declare an array
	var dstFruits [2]string
	src := []string{"watermelon", "pinnaple", "apple", "orange"}

	// error, must be a slice, turn it into slice
	// copy(dstFruits, src)
	n := copy(dstFruits[:], src)
	fmt.Printf("Coppied %d element\n", n)
	fmt.Println(dstFruits)

	// make a slice using "make" keyword
	var dstCars = make([]string, 3)
	// it's not an array, must convert into slice also
	var srcCars = [5]string{"Cars 1", "Cars 2", "Cars 3", "Cars 4", "Cars 5"}
	x := copy(dstCars, srcCars[:])

	fmt.Println(x)
	fmt.Println(dstCars)

}
