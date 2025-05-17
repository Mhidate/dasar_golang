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

	angka := 10
	fmt.Println("Sebelum:", angka)

	tambahLima(&angka)

	fmt.Println("Sesudah:", angka)

	user := User{"Budi"}
	ubahNama(&user, "Andi")
	fmt.Println(user.Nama)
}
