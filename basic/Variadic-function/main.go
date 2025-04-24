package main

import "fmt"

func tambah(angka ...int) int {
	total := 0
	for _, nilai := range angka {
		total += nilai
	}
	return total
}

func cetak(pesan ...string) {
	for i, p := range pesan {
		fmt.Println(i+1, p)
	}
}

func main() {
	fmt.Println("Belajar tetang variadic function")

	//Variadic function adalah fungsi yang bisa menerima parameter dengan sejumlah tak terbatas
	fmt.Println(tambah(4, 5, 3, 8))

	cetak("halo,jangan lupa cari kerja", ",jangan mencuri", "waulaupun lagi susah")

}
