package main

import "fmt"

//function dengan nama awal huruf kecil hanya bisa di panggil di dalam paket tempat dibuat
//sedangkan yang yang nama awal huruf besar bisa di panggil di luar paket
func salam() {
	fmt.Println("Halo,selamat siang")
}
func eksekusi(f func(string)) {
	f("Dewi")
}

func sapa(nama string) {
	fmt.Println("Hai", nama)
}

func pembuatSapa() func(string) {
	return func(nama string) {
		fmt.Println("Hai", nama)
	}
}

func main() {
	fmt.Println("Belajar tentang Function")

	salam()

	//Function bisa diperlakukan seperti data lainnya, misalnya:
	//Disimpan dalam variabel
	greet := func(nama string) {
		fmt.Println("Halo", nama)
	}

	greet("Andi")

	//Dikirim sebagai parameter ke fungsi lain
	eksekusi(sapa)

	//Dikembalikan dari fungsi
	sapa := pembuatSapa()
	sapa("Budi")

}
