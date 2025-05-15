package main

import "fmt"

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
}
