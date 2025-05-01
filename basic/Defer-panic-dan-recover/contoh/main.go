package main

import (
	"fmt"
	"os"
	"time"
)

func logAkhir() {
	fmt.Println("[LOG] Program selesai")
}

func simpanLog() {
	// Buka atau buat file log.txt
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[ERROR] Gagal buka file log:", err)
		return
	}
	defer file.Close()

	// Tambahkan timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] Menyimpan log ke file log.txt\n", timestamp)

	// Tulis ke file
	if _, err := file.WriteString(logMessage); err != nil {
		fmt.Println("[ERROR] Gagal tulis ke file log:", err)
	} else {
		fmt.Println("[LOG] Log berhasil disimpan ke log.txt")
	}
}

func tangkapPanic() {
	if r := recover(); r != nil {
		fmt.Println("[RECOVER] Ada panic ketangkep:", r)
		fmt.Println("[RECOVER] Program dilanjutkan dengan aman ")
	}
}

func bacaFile(namaFile string) {
	defer logAkhir()
	defer simpanLog()
	defer tangkapPanic()

	fmt.Println("Proses buka file:", namaFile)

	file, err := os.Open(namaFile)
	if err != nil {
		panic("[PANIC] File tidak ditemukan: " + namaFile)
	}
	defer file.Close()
	fmt.Println("Berhasil buka file:", namaFile)
}

func main() {
	fmt.Println("Program dimulai ")
	bacaFile("data.txt")
	fmt.Println("Program tetap berjalan lanjut ")
}
