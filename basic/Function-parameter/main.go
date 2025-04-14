package main

import "fmt"

//penulisan
// func namaFungsi(namaParam1 tipe1, namaParam2 tipe2) returnTipe {
// kode fungsi
// }

func ganjil(a int) {
	if a%2 != 0 {
		fmt.Println("Angka ", a, "adalah gnjil")
	}
}

func genap(a int) string {
	var hasil string
	if a%2 == 0 {
		hasil = "genap"

	} else {
		hasil = "ganjil"
	}
	return hasil
}

//menyikat parameter dengan tipe yang sama
func kurang(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Belajar function parameter")

	//Parameter adalah variabel sementara yang digunakan untuk menerima nilai dari luar ketika fungsi dipanggil
	ganjil(5)

	a := 8
	hasil := genap(a)

	fmt.Println("Angka", a, "adalah", hasil)
	c := 10
	d := 6
	b := kurang(c, d)
	fmt.Println(c, "-", d, "=", b)

}
