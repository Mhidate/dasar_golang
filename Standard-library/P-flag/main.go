package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Belajar paket flag")
	/*
		Paket flag di Go dipakai buat:
			Baca argumen dari command line
			Parsing nilai sesuai tipe data (string, int, bool, dll)
			Menentukan default value & deskripsi untuk setiap flag
			Bisa generate help otomatis
	*/

	// perintah untuk run go run main.go -nama=Andi -umur=25 -aktif=true tambahan1 tambahan2

	// Definisikan flag
	nama := flag.String("nama", "User", "Nama pengguna")
	umur := flag.Int("umur", 20, "Umur pengguna")
	aktif := flag.Bool("aktif", false, "Status aktif")

	// Parsing flag
	flag.Parse()

	// Cetak hasil parsing
	fmt.Println("Nama :", *nama)
	fmt.Println("Umur :", *umur)
	fmt.Println("Aktif:", *aktif)

	// Cek sisa argumen (jika ada)
	fmt.Println("Argumen tambahan:", flag.Args())
}
