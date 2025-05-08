package main

import "fmt"

// Aturan interface
type Hewan interface {
	Bersuara() // Semua yang disebut Hewan harus bisa Bersuara
}

// Ini struct Kucing
type Kucing struct{}

// Kucing mengikuti aturan Hewan (punya method Bersuara)
func (k Kucing) Bersuara() {
	fmt.Println("Meong!") // Cara Kucing bersuara
}

// Ini struct Anjing
type Anjing struct{}

// Anjing juga mengikuti aturan Hewan
func (a Anjing) Bersuara() {
	fmt.Println("Guk Guk!")
}

//Ini struct burung
type Burung struct{}

// Burung juga mengikuti aturan Hewan
func (b Burung) Bersuara() {
	fmt.Println("Cuit cuit!")
}

// Function SuaraHewan menerima parameter tipe Hewan (bukan struct)
func SuaraHewan(h Hewan) {
	h.Bersuara() // Panggil suara sesuai siapa yang dikirim
}

func main() {
	fmt.Println("Belajar tentang interface")
	//Interface di Go adalah tipe data abstrak yang mendefinisikan kumpulan method tanpa implementasi.

	kucing := Kucing{}
	anjing := Anjing{}

	SuaraHewan(kucing)
	SuaraHewan(anjing)

}
