package main

import (
	"fmt"
)

func main() {
	fmt.Println("Type data number")
	var a int = 100
	// uint tidak bisa menyimpan bilangan negatif
	var b uint = 200
	var c float64 = 10.5
	var d complex64 = complex(3, 4)

	fmt.Println("int:", a)
	fmt.Println("uint:", b)
	fmt.Println("float64:", c)
	fmt.Println("complex64:", d)
	fmt.Println()

	fmt.Println("Type data boolean")
	// Tipe data boolean
	var isActive bool = true
	var isDeleted bool = false

	fmt.Println("Status aktif:", isActive)
	fmt.Println("Status terhapus:", isDeleted)
	fmt.Println()

	fmt.Println("Type data string")
	// Tipe data string
	var name string = "Golang"
	var message string = "Hello, World!"

	fmt.Println("Nama:", name)
	fmt.Println("Pesan:", message)

}
