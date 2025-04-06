package main

import "fmt"

// Fungsi untuk menghitung perkalian secara rekursif
func multiply(n int, m int) int {
	if m == 0 {
		return 0
	}
	return n + multiply(n, m-1)
}

func main() {
	var n_103112400055, m int

	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n_103112400055)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)

	hasil := multiply(n_103112400055, m)

	fmt.Printf("Hasil dari %d x %d = %d\n", n_103112400055, m, hasil)
}
