package main

import "fmt"

func counter() func() int {
	angka := 0

	return func() int {
		angka++
		return angka
	}
}

func kelipatan(kali int) func(int) int {
	return func(nilai int) int {
		return nilai * kali
	}
}

func main() {
	fmt.Println("Belajar tentang closure")
	//Closure adalah fungsi yang bisa mengakses variabel di sekitarnya (di scope luar) meskipun scope itu sudah selesai dieksekusi.
	//closure cocok digunakan untuk membuat fungsi yang ingat nilai sebelumnya,buat private state di dalam function,buat fungsi konfigurasi dinamis

	nama := "Jhon"

	// closure function
	salam := func() {
		fmt.Println("Halo", nama)
	}

	salam()

	hitung := counter()

	fmt.Println(hitung())
	fmt.Println(hitung())
	fmt.Println(hitung())

	kaliDua := kelipatan(2)
	kaliTiga := kelipatan(3)

	fmt.Println(kaliDua(5))
	fmt.Println(kaliTiga(5))
}
