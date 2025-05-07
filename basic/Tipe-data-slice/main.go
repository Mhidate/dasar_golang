package main

import "fmt"

type Materi struct {
	Kode  string
	Judul string
}

//Fungsi untuk hapus elemen pada slice berdasarkan index
func hapusElemen(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	fmt.Println("Belajar tipe data slice")
	//Slice di Go adalah potongan dari array yang dinamis panjangnya dan tidak memerlukan alokasi ulang.

	buah := []string{"apel", "mangga", "jeruk"}

	//Cek panjang
	fmt.Println(len(buah))

	//Cek kapasitas
	fmt.Println(cap(buah))

	// Tambah elemen
	buah = append(buah, "pisang")

	// Akses elemen dengan index
	fmt.Println(buah[2])

	//Mengubah nilai elemen slice
	buah[0] = "jambu"

	// Looping
	fmt.Println("Sebelum ada index yang di hapus")
	for i, val := range buah {
		fmt.Printf("Index %d: %s\n", i, val)
	}
	fmt.Println()

	// Menghapus elemen pada slice
	index := 0
	buah = hapusElemen(buah, index)
	fmt.Println("Setelah ada index yang di hapus")
	fmt.Println(buah)
	fmt.Println()

	// Slicing
	bagian := buah[1:3]
	fmt.Println("Bagian:", bagian)

	cariBuah := "Anggur"
	for _, value := range buah {
		if value == cariBuah {
			fmt.Println("Buah yang anda cari ada di slice")
			break
		}
		fmt.Println("Buah yang anda cari tidak ada di slice")
	}

	//slice of struct
	materi := []Materi{
		{"06", "Multiple Main Function"},
		{"07", "Tipe Data: Number"},
		{"08", "Tipe Data: Boolean"},
	}

	fmt.Println(materi)
}
