package main

import (
	"fmt"

	"github.com/Mhidate/kalk"
	"github.com/gosimple/slug"
)

func main() {
	text := "Belajar Go Modules!"
	/*
		Go Modules adalah sistem dependency management bawaan Go yang diperkenalkan di Go 1.11 (stabil di Go 1.13).
		Modul adalah kumpulan kode Go dalam sebuah direktori yang berisi file go.mod.
		Fungsinya:
		Menyimpan informasi tentang dependencies (package lain yang dibutuhkan)
		Mengatur versi dependency yang dipakai
		Memudahkan pembuatan, pengelolaan, dan distribusi project Go
	*/

	/*
		Perintah Dasar Go Modules
		go mod init	,Inisialisasi modul
		go mod tidy,	Download missing deps & hapus unused
		go mod download,	Download semua dependency
		go mod verify,	Verifikasi checksum di go.sum
		go list -m all,	Lihat semua modul dan versinya
		go get	,Update/tambah dependency
	*/

	fmt.Println(slug.Make(text))

	text2 := slug.Make("Hellö Wörld хелло ворлд")
	fmt.Println(text2)

	fmt.Println(kalk.Penjumlahan(10, 45))

}
