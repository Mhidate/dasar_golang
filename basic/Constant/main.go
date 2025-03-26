package main

import "fmt"

func main() {
	fmt.Println("Belajar konstanta (const)")

	var name = "Alice" // Bisa diubah
	name = "Bob"

	const appName = "MyApp" // Tidak bisa diubah
	// appName = "NewApp" //  Error

	fmt.Println(name)
	fmt.Println(appName)

	//deklarasi const yang bertipe data
	const (
		a int     = 10
		b float64 = 3.14
		c string  = "Hello"
	)
	fmt.Println(a, b, c)

	//const tidak bisa di deklarasikan dengan menggunakan :=
	//const x := 10

	//tidak bisa menggunakan variabel sebagai nilai konstanta
	//var x = 10
	//const y = x

}
