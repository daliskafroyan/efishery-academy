package main

import "fmt"

// multiple return values function
func swap(x, y string) (string, string) {
	return y, x
}

// consecutive name parameters function
func add(x, y int) int {
	return x + y
}

// embedded struct
type person struct {
	name string
	age int
}

type student struct {
	grade int
	person
}

func main() {
	fmt.Printf("hello wrold\n\n")


	/* TYPE DATA */
	// decimal number
	var decimalNumber = 2.231729
	fmt.Printf("Decimal Number %f\n", decimalNumber)
	fmt.Printf("Decimal Number Formatted %.4f\n", decimalNumber)

	// boolean number
	var exist bool = true
	fmt.Printf("exist? %t\n", exist)

	// string
	var message	string= `Hi! wazzup I'm 
	not a robot`
	fmt.Printf(message)
	/* END OF TYPE DATA */


	/* VARIABLES */
	// declare variable without type data
	message1	:= "Ahoy sailor"
	fmt.Printf(message1)

	// declaration multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Printf("Declaration multivariable 1: %s, %s, %s", first, second, third)
	fmt.Printf("Declaration multivariable 2: %s, %s, %s", seventh, eight, ninth)

	// declaration underscore variables
	name, _ := "sukimin", "wiranto"
	fmt.Printf("underscore var", name)

	// constant
	const firstName string = "first name" /* can't declare without type data form */
	/* END OF VARIABLES */


	/* CONDITION */
	// if ... else
	var point = 8

	if point == 10 {
		fmt.Println("perfect")
	} else if point > 7 {
		fmt.Println("biasa aja")
	} else {
		fmt.Println("jelek")
	}

	// switch case
	var pointSwitch = 6

	switch pointSwitch {
		case 1:
			fmt.Println(`not the case`)
		default:
			fmt.Println(`it's the case`)
	}
	/* END OF CONDITION */


	/* LOOPING */
	// range
	var fruits = [4]string{"apple", "fruit", "banana", "lychee"}

	for i, fruit := range fruits {
		fmt.Printf("element %d : %s\n", i, fruit)
	}

	// for loop
	for i:=0; i < 10; i++ {
		if i % 2 == 1 {
			continue
		}
		if i > 5 {
			break
		}
	
		fmt.Println("Angka", i)
	}
	/* END OF LOOPING*/


	/* FUNCTION */
	// invoke function
	a, b := swap("za", "waarudo")
	fmt.Println(a, b)

	num := add(1, 2)
	fmt.Println(num)
	/* END OF FUNCTION */


	/* STRUCT */
	// struct demonstration
	var s1 = student{}
	s1.name = "owi"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name : ", s1.name)

	// combine struct with slice
	var allPersons = []person{
		{name: "Wick", age: 23},
		{name: "Wok", age: 19},
		{name: "Owi", age: 21},
	}
	fmt.Println("all person : ", allPersons)
	/* END OF STRUCT */


	/* DATA STRUCTURE */
	// map
	var chicken map[string]int

	chicken["Januari"] = 50
	chicken["Februari"] = 40

	fmt.Println("Januari", chicken["Januari"])

	// array
	var names [4]string
	names[0] = "trapager"
	names[1] = "lupi"
	
	fmt.Println(names[0], names[1])
	/* END OF DATA STRUCTURE */


	/* SLICE */
	var slice [4]string
	slice[0] = "trapager"
	slice[1] = "lupi"

	fmt.Println(slice[0], slice[1])
	/* END OF SLICE */
}