package main

import "fmt"

func main() {

	//Type Declaration di Golang adalah fitur yang memungkinkan kita untuk membuat alias atau mendefinisikan tipe data baru berdasarkan tipe data yang sudah ada.
	fmt.Println("Belajar type declarations")

	type name string
	type addres string
	type age int16

	var nama name = "Jhone"
	var alamat addres = "London"
	var umur age = 21

	fmt.Println("----------------")
	fmt.Println(nama)
	fmt.Println(alamat)
	fmt.Println(umur)
	fmt.Println("----------------")

	fmt.Println("Mencoba penjumlahan dengan tipe asli dan alias")

	//nama variabel tidak boleh sama dengan nama type declarations atau alias
	type angka1 int16
	type angka2 int16

	var angka12 angka1 = 5
	var angka22 angka2 = 6
	// var angka3 int16 = 8

	fmt.Println(angka12)
	fmt.Println(angka22)
	fmt.Println("---------------")
	// fmt.Println(angka12 + angka22) tidak bisa melakukan operasi dengan type declarations yang berbeda atau type data asli
	// fmt.Println(angka12 + angka3) variabel angka3 harus di ubah ke type angka1 agar bisa menjalankan operasi

	fmt.Println("Mencoba membuat type declarations dengan struct")

	type Animal struct {
		Nama    string
		Habitat string
	}

	var hewan Animal = Animal{Nama: "Ikan lele", Habitat: "Air Tawar"}
	fmt.Println(hewan.Nama, hewan.Habitat)
	fmt.Println("---------------")

	fmt.Println("Mencoba membuat type declarations dengan Interface")

	var s Speaker = Human{Name: "Alisia"}
	fmt.Println(s.Speak())
	fmt.Println("---------------")
	fmt.Println("---------------")

	fmt.Println("Mencoba membuat type declarations dengan Function")
	var operation AddFunc = add
	fmt.Println(operation(3, 4))
	fmt.Println("---------------")

	fmt.Println("Mencoba membuat type declarations dengan Slice dan Map")
	type Names []string         // Alias untuk slice string
	type Role map[string]string // Alias untuk map[string]string

	var pegawai Names = []string{"Alice", "Bob", "Charlie"}
	fmt.Println(pegawai)

	var staff Role = map[string]string{
		"backend":  "Backend Go",
		"frontend": "Frontend Vue.js",
	}
	fmt.Println(staff["backend"])

}

type Speaker interface {
	Speak() string
}

type Human struct {
	Name string
}

func (h Human) Speak() string {

	return "Hello, my name is " + h.Name
}

type AddFunc func(int, int) int // Alias untuk function yang menerima dua int dan mengembalikan int

func add(a int, b int) int {
	return a + b
}
