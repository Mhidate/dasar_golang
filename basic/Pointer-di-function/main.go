package main

import "fmt"

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	fmt.Println("Belajar tentang pointer di function")

	/*
		Di Go, kita bisa mengirim alamat memory (pointer) ke function supaya function tersebut bisa mengubah nilai variabel asli di luar function.
	*/

	a := 5
	b := 10

	fmt.Println("Sebelum pertukaran: a =", a, ", b =", b)
	swap(&a, &b) // Mengirim alamat memori a dan b ke fungsi swap
	fmt.Println("Setelah pertukaran: a =", a, ", b =", b)
}
