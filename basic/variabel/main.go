package main

import "fmt"

func main() {
	// Deklarasi dengan tipe data
	var name string = "jhon"
	var age int = 10

	// Deklarasi tanpa tipe data (Go akan infer tipe data)
	var city = "Bandung"

	fmt.Println(name, age, city)

}
