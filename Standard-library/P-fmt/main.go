package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Belajar paket fmt")

	/*
		fmt adalah paket bawaan di Go (Golang) yang dipakai untuk memformat dan mencetak teks ke output standar (seperti terminal), atau ke string, atau ke file, dll.
		Nama fmt sendiri singkatan dari format.
	*/

	fmt.Println("Apa kabar ?")

	//fungsi Printf
	/*
		%s → string
		%d → integer
		%f → float
		%t → boolean
		%v → any value
		%+v → any value with field names (untuk struct)
	*/
	umur := 25
	nama := "John"
	fmt.Printf("Halo, nama saya %s dan umur saya %d tahun.\n", nama, umur)

	// ===== Sprint, Sprintln, Sprintf (hasilnya string) =====

	s1 := fmt.Sprint("Halo, ", nama)
	s2 := fmt.Sprintln("Halo,", nama)
	s3 := fmt.Sprintf("Halo, nama saya %s.", nama)

	fmt.Println("Sprint result:", s1)
	fmt.Print("Sprintln result:", s2)
	fmt.Println("Sprintf result:", s3)

	// ===== Fprint, Fprintln, Fprintf (ke file) =====
	file, err := os.Create("hasil.txt")
	if err != nil {
		fmt.Println("Gagal buat file:", err)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, "Baris ini ditulis ke file.")
	fmt.Fprintf(file, "Nama saya %s, umur %d tahun.\n", nama, 25)

	// ===== Scan, Scanf, Scanln (input dari terminal) =====
	var input1 string
	var input2 int
	fmt.Print("Masukkan nama kamu: ")
	fmt.Scan(&input1)
	fmt.Println("Hai,", input1)

	fmt.Print("Masukkan umur kamu: ")
	fmt.Scanf("%d", &input2)
	fmt.Printf("Umur kamu: %d tahun\n", input2)

	// ===== Sscan, Sscanf, Sscanln (input dari string) =====
	str := "Budi 30"
	var nama2 string
	var umur2 int
	fmt.Sscan(str, &nama2, &umur2)
	fmt.Printf("Dari string: %s umur %d tahun\n", nama2, umur2)

	// Contoh pakai Sscanf
	data := "Laptop 7500000"
	var barang string
	var harga int
	fmt.Sscanf(data, "%s %d", &barang, &harga)
	fmt.Printf("Barang: %s, Harga: %d\n", barang, harga)

	// ===== Errorf =====
	errMsg := fmt.Errorf("data tidak valid: %s", strings.ToUpper("error"))
	fmt.Println("Pesan error:", errMsg)
}
