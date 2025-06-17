package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

func main() {

	fmt.Println("Belajar paket encoding")
	/*
		Di Go, encoding bukan 1 package langsung dipakai, tapi sekelompok sub-package di dalam standard library yang dipakai untuk:
			Encode data ke format tertentu
			Decode data dari format tertentu
			Umumnya buat keperluan data interchange kayak:
			Kirim data ke client (misal JSON, XML)
			Simpan config
			Transfer antar service via HTTP API
	*/

	//Encode struct ke JSON
	user := User{"Dian", 25}
	data, _ := json.Marshal(user)
	fmt.Println(string(data))

	//Decode JSON ke struct
	var jsonData = []byte(`{"nama":"Dian","umur":25}`)
	var u User
	json.Unmarshal(jsonData, &u)
	fmt.Println(u.Nama, u.Umur)

}
