package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Belajar file manipulation")

	//Buat dan tulis ke file
	file, _ := os.Create("contoh.txt")
	defer file.Close()

	file.WriteString("Halo Dunia!\n")

	//Baca isi file
	data, _ := os.ReadFile("contoh.txt")
	fmt.Println(string(data))

	//Salin file
	src, _ := os.Open("contoh.txt")
	defer src.Close()

	dst, _ := os.Create("copy.txt")
	defer dst.Close()

	io.Copy(dst, src)

	//Hapus file
	os.Remove("contoh.txt")

	//Rename / Pindahkan file
	os.Rename("copy.txt", "baru.txt")

	//Buat folder
	os.Mkdir("folderku", 0755)

	//Atau kalau parent-folder belum ada:
	//os.MkdirAll("data/log/2025", 0755)

	//Hapus folder beserta isinya
	os.RemoveAll("data")

	//Cek apakah file ada
	if _, err := os.Stat("folderku/baru.txt"); err == nil {
		fmt.Println("File ada")
	} else {
		fmt.Println("File tidak ada")
	}

	//Cek info file
	info, _ := os.Stat("folderku/baru.txt")
	fmt.Println("Nama:", info.Name())
	fmt.Println("Ukuran:", info.Size())
	fmt.Println("ModTime:", info.ModTime())
	fmt.Println("Mode:", info.Mode())

}
