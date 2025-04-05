package main

import "fmt"

func main() {
	fmt.Println("Belajar tipe data array")

	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// cara inisialisasi horizontal
	var fruits2 = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	var fruits3 = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
	fmt.Println("Isi semua element \t", fruits2)
	fmt.Println("Isi semua element \t", fruits3)

	// Gunakan ellipsis (...) untuk menentukan panjang otomatis
	data := [...]int{1, 2, 3, 4, 5} // panjang array otomatis = 5
	fmt.Println("data array \t:", data)
	fmt.Println("jumlah elemen \t:", len(data))

	// array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// memanggil array dengan for
	var fruits4 = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits4[i])
	}

	//dengan for range
	var fruits5 = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits5 {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// dengan underscore
	var fruits6 = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit2 := range fruits6 {
		fmt.Printf("nama buah : %s\n", fruit2)
	}

	// Deklarasi sekaligus alokasi kapasitas array juga bisa dilakukan lewat keyword make.
	var fruits7 = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits7)
}
