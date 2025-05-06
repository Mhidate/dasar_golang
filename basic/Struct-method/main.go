package main

import "fmt"

type Produk struct {
	Nama  string
	Harga int
}

// Struct Biodata
type Biodata struct {
	Nama   string
	Alamat string
	Umur   int
}

// Method value receiver
func (p Produk) Info() {
	fmt.Println("Produk:", p.Nama, "Harga:", p.Harga)
}

// Method pointer receiver
func (p *Produk) UbahHarga(hargaBaru int) {
	p.Harga = hargaBaru
}

// Method untuk menampilkan data biodata
func (b Biodata) Tampilkan() {
	fmt.Println("Nama  :", b.Nama)
	fmt.Println("Alamat:", b.Alamat)
	fmt.Println("Umur  :", b.Umur, "tahun")
	fmt.Println("-----------------------------")
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

	daftarBiodata := []Biodata{
		{Nama: "Budi", Alamat: "Jakarta", Umur: 15},
		{Nama: "John", Alamat: "London", Umur: 35},
		{Nama: "Karen", Alamat: "Texas", Umur: 29},
	}

	// Tambah biodata baru ke slice
	biodataBaru := Biodata{Nama: "Dewi", Alamat: "Meksiko", Umur: 22}
	daftarBiodata = append(daftarBiodata, biodataBaru)

	for _, biodata := range daftarBiodata {
		biodata.Tampilkan()
	}

	// Cari data dengan nama "Budi"
	namaDicari := "Budi"
	ditemukan := false
	fmt.Println("=====Mencari data=====")
	for _, b := range daftarBiodata {
		if b.Nama == namaDicari {
			fmt.Println("Data ditemukan:")
			fmt.Println("Nama  :", b.Nama)
			fmt.Println("Alamat:", b.Alamat)
			fmt.Println("Umur  :", b.Umur)
			ditemukan = true
			break
		}
	}

	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}

	fmt.Println("=====Menghapus data=====")
	namaDihapus := "John"
	indexDihapus := -1

	// Cari index data yang mau dihapus
	for i, b := range daftarBiodata {
		if b.Nama == namaDihapus {
			indexDihapus = i
			break
		}
	}

	if indexDihapus != -1 {
		// Hapus data dari slice
		daftarBiodata = append(daftarBiodata[:indexDihapus], daftarBiodata[indexDihapus+1:]...)
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}

	// Tampilkan data sisa
	fmt.Println("Data sisa:")
	for _, b := range daftarBiodata {
		fmt.Println(b)
	}
}
