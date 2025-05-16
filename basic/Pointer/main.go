package main

import "fmt"

// Struct Biodata
type Biodata struct {
	Nama   string
	Umur   int
	Alamat string
}

// Function ubah data (pakai pointer)
func ubahNama(b *Biodata, namaBaru string) {
	b.Nama = namaBaru
}

// Function hapus data dari slice
func hapusData(slice []Biodata, index int) []Biodata {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	fmt.Println("Belajar tentang pointer")
	//Pointer adalah variabel yang menyimpan alamat memori dari variabel lain.
	/*
		Dengan pointer, kamu bisa:
		 Mengakses atau mengubah nilai variabel lain lewat alamatnya.
		 Menghemat memori dan performa, karena cukup kirim alamatnya saja, bukan salinan datanya.
		Ada 2 operator penting:
			& → mengambil alamat memori
			* → mengambil nilai di alamat memori (dereference)
	*/

	nama := "Jhon"

	// Ambil alamat dari variabel nama
	var alamatPointer *string = &nama

	fmt.Println("Alamat memori nama:", alamatPointer)   // alamat memori
	fmt.Println("Isi alamat tersebut:", *alamatPointer) // isi di alamat tsb

	// Ubah isi dari alamat tsb
	*alamatPointer = "Andi"
	fmt.Println("Nama setelah diubah lewat pointer:", nama)

	// Slice of struct
	biodataList := []Biodata{
		{"Budi", 25, "Bandung"},
		{"Sari", 22, "Jakarta"},
		{"Andi", 28, "Surabaya"},
	}

	fmt.Println("Data Awal:")
	for _, data := range biodataList {
		fmt.Println(data)
	}

	// Ubah nama Andi jadi Rudi (pakai pointer)
	ubahNama(&biodataList[2], "Rudi")

	fmt.Println("\nSetelah Ubah Nama:")
	for _, data := range biodataList {
		fmt.Println(data)
	}

	// Hapus data ke-1 (Sari)
	biodataList = hapusData(biodataList, 1)

	fmt.Println("\nSetelah Hapus Data:")
	for _, data := range biodataList {
		fmt.Println(data)
	}
}
