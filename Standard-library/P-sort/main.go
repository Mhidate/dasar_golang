package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Nama string
	Skor int
}

func main() {

	fmt.Println("Belajar paket sort")
	/*
		Paket sort di Go menyediakan fungsi untuk mengurutkan slice:
			Ascending (naik)
			Descending (kita bisa custom)
			Berdasarkan kriteria tertentu (misal urut nama, skor, dll)
	*/

	angka := []int{4, 2, 7, 1, 9}
	nama := []string{"Dian", "Aldi", "Bima", "Citra"}

	sort.Ints(angka)
	sort.Strings(nama)

	fmt.Println("Urut angka:", angka)
	fmt.Println("Urut nama :", nama)

	// Urut descending
	sort.Slice(angka, func(i, j int) bool {
		return angka[i] > angka[j]
	})

	sort.Slice(nama, func(i, j int) bool {
		return nama[i] > nama[j] // alfabet terbesar dulu
	})

	fmt.Println("Angka descending:", angka)
	fmt.Println("Nama descending:", nama)

	pemain := []Player{
		{"Dian", 70},
		{"Aldi", 90},
		{"Bima", 80},
	}

	sort.Slice(pemain, func(i, j int) bool {
		return pemain[i].Skor > pemain[j].Skor // Descending
	})

	fmt.Println("Urut Skor:")
	for _, p := range pemain {
		fmt.Println(p.Nama, p.Skor)
	}
}
