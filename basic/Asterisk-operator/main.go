package main

import "fmt"

func tambahLima(nilai *int) {
	*nilai += 5
}

type User struct {
	Nama string
}

func ubahNama(u *User, namaBaru string) {
	u.Nama = namaBaru
}

func main() {
	fmt.Println("Belajar Asterisk Operator di pointer")

	/*
		Asterisk (*) operator dalam Go memiliki dua fungsi utama terkait pointer:
		 mendeklarasikan pointer dan mendereferensikan pointer.
	*/

	// Deklarasi variabel biasa
	var angka int = 10

	// Deklarasi pointer: p adalah pointer ke int
	var p *int = &angka

	fmt.Println("Alamat angka:", &angka) // Tampilkan alamat memori angka
	fmt.Println("Isi pointer (alamat angka):", p)
	fmt.Println("Nilai di alamat pointer (nilai angka):", *p)

	// Ubah nilai lewat dereference pointer
	*p = 20
	fmt.Println("Nilai angka setelah diubah lewat pointer:", angka)

	angka2 := 10
	fmt.Println("Sebelum:", angka2)

	tambahLima(&angka2)

	fmt.Println("Sesudah:", angka2)

	user := User{"Budi"}
	ubahNama(&user, "Andi")
	fmt.Println(user.Nama)
}
