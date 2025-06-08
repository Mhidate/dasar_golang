package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Belajar paket math")
	/*
		Paket math di Go berisi kumpulan fungsi-fungsi untuk operasi matematika:
			Pangkat
			Akar kuadrat
			Trigonometri
			Logaritma
			Pembulatan (ceil, floor, round)
			Konstanta matematika (π, e)
	*/
	x := 9.0
	y := 2.0

	// Nilai absolut
	fmt.Println("Abs(-9):", math.Abs(-9))

	// Pangkat
	fmt.Println("Pow(9, 2):", math.Pow(x, y))

	// Akar kuadrat
	fmt.Println("Sqrt(9):", math.Sqrt(x))

	// Akar pangkat 3
	fmt.Println("Cbrt(27):", math.Cbrt(27))

	// Pembulatan
	fmt.Println("Ceil(3.3):", math.Ceil(3.3))
	fmt.Println("Floor(3.7):", math.Floor(3.7))
	fmt.Println("Round(3.5):", math.Round(3.5))

	// Max & Min
	fmt.Println("Max(10, 20):", math.Max(10, 20))
	fmt.Println("Min(10, 20):", math.Min(10, 20))

	// Trigonometri (pakai radian)
	fmt.Println("Sin(90°):", math.Sin(math.Pi/2))

	// Logaritma
	fmt.Println("Log(e):", math.Log(math.E))
	fmt.Println("Log10(100):", math.Log10(100))

	// Eksponen
	fmt.Println("Exp(1):", math.Exp(1))
}
