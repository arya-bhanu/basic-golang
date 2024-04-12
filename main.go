package main

import (
	"fmt"

)

type student struct {
	name  string
	grade int
}

type person struct {
	name string
}

type employee struct {
	person
	nik string
}

func main() {

	var s1 student
	s1.grade = 2
	s1.name = "Arya"
	fmt.Println(s1)

	var s2 student
	var s3 = student{}
	// result same default value
	fmt.Println(s2)
	fmt.Println(s3)

	// declare and initialization
	var s4 = student{name: "Putu", grade: 20}
	fmt.Println(s4)

	// object pointer variable
	var student1 = student{name: "Putu Gde", grade: 20}
	var student2 *student = &student1
	var student3 = student1

	// object pointer doesn't need '&' to access properties
	// &student2.grade = 30
	student2.grade = 40

	// will not change student1 property value
	student3.grade = 100

	fmt.Println(student1)
	fmt.Println(student3)

	// embedded struct and sub-struct value

	// declare an object then initialize it
	var em1 employee
	em1.person.name = "Arya artha"
	em1.nik = "12310231239123"
	fmt.Println(em1)

	// declare and initialize object at once
	var em2 = employee{nik: "123132123", person: person{name: "Employee 2"}}
	fmt.Println(em2)

	// anonymous struct
	var CEO = struct {
		employee
		salary int32
	}{
		salary: 300_000_000,
		employee: employee{person: person{
			name: "Putu Gde Arya Bhanuartha",
		},
			nik: "2191023809128310923"},
	}

	fmt.Println(CEO)

	var allEmployee = []employee{
		{
			person: person{name: "Employee 1"},
			nik:    "10238129313",
		}, {
			person: person{name: "Employee 3"},
			nik:    "12938109238012",
		},
	}

	for _, val := range allEmployee {
		fmt.Println(val)
	}

}
