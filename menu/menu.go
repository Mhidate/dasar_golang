package menu

/*

v "dasar_golang/tampil" = di impor lalu di berikan alias,pemanggilannya sesuai alias v.<nama fungsi>

"dasar_golang/tampil" = di impor tanpa alias,pemanggilannya sesuai nama paket  variatipe.<nama fungsi>

*/

import (
	"dasar_golang/tampil"
	"fmt"
	"os"
)

func Pilihmenu() {

	mapMenu := map[string]func(){
		"1.Kalkulator": tampil.Prin,
		"2.Bangun datar": func() {
			fmt.Println("Bangun datar")
		},
		"0.Keluar": func() {
			fmt.Println("Terima kasih telah menggunakan program ini")
			os.Exit(0)
		},
	}

	for {
		fmt.Println()
		fmt.Println("====Hello World====")
		fmt.Println()
		fmt.Println("Menu")
		for key := range mapMenu {
			fmt.Println(key)
		}
	
		
		var pilihan string
		fmt.Println()
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilihan)
		
	
		switch pilihan {
		case "1":
			mapMenu["1.Kalkulator"]()
		case "2":
			mapMenu["2.Bangun datar"]()
		case "0":
			mapMenu["0.Keluar"]()
	
		}
	}
   

}

// for {
// 	var menu int
// 	fmt.Println("hello world")
// 	fmt.Println()
// 	fmt.Println("Menu")
// 	fmt.Println("1.Variabel")
// 	fmt.Print("Pilih Menu : ")
// 	fmt.Scanln(&menu)

// 	if menu == 1 {
// 		tampil.Prin()
// 	}else {
// 		break
// 	}

// }
