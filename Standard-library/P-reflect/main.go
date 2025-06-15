package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Nama  string
	Umur  int
	Email string
}

func main() {

	fmt.Println("Belajar paket reflect")
	/*
		Paket reflect di Go dipakai untuk:
			Mengakses informasi tipe data saat runtime
			Membaca dan memodifikasi nilai variabel secara dinamis
			Biasa dipakai untuk:
			serialization
			ORM
			validator struct field
			generic log & debugging tool
			Karena Go itu statically typed language, tapi kadang di runtime kita pengen bisa tahu tipe, nama, atau isi field dari suatu data tanpa hardcoded — nah di situ reflect dipakai.
	*/

	// Dapatkan Tipe & Value
	var x int = 42

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("Tipe :", t)
	fmt.Println("Value:", v)

	// Iterasi Field Struct Pakai Reflect
	u := User{"Dian", 25, "dian@mail.com"}

	val := reflect.ValueOf(u)
	typ := reflect.TypeOf(u)

	for i := 0; i < val.NumField(); i++ {
		fmt.Printf("Field %d: %s = %v\n", i, typ.Field(i).Name, val.Field(i))
	}

	// Set Value Dinamis (harus pointer)
	var y int = 10
	z := reflect.ValueOf(&y).Elem()

	fmt.Println("Sebelum:", y)
	z.SetInt(99)
	fmt.Println("Sesudah:", y)
}
