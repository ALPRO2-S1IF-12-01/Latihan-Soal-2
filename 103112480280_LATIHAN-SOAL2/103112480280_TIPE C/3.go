// Nama: Anggun Wahyu Widiyana
// NIM : 103112480280
package main

import "fmt"

// Fungsi rekursif untuk menghitung perkalian n x m menggunakan penjumlahan berulang
func perkalian(n, m, hasil_103112480280 int) int {
	if m == 0 {
		return hasil_103112480280
	}
	return perkalian(n, m-1, hasil_103112480280+n)
}

func main() {
	var n, m int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)

	result := perkalian(n, m, 0)

	fmt.Printf("Hasil dari %d x %d - %d\n", n, m, result)
}