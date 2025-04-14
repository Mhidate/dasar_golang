package main

import "fmt"

//penulisan
// func namaFungsi(namaParam1 tipe1, namaParam2 tipe2) returnTipe {
// kode fungsi
// }

func ganjil(a int) {
	if a%2 != 0 {
		fmt.Println("Angka ", a, "adalah gnjil")
	}
}

func genap(a int) string {
	var hasil string
	if a%2 == 0 {
		hasil = "genap"

	} else {
		hasil = "ganjil"
	}
	return hasil
}

//menyikat parameter dengan tipe yang sama
func kurang(a, b int) int {
	return a - b
}

//struct sebagai parameter
type User struct {
	Nama string
	Id   int
}

func cetakUser(u User) {
	fmt.Println("Nama:", u.Nama)
	fmt.Println("ID:", u.Id)
}

//slice sebagai parameter
func cetakSlice(data []string) {
	for i, v := range data {
		fmt.Printf("Index %d: %s\n", i, v)
	}
}

//array sebagai parameter
func cetakArray(data [3]int) {
	for i, v := range data {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}

//map sebagai parameter
func cetakMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
}

//interface sebagai parameter
func cetakInterface(data interface{}) {
	fmt.Println("Isi interface:", data)
}

//slice dan interface sebagai parameter
func cetakSemua(data []interface{}) {
	for i, v := range data {
		fmt.Printf("Index %d: %v\n", i, v)
	}
}

func main() {
	fmt.Println("Belajar function parameter")

	//Parameter adalah variabel sementara yang digunakan untuk menerima nilai dari luar ketika fungsi dipanggil
	ganjil(5)

	a := 8
	hasil := genap(a)

	fmt.Println("Angka", a, "adalah", hasil)
	c := 10
	d := 6
	b := kurang(c, d)
	fmt.Println(c, "-", d, "=", b)

	user := User{
		Nama: "Andi",
		Id:   123,
	}
	cetakUser(user)

	nama := []string{"Andi", "Budi", "Citra"}
	cetakSlice(nama)

	angka1 := [3]int{10, 20, 30}
	cetakArray(angka1)

	data := map[string]string{
		"nama":  "Dewi",
		"email": "dewi@example.com",
	}
	cetakMap(data)

	cetakInterface("Halo Dunia")
	cetakInterface(123)
	cetakInterface([]int{1, 2, 3})

	koleksi := []interface{}{"Teks", 42, true}
	cetakSemua(koleksi)

}
