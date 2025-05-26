package main

import (
	"dasar_golang/basic/Access-modifier/person"
	"fmt"
)

func main() {
	fmt.Println("Belajar Access modifier")
	/*
			Access Modifier,ini konsep penting soal siapa yang boleh mengakses variabel, fungsi, struct, dan method di dalam atau luar package.
			Go itu sederhana banget soal access modifier, nggak kayak bahasa lain yang ada public, private, protected — di Go cuma dua aturan, yaitu berdasarkan huruf kapital di awal nama.
		    Cara Kerja Access Modifier di Go :
			Nama dengan huruf Kapital di awal artinya Exported / Public,huruf kecil di awal artinya Unexported / Private.

	*/

	// Akses variabel public
	fmt.Println("Nama public:", person.PublicNama)

	// Akses variabel private — ERROR
	// fmt.Println(person.privateAlamat)

	// Panggil fungsi public
	person.SayHello()

	// Panggil fungsi private — ERROR
	// person.sayAlamat()

	// Buat struct public
	p := person.Person{Nama: "Andi"}

	// Akses field public
	fmt.Println("Nama struct:", p.Nama)

	// Akses field private — ERROR
	// fmt.Println(p.umur)

	// Panggil method public
	p.CetakNama()

	// Panggil method private — ERROR
	// p.cetakUmur()
}
