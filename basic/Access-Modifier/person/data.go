package person

import "fmt"

// Exported (bisa diakses luar package)
var PublicNama string = "Budi"

// Unexported (hanya bisa diakses dalam package person)
var privateAlamat string = "Jakarta"

// Exported struct
type Person struct {
	Nama string // Exported field
	umur int    // Unexported field
}

// Exported function
func SayHello() {
	fmt.Println("Halo dari package person!")
}

// Unexported function
func sayAlamat() {
	fmt.Println("Alamat:", privateAlamat)
}

// Method (Exported)
func (p Person) CetakNama() {
	fmt.Println("Nama saya", p.Nama)
}

// Method (Unexported)
func (p Person) cetakUmur() {
	fmt.Println("Umur saya", p.umur)
}
