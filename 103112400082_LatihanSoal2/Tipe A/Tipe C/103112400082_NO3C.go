// Rizkina Azizah_103112400082
package main

import "fmt"

func perkalianIteratif(n, m int) int {
	hasil := 0
	for i := 0; i < n; i++ {
		hasil += m
	}
	return hasil
}

// Fungsi tail-recursive
func perkalianTailRekursif(n, m, akumulasi int) int {
	if n == 0 {
		return akumulasi
	}
	return perkalianTailRekursif(n-1, m, akumulasi+m)
}

func main() {
	var n, m int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)

	hasil1 := perkalianIteratif(n, m)
	hasil2 := perkalianTailRekursif(n, m, 0)
	fmt.Printf("Hasil perkalian iteratif %d x %d = %d\n", n, m, hasil1)
	fmt.Printf("Hasil perkalian rekursif %d x %d = %d\n", n, m, hasil2)
}
