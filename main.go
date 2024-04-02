package main

import "fmt"

func main() {
	var angka = 8
	// golang tidak mendukung 1 baris if - else
	// if angka > 4 fmt.Println("Angka lebih dari 4")

	if angka > 4 {
		fmt.Println("Angka lebih dari 4")
	}

	// variable temporary
	if calculate := (angka / 2); calculate > 3 {
		fmt.Println("Angka temporary lebih dari 3")
	} else {
		fmt.Println("Angka temporary kurang dari sama dengan 3")
	}

	// switch case
	// automatically add "break" statement on each case
	angka = 2
	switch angka {
	case 1:
		fmt.Println("Angka 1")
	case 2:
		fmt.Println("Angka 2")
	default:
		{
			fmt.Println("Bukan angka 1 maupun 2")
		}
	}

	// if - else in switch case form
	angka = 2
	switch {
	case angka > 1:
		fmt.Println("Angka more than 1")
		fmt.Println("=================")
		// remove the default "break" behaviour
		fallthrough
	case angka > 2:
		fmt.Println("Angka more than 2")
		fmt.Println("=================")
		fallthrough
	case angka > 3:
		{
			fmt.Println("Angka more than 3")
			fmt.Println("=================")
		}
	default:
		fmt.Println("None of the above")
	}
}
