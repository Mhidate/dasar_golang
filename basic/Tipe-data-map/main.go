package main

import "fmt"

func main() {
	fmt.Println("Belajar tipe data map")
	//Map adalah tipe data asosiatif yang ada di Go yang berbentuk key-value pair.map memiliki key dan value ,keynya harus
	// var namaMap map[TipeKey]TipeValue

	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(chicken1)

	// Deklarasi dan inisialisasi vertical
	warga := map[string]string{
		"nama":   "Jhon",
		"gender": "Laki-laki",
		"RT":     "01",
	}

	// Tambah data
	warga["status"] = "belum menikah"

	// Ubah data
	warga["RT"] = "05"

	// Akses
	fmt.Println("Nama:", warga["nama"])

	// Cek key
	if val, ok := warga["status"]; ok {
		fmt.Println("Status:", val)
	}

	// Loop
	for k, v := range warga {
		fmt.Println(k, ":", v)
	}

	// Hapus
	delete(warga, "RT")
}
