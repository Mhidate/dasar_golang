package main

import "fmt"

// fungsi cekTipe menerima interface kosong
func cekTipe(data interface{}) {
	switch value := data.(type) {
	case string:
		fmt.Println("Tipe data String:", value)
	case int:
		fmt.Println("Tipe data Integer:", value)
	case bool:
		fmt.Println("Tipe data Boolean:", value)
	case float64:
		fmt.Println("Tipe data Float64:", value)
	case []string:
		fmt.Println("Tipe data Slice of String:", value)
	case map[string]int:
		fmt.Println("Tipe data Map string ke int:", value)
	case nil:
		fmt.Println("Data nil")
	default:
		fmt.Println("Tipe data tidak dikenal")
	}
}

// Struct pertama
type User struct {
	Nama string
	Umur int
}

// Struct kedua
type Produk struct {
	Nama  string
	Harga int
}

// Fungsi cekData menerima interface kosong
func cekData(data interface{}) {
	switch value := data.(type) {
	case User:
		fmt.Println("Ini data User")
		fmt.Println("Nama:", value.Nama)
		fmt.Println("Umur:", value.Umur)
	case Produk:
		fmt.Println("Ini data Produk")
		fmt.Println("Nama Produk:", value.Nama)
		fmt.Println("Harga:", value.Harga)
	default:
		fmt.Println("Tipe data tidak dikenal")
	}
}

func main() {
	fmt.Println("Belajar Type assertions")

	/*
		    Type assertion di Go digunakan untuk mengakses nilai asli dari sebuah interface.
			Karena kalau kamu simpan data ke dalam interface{}, tipenya jadi "disembunyikan", nah Type Assertion dipakai buat "membuka bungkus"-nya.
	*/

	cekTipe("Halo Go")
	cekTipe(123)
	cekTipe(true)
	cekTipe(3.14)
	cekTipe([]string{"apel", "jeruk", "mangga"})
	cekTipe(map[string]int{"a": 1, "b": 2})
	cekTipe(nil)

	// Kirim struct ke fungsi cekData
	cekData(User{"Budi", 25})
	cekData(Produk{"Laptop", 7500000})
	cekData("string biasa")
}
