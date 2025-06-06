package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Belajar paket strings")
	/*
		Paket strings di Go berisi fungsi-fungsi utility untuk:
			Manipulasi string
			Mencari substring
			Mengganti string
			Split string jadi slice
			Join slice jadi string
			Ubah huruf besar-kecil
			Dll
	*/
	text := "Halo Dunia"

	// Contains
	fmt.Println(strings.Contains(text, "Dunia")) // true

	// Count
	fmt.Println(strings.Count(text, "a")) // 2

	// HasPrefix & HasSuffix
	fmt.Println(strings.HasPrefix(text, "Halo"))  // true
	fmt.Println(strings.HasSuffix(text, "Dunia")) // true

	// Index & LastIndex
	fmt.Println(strings.Index(text, "a"))     // 3
	fmt.Println(strings.LastIndex(text, "a")) // 9

	// Replace
	fmt.Println(strings.Replace(text, "a", "@", -1)) // H@lo Duni@

	// Split
	parts := strings.Split(text, " ")
	fmt.Println(parts) // [Halo Dunia]

	// Join
	joined := strings.Join(parts, "-")
	fmt.Println(joined) // Halo-Dunia

	// ToUpper & ToLower
	fmt.Println(strings.ToUpper(text)) // HALO DUNIA
	fmt.Println(strings.ToLower(text)) // halo dunia

	// Trim & TrimSpace
	spaced := "   Halo Dunia   "
	fmt.Println(strings.TrimSpace(spaced)) // "Halo Dunia"

	// Repeat
	fmt.Println(strings.Repeat("Go ", 3)) // Go Go Go
}
