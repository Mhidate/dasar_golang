package main

import "fmt"

func main() {
	// Deklarasi dengan tipe data
	var name string = "jhon"
	var age int = 10

	// Deklarasi tanpa tipe data (Go akan infer tipe data)
	var city = "Bandung"

	fmt.Println(name, age, city)

	//short declaration
	name2 := "Doe"
	age2 := 10
	city2 := "Medan"

	fmt.Println(name2, age2, city2)

	//deklarasi banyak variabel sekaligus
	var (
		firstName3 = "John"
		lastName3  = "Doe"
		age3       = 30
	)

	fmt.Println(firstName3, lastName3, age3)

}
