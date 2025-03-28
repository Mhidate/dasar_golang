package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Belajar konversi tipe data")

	//float ke int
	var num float64 = 9.99
	var integer int = int(num) // Hanya mengambil bagian bulatnya

	fmt.Println(integer)

	str := "100"
	num2, err := strconv.Atoi(str) // Konversi string ke int
	if err == nil {
		fmt.Println(num2) // Output: 100
	}

}
