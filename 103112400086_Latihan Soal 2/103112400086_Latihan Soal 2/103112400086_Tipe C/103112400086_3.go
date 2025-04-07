package main

import "fmt"

//rekursif u/menghitung perkalian tanpa '*'
func perkalian(a, b int) int {
	if b == 0 {
		return 0
	}
	return a + perkalian(a, b-1)
}

func main() {
	var n, m int
	fmt.Print("masukkan angka pertama (n): ")
	fmt.Scan(&n)
	fmt.Print("masukkan angka kedua (m): ")
	fmt.Scan(&m)

	hasil := perkalian(n, m)
	fmt.Printf("hasil dari %d x %d = %d\n", n, m, hasil)
	fmt.Print("Sheila Stephanie A [103112400086]")
}
