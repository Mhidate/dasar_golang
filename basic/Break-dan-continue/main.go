package main

import "fmt"

func main() {
	fmt.Println("Belajar break dan continue")
	//Break digunakan untuk menghentikan loop atau switch secara paksa, bahkan jika kondisinya masih true.
	//Continue digunakan untuk melewati sisa kode dalam iterasi saat ini dan lanjut ke iterasi berikutnya.

	//Break Menghentikan seluruh loop, tidak lanjut ke iterasi berikut
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	//Continue melewati satu iterasi saja,langsung lanjut ke iterasi berikut
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i) //dilewati saat i=2
	}

	//Gunakan label agar bisa break/continue loop luar dari loop dalam,contoh:

outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				break outer // keluar dari semua loop
			}
			fmt.Println(i, j)
		}
	}

outer2:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue outer2 // lanjut ke iterasi berikutnya dari loop `outer`
			}
			fmt.Println(i, j)
		}
	}

}
