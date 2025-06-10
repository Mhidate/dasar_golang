package main

import (
	"container/ring"
	"fmt"
)

func main() {

	fmt.Println("Belajar container ring")
	/*
		Paket container/ring adalah implementasi circular linked list di Go.
			Jadi:
			Bentuknya kayak lingkaran
			Tiap node (*Ring) punya pointer ke elemen berikutnya
			Kalau sudah di ujung, lanjut ke awal lagi (muter terus)
	*/

	// Buat ring isi 5 elemen
	r := ring.New(5)

	// Isi nilai
	for i := 0; i < r.Len(); i++ {
		r.Value = i + 1
		r = r.Next()
	}

	// Iterasi isinya
	fmt.Println("Isi ring:")
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})

	// Akses muter-muter
	fmt.Println("\nMuter 7 kali:")
	for i := 0; i < 7; i++ {
		fmt.Println("Posisi ke", i+1, ":", r.Value)
		r = r.Next()
	}
}
