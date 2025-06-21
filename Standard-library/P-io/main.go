package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Belajar paket io")
	/*
		 Paket io di Go berisi interface dan fungsi-fungsi dasar untuk operasi input/output (I/O).
			Biasanya dipakai untuk:
			baca dari file, socket, atau buffer
			tulis ke file, socket, atau buffer
			copy data antar stream
			limit data stream
			dan lain-lain
	*/

	//Baca File Pakai io.Reader
	file, _ := os.Open("data.txt")
	defer file.Close()

	data, _ := io.ReadAll(file)
	fmt.Println(string(data))

	//Copy File Pakai io.Copy
	src, _ := os.Open("data.txt")
	dst, _ := os.Create("copy.txt")

	io.Copy(dst, src)

	src.Close()
	dst.Close()

	//Limit Baca Maksimal 5 Byte
	file2, _ := os.Open("data.txt")
	defer file2.Close()

	limited := io.LimitReader(file2, 5)
	data2, _ := io.ReadAll(limited)
	fmt.Println(string(data2))

	//TeeReader ,saat data dibaca dari source (Reader) maka langsung dicopy juga ke writer lain (misal: log file, buffer, dsb)
	//TeeReader bukan buffer. Dia harus dibaca sampai habis baru prosesnya jalan.
	src2, _ := os.Open("data.txt")
	defer src2.Close()

	logFile, _ := os.Create("log.txt")
	defer logFile.Close()

	tee := io.TeeReader(src2, logFile)
	data3, _ := io.ReadAll(tee)

	fmt.Println(string(data3))
}
