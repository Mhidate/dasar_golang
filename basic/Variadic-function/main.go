package main

import "fmt"

func tambah(angka ...int) int {
	total := 0
	for _, nilai := range angka {
		total += nilai
	}
	return total
}

func cetak(pesan ...string) {
	for i, p := range pesan {
		fmt.Println(i+1, p)
	}
}

func biodata(umur int, bio ...string) {

	fmt.Println("Nama   :", bio[0])
	fmt.Println("Alamat :", bio[1])
	fmt.Println("Status :", bio[2])
	fmt.Println("Umur : ", umur)

}

func main() {
	fmt.Println("Belajar tetang variadic function")

	//Variadic function adalah fungsi yang bisa menerima parameter dengan sejumlah tak terbatas.
	//Parameter variadic harus berada di akhir.
	//Parameter variadic tidak boleh lebih dari satu.

	fmt.Println(tambah(4, 5, 3, 8))

	cetak("halo,jangan lupa cari kerja", ",jangan mencuri", "waulaupun lagi susah")

	biodata(30, "John", "London", "Lajang")

}
