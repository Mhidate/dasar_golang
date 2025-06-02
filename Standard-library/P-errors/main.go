package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("data tidak ditemukan")

func getData() error {
	return fmt.Errorf("query gagal: %w", ErrNotFound)
}

type CustomError struct {
	Msg string
}

func (c CustomError) Error() string {
	return c.Msg
}

func main() {
	fmt.Println("Belajar paket errors")
	/*
		errors adalah paket standar di Go untuk membuat dan memanipulasi error.
		Error di Go itu tipe data interface yang didefinisikan seperti ini:
			type error interface {
				Error() string
			}
	*/

	err := errors.New("data tidak valid")
	fmt.Println(err)

	//fmt.Errorf + %w (wrap error),kamu bisa membungkus error lain pakai fmt.Errorf + %w
	originalErr := errors.New("file tidak ditemukan")
	wrappedErr := fmt.Errorf("gagal proses data: %w", originalErr)
	fmt.Println(wrappedErr)

	/*
		Kamu bikin error baru: gagal proses: [wrapped error]
		Yang di-wrap itu cuma errors.New("file tidak ada")
		fmt.Errorf dengan %w membungkus error lama di dalam error baru
	*/
	err2 := fmt.Errorf("gagal proses: %w", errors.New("file tidak ada"))
	cause := errors.Unwrap(err2)
	fmt.Println(cause)

	//errors.Is,untuk membandingkan apakah sebuah error sama atau termasuk di chain-nya:
	err3 := getData()
	if errors.Is(err3, ErrNotFound) {
		fmt.Println("Error karena data tidak ditemukan.")
	}

	//errors.As,untuk mengecek tipe error dan assign ke variabel kalau cocok. Biasanya dipakai kalau kamu pakai custom error type.
	err4 := CustomError{"terjadi kesalahan"}
	var customErr CustomError
	if errors.As(err4, &customErr) {
		fmt.Println("Custom error:", customErr.Msg)
	}
}
