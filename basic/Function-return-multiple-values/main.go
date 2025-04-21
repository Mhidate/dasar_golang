package main

import "fmt"

func kali(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a * b, true
}

func main() {
	fmt.Println("Belajar fungsi yang mengembalikan banyak nilai")

	hasil, sukses := kali(10, 8)
	if sukses {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Pembagian tidak valid")
	}
}
