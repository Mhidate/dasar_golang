package main

import (
	"fmt"
	"time"
)

// LogSystem struct dengan slice interface{}
type LogSystem struct {
	Logs []interface{}
}

// Method untuk menambahkan log
func (ls *LogSystem) TambahLog(data interface{}) {
	waktu := time.Now().Format("2006-01-02 15:04:05")
	log := fmt.Sprintf("[%s] %v", waktu, data)
	ls.Logs = append(ls.Logs, log)
}

// Method untuk menampilkan semua log
func (ls LogSystem) TampilLog() {
	fmt.Println("==== Log Aktif ====")
	for _, log := range ls.Logs {
		fmt.Println(log)
	}
}

func main() {
	fmt.Println("Belajar interface kosong")
	// interface yang tidak memiliki method sama sekali ,jadi bisa menampung semua jenis tipe data.

	dataCampuran := []interface{}{"String", 100, true, 3.14}

	for _, data := range dataCampuran {
		fmt.Println(data)
	}

	person := map[string]interface{}{
		"nama":   "Adit",
		"umur":   21,
		"status": true,
	}

	fmt.Println(person["nama"])
	fmt.Println(person["umur"])
	fmt.Println(person["status"])

	// Inisialisasi log system
	logSystem := LogSystem{}

	// Tambahkan beberapa log
	logSystem.TambahLog("Program dimulai")
	logSystem.TambahLog(404)
	logSystem.TambahLog(true)
	logSystem.TambahLog(3.14)

	// Tampilkan semua log
	logSystem.TampilLog()
}
