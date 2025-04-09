package main

import "fmt"

func main() {
	fmt.Println("Belajar menggunakan switch")
	//switch di go memiliki perbedaan dengan bahasa pemrograman lain,jika kondisi case sudah terpenuhi maka tidak akan lanjut untuk cek case selanjutnya
	//jika ingin case selanjutnya tetap di cek walapun pada case sebelumnya kondisi sudah terpenuhi maka gunakan fallthrough
	//efeknya adalah case di pengecekan selanjutnya selalu dianggap true (meskipun aslinya bisa saja kondisi tersebut tidak terpenuhi, akan tetap dianggap true).

	lampu := "merah"

	switch lampu {
	case "hijau":
		fmt.Println("Jalan")
	case "merah":
		fmt.Println("Berhenti")
	case "Kuning":
		fmt.Println("Hati-hati")
	}

	//contoh case untuk banyak kondisi dan kurung Kurawal Pada Keyword case & default
	var point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		{
			fmt.Println("awesome")
			fmt.Println("Lebih giat lagi")
		}
	default:
		{

			fmt.Println("not bad")
			fmt.Println("Semangat!!!!!")
		}
	}

	//switch gaya if else
	var point2 = 6

	switch {
	case point2 == 8:
		fmt.Println("perfect2")
	case (point2 < 8) && (point2 > 3):
		fmt.Println("awesome2")
	default:
		{
			fmt.Println("not bad2")
			fmt.Println("you need to learn more2")
		}
	}

	//contoh penggunaan fallthrough
	var point3 = 6

	switch {
	case point3 == 8:
		fmt.Println("perfect3")
	case (point3 < 8) && (point3 > 3):
		fmt.Println("awesome3")
		fallthrough
	case point3 < 5:
		fmt.Println("you need to learn more3")
	default:
		{
			fmt.Println("not bad3")
			fmt.Println("you need to learn more3")
		}
	}

	//Short Statement di Switch
	//variabel angka hanya berlaku di dalam blok switch tersebut.
	switch angka := 5; angka {
	case 1:
		fmt.Println("Satu")
	case 5:
		fmt.Println("Lima")
	default:
		fmt.Println("Tidak diketahui")
	}

}
