package main

import "fmt"

func cetakHasil(angka int, proses func(int) int) {
	fmt.Println("hasil ", angka, "x", "9 =", proses(angka))
}

func operasi(a, b int, proses func(int, int) int) int {
	return proses(a, b)
}

func main() {
	fmt.Println("Belajar function as parameter")
	//Fungsi yang menerima parameter berupa fungsi lain.

	kaliSembilan := func(n int) int {
		return n * 9
	}

	cetakHasil(6, kaliSembilan)

	tambah := func(x, y int) int {
		return x + y
	}

	hasil := operasi(5, 3, tambah)
	fmt.Println("Hasil tambah:", hasil)

	// atau pakai anonymous function langsung
	hasilKali := operasi(5, 3, func(x, y int) int {
		return x * y
	})
	fmt.Println("Hasil kali:", hasilKali)

}
