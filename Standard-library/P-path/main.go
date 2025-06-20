package main

import (
	"fmt"
	"path"
)

func main() {

	fmt.Println("Belajar paket path")
	/*
		path adalah package di Go yang berisi fungsi-fungsi untuk manipulasi string path (bukan file/directory-nya langsung).
			Jadi ini buat operasi string-nya aja, seperti:
			join path
			ambil nama file
			ambil directory-nya
			normalize path (hapus . atau ..)
			split path
			Perlu dicatat:
			path dipakai untuk path berbasis slash (/), biasanya untuk URL atau UNIX-style path.
			Kalau untuk path di OS (Windows/Linux), pakai path/filepath
	*/

	//Ambil nama file / Base
	fmt.Println(path.Base("/folder/file.txt"))

	//Ambil nama direktori
	fmt.Println(path.Dir("/folder/file.txt"))

	//Ambil ekstensi file
	fmt.Println(path.Ext("/folder/file.txt"))

	//Gabungkan beberapa path
	p := path.Join("folder", "subfolder", "file.txt")
	fmt.Println(p)

	//Normalisasi path (bersihkan . ..)
	p2 := path.Clean("/folder/../file.txt")
	fmt.Println(p2)

	//Pisahkan directory & file
	dir, file := path.Split("/folder/file.txt")
	fmt.Println("Dir :", dir)
	fmt.Println("File:", file)
}
