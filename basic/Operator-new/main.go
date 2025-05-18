package main

import "fmt"

type User struct {
	Nama string
	Umur int
}

func main() {
	fmt.Println("Belajar Operator new di pointer")

	/*
		new adalah built-in function di Go yang dipakai untuk mengalokasikan memory baru untuk tipe data tertentu dan mengembalikan pointer ke memory itu.
	*/

	// Menggunakan operator new untuk buat pointer ke int
	var p *int = new(int)

	// Nilai default dari int adalah 0
	fmt.Println("Nilai awal pointer:", *p) // dereference

	// Ubah nilai lewat dereference
	*p = 100

	fmt.Println("Nilai setelah diubah:", *p)

	user1 := User{"Dewi", 22}
	ptrUser1 := &user1

	fmt.Println("\nNama user1:", ptrUser1.Nama)
	fmt.Println("Umur user1:", ptrUser1.Umur)

	// Ubah lewat pointer
	ptrUser1.Umur = 23
	fmt.Println("Umur user1 setelah diubah:", user1.Umur)

	ptrUser2 := new(User) // User{} zero value → Nama "" Umur 0

	fmt.Println("\nNilai awal Nama user2:", ptrUser2.Nama)
	fmt.Println("Nilai awal Umur user2:", ptrUser2.Umur)

	// Isi data lewat pointer
	ptrUser2.Nama = "Rina"
	ptrUser2.Umur = 30

	fmt.Println("Nama user2 setelah diisi:", ptrUser2.Nama)
	fmt.Println("Umur user2 setelah diisi:", ptrUser2.Umur)
}
