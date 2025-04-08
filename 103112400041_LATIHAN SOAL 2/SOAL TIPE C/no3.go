package main

import "fmt"

func PerkalianPenjumlahan(n , m int) int {
	if m == 0 {
		return 0
	}
	return n + PerkalianPenjumlahan(n, m-1)
}

func main() {
	var n_103112400041, m int
	fmt.Print("Masukkan Bilangan n: ")
	fmt.Scanln(&n_103112400041)
	fmt.Print("Masukkan Bilangan m: ")
	fmt.Scanln(&m)
	fmt.Printf("Hasil dari %d x %d = %d", n_103112400041, m, PerkalianPenjumlahan(n_103112400041,m))
}