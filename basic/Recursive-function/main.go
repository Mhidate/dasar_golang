package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1 // base case
	} else {
		return n * factorial(n-1) // recursive call
	}
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println("Belajar recursive function")
	//Recursive function adalah fungsi yang memanggil dirinya sendiri di dalam blok kodenya

	fmt.Println("5! =", factorial(5))
	fmt.Println("Fibonacci ke-7 =", fibonacci(7))

}
