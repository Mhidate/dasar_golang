package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Belajar paket bufio")
	/*
		bufio (Buffered I/O) adalah package di Go yang menyediakan I/O wrapper dengan buffer memory.
			Artinya:
			Saat baca/tulis file, socket, stdin, dsb
			bufio akan menyimpan data sementara di buffer memory
			Baru proses / kirim data itu sekaligus, bukan byte per byte
	*/

	//Membaca file per baris (pakai bufio.Reader)
	file, _ := os.Open("data.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}

	//Menulis ke file dengan buffer (pakai bufio.Writer)
	file2, _ := os.Create("output.txt")
	writer := bufio.NewWriter(file2)

	writer.WriteString("Halo Dunia!\n")
	writer.WriteString("Baris ke-2\n")

	writer.Flush() // WAJIB, biar buffer ditulis ke file
	file.Close()

	//Membaca input dari terminal (stdin)
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Ketik sesuatu: ")
	text, _ := reader2.ReadString('\n')
	fmt.Println("Kamu ketik:", text)

	//Baca per baris pakai bufio.Scanner,lebih ringan dan aman daripada ReadString
	file3, _ := os.Open("data.txt")
	defer file3.Close()

	scanner := bufio.NewScanner(file3)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
