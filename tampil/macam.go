package tampil

import"fmt"

/*
macam macam variabel dan tipe data
*/

func Prin(){
	fmt.Println("====BELAJAR VARIABEL (MANIFEST TYPING)====")
	var nama string = "jhone"
	var umur int = 12
	var status bool = false
	var tinggi float64 = 12.2
	fmt.Printf("halo nama saya %s,umur saya %d,status %t,dan tinggi saya %.2f\n", nama, umur, status, tinggi)
	fmt.Println("=========================================")

	fmt.Println("====BELAJAR VARIABEL (TYPE INFERENCE)====")
	var firstname = "jhon" //menggunakan var 
	lastname := "kim" //tanpa var
	age := 14
	statuss := true
	high := 23.3
	fmt.Println("Nama :",firstname,lastname,"umur :",age,"nikah :",statuss,"tinggi",high)
	fmt.Println("========================================")
	
	fmt.Println("====DEKLARASI MULTI VARIABEL====")
	a, b, c := 1, 2, 3
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(a, b, c)
	fmt.Println(one,isFriday,twoPointTwo,say)
	fmt.Println("========================================")
}