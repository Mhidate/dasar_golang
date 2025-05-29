package main

import (
	"fmt"
	"log"
	"os"
)

// Custom Error untuk form
type FormError struct {
	Fields map[string]string
}

// Wajib implement method Error() string
func (e FormError) Error() string {
	return "Ada error pada form input"
}

// Function buat validasi
func validasiForm(nama string, umur int, email string) error {
	errors := make(map[string]string)

	if nama == "" {
		errors["Nama"] = "Nama tidak boleh kosong"
	}
	if umur < 17 {
		errors["Umur"] = "Umur minimal harus 17 tahun"
	}
	if email == "" {
		errors["Email"] = "Email wajib diisi"
	}

	// Kalau ada error, return custom error
	if len(errors) > 0 {
		return FormError{Fields: errors}
	}

	// Kalau gak ada error
	return nil
}

func main() {
	fmt.Println("Belajar custom error")

	/*
		Custom error dalam Go digunakan untuk memberikan informasi lebih spesifik tentang kesalahan yang terjadi dalam aplikasi.
	*/

	// Setup log file
	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Gagal buat file log:", err)
		return
	}
	defer file.Close()

	log.SetOutput(file) // Arahkan log ke file

	// Simulasi input
	nama := ""
	umur := 11
	email := ""

	// Validasi input
	err = validasiForm(nama, umur, email)
	if err != nil {
		fmt.Println(err)

		// Cek kalau error-nya custom FormError
		if formErr, ok := err.(FormError); ok {
			fmt.Println("Detail errornya:")
			for field, msg := range formErr.Fields {
				fmt.Printf("- %s: %s\n", field, msg)
				log.Println(fmt.Sprintf("%s: %s", field, msg)) // Simpan ke log file
			}
		}
		return
	}

	fmt.Println("Form valid, data disimpan!")
}
