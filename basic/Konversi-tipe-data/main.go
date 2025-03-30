package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Belajar konversi tipe data")

	//float ke int
	var num float64 = 9.99
	var integer int = int(num) // Hanya mengambil bagian bulatnya

	fmt.Println(integer)

	str := "100"
	num2, err := strconv.Atoi(str) // Konversi string ke int
	if err == nil {
		fmt.Println(num2) // Output: 100
	}

	num3 := 100
	str2 := strconv.Itoa(num3) // Konversi int ke string

	fmt.Println(str2) // Output: "100"

	str3 := "3.14"
	num4, err := strconv.ParseFloat(str3, 64) // Konversi string ke float64
	if err == nil {
		fmt.Println(num4) // Output: 3.14
	}

	num5 := 3.14
	str4 := strconv.FormatFloat(num5, 'f', 2, 64) // Konversi float ke string

	fmt.Println(str4) // Output: "3.14"

	val := true
	str6 := strconv.FormatBool(val) // Konversi bool ke string

	fmt.Println(str6) // Output: "true"

	str7 := "true"
	val2, err := strconv.ParseBool(str7) // Konversi string ke bool
	if err == nil {
		fmt.Println(val2) // Output: true
	}

	str8 := "Hello"
	bytes := []byte(str8) // Konversi string ke byte slice
	fmt.Println(bytes)    // Output: [72 101 108 108 111]

	bytes2 := []byte{72, 101, 108, 108, 111}
	str9 := string(bytes2) // Konversi byte slice ke string
	fmt.Println(str9)      // Output: "Hello"

	var a interface{} = Animal{Name: "Dog"}
	animal := a.(Animal)     // Type assertion dari interface ke struct
	fmt.Println(animal.Name) // Output: Dog

}

type Animal struct {
	Name string
}
