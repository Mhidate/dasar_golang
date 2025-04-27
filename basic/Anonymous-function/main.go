package main

import "fmt"

func tampilkan(angka int, f func(int) int) {
	fmt.Println("Hasil:", f(angka))
}

func main() {
	fmt.Println("Belajar anonymous function")
	//Anonymous function adalah fungsi yang tidak memiliki nama
	// secara umum digunakan dengan cara langsung dieksekusi,disimpan ke variabel atau dikirim sebagai parameter

	//contoh 1
	func() {
		fmt.Println("Halo, ini anonymous function!")
	}()

	//contoh 2
	sapa := func(nama string) {
		fmt.Println("Halo", nama)
	}

	sapa("Dimas")
	sapa("Lina")

	//contoh 3
	tampilkan(10, func(x int) int {
		return x * 2
	})

	//contoh 4
	hasil := func(a, b int) int {
		return a + b
	}(3, 5)

	fmt.Println("Hasil:", hasil)

}
