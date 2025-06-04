package main

import (
	"fmt"
	"os"
	"os/exec"
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

	// 10. Ambil lokasi file notepad.exe
	cmd := exec.Command("where", "notepad.exe")
	output, _ := cmd.Output()
	fmt.Println(string(output))

	// 11. Buka Notepad
	cmd2 := exec.Command("notepad.exe")
	err2 := cmd2.Start()
	if err2 != nil {
		fmt.Println("Gagal buka Notepad:", err2)
		return
	}
	fmt.Println("Notepad dibuka.")

	// 12. Buka Kalkulator
	cmd3 := exec.Command("calc.exe")
	err3 := cmd3.Start()
	if err3 != nil {
		fmt.Println("Gagal buka Kalkulator:", err3)
		return
	}
	fmt.Println("Kalkulator dibuka.")

	// 13. Buka folder pakai Explorer
	path2 := "C:\\Windows"
	cmd4 := exec.Command("explorer.exe", path2)
	err4 := cmd4.Start()
	if err != nil {
		fmt.Println("Gagal buka folder:", err4)
		return
	}
	fmt.Println("Explorer dibuka ke folder:", path2)
}
