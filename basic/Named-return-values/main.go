package main

import "fmt"

func kalkulator(a, b int) (kali, bagi, modulus int) {

	kali = a * b
	bagi = a / b
	modulus = a % b

	// here you have return keyword
	// without any resultant parameters
	return
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Belajar named return values")
	kali, bagi, modulus := kalkulator(105, 7)

	fmt.Println("105 x 7 = ", kali)
	fmt.Println("105 / 7 = ", bagi)
	fmt.Println("105 % 7 = ", modulus)

	fmt.Println(split(20))
}
