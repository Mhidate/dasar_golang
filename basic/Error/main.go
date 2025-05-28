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

		// Contoh 1
		return 0, errors.New("tidak bisa membagi dengan nol")

		// Contoh 2
		//return 0, fmt.Errorf("error: tidak bisa membagi %d dengan %d (pembagi nol)", a, b)

		//Perbedaan errors.New() vs fmt.Errorf()
		//errors.New(), pesan error statis (tidak pakai variabel)
		// fmt.Errorf(),bisa buat pesan error dinamis (pakai variabel)
	}
	// kalau aman, return hasil dan nil (tidak ada error)
	return a / b, nil
}

// Function validasi nama
func cekNama(nama string) error {
	if nama == "" {
		return errors.New("error: nama tidak boleh kosong")
	}
	return nil
}

// Function validasi umur
func cekUmur(umur int) error {
	if umur < 0 {
		return fmt.Errorf("error: umur tidak boleh negatif, tapi %d", umur)
	}
	if umur > 120 {
		return fmt.Errorf("error: umur %d tidak masuk akal", umur)
	}
	return nil
}

func main() {
	fmt.Println("Belajar tentang error ")

	/*
	   Di Go, error adalah nilai (value) yang mengimplementasikan interface error. Jadi, error itu kayak tipe data khusus yang menyimpan pesan kesalahan.
	   Go tidak pakai try-catch, tapi menggunakan return value untuk mengembalikan error dari sebuah function.
	*/

	nama := "Andi"
	umur := 25
	angka1 := 100
	angka2 := 0

	// Cek nama
	if err := cekNama(nama); err != nil {
		fmt.Println(err)
		return
	}

	// Cek umur
	if err := cekUmur(umur); err != nil {
		fmt.Println(err)
		return
	}

	hasil, err := bagi(angka1, angka2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Halo %s, umur kamu %d tahun.\n", nama, umur)
	fmt.Printf("Hasil %d / %d = %d\n", angka1, angka2, hasil)

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
