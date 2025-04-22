package main

import "fmt"

func kali(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a * b, true
}

func biodata(nama, alamat, status string) (string, string, string) {
	return nama, alamat, status
}

func cariData(index int) (string, error) {
	data := []string{"A", "B", "C"}
	if index < 0 || index >= len(data) {
		return "", fmt.Errorf("Index out of range")
	}
	return data[index], nil
}

func main() {
	fmt.Println("Belajar fungsi yang mengembalikan banyak nilai")
	//Go memungkinkan sebuah fungsi untuk mengembalikan lebih dari satu nilai

	hasil, sukses := kali(10, 8)
	if sukses {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Pembagian tidak valid")
	}

	nama, alamat, status := biodata("Jhone", "London", "Lajang")
	fmt.Println("Nama lengkap:", nama)
	fmt.Println("Alamat:", alamat)
	fmt.Println("Status:", status)

	nilai, err := cariData(1)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	} else {
		fmt.Println("Data ditemukan:", nilai)
	}
}
