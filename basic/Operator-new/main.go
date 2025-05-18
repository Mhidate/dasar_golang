package main

import "fmt"

func main() {
	fmt.Println("Belajar Operator new di pointer")

	/*
		new adalah built-in function di Go yang dipakai untuk mengalokasikan memory baru untuk tipe data tertentu dan mengembalikan pointer ke memory itu.
	*/

	// Menggunakan operator new untuk buat pointer ke int
	var p *int = new(int)

	// Nilai default dari int adalah 0
	fmt.Println("Nilai awal pointer:", *p) // dereference

	// Ubah nilai lewat dereference
	*p = 100

	fmt.Println("Nilai setelah diubah:", *p)
}
