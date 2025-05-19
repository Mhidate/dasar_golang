package main

import "fmt"

type User struct {
	Nama string
	Umur int
}

// Function untuk ubah nilai string lewat pointer
func ubahNama(nama *string) {
	*nama = "Dino"
}

// Function untuk tambah umur lewat pointer
func tambahUmur(umur *int) {
	*umur += 1
}

// Function untuk ubah struct lewat pointer
func updateUser(u *User) {
	u.Nama = "Eka"
	u.Umur = 30
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

	///contoh dengan fungsi
	nama := "Budi"
	fmt.Println("Nama sebelum diubah:", nama)

	ubahNama(&nama)
	fmt.Println("Nama setelah diubah:", nama)

	umur := new(int)
	*umur = 25
	fmt.Println("\nUmur sebelum tambah:", *umur)

	tambahUmur(umur)
	fmt.Println("Umur setelah tambah:", *umur)

	user2 := User{"Dewi", 22}
	fmt.Println("\nUser1 sebelum update:", user2)

	updateUser(&user2)
	fmt.Println("User1 setelah update:", user2)

	user3 := new(User)
	user3.Nama = "Rina"
	user3.Umur = 20

	fmt.Println("\nUser2 sebelum update:", *user3)

	updateUser(user3)
	fmt.Println("User2 setelah update:", *user3)
}
