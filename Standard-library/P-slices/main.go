package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Println("Belajar paket slices")
	/*
		Menyediakan fungsi utility untuk operasi slice yang sebelumnya harus kita bikin manual.
			Sekarang bisa:
			Cari index item
			Cek apakah ada item
			Hapus item
			Bandingkan slice
			Copy slice
			Sort slice
			Reverse slice
			dll,
			tanpa repot bikin loop sendiri.
	*/

	//Cek apakah 2 slice sama
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	fmt.Println(slices.Equal(a, b))

	//Cek apakah slice berisi item
	items := []string{"apel", "jeruk", "mangga"}

	fmt.Println(slices.Contains(items, "jeruk"))
	fmt.Println(slices.Contains(items, "durian"))

	//Hapus item di index tertentu
	angka := []int{1, 2, 3, 4, 5}
	// hapus index 1-3 (yaitu 2,3,4)
	angka = slices.Delete(angka, 1, 4)
	fmt.Println(angka)

	//Urutkan dan Balik Urutan
	angka2 := []int{5, 2, 8, 1}
	slices.Sort(angka2)
	fmt.Println(angka2)

	slices.Reverse(angka2)
	fmt.Println(angka2)

	//Cari Nilai Maksimum & Minimum (Go 1.22 ke atas)
	angka3 := []int{5, 2, 8, 1}
	fmt.Println(slices.Max(angka3))
	fmt.Println(slices.Min(angka3))
}
