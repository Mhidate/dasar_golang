package main

import "fmt"

type Materi struct {
	Kode  string
	Judul string
}

func main() {
	fmt.Println("Belajar tipe data slice")

	buah := []string{"apel", "mangga", "jeruk"}

	// Tambah elemen
	buah = append(buah, "pisang")

	// Akses elemen
	fmt.Println(buah[2]) // jeruk

	// Looping
	for i, val := range buah {
		fmt.Printf("Index %d: %s\n", i, val)
	}

	// Slicing
	bagian := buah[1:3] // ["mangga", "jeruk"]
	fmt.Println("Bagian:", bagian)

	//slice of struct
	materi := []Materi{
		{"06", "Multiple Main Function"},
		{"07", "Tipe Data: Number"},
		{"08", "Tipe Data: Boolean"},
	}

	fmt.Println(materi)
}
