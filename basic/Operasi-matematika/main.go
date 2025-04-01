package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Belajar operasi matematika")

	a := 10
	b := 3
	c := 3.0
	d := 2.0

	fmt.Println("Penjumlahan:", a+b)
	fmt.Println("Pengurangan:", a-b)
	fmt.Println("Perkalian:", a*b)
	fmt.Println("Pembagian:", c/d)
	fmt.Println("Modulus:", a%b)
	fmt.Println("------------------")

	fmt.Println("Operasi matematika menggunakan paket math")

	fmt.Println("Nilai absolut:", math.Abs(-10.5))
	fmt.Println("Pangkat:", math.Pow(2, 3))
	fmt.Println("Akar kuadrat:", math.Sqrt(25))
	fmt.Println("Pembulatan ke atas:", math.Ceil(4.3))
	fmt.Println("Pembulatan ke bawah:", math.Floor(4.3))
	fmt.Println("Pembulatan ke terdekat:", math.Round(4.3))

}
