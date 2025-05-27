package main

import (
	"dasar_golang/basic/Package-initialization/helper"
	"fmt"
)

func main() {
	fmt.Println("Belajar tentang Package initialization")

	/*
		Package-Initialization atau inisialisasi paket adalah proses di mana sebuah paket Go diinisialisasi sebelum kode di dalamnya dapat digunakan. Proses ini dilakukan melalui fungsi init(), yang dapat digunakan untuk menjalankan kode inisialisasi awal, seperti menginisialisasi variabel, melakukan pengaturan, atau membuat koneksi ke database, sebelum kode lainnya dalam paket tersebut dieksekusi.
		Cara Kerja Package Initialization
			Di Go, kalau sebuah package diimport:
			Kode di luar function akan dijalankan lebih dulu.
			Function khusus init() juga akan otomatis dipanggil sebelum program berjalan.
			init() bisa lebih dari satu di satu package, dan akan dieksekusi secara urut sesuai urutan definisinya di file.
	*/

	fmt.Println("Program mulai...")

	// Pakai function dari helper
	helper.SayHello()

	fmt.Println("Nama dari helper:", helper.Nama)

}
