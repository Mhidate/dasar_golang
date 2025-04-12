package main

import "fmt"

//function dengan nama awal huruf kecil hanya bisa di panggil di dalam paket tempat dibuat
//sedangkan yang yang nama awal huruf besar bisa di panggil di luar paket
func salam() {
	fmt.Println("Halo,selamat siang")
}
func main() {
	fmt.Println("Belajar tentang Function")
	salam()
}
