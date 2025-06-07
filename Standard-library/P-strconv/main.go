package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Belajar paket strconv")
	/*
			Paket strconv berisi fungsi-fungsi untuk konversi antara string dan tipe data numerik (int, float, bool, dll).
		    Jadi kalau kamu mau convert string ke angka, angka ke string, atau string ke bool, inilah paketnya.
	*/

	// String ke int
	numStr := "42"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Atoi:", num)
	}

	// Int ke string
	intVal := 100
	strVal := strconv.Itoa(intVal)
	fmt.Println("Itoa:", strVal)

	// String ke bool
	boolStr := "true"
	boolVal, _ := strconv.ParseBool(boolStr)
	fmt.Println("ParseBool:", boolVal)

	// String ke float
	floatStr := "3.14"
	floatVal, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Println("ParseFloat:", floatVal)

	// String ke int pakai base dan bitSize
	hexStr := "ff"
	hexVal, _ := strconv.ParseInt(hexStr, 16, 64)
	fmt.Println("ParseInt (hex):", hexVal)

	// Bool ke string
	boolToStr := strconv.FormatBool(true)
	fmt.Println("FormatBool:", boolToStr)

	// Float ke string
	floatToStr := strconv.FormatFloat(3.1415, 'f', 2, 64)
	fmt.Println("FormatFloat:", floatToStr)

	// Int ke string dengan base 10
	intToStr := strconv.FormatInt(255, 10)
	fmt.Println("FormatInt:", intToStr)

	// Uint ke string dengan base 16 (hex)
	uintToStr := strconv.FormatUint(255, 16)
	fmt.Println("FormatUint (hex):", uintToStr)
}
