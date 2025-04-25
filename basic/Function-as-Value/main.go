package main

import "fmt"

func operasi(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	fmt.Println("Belajar Function as value")

	//Function as Value artinya kita bisa menyimpan sebuah fungsi ke dalam variabel dan memanggilnya lewat variabel tersebut.
	//Bisa disimpan ke variabel, bisa dikirim jadi parameter ke fungsi lain, bisa dikembalikan dari fungsi lain

	kali := func(a, b int) int {
		return a * b
	}

	fmt.Println("Hasil kali:", kali(5, 3))

	tambah := func(x, y int) int {
		return x + y
	}

	hasil := operasi(4, 5, tambah)
	fmt.Println("Hasil operasi:", hasil)

}
