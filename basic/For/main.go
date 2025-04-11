package main

import "fmt"

func main() {
	fmt.Println("Belajar for loop atau perulangan")

	for a := 0; a < 10; a++ {
		fmt.Println("Nilai variabel a=", a)

	}

	//for tanpa kondisi
	fmt.Println("Perulangan tanpa henti")
	i := 1
	for {
		fmt.Println("Nilai variabel i=", i)
		if i == 3 {
			break
		}
		i++
		// break //agar tidak melakukan perulangan secara terus menerus
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

	//hanya kondisi(seperti while)
	w := 1
	for w <= 5 {
		fmt.Println("Nilai variabel w=", w)
		w++
	}

	// penjumlahan bilangan dari 1-100 ,output yang diharapkan 5050
	b1 := 0
	for b2 := 1; b2 <= 100; b2++ {
		b1 += b2
	}
	fmt.Println("Hasil penjumlahan =", b1)

	//tampilkan hanya bilangan genap dari 1-20
	for g := 1; g <= 20; g++ {
		p := 2
		if g%p == 0 {
			fmt.Println("Bilangan genap dari 1-20=", g)
		}
	}

	//tabel perkalian 1-3 (output harus berhenti di 3x3=9)
	for k := 1; k <= 3; k++ {
		fmt.Println("Perkalian ", k)
		for f := 1; f <= 9; f++ {
			fmt.Println(k, "x", f, "=", k*f)
			if k*f == 9 {
				break
			}

		}
	}

	//perkalian 1-3 hanya sampai 3 ,contoh 1x3,2x3,3x3
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
	}
}
