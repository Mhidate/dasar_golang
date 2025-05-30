package main

import "fmt"

func main() {

	fmt.Println("Belajar label")
	/*
		Label di Go adalah nama identifier yang bisa ditempelkan di:
		Sebuah statement (biasanya for, switch, atau select)
		Dipanggil menggunakan break, continue, atau goto


		Apakah penamaan label di Go bebas?
		Bebas dalam artian:
		- Boleh pakai huruf apa saja (asal sesuai aturan penamaan identifier di Go)
		- Tidak bentrok sama keyword Go (kayak for, if, break, dll.)
		- Harus unik di dalam scope-nya
		- Biasanya ditulis dengan huruf kecil atau camelCase, meskipun tidak wajib.
		- Karena dia nggak dipakai di luar package (nggak kayak export struct, fungsi, dll)
		Aturan penamaan identifier di Go:
		- Boleh mulai dengan huruf (A-Z, a-z) atau underscore (_)
		- Bisa diikuti huruf, angka, atau underscore
		- Tidak boleh pakai spasi atau karakter khusus(kaya @, %, $, -, ! dll.)

	*/

outerLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i, j)
			if i == 2 && j == 2 {
				break outerLoop // keluar dari kedua loop
			}
		}
	}
	fmt.Println("Selesai")

outerLoop2:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue outerLoop2 // lanjut ke iterasi selanjutnya di outerLoop
			}
			fmt.Println(i, j)
		}
	}
}
