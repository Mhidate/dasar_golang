package main

import (
	"fmt"
)

func main() {
	fmt.Println("Type data number")
	var a int = 100
	var b uint = 200
	var c float64 = 10.5
	var d complex64 = complex(3, 4)

	fmt.Println("int:", a)
	fmt.Println("uint:", b)
	fmt.Println("float64:", c)
	fmt.Println("complex64:", d)

}
