package main

import "fmt"

// Variabel global
var globalVar = "Saya bisa diakses di mana saja"

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

	// Variabel lokal
	var localVar = "Saya hanya bisa diakses di dalam main()"

	fmt.Println(globalVar)
	fmt.Println(localVar)

	var a int     // Default: 0
	var b string  // Default: ""
	var c bool    // Default: false
	var d float64 // Default: 0.0

	fmt.Println(a, b, c, d)
}
