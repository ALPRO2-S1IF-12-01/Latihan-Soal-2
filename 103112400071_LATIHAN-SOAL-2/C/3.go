// Raihan Adi Arba_103112400071
package main

import "fmt"

func perkalian(n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}
	if m < 0 {
		return -perkalian(n, -m)
	}
	return n + perkalian(n, m-1)
}

func main() {
	var n, m int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)

	hasil := perkalian(n, m)
	fmt.Printf("Hasil dari %d x %d = %d\n", n, m, hasil)
}
