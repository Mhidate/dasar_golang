package main

import "fmt"

type User struct {
	Nama string
}

func (u User) GantiNamaBaru(nama string) {
	u.Nama = nama // cuma mengubah salinan, bukan struct aslinya
}

func (u *User) GantiNamaBaru2(nama string) {
	u.Nama = nama // langsung ke alamat aslinya
}

func main() {
	fmt.Println("Belajar pointer di method")

	/*
			 Di Go, method itu function yang menempel ke struct.
		     Method bisa punya receiver:
		     Value receiver → method menerima salinan struct.
		     Pointer receiver → method menerima alamat struct, jadi bisa ubah datanya langsung.
	*/

	// Contoh Value Receiver (tidak mengubah data asli)
	user := User{"John"}
	user.GantiNamaBaru("Jenny")
	fmt.Println("Nama setelah method:", user.Nama) // tetap "John"

	/*
			 Kenapa?
		     Karena GantiNamaBaru pakai value receiver (u User), jadi yang diproses itu salinan user, bukan data aslinya.
	*/

	// Contoh Pointer Receiver (mengubah data asli)
	user2 := User{"Jhonas"}
	user2.GantiNamaBaru2("Juny")
	fmt.Println("Nama setelah method:", user2.Nama) // jadi "Juny"

}
