package main

import "fmt"

func main() {
	fmt.Println("Belajar for loop atau perulangan")

	for a := 0; a < 10; a++ {
		fmt.Println("Nilai variabel a=", a)

	}

	//for tanpa kondisi
	for {
		fmt.Println("Perulangan tanpa henti")
		break //agar tidak melakukan perulangan secara terus menerus
	}

	// perulangan slice dan map
	bulan := []string{"Januari", "Februari", "Maret"}

	for i, v := range bulan {
		fmt.Println(i, v)
	}

	angka := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range angka {
		fmt.Println(k, v)
	}
}
