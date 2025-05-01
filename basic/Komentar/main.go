package main

import "fmt"

// Add mengembalikan hasil penjumlahan dua bilangan integer.
func Add(a, b int) int {
	return a + b
}

// Subtract mengembalikan hasil pengurangan b dari a.
func Subtract(a, b int) int {
	return a - b
}

// Multiply mengembalikan hasil perkalian dua bilangan integer.
func Multiply(a, b int) int {
	return a * b
}

// Divide mengembalikan hasil pembagian a dengan b.
// Jika b = 0, fungsi akan mengembalikan 0 untuk menghindari pembagian dengan nol.
func Divide(a, b int) int {
	if b == 0 {
		// Hindari pembagian dengan nol
		return 0
	}
	return a / b
}

func main() {

	fmt.Println("Belajar tentang komentar")
	/*
			   Di Go, ada konvensi penting soal komentar — komentar untuk dokumentasi (Godoc).
		       Format Godoc:
		       Komentar dokumentasi biasanya langsung ditulis di atas deklarasi fungsi, variabel, struct, atau package, dan harus diawali dengan nama item yang dijelaskan.
	           Go tidak mendukung komentar dalam komentar.
	*/

	// Ini adalah komentar satu baris
	fmt.Println("Hello, world!") // Ini juga komentar satu

	/*
		Ini adalah komentar
		lebih dari satu baris
	*/
	fmt.Println("Hello, world!")

	// Contoh penggunaan fungsi Add
	resultAdd := Add(10, 5)
	fmt.Println("Hasil penjumlahan:", resultAdd)

	// Contoh penggunaan fungsi Subtract
	resultSubtract := Subtract(10, 5)
	fmt.Println("Hasil pengurangan:", resultSubtract)

	// Contoh penggunaan fungsi Multiply
	resultMultiply := Multiply(10, 5)
	fmt.Println("Hasil perkalian:", resultMultiply)

	// Contoh penggunaan fungsi Divide
	resultDivide := Divide(10, 0)
	fmt.Println("Hasil pembagian:", resultDivide)

}
