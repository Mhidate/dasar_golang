package main

import "fmt"

func main() {
	fmt.Println("Belajar interface kosong")
	// interface yang tidak memiliki method sama sekali → jadi bisa menampung semua jenis tipe data.

	dataCampuran := []interface{}{"String", 100, true, 3.14}

	for _, data := range dataCampuran {
		fmt.Println(data)
	}

	person := map[string]interface{}{
		"nama":   "Adit",
		"umur":   21,
		"status": true,
	}

	fmt.Println(person["nama"])
	fmt.Println(person["umur"])
	fmt.Println(person["status"])
}
