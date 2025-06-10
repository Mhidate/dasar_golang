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

	// Buat queue ,FIFO — First In First Out
	queue := list.New()

	// Enqueue: masuk ke belakang
	queue.PushBack("A")
	queue.PushBack("B")
	queue.PushBack("C")

	// Tampilkan isi queue
	fmt.Println("Isi Queue:")
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Dequeue: keluar dari depan
	front := queue.Front()
	if front != nil {
		fmt.Println("\nDequeue:", front.Value)
		queue.Remove(front)
	}

	// Tampilkan isi queue setelah dequeue
	fmt.Println("\nIsi Queue setelah Dequeue:")
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Buat stack,  LIFO — Last In First Out
	stack := list.New()

	// Push: masuk ke belakang
	stack.PushBack("1")
	stack.PushBack("2")
	stack.PushBack("3")

	// Tampilkan isi stack
	fmt.Println("Isi Stack:")
	for e := stack.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	// Pop: ambil dari belakang
	top := stack.Back()
	if top != nil {
		fmt.Println("\nPop:", top.Value)
		stack.Remove(top)
	}

	// Tampilkan isi stack setelah pop
	fmt.Println("\nIsi Stack setelah Pop:")
	for e := stack.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
