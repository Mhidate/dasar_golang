package main

import "fmt"

func tambah(a int, b int) int {
	return a + b
}
func salam(nama string) string {
	return "Halo, selamat siang " + nama
}

func main() {
	fmt.Println("Belajar function return value")

	//Fungsi di Go bisa mengembalikan nilai setelah dijalankan. Nilai ini disebut return value.
	//Fungsi dengan return harus menggunakan keyword return.

	hasil := tambah(3, 4)
	fmt.Println("Hasil:", hasil)

	fmt.Println(salam("Jhone"))

}
