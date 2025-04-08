package main

import "fmt"

func main() {
	fmt.Println("Belajar percabangan atau IF")

	// if mengeksekusi kondisi tertentu

	nilai := 66

	if nilai >= 70 {
		fmt.Println("C")
	} else if nilai >= 80 {
		fmt.Println("B")
	} else if nilai > 90 {
		fmt.Println("A")
	} else {
		fmt.Println("D")
	}

	//IF dengan short statement
	if skor := 70; skor >= 60 {
		fmt.Println("Lulus")
	}

	if panjang := len("golang"); panjang > 5 {
		fmt.Println("Panjang string lebih dari 5:", panjang)
	} else {
		fmt.Println("Pendek")
	}

	data := map[string]string{
		"username": "johndoe",
		"email":    "john@example.com",
	}

	if val, ok := data["email"]; ok {
		fmt.Println("Email ditemukan:", val)
	} else {
		fmt.Println("Email tidak ditemukan")
	}

}
