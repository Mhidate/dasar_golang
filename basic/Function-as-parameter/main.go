package main

import "fmt"

func cetakHasil(angka int, proses func(int) int) {
	fmt.Println("hasil ", angka, "x", "9 =", proses(angka))
}

func main() {
	fmt.Println("Belajar function as parameter")
	//Fungsi yang menerima parameter berupa fungsi lain.

	kaliSembilan := func(n int) int {
		return n * 9
	}

	cetakHasil(6, kaliSembilan)

}
