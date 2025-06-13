package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Belajar paket time")
	/*
		Paket time menyediakan fungsi dan tipe untuk:
			Mengambil waktu sekarang
			Mengatur format dan parsing waktu
			Menghitung durasi antar waktu
			Delay/timer
			Scheduling event berdasarkan waktu
	*/

	// Ambil Waktu Sekarang
	now := time.Now()
	fmt.Println("Sekarang:", now)

	// Format Waktu,Go pakai layout unik: Mon Jan 2 15:04:05 MST 2006
	now2 := time.Now()
	fmt.Println("Tgl Custom:", now2.Format("2006-01-02 15:04:05"))

	// Sleep (Delay)
	fmt.Println("Mulai")
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai setelah 2 detik")

	// Durasi Antara Dua Waktu
	start := time.Now()
	time.Sleep(1 * time.Second)
	selisih := time.Since(start)
	fmt.Println("Durasi:", selisih)

	// Parsing String ke Waktu
	t, _ := time.Parse("2006-01-02", "2025-06-13")
	fmt.Println("Hasil parsing:", t)

	// Timer & Ticker,timer (1x trigger setelah delay):
	timer := time.After(3 * time.Second)
	fmt.Println("Tunggu 3 detik...")
	<-timer
	fmt.Println("Timer selesai")

	// Ticker (trigger berkala):
	ticker := time.Tick(1 * time.Second)
	for i := 0; i < 3; i++ {
		<-ticker
		fmt.Println("1 detik berlalu")
	}
}
