package main

import (
	"errors"
	"fmt"
)

// Function pembagian dengan error handling
// Kenapa Pakai (int, error)?
// Karena Go gak pakai try-catch.
// Semua function yang rawan error lebih baik return value hasil + error. Nanti di tempat pemanggilan tinggal dicek.
func bagi(a, b int) (int, error) {
	if b == 0 {
		// kalau b nol, return error
		return 0, errors.New("tidak bisa membagi dengan nol")
	}
	// kalau aman, return hasil dan nil (tidak ada error)
	return a / b, nil
}

func main() {
	fmt.Println("Belajar tentang error ")

	/*
	   Di Go, error adalah nilai (value) yang mengimplementasikan interface error. Jadi, error itu kayak tipe data khusus yang menyimpan pesan kesalahan.
	   Go tidak pakai try-catch, tapi menggunakan return value untuk mengembalikan error dari sebuah function.
	*/

	hasil1, err1 := bagi(10, 2)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Hasil 10 / 2 =", hasil1)
	}

	hasil2, err2 := bagi(5, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Hasil 5 / 0 =", hasil2)
	}
}
