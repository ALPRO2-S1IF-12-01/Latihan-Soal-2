// M.HANIF AL FAIZ
// 103112400042
package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)
	hasil := perkalianRekursif(n, m)
	fmt.Printf("Hasil dari %d x %d = %d\n", n, m, hasil)
}
func perkalianRekursif(n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}
	if n == 1 {
		return m
	}
	if m == 1 {
		return n
	}

	// Rekursi tail-end
	if n > 0 {
		return m + perkalianRekursif(n-1, m)
	}
	// Untuk kasus bilangan negatif
	return -perkalianRekursif(-n, m)
}
