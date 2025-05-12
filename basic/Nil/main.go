package main

import "fmt"

func cariData(data map[string]string, key string) string {
	if data == nil {
		return "Data tidak tersedia"
	}

	if value, exists := data[key]; exists {
		return value
	} else {
		return "Data tidak ditemukan"
	}
}

func main() {
	fmt.Println("Belajar tentang nil")
	/*
		Nil adalah representasi nilai kosong atau tidak ada nilai untuk beberapa tipe data tertentu.
		 Kalau di bahasa lain mungkin kayak null, None, atau undefined.
	*/

	var data map[string]string

	fmt.Println(cariData(data, "nama"))

	data = map[string]string{
		"nama": "Budi",
	}

	fmt.Println(cariData(data, "nama"))
}
