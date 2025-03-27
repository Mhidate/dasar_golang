package main

import "fmt"

func main() {

	fmt.Println("Belajar konversi tipe data")

	//float ke int
	var num float64 = 9.99
	var integer int = int(num) // Hanya mengambil bagian bulatnya

	fmt.Println(integer)
}
