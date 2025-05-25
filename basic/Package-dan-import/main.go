package main

import (
	// Contoh import biasa
	"dasar_golang/basic/Package-dan-import/kalkulator"

	// Contoh import dengan alias
	// kalku "dasar_golang/basic/Package-dan-import/kalkulator"
	"fmt"
)

func main() {
	fmt.Println("Belajar tentang Package dan import")
	/*
		Package di Go adalah kumpulan file Go yang berada di dalam satu folder yang sama, dan memiliki deklarasi package di awal file-nya.
		Contoh Package
		- package main
		- package kalkulator


		Import digunakan untuk menggunakan kode atau fungsi yang berada di package lain.
		Contoh Import
		- import "fmt"
		- import "projectku/kalkulator"

	*/

	hasil := kalkulator.Tambah(5, 3)
	fmt.Println("Hasil:", hasil)
}
