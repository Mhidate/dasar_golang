package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("Belajar paket regexp")
	/*
		Paket regexp di Go dipakai buat:
			Mencari teks yang cocok dengan pola tertentu
			Memvalidasi format string (misal: email, nomor hp, dsb)
			Split string berdasarkan pola
			Replace string berdasarkan pola
			Pola yang dipakai disebut Regular Expression (regex)
			Contoh regex:
			^[A-Za-z]+$ → hanya huruf
			^[0-9]+$ → hanya angka
			^.+@.+\..+$ → format email sederhana
	*/

	//Cek Kecocokan Pola
	pola := regexp.MustCompile("^[a-z]+$")

	fmt.Println(pola.MatchString("dian"))
	fmt.Println(pola.MatchString("Dian123"))

	//Ambil Semua Kecocokan
	teks := "Halo 123, ini 456, itu 789"
	pola2 := regexp.MustCompile("[0-9]+")

	angka := pola2.FindAllString(teks, -1) //-1, artinya:Ambil semua hasil yang cocok tanpa batas,kalo menggunakan 2, berati 2 hasil yg cocok
	fmt.Println(angka)

	//Replace String
	teks2 := "password: rahasia123"
	pola3 := regexp.MustCompile("[0-9]+")

	hasil := pola3.ReplaceAllString(teks2, "****")
	fmt.Println(hasil)

	//Split String Pakai Regex
	teks3 := "apel, jeruk; mangga pisang"
	pola4 := regexp.MustCompile("[,; ]+")

	hasil2 := pola4.Split(teks3, -1)
	fmt.Println(hasil2)

}
