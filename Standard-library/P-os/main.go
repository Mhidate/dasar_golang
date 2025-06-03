package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Belajar paket OS")
	/*
		 Paket os di Go menyediakan berbagai fungsi dan tipe data untuk:
			Mengakses file & direktori
			Menangani environment variable
			Baca argumen program
			Info tentang OS & sistem
			Exit program
			Bikin/rename/hapus file & folder
			Dan banyak lagi
	*/

	// 1. Create file
	file, err := os.Create("contoh.txt")
	if err != nil {
		fmt.Println("Gagal buat file:", err)
		return
	}
	defer file.Close()

	// 2. Write ke file
	file.WriteString("Halo dari Go!\n")

	// 3. Baca file
	readFile, err := os.Open("contoh.txt")
	if err != nil {
		fmt.Println("Gagal buka file:", err)
		return
	}
	defer readFile.Close()

	// Buffer untuk baca
	buffer := make([]byte, 100)
	n, _ := readFile.Read(buffer)
	fmt.Println("Isi file:")
	fmt.Println(string(buffer[:n]))

	// 4. Rename file
	err = os.Rename("contoh.txt", "data-baru.txt")
	if err != nil {
		fmt.Println("Gagal rename:", err)
	}

	// 5. Buat direktori
	err = os.Mkdir("data", 0755)
	if err != nil {
		fmt.Println("Gagal buat folder:", err)
	}

	// 6. Hapus file
	err = os.Remove("data-baru.txt")
	if err != nil {
		fmt.Println("Gagal hapus file:", err)
	}

	// 7. Ambil environment variable
	path := os.Getenv("PATH")
	fmt.Println("PATH env:", path)

	// 8. Argumen program
	fmt.Println("Argumen CLI:", os.Args)

	// 9. Working directory
	dir, _ := os.Getwd()
	fmt.Println("Current directory:", dir)
}
