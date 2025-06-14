package main

import (
	"fmt"
	"reflect"
)

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
	var x int = 42

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("Tipe :", t)
	fmt.Println("Value:", v)
}
