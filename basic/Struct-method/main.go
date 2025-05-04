package main

import "fmt"

type Produk struct {
	Nama  string
	Harga int
}

// Method value receiver
func (p Produk) Info() {
	fmt.Println("Produk:", p.Nama, "Harga:", p.Harga)
}

// Method pointer receiver
func (p *Produk) UbahHarga(hargaBaru int) {
	p.Harga = hargaBaru
}

func main() {

	fmt.Println("Belajar struct method")
	/* Struct Method adalah function yang “menempel” pada struct. Artinya, method itu bisa dipanggil oleh object struct tertentu dan bisa mengakses field yang ada di struct tersebut.
	   Di Go, method didefinisikan dengan menambahkan receiver di antara func dan nama method-nya.
	*/

	p := Produk{Nama: "Laptop", Harga: 10000000}
	p.Info()

	p.UbahHarga(9500000)
	p.Info()
}
