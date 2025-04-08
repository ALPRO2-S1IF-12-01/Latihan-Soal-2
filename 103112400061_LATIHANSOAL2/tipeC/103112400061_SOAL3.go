package main

import "fmt"

func main() {
	var n_103112400061, m int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n_103112400061)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)
	fmt.Printf("Hasil dari %v x %v = %v\n", n_103112400061, m, perkalianJumlah(n_103112400061, m))
}

func perkalianJumlah(n, m int) int {
	if (m == 0) {
		return 0
	}
	return (n + perkalianJumlah(n, m-1))
}