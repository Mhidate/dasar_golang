package main

import "fmt"

func isEven(number int) bool {
	return number%2 == 0
}

func main() {
	fmt.Println("Belajar operasi boolean atau operator logika")

	//Dalam Go, jika variabel boolean tidak diinisialisasi, maka nilainya default false.
	a := true
	b := false

	fmt.Println("AND:", a && b)
	fmt.Println("OR:", a || b)
	fmt.Println("NOT a:", !a)

	c := 10
	d := 20
	e := 30

	fmt.Println((c < d) && (d < e))
	fmt.Println((c > d) || (d < e))
	fmt.Println(!(c == e))

	fmt.Println(isEven(10))
	fmt.Println(isEven(7))

	status := true

	switch status {
	case true:
		fmt.Println("Status aktif")
	case false:
		fmt.Println("Status tidak aktif")
	}

}
