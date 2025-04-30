package main

import "fmt"

func tangkapPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recover menangkap panic:", r)
	}
}

func main() {

	fmt.Println("Belajar Defer-Panic-dan-Recover")
	//DEFER adalah perintah untuk menunda eksekusi sebuah statement/fungsi sampai function tempatnya dipanggil selesai.
	//Semua defer dieksekusi setelah function selesai, dieksekusi secara LIFO (Last In, First Out).
	//contoh sederhana defer
	defer fmt.Println("Akhir")
	fmt.Println("Awal")

	//contoh defer sederhana 2
	defer fmt.Println("Defer pertama")
	defer fmt.Println("Defer kedua")
	fmt.Println("Main function")

	//PANIC digunakan untuk menghentikan program secara mendadak, biasanya saat kondisi error fatal terjadi.
	//Saat panic dipanggil:
	//Program akan langsung berhenti,tapi defer tetap dijalankan dulu sebelum program benar-benar berhenti
	//contoh sederhana panic
	defer fmt.Println("Defer tetap jalan")
	panic("Ada masalah besar!")
	fmt.Println("Ini tidak akan dijalankan")

	//RECOVER dipakai untuk menangkap panic, sehingga program bisa tetap lanjut tanpa benar-benar crash.
	//recover hanya efektif kalau dipanggil di dalam defer
	//contoh recover
	defer tangkapPanic()
	panic("Terjadi error fatal")
	fmt.Println("Program selesai") // tidak jalan
}
