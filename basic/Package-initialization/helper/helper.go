package helper

import "fmt"

// kode ini otomatis jalan saat package di-import
var Nama string = "Budi"

func init() {
	fmt.Println("Package helper sudah di-inisialisasi")
	Nama = "Budi diubah saat init"
}

// Function biasa
func SayHello() {
	fmt.Println("Halo dari helper:", Nama)
}
