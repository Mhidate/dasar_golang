package main

import (
	"container/list"
	"fmt"
)

func main() {

	fmt.Println("Belajar container list")
	/*
		Paket container/list adalah implementasi doubly linked list di Go.
		Tiap elemen (Element) saling terhubung ke depan dan ke belakang.
		Bisa push/pop di depan (Front) atau belakang (Back).
		Cocok buat struktur data yang sering insert & delete di awal/akhir tanpa harus geser elemen seperti slice.
	*/
	// Bikin list baru
	myList := list.New()

	// Tambah elemen ke belakang
	myList.PushBack("Go")
	myList.PushBack("Golang")
	myList.PushFront("Belajar")

	// Tampilkan semua isi list
	fmt.Println("Isi list:")
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Insert setelah elemen pertama
	first := myList.Front()
	myList.InsertAfter("Setelah Belajar", first)

	// Hapus elemen kedua
	second := first.Next()
	myList.Remove(second)

	fmt.Println("\nIsi list setelah perubahan:")
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Cek jumlah elemen
	fmt.Println("\nJumlah elemen:", myList.Len())
}
