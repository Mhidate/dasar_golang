package main

import "fmt"

type Alamat struct {
	Jalan string
	Kota  string
}

type User struct {
	Nama   string
	Umur   int
	Alamat Alamat
}

func cetakUser(u User) {
	fmt.Println("Nama:", u.Nama)
	fmt.Println("Umur:", u.Umur)
}

func main() {

	fmt.Println("Belajar tentang struct")
	/*
		Struct (singkatan dari Structure) adalah tipe data bentukan (composite type) di Go yang digunakan untuk mengelompokkan beberapa data (field) yang berbeda tipe ke dalam satu kesatuan.
		Mirip kayak object di bahasa lain, tapi di Go nggak ada class — struct lah yang dipakai buat representasi data yang kompleks.
	*/

	u := User{
		Nama: "Jhon",
		Umur: 24,
		Alamat: Alamat{
			Jalan: "Jl. Batu",
			Kota:  "London",
		},
	}
	fmt.Println(u)

	fmt.Println(u.Nama)
	fmt.Println(u.Alamat.Jalan)

	cetakUser(u)

	//struct anonim
	animal := struct {
		Nama    string
		Habitat string
	}{
		Nama:    "Hiu",
		Habitat: "Laut",
	}

	fmt.Println(animal)

}
