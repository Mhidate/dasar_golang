package main

import "fmt"

func main() {
	fmt.Println("Belajar type declarations")

	type name string
	type addres string
	type age int16

	var nama name = "Jhone"
	var alamat addres = "London"
	var umur age = 21

	fmt.Println("----------------")
	fmt.Println(nama)
	fmt.Println(alamat)
	fmt.Println(umur)
	fmt.Println("----------------")

	fmt.Println("Mencoba penjumlahan dengan tipe asli dan type declarations")

	//nama variabel tidak boleh sama dengan nama type declarations
	type angka1 int16
	type angka2 int16

	var angka12 angka1 = 5
	var angka22 angka2 = 6
	// var angka3 int16 = 8

	fmt.Println(angka12)
	fmt.Println(angka22)
	// fmt.Println(angka12 + angka22) tidak bisa melakukan operasi dengan type declarations yang berbeda atau type data asli
	// fmt.Println(angka12 + angka3)

}
