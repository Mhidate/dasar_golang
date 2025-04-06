package main

import "fmt"

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
}
