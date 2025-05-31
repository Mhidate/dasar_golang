package main

import "fmt"

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
	nama := "Andi"
	fmt.Printf("Halo, nama saya %s dan umur saya %d tahun.\n", nama, umur)

}
