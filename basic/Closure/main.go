package main

import "fmt"

func counter() func() int {
	angka := 0 // <-- ini disebut variabel di scope luar oleh closure-nya (function anonymous di bawah)

	return func() int { // closure mengakses variabel 'angka' yang berada di scope fungsi buatCounter
		angka++
		return angka
	}
}

func kelipatan(kali int) func(int) int {
	return func(nilai int) int {
		return nilai * kali
	}
}

func luar() func() func() int {
	angka := 0
	fmt.Println("Scope luar dibuat")

	// Function tengah
	return func() func() int {
		fmt.Println("Scope tengah dibuat")

		// Function dalam (closure)
		return func() int {
			angka++
			return angka
		}
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

	fmt.Printf("Angka %d \n", hitung())
	fmt.Printf("Angka %d \n", hitung())
	fmt.Printf("Angka %d \n", hitung())

	kaliDua := kelipatan(2)
	kaliTiga := kelipatan(3)

	fmt.Println(kaliDua(5))
	fmt.Println(kaliTiga(5))

	// Panggil function luar
	tengah := luar()  // bikin scope luar
	dalam := tengah() // bikin scope tengah

	// Eksekusi closure
	fmt.Println("Hasil pertama:", dalam())
	fmt.Println("Hasil kedua  :", dalam())
	fmt.Println("Hasil ketiga :", dalam())
}
